package api

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/tracker/models"
	"github.com/spf13/viper"
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

var syncSessionLk sync.RWMutex

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

	// 在文件上传完成时
	SyncStratage := viper.GetInt("SyncStratage")
	if SyncStratage < 0 {
		return
	}
	// 添加文件同步任务 到 队列
	group, err := leveldb.GetOne[models.Group](fileinfo.Storage.Group)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	targets := group.GetSyncStorages(fileinfo.Storage)
	if targets == nil || len(targets) <= 0 {
		logger.Logger.Warn("No Valid Sync Storages")
		return
	}
	dst := make([]models.SyncDst, 0)
	for _, t := range targets {
		dst = append(dst, models.SyncDst{
			Storage: &t,
			Status:  "等待上传",
			Percent: 0,
		})
	}
	syncSession := &models.SyncSession{
		ID:        util.GetUnixString(),
		Src:       fileinfo.Storage,
		Dst:       dst,
		FID:       fileinfo.ID,
		Status:    "等待同步",
		CreatedAt: time.Now().Unix(),
	}
	if SyncStratage == 0 {
		syncSession.BeginAt = syncSession.CreatedAt
	} else if SyncStratage == 1 {
		syncSession.BeginAt = time.Now().Add(time.Hour).Unix()
	} else {
		t := time.Now()
		// 第二天凌晨
		syncSession.BeginAt = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, 1).Unix()
	}
	syncSession.BeginAt = syncSession.CreatedAt
	syncSessionLk.Lock()
	defer syncSessionLk.Unlock()

	err = leveldb.UpdataOne(syncSession)
	if err != nil {
		util.Response.Error(c, nil, "Sync Err: "+err.Error())
		return
	}
	// util.Response.Success(c, nil, "同步队列加入成功")

}
