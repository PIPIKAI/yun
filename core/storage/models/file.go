package models

import (
	"strconv"

	"github.com/pipikai/yun/common/models"
)

// StorageFileDb
const StorageFileDb = "storage_file_db"

// File
type File struct {
	ID string
	models.FileMeta
	Name        string
	Status      bool
	Path        string
	BlockMd5    []string
	BlockStatus []bool
	CreatedTime int64
}

func (d File) GetPath() string {
	return d.Md5 + "-" + strconv.Itoa(int(d.ModTime))
}
func (d File) GetStatus() bool {
	for _, v := range d.BlockStatus {
		if !v {
			return false
		}
	}
	return true
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
	return d.ID
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
