package tracker

import (
	"time"

	"github.com/pipikai/yun/common/consts"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/models"
)

func (t tracker) FreshStorage() {
	ldb, err := leveldb.NewLDB(consts.Group_Storage_DB)
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	logger.Logger.Info("FreshStorage")
	groups, _ := ldb.GetAllGroups()
	for _, group := range groups {
		for k, v := range group.Storages {
			if time.Now().Unix()-v.UpdataTime > 60 && v.Status == "work" {
				// groups[i].Storages[k].Status =
				Storage := models.Storage{
					Group:      v.Group,
					ServerAddr: v.ServerAddr,
					Status:     "Expired",
					Cap:        v.Cap,
					UpdataTime: v.UpdataTime,
				}
				logger.Logger.Warnf("storage %s Expired", k)
				if err := ldb.UpdateGroup(Storage); err != nil {
					logger.Logger.Error(err)
				}

			}
		}
	}

}
