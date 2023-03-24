package baidu_netdisk

import "github.com/pipikai/yun/core/storage/drivers/vo"

type Addition struct {
	RefreshToken   string
	RootPath       string
	OrderBy        string
	OrderDirection string
	DownloadAPI    string
	ClientID       string
	ClientSecret   string
}

var config = vo.Config{
	Name:        "BaiduNetdisk",
	DefaultRoot: "/",
}
var DefaultDir = "/yun"

func New() *BaiduNetdisk {
	// op.RegisterDriver(func() driver.Driver {
	// 	return &BaiduNetdisk{}
	// })
	return &BaiduNetdisk{}
}
