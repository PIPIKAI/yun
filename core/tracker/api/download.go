package api

import (
	"context"
	"errors"
	"fmt"
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

		c.Writer.Header().Set("Content-Type", fileinfo.Type)
		c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%s", fileinfo.Name))
		c.Writer.WriteHeader(http.StatusOK)
		groups, err := leveldb.GetOne[models.Group](fileinfo.Group)
		if err != nil {
			util.Response.Error(c, nil, "DB err")
			return
		}
		// content := make([]byte, 0)
		logger.Logger.Info(fileinfo.Type, "fileinfo size :", fileinfo.Size)

		flusher, ok := c.Writer.(http.Flusher)
		if !ok {
			// 如果 ResponseWriter 没有实现 Flusher 接口，返回 500 状态码
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		for idx := 0; idx < int(fileinfo.BlockSize); idx++ {
			ok := false
			BytesChan := make(chan []byte, 1)
			ErrChan := make(chan error)
			storageCtx, cancel := context.WithCancel(context.Background())
			// 当第一个协程完成传输，调用cancel
			defer cancel()
			// 使用协程对每一个在工作中的存储服务器进行下载，保证下载速度
			for _, storage := range groups.Storages {
				if storage.Status != "work" {
					continue
				}
				ok = true
				go func(ctx context.Context, storage models.Storage) {
					resp, err := Dial(storage.ServerAddr, func(client pb.StorageClient) (interface{}, error) {
						return client.Download(ctx, &pb.DownloadRequest{
							Md5: fileinfo.BlockMd5[idx],
						})
					})
					select {
					case <-ctx.Done():
						return
					default:
						if err == nil && resp != nil {
							BytesChan <- resp.(*pb.DownloadReply).Content
						} else {
							if err != nil {
								ErrChan <- err
							} else {
								ErrChan <- errors.New("other errs")
							}
						}
						return
					}
				}(storageCtx, storage)

			}
			if !ok {
				util.Response.Error(c, nil, "storages Not work")
				return
			}

			select {
			case <-ErrChan:
				util.Response.Error(c, nil, err.Error())
				return
			case res := <-BytesChan:
				// content = append(content, res...)
				_, err := c.Writer.Write(res)
				if err != nil {
					return
				}
				flusher.Flush()
			}
		}
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
