package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/logger"
	common_models "github.com/pipikai/yun/common/models"
	"github.com/pipikai/yun/core/tracker/models"
)

// HanldeStorageServerReport
//
//	@param c
func HanldeStorageServerReport(c *gin.Context) {
	var req common_models.Report
	err := c.ShouldBind(&req)
	if err != nil {
		logger.Logger.Error(err)
	}
	nowStorage := models.Storage{
		Group:        req.Group,
		ServerAddr:   req.IpAddr + req.RpcPort,
		DownloadAddr: req.IpAddr + req.HttpPort,
		Cap:          req.Cap,
		Status:       req.Status,
		UpdataTime:   time.Now().Unix(),
	}

	group, err := leveldb.GetOne[models.Group](req.Group)
	if err != nil {
		newGroup := models.Group{
			Name:   req.Group,
			Cap:    req.Cap,
			Status: req.Status,
			Storages: map[string]models.Storage{
				nowStorage.GetClientKey(): nowStorage,
			},
		}
		err = leveldb.UpdataOne(newGroup)
		if err != nil {
			logger.Logger.Errorf("Db  Update Grop err: %v", err)
			return
		}
		return
	}
	group.Storages[nowStorage.GetClientKey()] = nowStorage
	err = leveldb.UpdataOne(*group)
	if err != nil {
		logger.Logger.Errorf("Db  Update Grop err: %v", err)
	}
}
