package api

import (
	"context"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/tracker/models"
	"github.com/pipikai/yun/pb"
)

type MergeReq struct {
	SessionID string `json:"session_id"`
	Size      int64  `json:"size"`
	BlockSize int64  `json:"block_size"`
}

func Merge(c *gin.Context) {

	var req MergeReq
	if err := c.ShouldBind(&req); err != nil {
		util.Response.Error(c, nil, "参数错误")
		return
	}

	session, err := leveldb.GetOne[models.UploadSession](req.SessionID)

	if err != nil {
		util.Response.Error(c, nil, "DB Error :"+err.Error())
		return
	}
	//  client.upload()
	rpc_res, err := Dial(session.Storage.ServerAddr, func(client pb.StorageClient) (interface{}, error) {

		return client.Merge(context.Background(), &pb.MergeRequest{
			SessionId: req.SessionID,
			Size:      req.Size,
			BlockSize: req.BlockSize,
		})
	})

	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	session.Status = "上传成功"
	err = leveldb.UpdataOne(session)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	res := rpc_res.(*pb.MergeReply)
	fileinfo := models.FileInfo{
		FileMeta: models.FileMeta{
			Size:        session.Size,
			Format:      path.Ext(session.FileName),
			Name:        session.FileName,
			CreatedTime: time.Now(),
			Md5:         res.Md5,
		},
		ID:      session.ID,
		Storage: session.Storage.Group,
	}
	err = leveldb.UpdataOne(fileinfo)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	util.Response.Success(c, gin.H{"data": fileinfo}, "success")
}
