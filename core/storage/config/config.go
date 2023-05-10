package config

import "github.com/spf13/viper"

type StorageConfig struct {
	Dev           bool
	IpAddr        string
	RpcPort       string
	HttpPort      string
	Group         string
	DriverName    string
	Trackers      []string
	FileSizeLimit int64
	DriverAddtion map[string]string
}

func NewStorageConfig() *StorageConfig {

	return &StorageConfig{
		Dev:           viper.GetBool("Dev"),
		IpAddr:        viper.GetString("IpAddr"),
		RpcPort:       viper.GetString("RpcPort"),
		HttpPort:      viper.GetString("HttpPort"),
		Group:         viper.GetString("Group"),
		DriverName:    viper.GetString("DriverName"),
		FileSizeLimit: viper.GetInt64("FileSizeLimit"),
		Trackers:      viper.GetStringSlice("Trackers"),
		DriverAddtion: viper.GetStringMapString("DriverAddtion"),
	}
}
