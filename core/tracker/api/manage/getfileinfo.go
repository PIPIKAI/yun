package manage

import (
	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/tracker/models"
)

// GetFileInfos
//
//	@param c
func GetFileInfos(c *gin.Context) {
	type Res struct {
		FileInfo []models.File `json:"file"`
	}

	var res Res

	dbdata, err := leveldb.GetAll[models.File]()

	if err != nil {
		util.Response.Error(c, nil, err.Error())
	}
	res.FileInfo = dbdata
	util.Response.Success(c, gin.H{
		"data": res,
	}, "ok")
}
