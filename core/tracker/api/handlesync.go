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
	var data common_models.SyncReport
	err := c.ShouldBind(&data)
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	logger.Logger.Info(data)
	syncSessionLk.Lock()
	sesssion, err := leveldb.GetOne[models.SyncSession](data.SessionID)
	syncSessionLk.Unlock()

	if err != nil {
		logger.Logger.Error(err)
		return
	}
	succed := 0
	for idx, dst := range sesssion.Dst {
		for _, reqStatus := range data.SyncDetails {
			if dst.Storage.ServerAddr == reqStatus.ServerAddr {
				fileinfo, _ := leveldb.GetOne[models.File](sesssion.FID)
				if int(fileinfo.BlockSize) <= reqStatus.Percent {
					reqStatus.Percent = 100
					sesssion.Dst[idx].Status = "同步完成"
					succed++
				}
			}
		}
	}
	if succed == len(sesssion.Dst) {
		sesssion.Status = "同步完成"
	}
	sesssion.UpdataAt = time.Now().Unix()
	reportlk.Lock()
	leveldb.UpdataOne(sesssion)
	reportlk.Unlock()

}
