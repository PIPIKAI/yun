package models

import "time"

var FileMetaDB = "FileMeta"

type FileMeta struct {
	ID          string
	Storage     string
	Size        int64
	Format      string
	Name        string
	ModTime     time.Time
	CreatedTime time.Time
	Md5         string
}
