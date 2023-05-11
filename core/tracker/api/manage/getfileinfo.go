package manage

import (
	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/models"
	"github.com/pipikai/yun/common/util"
)

// GetFileInfos
//
//	@param c
func GetFileInfos(c *gin.Context) {
	type Res struct {
		FileInfo []models.FileInfo `json:"file_info"`
	}

	var res Res

	dbdata, err := leveldb.GetAll[models.FileInfo]()

	if err != nil {
		util.Response.Error(c, nil, err.Error())
	}
	res.FileInfo = dbdata
	util.Response.Success(c, gin.H{
		"data": res,
	}, "ok")
}
