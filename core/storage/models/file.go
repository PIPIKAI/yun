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

func (d File) GetDB() string {
	return StorageFileDb
}

func (d File) GetID() string {
	return d.ID
}

func (d File) GetMd5() string {
	return d.Md5
}

func (d File) GetName() string {
	return d.Name
}
