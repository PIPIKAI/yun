package models

type ServerReport struct {
	Group  string `json:"group"`
	IpAddr string `json:"ip_addr"`
	Status string `json:"status"`
	Driver string `json:"driver"`
	Cap    int64  `json:"cap"`
}
