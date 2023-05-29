package models

// Report json body
type Report struct {
	Group    string `json:"group"`
	IpAddr   string `json:"ip_addr"`
	RpcPort  string `json:"rpc_port"`
	HttpPort string `json:"http_port"`
	Status   string `json:"status"`
	Driver   string `json:"driver"`
	Cap      int64  `json:"cap"`
	NowTime  int64  `json:"now_time"`
}

var commom_sync_db = "commom_sync"

type SyncDetail struct {
	ServerAddr string `json:"server_addr"`
	Status     string `json:"status"`
	Percent    int    `json:"percent"`
}

type SyncReport struct {
	SessionID   string       `json:"session_id"`
	FID         string       `json:"f_id"`
	SyncDetails []SyncDetail `json:"sync_details"`
	BlockMd5    []string     `json:"block_md5"`
}

func (s SyncReport) GetID() string {
	return s.SessionID
}
func (s SyncReport) GetDB() string {
	return commom_sync_db
}
