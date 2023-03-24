package drivers

import (
	"github.com/pipikai/yun/core/storage/drivers/baidu_netdisk"
	"github.com/pipikai/yun/core/storage/drivers/local"
	"github.com/pipikai/yun/core/storage/drivers/vo"
)

var DriverCenter = map[string]vo.Driver{
	"Local":        local.New(),
	"BaiduNetdisk": baidu_netdisk.New(),
}

func GetDriver(DriverName string) vo.Driver {
	return DriverCenter[DriverName]
}
