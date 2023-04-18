package models

import (
	"time"
)

// k : block md5 , v :
var UploadSessionDB = "upload_session_db"

var BlockDB = "block_db"

// k: md5
type Block struct {
	Md5         string
	CreatedTime time.Time
	Size        int64
	Status      string
	Path        string
}

type UploadSession struct {
	ID          string
	Storage     Storage
	CreatedTime time.Time
	// uploading done failed stoped
	BlockSize int64
	BlockMD5  []string
	Status    string
	Percent   float32
	FileName  string
	Size      int64
}

func (u *UploadSession) GetStatus() string {
	return u.Status
}
func (d UploadSession) GetDB() string {
	return UploadSessionDB
}
func (d UploadSession) GetID() string {
	return d.ID
}
