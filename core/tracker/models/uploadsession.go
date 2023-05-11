package models

// k : block md5 , v :
var UploadSessionDB = "upload_session_db"

// UploadSession
type UploadSession struct {
	ID          string  `json:"id"`
	FileID      string  `json:"file_id"`
	CreatedTime int64   `json:"cteated_time"`
	UpdataTime  int64   `json:"update_time"`
	Status      string  `json:"status"`
	Percent     float32 `json:"percent"`
	BlockSize   int64   `json:"block_size"`
}

// GetStatus
//
//	@receiver u
//	@return string
func (u *UploadSession) GetStatus() string {
	return u.Status
}

// GetDB
//
//	@receiver d
//	@return string
func (d UploadSession) GetDB() string {
	return UploadSessionDB
}

// GetID
//
//	@receiver d
//	@return string
func (d UploadSession) GetID() string {
	return d.ID
}

// UpdataPercent
//
//	@receiver u
//	@param blockStatus
func (u *UploadSession) UpdataPercent(blockStatus []bool) {
	cnt := 0
	for _, v := range blockStatus {
		if v {
			cnt = cnt + 1
		}
	}
	u.Percent = 100.0 * float32(cnt) / float32(len(blockStatus))
}

// AddPercent
//
//	@receiver u
func (u *UploadSession) AddPercent() {
	u.Percent = u.Percent + 100*(1.0/float32(u.BlockSize))

}
