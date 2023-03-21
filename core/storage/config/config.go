package config

import "github.com/spf13/viper"

type StorageConfig struct {
	Dev           bool
	ListenOn      string
	Group         string
	DriverName    string
	DriverAddtion map[string]string
	Trackers      []string
	FileSizeLimit int64
}

func NewStorageConfig() *StorageConfig {

	return &StorageConfig{
		Dev:           viper.GetBool("Dev"),
		ListenOn:      viper.GetString("ListenOn"),
		Group:         viper.GetString("Group"),
		DriverName:    viper.GetString("DriverName"),
		FileSizeLimit: viper.GetInt64("FileSizeLimit"),
		Trackers:      viper.GetStringSlice("Trackers"),
		DriverAddtion: viper.GetStringMapString("DriverAddtion"),
	}
}
