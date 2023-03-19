package core

import (
	"fmt"

	"github.com/pipikai/yun/common/config"
	"github.com/pipikai/yun/core/storage"
	"github.com/pipikai/yun/core/tracker"
	"github.com/spf13/viper"
)

func Start() {
	config.InitViper()
	switch viper.GetString("ServiceType") {
	case "storage":
		storage.Run()
	case "tracker":
		tracker.Run()
	default:
		panic(fmt.Sprintf("No Such %v ServiceType  ", viper.GetString("ServiceType")))
	}
}
