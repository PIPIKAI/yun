// package tracker core api
package api

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/logger"
	common_models "github.com/pipikai/yun/common/models"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/tracker/models"
	"github.com/pipikai/yun/pb"
)

// BeforeUploadReq
type BeforeUploadReq struct {
	Group     string   `json:"group"`
	Filename  string   `json:"filename"`
	Size      int64    `json:"size"`
	Md5       string   `json:"md5"`
	ModTime   int64    `json:"mod_time"`
	BlockSize int64    `json:"block_size"`
	BlockMd5  []string `json:"block_md5s"`
	Type      string   `json:"type"`
}

// BeforeUpload
//
//	@param c
func BeforeUpload(c *gin.Context) {
	type Res struct {
		SessionID   string `json:"session_id"`
		Code        int    `json:"code"`
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

	newFileinfo := &models.File{
		FileMeta: common_models.FileMeta{
			Size:    req.Size,
			ModTime: req.ModTime,
			Md5:     req.Md5,
		},
		Storage:     *storage,
		Name:        req.Filename,
		Status:      0,
		CreatedTime: time.Now().Unix(),
	}

	res := &Res{
		BlockStatus: make([]bool, 0),
	}
	session := &models.UploadSession{
		ID:          strconv.Itoa(int(time.Now().Unix())),
		FileID:      newFileinfo.GetID(),
		CreatedTime: time.Now().Unix(),
		Status:      "上传中",
		Percent:     0,
		BlockSize:   req.BlockSize,
	}

	// 秒传
	irpc_res, err := Dial(storage.ServerAddr, func(client pb.StorageClient) (interface{}, error) {
		return client.PreUpload(context.Background(), &pb.PreUploadRequest{
			SessionId: session.ID,
			Filemata: &pb.FileMeta{
				Size:    req.Size,
				Name:    req.Filename,
				ModTime: req.ModTime,
				Md5:     req.Md5,
			},
			BlockMd5: req.BlockMd5,
		})
	})
	if err != nil {
		util.Response.Error(c, nil, "rpc PreUpload error")
		return
	}
	rpc_res := irpc_res.(*pb.PreUploadReply)
	if rpc_res.Code == 2 {
		session.Status = "秒传成功"
		session.Percent = 100
	}
	res.Code = int(rpc_res.Code)
	res.BlockStatus = rpc_res.Blockstatus
	res.SessionID = session.ID

	if of, err := leveldb.GetOne[models.File](newFileinfo.GetID()); err != nil {
		err = leveldb.UpdataOne(newFileinfo)
		if err != nil {
			util.Response.Error(c, nil, err.Error())
			return
		}
	} else {
		if of.Name == req.Filename {
			util.Response.Error(c, nil, "相同的文件已存在")
			return
		}
	}

	err = leveldb.UpdataOne(session)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	util.Response.Success(c, gin.H{"data": res}, "before upload")

}
