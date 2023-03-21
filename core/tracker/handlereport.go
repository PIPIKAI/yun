package tracker

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/consts"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/models"
)

func (t *tracker) HanldeStorageServerReport(c *gin.Context) {
	var req models.ServerReport
	c.ShouldBind(&req)

	nowStorage := models.Storage{
		Group:      req.Group,
		ServerAddr: req.IpAddr,
		Cap:        req.Cap,
		Status:     req.Status,
		UpdataTime: time.Now().Unix(),
	}

	// 更新数据库
	ldb, err := leveldb.NewLDB(consts.Group_Storage_DB)
	if err != nil {
		logger.Logger.Errorf("Db  connect: %v", err)
		return
	}

	err = ldb.UpdateGroup(nowStorage)
	if err != nil {
		logger.Logger.Errorf("Db  Update Grop err: %v", err)
	}
}
