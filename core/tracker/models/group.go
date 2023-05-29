package models

import (
	"crypto/md5"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/logger"
)

// GroupDB
var GroupDB = "group_db"

// Storage
type Storage struct {
	Group        string `json:"group"`
	ServerAddr   string `json:"server_addr"`
	DownloadAddr string `json:"download_addr"`
	// work is good else is err info
	Status     string `json:"status"`
	Cap        int64  `json:"cap"`
	Delay      int64  `json:"delay"`
	UpdataTime int64  `json:"updata_time"`
}

// GetClientKey
//
//	@receiver s
//	@return string
func (s Storage) GetClientKey() string {
	return s.Group + s.ServerAddr
}

// Group
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
func (g Group) GetMinDelayStorage() *Storage {
	var res Storage
	nowdelay := int64(99999)
	for _, v := range g.Storages {
		if v.Delay < nowdelay && v.Status == "work" {
			res = v
		}
	}
	return &res
}
func (d Group) GetDB() string {
	return GroupDB
}

func (d Group) GetID() string {
	return d.Name
}

func (g Group) GetStorageByHash(c *gin.Context) (*Storage, error) {
	storages := g.GetValidStorages()
	if storages == nil {
		return nil, errors.New("no avalid storage server")
	}
	// calculate ip hash , find the storage server
	signByte := []byte(c.ClientIP())
	hash := md5.New()
	hash.Write(signByte)
	md5Hex := hash.Sum(nil)
	hashIndex := int(md5Hex[len(md5Hex)-1]) % len(storages)
	vsm := storages[hashIndex]
	return &vsm, nil
}

func (g Group) GetSyncStorages(master *Storage) []Storage {
	res := make([]Storage, 0)
	var tp Storage
	nowCap := int64(0)
	storages := g.GetValidStorages()
	if len(storages) == 0 {
		return nil
	}
	for _, v := range storages {
		if v.GetClientKey() != master.GetClientKey() && v.Cap >= nowCap {
			tp = v
			nowCap = v.Cap
		}
	}
	res = append(res, tp)
	return res
}
