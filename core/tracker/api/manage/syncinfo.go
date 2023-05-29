package manage

import (
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/tracker/models"
)

// Res
type SyncInfoRes struct {
	Session  models.SyncSession `json:"upload_session"`
	FileInfo models.File        `json:"file_info"`
	Status   string             `json:"status"`
}

func getSyncSession(conditon func(SyncInfoRes) bool) ([]SyncInfoRes, error) {
	res := make([]SyncInfoRes, 0)
	dbdata, err := leveldb.GetAll[models.SyncSession]()
	if err != nil {
		return nil, err
	}
	for _, v := range dbdata {
		fi, err := leveldb.GetOne[models.File](v.FID)
		fi.ID = fi.GetID()
		if err != nil {
			return nil, err
		}

		element := SyncInfoRes{
			Session:  v,
			FileInfo: *fi,
		}
		if conditon(element) {
			res = append(res, element)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Session.CreatedAt > res[j].Session.CreatedAt
	})
	return res, nil
}

// GetUploading
//
//	@param c
func GetSyncing(c *gin.Context) {
	res, err := getSyncSession(func(r SyncInfoRes) bool { return true })
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	util.Response.Success(c, gin.H{
		"data": res,
	}, "ok")
}
