package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/tracker/models"
	"github.com/pipikai/yun/pb"
)

// MergeReq
type MergeReq struct {
	SessionID string `json:"session_id"`
	Size      int64  `json:"size"`
	BlockSize int64  `json:"block_size"`
}

// Merge
//
//	@param c
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
	fileinfo, err := leveldb.GetOne[models.File](session.FileID)
	if err != nil {
		util.Response.Error(c, nil, "DB Error :"+err.Error())
		return
	}

	//  client.upload()
	rpc_res, err := Dial(fileinfo.Storage.ServerAddr, func(client pb.StorageClient) (interface{}, error) {

		return client.Merge(context.Background(), &pb.MergeRequest{
			Md5:       fileinfo.Md5,
			BlockSize: req.BlockSize,
		})
	})

	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	session.Status = "上传成功"
	session.Percent = 100
	err = leveldb.UpdataOne(session)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	res := rpc_res.(*pb.MergeReply)
	fileinfo.ID = fileinfo.GetID()
	fileinfo.Link = &models.Link{
		Path: res.Path,
	}
	fileinfo.Status = 1
	err = leveldb.UpdataOne(fileinfo)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	util.Response.Success(c, gin.H{"data": fileinfo}, "success")
}
