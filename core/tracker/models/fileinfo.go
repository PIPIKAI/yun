package models

import (
	"time"
)

var FileInfoDB = "file_info_db"

type FileMeta struct {
	Size        int64
	Format      string
	Name        string
	ModTime     time.Time
	CreatedTime time.Time
	Md5         string
}

type FileInfo struct {
	FileMeta
	ID      string
	Storage string
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
