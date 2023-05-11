package models

import (
	"github.com/pipikai/yun/common/models"
)

// StorageFileDb
const StorageFileDb = "tracker_file_db"

// File
type File struct {
	PreID           string `json:"pre_id"`
	ID              string `json:"id"`
	models.FileMeta `json:"file_meta"`
	Storage         Storage `json:"storage"`
	Name            string  `json:"name"`
	// 0: 正在上传 1:上传完成 -1:已经删除
	Status      int    `json:"status"`
	Path        string `json:"path"`
	Dir         bool   `json:"dir"`
	Link        *Link  `json:"link"`
	CreatedTime int64  `json:"created_time"`
}

func (d File) GetPath() string {
	return d.Path
}

func (d File) GetDB() string {
	return StorageFileDb
}

func (d File) GetID() string {
	return d.GetMd5() + ";" + d.Storage.GetClientKey()
}

func (d File) GetMd5() string {
	return d.Md5
}

func (d File) GetName() string {
	return d.Name
}
