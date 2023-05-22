package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/tracker/models"
)

// MergeReq
type CreateReq struct {
	SessionID string `json:"session_id"`
	Size      int64  `json:"size"`
	BlockSize int64  `json:"block_size"`
}

// Create
//
//	@param c

func Create(c *gin.Context) {

	var req CreateReq
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

	lk.Lock()
	defer lk.Unlock()
	session, err = leveldb.GetOne[models.UploadSession](req.SessionID)
	if err != nil {
		util.Response.Error(c, nil, "DB Error :"+err.Error())
		return
	}
	percent, err := session.GetPercent()
	if err != nil {
		session.Status = err.Error()
	}
	session.Percent = percent
	if percent < 100 {
		util.Response.Error(c, nil, "Still Uploading")
		return
	}
	session.Status = "上传成功"
	err = leveldb.UpdataOne(session)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	fileinfo, _ = leveldb.GetOne[models.File](session.FileID)
	fileinfo.Status = 1
	err = leveldb.UpdataOne(fileinfo)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	util.Response.Success(c, gin.H{"data": fileinfo}, "success")
}
