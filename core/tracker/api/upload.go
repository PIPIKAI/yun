package api

import (
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/tracker/models"
	"github.com/pipikai/yun/pb"
)

// Upload logic :
// select Group -> select storage -> get driver -> driver upload ->
// get Link , gen token -> save db token:Link , token:FileInfo
//

func Upload(c *gin.Context) {
	type Req struct {
		SessionID string `json:"session_id"`
		BlockSeq  int    `json:"block_seq"`
		RawData   []byte `json:"raw_data"`
	}
	var req Req
	if err := c.ShouldBind(&req); err != nil {
		util.Response.Error(c, nil, "参数错误")
		return
	}
	session_db, err := leveldb.NewLDB(models.UploadSessionDB)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	var session models.UploadSession
	session_data, err := session_db.Do(req.SessionID)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	err = json.Unmarshal(session_data, &session)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}

	//  client.upload()
	_, err = Dial(session.Storage.ServerAddr, func(client pb.StorageClient) (interface{}, error) {

		stream, err := client.Upload(context.Background())
		if err != nil {
			return nil, err
		}

		err = stream.Send(&pb.UploadRequest{
			File: &pb.File{
				Content: req.RawData,
			},
		})
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	//  client.upload()

	// client.create()

	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	// res.

}
