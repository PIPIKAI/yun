package models

import (
	"time"
)

var FileInfoDB = "file_info_db"

type FileMeta struct {
	Size        int64     `json:"size"`
	Format      string    `json:"format"`
	Name        string    `json:"anme"`
	ModTime     time.Time `json:"modtime"`
	CreatedTime time.Time `json:"create_time"`
	Md5         string    `json:"md5"`
	Dir         bool      `json:"dir"`
}

type FileInfo struct {
	FileMeta
	ID      string `json:"id"`
	Storage string `json:"storage"`
}

func (d FileInfo) GetDB() string {
	return FileInfoDB
}
func (d FileInfo) GetID() string {
	return d.ID
}
func (d FileInfo) GetSize() int64 {
	return d.Size
}
func (d FileInfo) GetName() string {
	return d.Name
}
func (d FileInfo) GetModTime() time.Time {
	return d.ModTime
}
