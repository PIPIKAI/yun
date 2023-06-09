// package core
package svc

import (
	"context"
	"time"

	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/common/models"
	"github.com/pipikai/yun/common/util"
)

// ReportSchedule
var ReportSchedule = "*/5 * * * * *"

var ReportSyncQueue []models.SyncReport

func (s *Server) ReportStatus() {
	data := models.Report{
		Group:    s.Config.Group,
		IpAddr:   s.Config.IpAddr,
		RpcPort:  s.Config.RpcPort,
		HttpPort: s.Config.HttpPort,
		Status:   "work",
		Driver:   s.Config.DriverName,
		Cap:      0,
		NowTime:  time.Now().UTC().UnixMilli(),
	}
	data.Cap, _ = s.Driver.GetCap(context.Background())
	for _, ip := range s.Config.Trackers {
		if _, err := util.PostJSON(ip+"/report-status", data, nil); err != nil {
			logger.Logger.Warnf("Report Err :%s ", string(ip))
		}
		if len(ReportSyncQueue) != 0 {
			for _, syncsession := range ReportSyncQueue {
				logger.Logger.Info("syncsession : %v ", syncsession)

				if _, err := util.PostJSON(ip+"/report-sync", syncsession, nil); err != nil {
					logger.Logger.Warnf("Report Err :%s ", string(ip))
				}
			}
			syncLock.Lock()
			ReportSyncQueue = ReportSyncQueue[0:0]
			syncLock.Unlock()
		}

	}

}
