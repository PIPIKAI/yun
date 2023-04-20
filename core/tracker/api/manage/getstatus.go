package manage

import (
	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/tracker/models"
)

func GetStatus(c *gin.Context) {
	type Res struct {
		StartTime int64          `json:"start_time"`
		Group     []models.Group `json:"groups"`
	}
	var res Res

	dbdata, err := leveldb.GetAll[models.Group]()

	if err != nil {
		util.Response.Error(c, nil, err.Error())
	}
	res.StartTime = StartTime
	res.Group = dbdata

	util.Response.Success(c, gin.H{
		"data": res,
	}, "ok")
}
