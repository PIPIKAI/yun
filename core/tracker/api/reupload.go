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

	fileinfo, err := leveldb.GetOne[models.File](session.FileID)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	if fileinfo.BlockSize != int64(len(req.BlockMd5s)) {
		util.Response.Error(c, nil, "文件错误")
		return
	}
	for i := 0; i < int(fileinfo.BlockSize); i++ {
		if fileinfo.BlockMd5[i] != req.BlockMd5s[i] {
			util.Response.Error(c, nil, "文件错误")
			return
		}
	}
	session.Status = "上传中"
	session.UpdataTime = time.Now().Unix()

	err = leveldb.UpdataOne(session)

	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	block_status := make([]bool, 0)
	for _, v := range req.BlockMd5s {
		sor, err := leveldb.GetOne[models.BlockStorage](v)
		if err == nil && len(sor.Mark) > 0 {
			block_status = append(block_status, false)
		} else {
			block_status = append(block_status, true)
		}
	}
	util.Response.Success(c, gin.H{"block_status": block_status}, "success")
}
