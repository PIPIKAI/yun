package api

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/tracker/models"
	"github.com/pipikai/yun/pb"
)

// Upload logic :
// select Group -> select storage -> get driver -> driver upload ->
// get Link , gen token -> save db token:Link , token:FileInfo
//
//	@param c
var sl sync.RWMutex

func Upload(c *gin.Context) {
	// 获取请求数据
	session_id := c.PostForm("session_id")                // 获取 URL 中的 ID 参数
	block_seq, _ := strconv.Atoi(c.PostForm("block_seq")) // 获取 URL 中的 ID 参数
	file, err := c.FormFile("file")                       // 获取上传文件
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	src, err := file.Open() // 打开上传文件
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	defer src.Close()
	file_raw, err := io.ReadAll(src) // 读取上传文件的内容为字节数组
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	hasher := md5.New()
	hasher.Write(file_raw)
	raw_md5 := hex.EncodeToString(hasher.Sum(nil))
	// 验证数据
	session, err := leveldb.GetOne[models.UploadSession](session_id)

	if err != nil {
		util.Response.Error(c, nil, "No found Session")
		return
	}
	fileinfo, err := leveldb.GetOne[models.File](session.FileID)
	if err != nil {
		util.Response.Error(c, nil, "No found File")
		return
	}

	if block_seq >= int(fileinfo.BlockSize) || raw_md5 != fileinfo.BlockMd5[block_seq] {
		util.Response.Error(c, nil, "File Err")
		return
	}
	group, err := leveldb.GetOne[models.Group](fileinfo.Group)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	upload_res := make(map[string]string)
	uploaded_one := false
	var wg sync.WaitGroup
	// 发送数据到存储服务器
	for key, storage := range group.Storages {
		if storage.Status != "work" {
			upload_res[key] = storage.Status
			continue
		}
		wg.Add(1)
		go func(storage models.Storage, key string) {
			defer wg.Done()
			err := util.Retry(3, func() error {
				_, err := Dial(storage.ServerAddr, func(client pb.StorageClient) (interface{}, error) {
					return client.Upload(context.Background(), &pb.UploadRequest{
						Md5:     raw_md5,
						RawData: file_raw,
					})
				})
				return err
			})
			if err != nil {
				upload_res[key] = err.Error()
			} else {
				uploaded_one = true
				upload_res[key] = "ok"
			}
		}(storage, key)
	}
	//  client.upload()
	wg.Wait()
	// 至少上传一个成功
	if !uploaded_one {
		util.Response.Error(c, gin.H{"data": upload_res}, "error")
	}

	sl.Lock()
	session, err = leveldb.GetOne[models.UploadSession](session_id)
	if err != nil {
		util.Response.Error(c, nil, "No found Session")
		return
	}

	session.UpdataTime = time.Now().Unix()
	session.AddPercent()
	if session.Percent >= 100 {
		session.Status = "上传成功"
	}
	err = leveldb.UpdataOne(session)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	sl.Unlock()
	util.Response.Success(c, gin.H{}, "success")
}
