// package vipper init
package config

import (
	"os"

	"github.com/pipikai/yun/cmd/flags"
	"github.com/spf13/viper"
)

// InitViper
func InitViper() {
	workDir, _ := os.Getwd()
	viper.SetConfigName(flags.ConfigFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("read config file err: " + err.Error())
	}
}
