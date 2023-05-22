package api

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/tracker/models"
	"github.com/pipikai/yun/pb"
)

// download logic
// get group ,token -> search db -> get link -> expire ? yes -> get group , Link() -> get link  -> proxy
func Download() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "GET" {
			c.Next()
		}
		matched, _ := regexp.MatchString("^/file/", c.Request.RequestURI)

		if !matched {
			return
		}

		fileId := strings.Split(c.Request.RequestURI, "/")[2]

		fileinfo, err := leveldb.GetOne[models.File](fileId)
		if err != nil {
			util.Response.ResponsFmt(c, http.StatusNotFound, 404, nil, "File Not Found")
			return
		}

		groups, err := leveldb.GetOne[models.Group](fileinfo.Group)
		if err != nil {
			util.Response.Error(c, nil, "DB err")
			return
		}
		logger.Logger.Info(fileinfo, groups)
		content := make([]byte, fileinfo.Size)
		for idx := 0; idx < int(fileinfo.BlockSize); idx++ {
			BytesChan := make(chan []byte, 1)
			storageCtx, cancel := context.WithCancel(context.Background())
			// 当第一个协程完成传输，调用cancel
			defer cancel()
			// 使用协程对每一个在工作中的存储服务器进行下载，保证下载速度
			for _, storage := range groups.Storages {
				if storage.Status != "ok" {
					continue
				}
				go func(storageCtx context.Context, storage models.Storage) {
					resp, err := Dial(storage.ServerAddr, func(client pb.StorageClient) (interface{}, error) {
						return client.Download(storageCtx, &pb.DownloadRequest{
							Md5: fileinfo.BlockMd5[idx],
						})
					})
					if err == nil && resp != nil {
						BytesChan <- resp.(*pb.DownloadReply).Content
					}
				}(storageCtx, storage)

			}
			// 等待第一个完成传输的storage
			content = append(content, <-BytesChan...)
			close(BytesChan)
		}
		c.Data(http.StatusOK, fileinfo.Type, content)
	}
}

// HTTPProxy ,http 反向代理
func HTTPProxy(c *gin.Context, Scheme, Host string) bool {

	remote, err := url.Parse(Scheme + "://" + Host)
	if err != nil {
		return false
	}

	logger.Logger.Info("HTTPProxy: ", remote)
	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.ServeHTTP(c.Writer, c.Request)
	return true
}
