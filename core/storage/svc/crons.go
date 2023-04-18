package svc

import (
	"github.com/pipikai/yun/common/consts"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/common/util"
)

var ReportSchedule = "*/5 * * * * *"

type ReportReq struct {
	Group  string `json:"group"`
	IpAddr string `json:"ip_addr"`
	Status string `json:"status"`
	Driver string `json:"driver"`
	Cap    int64  `json:"cap"`
}

func (s *Server) ReportStatus() {
	var data ReportReq

	data.Group = s.Config.Group
	data.IpAddr = s.Config.ListenOn
	data.Status = s.Status
	data.Driver = s.Config.DriverName

	for _, ip := range s.Config.Trackers {
		if _, err := util.PostJSON(ip+"/report-status", data, nil, consts.TimeOut); err != nil {
			logger.Logger.Warnf("Report Err :%s ", string(ip))
		}
	}

}
