package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/tracker/models"
)

// ReUploadSession
//
//	@param c
func ReUploadSession(c *gin.Context) {
	type Req struct {
		SessionID string   `json:"session_id"`
		BlockMd5s []string `json:"block_md5s"`
	}
	var req Req
	err := c.ShouldBind(&req)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	session, err := leveldb.GetOne[models.UploadSession](req.SessionID)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	// fileinfo, err := leveldb.GetOne[models.File](session.FileID)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	// 秒传
	// irpc_res, err := Dial(fileinfo.Storage.ServerAddr, func(client pb.StorageClient) (interface{}, error) {
	// 	return client.PreUpload(context.Background(), &pb.PreUploadRequest{
	// 		SessionId: session.ID,
	// 		Filemata: &pb.FileMeta{
	// 			Size:    fileinfo.Size,
	// 			Name:    fileinfo.Name,
	// 			ModTime: fileinfo.ModTime,
	// 			Md5:     fileinfo.Md5,
	// 		},
	// 		BlockMd5: req.BlockMd5s,
	// 	})
	// })
	if err != nil {
		util.Response.Error(c, nil, "rpc PreUpload error")
		return
	}
	// rpc_res := irpc_res.(*pb.PreUploadReply)
	// logger.Logger.Debug(rpc_res)
	session.Status = "上传中"
	session.UpdataTime = time.Now().Unix()

	// session.UpdataPercent(rpc_res.Blockstatus)

	err = leveldb.UpdataOne(session)

	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	// util.Response.Success(c, gin.H{"data": rpc_res.Blockstatus}, "ok")
}
