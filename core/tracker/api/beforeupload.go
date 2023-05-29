// package tracker core api
package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/strategy"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/tracker/models"
)

// BeforeUploadReq
type BeforeUploadReq struct {
	Group     string   `json:"group"`
	Filename  string   `json:"filename"`
	Size      int64    `json:"size"`
	Md5       string   `json:"md5"`
	ModTime   int64    `json:"mod_time"`
	BlockSize int64    `json:"block_size"`
	BlockMd5  []string `json:"block_md5s"`
	Type      string   `json:"type"`
}

// BeforeUpload
//
//	@param c
func BeforeUpload(c *gin.Context) {
	type Res struct {
		SessionID string `json:"session_id"`
		FileID    string `json:"File_id"`
		Code      int    `json:"code"`
	}
	var req BeforeUploadReq

	if err := c.ShouldBind(&req); err != nil {
		util.Response.Error(c, nil, "参数错误")
		return
	}

	fileId := strategy.Fid.GenFID()
	filehashid := strategy.GenFileHash(req.Md5, req.Size)
	session := &models.UploadSession{
		ID:          fileId,
		FileID:      fileId + util.Md5hex(req.Filename),
		CreatedTime: time.Now().Unix(),
		UpdataTime:  time.Now().Unix(),
		Status:      "上传中",
		BlockSize:   req.BlockSize,
	}
	// defer leveldb.UpdataOne(session)
	// 检测是否可以秒传
	v, err := leveldb.GetOne[models.FileHash](filehashid)
	if err == nil {
		if v.Ok {
			session.Status = "秒传成功"
			leveldb.UpdataOne(session)
			util.Response.Success(c, gin.H{
				"data": Res{
					Code:      2,
					SessionID: session.ID,
					FileID:    v.FID,
				},
			}, "秒传成功")
		} else {
			util.Response.Error(c, nil, "相同文件已在上传")
		}
		return
	}

	group, err := leveldb.GetOne[models.Group](req.Group)
	if err != nil || group.Status != "work" {
		util.Response.Error(c, nil, "Group Not Avaliable")
		return
	}

	selectedStorage := group.GetMinDelayStorage()

	newFileinfo := &models.File{
		ID: session.FileID,
		FileMeta: models.FileMeta{
			Size:    req.Size,
			ModTime: req.ModTime,
			Md5:     req.Md5,
			Type:    req.Type,
		},
		Storage:    selectedStorage,
		Name:       req.Filename,
		BlockSize:  req.BlockSize,
		BlockMd5:   req.BlockMd5,
		Status:     0,
		UpdataTime: time.Now().Unix(),
	}

	err = leveldb.UpdataOne(newFileinfo)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	err = leveldb.UpdataOne(session)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	err = leveldb.UpdataOne(models.FileHash{
		ID:  filehashid,
		FID: fileId + util.Md5hex(req.Filename),
		Ok:  false,
	})
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	block_status := make([]bool, 0)
	for _, v := range req.BlockMd5 {
		sor, err := leveldb.GetOne[models.BlockStorage](v)
		if err == nil && len(sor.Mark) > 0 {
			block_status = append(block_status, false)
		} else {
			block_status = append(block_status, true)
		}
	}
	util.Response.Success(c, gin.H{"session": session, "block_status": block_status}, "ok")

}
