package crons

import (
	"time"

	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/core/tracker/models"
)

// FreshStorageSpec
var FreshStorageSpec = "*/10 * * * * *"
var TimeOutTime = int64(60)

// UpdateStorageStatus
func UpdateStorageStatus() {
	groups, err := leveldb.GetAll[models.Group]()
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	logger.Logger.Info("FreshStorage")
	for _, group := range groups {
		group_cap := int64(0)
		worked_status := "died"
		for k, v := range group.Storages {
			group_cap += v.Cap
			// 检测超时
			if time.Now().Unix()-v.UpdataTime > TimeOutTime && v.Status == "work" {
				new_storage := v
				new_storage.Status = "Expired"
				logger.Logger.Warnf("storage %s Expired", k)
				group.Storages[v.GetClientKey()] = new_storage
			}
			logger.Logger.Info("v: ", v)

			if v.Status == "work" {
				worked_status = "work"
			}
		}
		group.Status = worked_status
		group.Cap = group_cap
		err := leveldb.UpdataOne(group)
		if err != nil {
			logger.Logger.Error(err)
		}
	}

}
