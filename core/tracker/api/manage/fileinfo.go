package manage

import (
	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/tracker/models"
)

// GetFiles
//
//	@param c
func GetFiles(c *gin.Context) {
	var res []models.File

	dbdata, err := leveldb.GetAll[models.File]()

	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	for _, v := range dbdata {
		if v.Status == 1 {
			res = append(res, v)
		}
	}
	util.Response.Success(c, gin.H{
		"data": res,
	}, "ok")
}
