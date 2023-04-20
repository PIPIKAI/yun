package manage

import (
	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/tracker/models"
)

func GetSession(c *gin.Context) {
	type Res struct {
		Sessions []models.UploadSession `json:"upload_session"`
	}
	var res Res
	dbdata, err := leveldb.GetAll[models.UploadSession]()

	if err != nil {
		util.Response.Error(c, nil, err.Error())
	}
	res.Sessions = dbdata
	util.Response.Success(c, gin.H{
		"data": res,
	}, "ok")
}
