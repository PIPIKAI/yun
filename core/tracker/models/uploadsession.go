package models

import (
	"github.com/pipikai/yun/common/leveldb"
)

// k : block md5 , v :
var UploadSessionDB = "upload_session_db"

// UploadSession
type UploadSession struct {
	ID          string  `json:"id"`
	FileID      string  `json:"file_id"`
	CreatedTime int64   `json:"cteated_time"`
	UpdataTime  int64   `json:"update_time"`
	Status      string  `json:"status"`
	BlockSize   int64   `json:"block_size"`
	Percent     float32 `json:"percent"`
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

func (u *UploadSession) GetPercent() (float32, error) {
	fileinfo, err := leveldb.GetOne[File](u.FileID)
	if err != nil {
		return 0, err
	}
	ok := 0
	for _, md5 := range fileinfo.BlockMd5 {
		v, err := leveldb.GetOne[BlockStorage](md5)
		if err == nil && len(v.Mark) > 0 {
			ok++
		}
	}
	return 100.0 * float32(ok) / float32(u.BlockSize), nil
}
