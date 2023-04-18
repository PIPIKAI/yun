package models

import (
	"time"
)

// k : block md5 , v :
var UploadSessionDB = "UploadSession"

var BlockDB = "Block"

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
	Status   string
	Percent  float32
	FileName string
	Size     int64
}

func (u *UploadSession) GetStatus() string {
	return u.Status
}

// func (u *UploadSession) GetPercent() float32 {
// 	done := float32(0)
// 	for _, v := range u.BlockStatus {
// 		if v {
// 			done += 1
// 		}
// 	}
// 	return done / float32(u.BlockSize)
// }
