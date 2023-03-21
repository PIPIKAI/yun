package tracker

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/consts"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/pb"
	"github.com/spf13/viper"
)

func (t *tracker) Upload(c *gin.Context) {
	// select gourp
	group, err := t.SelectMaxCapGroup()
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	logger.Logger.Info(group)

	storage, err := t.SelectStorage(c, *group)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	f, _ := c.FormFile("file")
	tempFile := viper.GetString("TempDir") + "/_tmp/" + f.Filename
	defer os.Remove(tempFile)

	if storage.Cap != 0 && storage.Cap < f.Size {
		util.Response.Error(c, nil, "rest cap is less than free ")
		return
	}

	if err := c.SaveUploadedFile(f, tempFile); err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	content, err := os.ReadFile(tempFile)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	pbFile := &pb.File{
		FileName: f.Filename,
		Size:     f.Size,
		Content:  content,
	}

	// Todo : change to stream rpc req
	res, err := t.Dial(storage.ServerAddr, func(client pb.StorageClient) (interface{}, error) {
		return client.Upload(context.Background(), &pb.UploadRequest{
			File: pbFile,
		})
	})
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	ldb, err := leveldb.NewLDB(consts.File_List_DB)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	logger.Logger.Info(res)
	url := res.(*pb.UploadReply)
	token := strconv.Itoa(int(time.Now().Unix() + pbFile.Size))
	ldb.Do(token, []byte(url.Url))
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	util.Response.Success(c, gin.H{
		"res": "http://" + t.config.ListenOn + "/group/" + storage.Group + "/" + token,
	}, "success")
}
