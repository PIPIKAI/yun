package baidu_netdisk

import "github.com/pipikai/yun/core/storage/drivers/vo"

// Addition
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

// DefaultDir
var DefaultDir = "/yun"

// New
//
//	@return *BaiduNetdisk
func New() *BaiduNetdisk {
	// op.RegisterDriver(func() driver.Driver {
	// 	return &BaiduNetdisk{}
	// })
	return &BaiduNetdisk{}
}
