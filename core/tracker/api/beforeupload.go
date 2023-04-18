package api

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/consts"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/common/strategy"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/tracker/models"
)

type BeforeUploadReq struct {
	Group     string   `json:"group"`
	Filename  string   `json:"filename"`
	Size      int64    `json:"size"`
	Md5       string   `json:"md5"`
	ModTime   int64    `json:"mod_time"`
	BlockSize int64    `json:"block_size"`
	Blocks    []byte   `json:"blocks"`
	BlockMd5  []string `json:"block_md5s"`
}

func BeforeUpload(c *gin.Context) {
	type Res struct {
		SessionID   string
		BlockStatus []bool `json:"blocks"`
	}
	var req BeforeUploadReq

	if err := c.ShouldBind(&req); err != nil {
		util.Response.Error(c, nil, "参数错误")
		return
	}

	// select gourp
	var group models.Group
	ldb, err := leveldb.NewLDB(consts.Group_Storage_DB)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	group_bytes, err := ldb.Do(req.Group)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	err = json.Unmarshal(group_bytes, &group)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	logger.Logger.Info(group)

	storage, err := SelectStorage(c, group)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	if storage.Cap < req.Size {
		util.Response.Error(c, nil, "rest cap is less than free ")
		return
	}

	res := &Res{
		BlockStatus: make([]bool, 0),
	}
	// 秒传
	// 匹配哈希
	file_db, err := leveldb.NewLDB(models.FileInfoDB)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	db_file, err := file_db.Do(strategy.GenFileUid(req.Md5, time.UnixMicro(req.ModTime)))

	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	if db_file != nil {
		util.Response.ResponsFmt(c, 200, 2, gin.H{}, "file founded 可以秒传")
		return
	}

	// 不能秒传
	// 分片
	session := &models.UploadSession{
		Storage:     *storage,
		CreatedTime: time.Now(),
		Status:      "uploading",
		FileName:    req.Filename,
		Size:        req.Size,
	}
	SessionID := req.Md5 + session.CreatedTime.GoString()

	res.SessionID = SessionID

	block_db, err := leveldb.NewLDB(models.BlockDB)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	// 上传
	for i := 0; i < int(req.BlockSize); i++ {
		b_recorde, err := block_db.Do(strategy.GenBlockUid(req.BlockMd5[i], req.BlockSize))
		if err != nil {
			util.Response.Error(c, nil, err.Error())
			return
		}
		if b_recorde != nil {
			res.BlockStatus = append(res.BlockStatus, true)
		} else {
			res.BlockStatus = append(res.BlockStatus, false)
		}
	}

	upload_session_db, err := leveldb.NewLDB(models.UploadSessionDB)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	session_bytes, err := json.Marshal(session)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	_, err = upload_session_db.Do(res.SessionID, session_bytes)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	util.Response.ResponsFmt(c, 200, 1, gin.H{"data": res}, "")

}
