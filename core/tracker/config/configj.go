package config

import "github.com/spf13/viper"

type TrackerConfig struct {
	Mode     string
	TempDir  string
	ListenOn string
	LevelDb  struct {
		StoreRoot string
	}
}

func NewTrackerConfig() *TrackerConfig {
	return &TrackerConfig{
		Mode:     viper.GetString("Mode"),
		TempDir:  viper.GetString("TempDir"),
		ListenOn: viper.GetString("ListenOn"),
		LevelDb: struct{ StoreRoot string }{
			StoreRoot: viper.GetString("LevelDb.storeRoot"),
		},
	}
}
