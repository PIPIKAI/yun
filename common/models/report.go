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
