package models

import (
	"time"
)

type FileInfo struct {
	Storage string
	ID      string
	Size    int64
	Name    string
	modTime time.Time
	Md5     string
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
func (d FileInfo) ModTime() time.Time {
	return d.modTime
}

type IStreamFile interface {
	GetSize() int64
	GetName() string
	GetContent() []byte
}
