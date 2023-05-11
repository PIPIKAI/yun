package api

import (
	"context"
	"io"
	"strconv"
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
func Upload(c *gin.Context) {
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
	session, err := leveldb.GetOne[models.UploadSession](session_id)

	if err != nil {
		util.Response.Error(c, nil, "DB Error :"+err.Error())
		return
	}
	fileinfo, err := leveldb.GetOne[models.File](session.FileID)
	if err != nil {
		util.Response.Error(c, nil, "DB Error :"+err.Error())
		return
	}

	//  client.upload()
	rpc_res, err := Dial(fileinfo.Storage.ServerAddr, func(client pb.StorageClient) (interface{}, error) {
		return client.Upload(context.Background(), &pb.UploadRequest{
			FileId:  fileinfo.Md5,
			RawData: file_raw,
			BlockId: int64(block_seq),
		})
	})
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	res := rpc_res.(*pb.UploadReply)

	session.UpdataTime = time.Now().Unix()
	session.AddPercent()
	err = leveldb.UpdataOne(session)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	util.Response.Success(c, gin.H{"md5": res.Md5}, "success")
}
