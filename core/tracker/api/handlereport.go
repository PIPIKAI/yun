package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/core/tracker/models"
)

func HanldeStorageServerReport(c *gin.Context) {
	var req models.ServerReport
	c.ShouldBind(&req)

	nowStorage := models.Storage{
		Group:      req.Group,
		ServerAddr: req.IpAddr,
		Cap:        req.Cap,
		Status:     req.Status,
		UpdataTime: time.Now().Unix(),
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
