package api

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/common/strategy"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/tracker/models"
	"github.com/pipikai/yun/pb"
)

type BeforeUploadReq struct {
	Group     string   `json:"group"`
	Filename  string   `json:"filename"`
	Size      int64    `json:"size"`
	Md5       string   `json:"md5"`
	ModTime   int64    `json:"mod_time"`
	BlockSize int64    `json:"block_size"`
	Blocks    []byte   `json:"blocks"`
	BlockMd5  []string `json:"block_md5s"`
}

func BeforeUpload(c *gin.Context) {
	type Res struct {
		SessionID   string
		BlockStatus []bool `json:"blocks"`
	}
	var req BeforeUploadReq

	if err := c.ShouldBind(&req); err != nil {
		util.Response.Error(c, nil, "参数错误")
		return
	}

	// select gourp
	group, err := leveldb.GetOne[models.Group](req.Group)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	logger.Logger.Info(group)

	storage, err := SelectStorage(c, group)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	if storage.Cap < req.Size {
		util.Response.Error(c, nil, "rest cap is less than free ")
		return
	}

	res := &Res{
		BlockStatus: make([]bool, 0),
	}
	session := &models.UploadSession{
		ID:          strategy.GenFileUid(req.Md5, time.UnixMicro(req.ModTime)),
		Storage:     *storage,
		BlockSize:   req.BlockSize,
		BlockMD5:    req.BlockMd5,
		CreatedTime: time.Now(),
		Status:      "uploading",
		FileName:    req.Filename,
		Size:        req.Size,
	}

	// 秒传
	irpc_res, err := Dial(session.Storage.ServerAddr, func(client pb.StorageClient) (interface{}, error) {
		return client.PreUpload(context.Background(), &pb.PreUploadRequest{
			SessionId: session.ID,
			Md5:       req.Md5,
			BlockMd5:  req.BlockMd5,
		})
	})
	if err != nil {
		util.Response.Error(c, nil, "rpc PreUpload error")

	}
	rpc_res := irpc_res.(*pb.PreUploadReply)
	if rpc_res.Code == 2 {
		util.Response.ResponsFmt(c, 200, 2, gin.H{}, "file founded 可以秒传")
		session.Status = "秒传成功"
		if err := leveldb.UpdataOne(session); err != nil {
			util.Response.Error(c, nil, err.Error())
			return
		}
		return
	}

	res.BlockStatus = rpc_res.Blockstatus
	res.SessionID = session.ID
	util.Response.ResponsFmt(c, 200, 1, gin.H{"data": res}, "")

}
