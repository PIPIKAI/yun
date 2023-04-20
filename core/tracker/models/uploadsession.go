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
	ID          string    `json:"id"`
	Storage     Storage   `json:"storage"`
	CreatedTime time.Time `json:"cteated_time"`
	// uploading done failed stoped
	BlockSize int64    `json:"block_size"`
	BlockMD5  []string `json:"block_md5"`
	Status    string   `json:"status"`
	Percent   float32  `json:"percent"`
	FileName  string   `json:"file_name"`
	Size      int64    `json:"size"`
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
