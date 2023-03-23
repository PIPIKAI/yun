package storage

import (
	"github.com/pipikai/yun/common/consts"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/storage/drivers/vo"
	"github.com/pipikai/yun/models"
	"github.com/spf13/viper"
)

// 一些定时任务

func SendStatus(trackers []string, response interface{}) {
	for _, trackerIP := range trackers {
		util.PostJSON(trackerIP, response, nil, viper.GetDuration("DriverConfig.TimeOut"))
	}
}

func FreshStatus(driver vo.Driver) {

}
func (s *storage) ReportStatus() {
	var data models.ServerReport

	data.Group = s.Config.Group
	data.IpAddr = s.Config.ListenOn
	data.Status = s.status
	data.Driver = s.Config.DriverName

	for _, ip := range s.Config.Trackers {
		if _, err := util.PostJSON(ip+"/report-status", data, nil, consts.TimeOut); err != nil {
			logger.Logger.Warnf("Report Err :%s ", string(ip))
		}
	}

}
