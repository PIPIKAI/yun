package models

var sync_session_db = "sync_session"

type SyncDst struct {
	Storage *Storage
	Status  string `json:"status"`
	Percent int    `json:"percent"`
}

type SyncSession struct {
	ID  string    `json:"id"`
	Src *Storage  `json:"src"`
	Dst []SyncDst `json:"dst"`
	FID string    `json:"fid"`
	// 等待同步 ；正在同步；异常
	Status    string `json:"status"`
	CreatedAt int64  `json:"ctreated_at"`
	BeginAt   int64  `json:"begin_at"`
	UpdataAt  int64  `json:"upldata_at"`
}

func (s SyncSession) GetID() string {
	return s.ID
}

func (s SyncSession) GetDB() string {
	return sync_session_db
}

func (s SyncSession) GetTargets() []string {
	var res []string
	for _, v := range s.Dst {
		res = append(res, v.Storage.ServerAddr)
	}
	return res
}
