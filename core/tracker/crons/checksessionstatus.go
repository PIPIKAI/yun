package crons

import (
	"time"

	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/core/tracker/models"
)

// CheckSessionStatusSpec
var CheckSessionStatusSpec = "*/5 * * * * *"

// CheckSessionStatus
func CheckSessionStatus() {

	go CheckSyncSession()
	sessions, err := leveldb.GetAll[models.UploadSession]()
	if err != nil {
		logger.Logger.Error(err)
		return
	}

	for _, session := range sessions {
		if session.Status == "上传中" && time.Since(time.Unix(session.UpdataTime, 0)).Minutes() >= 1 {
			session.Status = "异常"
			err = leveldb.UpdataOne(session)
			if err != nil {
				logger.Logger.Error(err)
				return
			}
		}
	}

}
