package crons

import (
	"context"
	"time"

	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/core/tracker/api"
	"github.com/pipikai/yun/core/tracker/models"
	"github.com/pipikai/yun/pb"
)

func CheckSyncSession() {
	sessions, err := leveldb.GetAll[models.SyncSession]()
	if err != nil {
		logger.Logger.Error(err)
		return
	}

	for _, session := range sessions {
		group, _ := leveldb.GetOne[models.Group](session.Src.Group)
		if session.Status == "正在同步" && time.Since(time.Unix(session.UpdataAt, 0)).Minutes() >= 10 {
			session.Status = "异常"
			err = leveldb.UpdataOne(session)
			if err != nil {
				logger.Logger.Error(err)
				return
			}
		}
		if session.Status == "等待同步" && time.Now().Unix() >= session.BeginAt {
			if group.Status != "work" || group.Storages[session.Src.GetClientKey()].Status != "work" {
				continue
			}

			session.Status = "正在同步"

			err = leveldb.UpdataOne(session)
			if err != nil {
				logger.Logger.Error(err)
				return
			}
			fileinfo, err := leveldb.GetOne[models.File](session.FID)
			if err != nil {
				logger.Logger.Error(err)
				return
			}

			api.Dial(session.Src.ServerAddr, func(client pb.StorageClient) (interface{}, error) {
				return client.Sync(context.Background(), &pb.SyncRequest{
					SessionId: session.ID,
					Fid:       session.FID,
					Md5S:      fileinfo.BlockMd5,
					Target:    session.GetTargets(),
				})
			})

			err = leveldb.UpdataOne(session)
			if err != nil {
				logger.Logger.Error(err)
				return
			}
		}
	}
}
