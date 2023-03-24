package models

import (
	"io"
	"net/http"
	"time"
)

type DirInfo struct {
	ID      string
	Size    int64
	Name    string
	modTime time.Time
	Md5     string
	isDir   bool
	Path    string
}

func (d DirInfo) GetID() string {
	return d.ID
}
func (d DirInfo) GetSize() int64 {
	return d.Size
}
func (d DirInfo) GetName() string {
	return d.Name
}
func (d DirInfo) ModTime() time.Time {
	return d.modTime
}
func (d DirInfo) IsDir() bool {
	return d.isDir
}
func (d DirInfo) GetPath() string {
	return d.Path
}

type IStreamFile interface {
	GetSize() int64
	GetName() string
	GetContent() []byte
}
type Link struct {
	URL        string         `json:"url"`
	Header     http.Header    `json:"header"` // needed header
	Data       io.ReadCloser  // return file reader directly
	Status     int            // status maybe 200 or 206, etc
	FilePath   *string        // local file, return the filepath
	Expiration *time.Duration // url expiration time
}
