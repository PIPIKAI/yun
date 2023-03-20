package storage

import (
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/storage/drivers"
	"github.com/spf13/viper"
)

// 一些定时任务

func SendStatus(trackers []string, response interface{}) {

	for _, trackerIP := range trackers {
		util.PostJSON(trackerIP, response, nil, viper.GetDuration("DriverConfig.TimeOut"))
	}
}

func FreshStatus(driver drivers.Driver) {

}
