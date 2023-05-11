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

// GetPath
//
//	@receiver d
//	@return string
func (d File) GetPath() string {
	return d.Path
}

// GetDB
//
//	@receiver d
//	@return string
func (d File) GetDB() string {
	return StorageFileDb
}

// GetID
//
//	@receiver d
//	@return string
func (d File) GetID() string {
	return d.GetMd5() + ";" + d.Storage.GetClientKey()
}

// GetMd5
//
//	@receiver d
//	@return string
func (d File) GetMd5() string {
	return d.Md5
}

// GetName
//
//	@receiver d
//	@return string
func (d File) GetName() string {
	return d.Name
}
