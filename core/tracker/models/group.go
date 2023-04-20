package models

import (
	"time"

	"github.com/pipikai/yun/common/logger"
)

var GroupDB = "group_db"

type Storage struct {
	Group      string `json:"group"`
	ServerAddr string `json:"server_addr"`
	Status     string `json:"status"`
	Cap        int64  `json:"cap"`
	UpdataTime int64  `json:"updata_time"`
}

func (s Storage) GetClientKey() string {
	return s.Group + s.ServerAddr
}

type Group struct {
	Name     string             `json:"name"`
	Cap      int64              `json:"cap"`
	Status   string             `json:"status"`
	Storages map[string]Storage `json:"storages"`
}

func (g Group) GetValidStorages() (storages []Storage) {

	logger.Logger.Info(g.Storages)
	for _, v := range g.Storages {
		if v.Status == "work" {
			storages = append(storages, v)
		}
	}
	return
}
func (g Group) GetLongLivedStorage() (storage Storage) {
	storage.UpdataTime = time.Now().Unix()
	for _, v := range g.Storages {
		if v.UpdataTime < storage.UpdataTime && v.Status == "work" {
			storage = v
		}
	}
	return
}

func (d Group) GetDB() string {
	return GroupDB
}
func (d Group) GetID() string {
	return d.Name
}
