package crons

import (
	"time"

	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/core/tracker/models"
)

var CheckSessionStatusSpec = "*/5 * * * * *"

func CheckSessionStatus() {
	sessions, err := leveldb.GetAll[models.UploadSession]()
	if err != nil {
		logger.Logger.Error(err)
		return
	}

	for _, session := range sessions {
		if session.Status == "上传中" && time.Since(time.Unix(session.UpdataTime, 0)).Minutes() >= 1 {
			session.Status = "异常"
			fileinfo, err := leveldb.GetOne[models.File](session.FileID)
			if err != nil {
				logger.Logger.Error(err)
				return
			}
			err = leveldb.UpdataOne(session)
			if err != nil {
				logger.Logger.Error(err)
				return
			}
			err = leveldb.UpdataOne(fileinfo)
			if err != nil {
				logger.Logger.Error(err)
				return
			}
		}
	}

}
