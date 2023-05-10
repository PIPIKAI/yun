package svc

import (
	"context"

	"github.com/pipikai/yun/common/consts"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/common/models"
	"github.com/pipikai/yun/common/util"
)

var ReportSchedule = "*/5 * * * * *"

func (s *Server) ReportStatus() {
	data := models.Report{
		Group:    s.Config.Group,
		IpAddr:   s.Config.IpAddr,
		RpcPort:  s.Config.RpcPort,
		HttpPort: s.Config.HttpPort,
		Status:   "work",
		Driver:   s.Config.DriverName,
		Cap:      0,
	}
	data.Cap, _ = s.Driver.GetCap(context.Background())
	for _, ip := range s.Config.Trackers {
		if _, err := util.PostJSON(ip+"/report-status", data, nil, consts.TimeOut); err != nil {
			logger.Logger.Warnf("Report Err :%s ", string(ip))
		}
	}

}
