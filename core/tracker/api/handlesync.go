package api

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	// MergeReq

	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/logger"
	common_models "github.com/pipikai/yun/common/models"
	"github.com/pipikai/yun/core/tracker/models"
)

var reportlk sync.RWMutex

func HandleSync(c *gin.Context) {
	var req common_models.SyncReport
	err := c.ShouldBind(&req)
	if err != nil {
		logger.Logger.Error(err)
		return
	}

	sesssion, err := leveldb.GetOne[models.SyncSession](req.SessionID)
	if err != nil {
		logger.Logger.Error(err)
		return
	}

	for idx, dst := range sesssion.Dst {
		for _, reqStatus := range req.SyncDetails {
			if dst.Storage.ServerAddr == reqStatus.ServerAddr {
				sesssion.Dst[idx].Percent = reqStatus.Percent
				sesssion.Dst[idx].Status = reqStatus.Status
			}
		}
	}
	sesssion.UpdataAt = time.Now().Unix()
	reportlk.Lock()
	leveldb.UpdataOne(sesssion)
	reportlk.Unlock()

}
