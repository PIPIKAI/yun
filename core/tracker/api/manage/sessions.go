package manage

import (
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/tracker/models"
)

// Res
type Res struct {
	Session  models.UploadSession `json:"upload_session"`
	FileInfo models.File          `json:"file_info"`
	Status   string               `json:"status"`
}

func getSession(conditon func(Res) bool) ([]Res, error) {
	res := make([]Res, 0)
	dbdata, err := leveldb.GetAll[models.UploadSession]()
	if err != nil {
		return nil, err
	}
	for _, v := range dbdata {
		fi, err := leveldb.GetOne[models.File](v.FileID)
		fi.ID = fi.GetID()
		if err != nil {
			return nil, err
		}
		element := Res{
			Session:  v,
			FileInfo: *fi,
		}
		if conditon(element) {
			res = append(res, element)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Session.CreatedTime > res[j].Session.CreatedTime
	})
	return res, nil
}

// GetUploading
//
//	@param c
func GetUploading(c *gin.Context) {
	res, err := getSession(func(r Res) bool { return true })
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	util.Response.Success(c, gin.H{
		"data": res,
	}, "ok")
}

// GetUploaded
//
//	@param c
func GetUploaded(c *gin.Context) {
	res, err := getSession(func(r Res) bool { return r.Session.Percent >= 100 })
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	util.Response.Success(c, gin.H{
		"data": res,
	}, "ok")
}

// DelSession
//
//	@param c
func DelSession(c *gin.Context) {
	type Req struct {
		SessionID string `json:"session_id"`
	}
	var req Req
	err := c.ShouldBind(&req)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	dbsession, err := leveldb.GetOne[models.UploadSession](req.SessionID)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	if dbsession.Status != "上传成功" {
		err = leveldb.DeleteOne[models.File](dbsession.FileID)
		if err != nil {
			util.Response.Error(c, nil, err.Error())
			return
		}
	}
	err = leveldb.DeleteOne[models.UploadSession](req.SessionID)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	util.Response.Success(c, nil, "ok")
}
