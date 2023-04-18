package api

import (
	"context"

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
		BlockSeq  int64  `json:"block_seq"`
		RawData   []byte `json:"raw_data"`
	}
	var req Req
	if err := c.ShouldBind(&req); err != nil {
		util.Response.Error(c, nil, "参数错误")
		return
	}

	session, err := leveldb.GetOne[models.UploadSession](req.SessionID)

	if err != nil {
		util.Response.Error(c, nil, "DB Error :"+err.Error())
		return
	}

	if err := c.ShouldBind(&req); err != nil {
		util.Response.Error(c, nil, "参数错误")
		return
	}
	//  client.upload()
	rpc_res, err := Dial(session.Storage.ServerAddr, func(client pb.StorageClient) (interface{}, error) {

		return client.Upload(context.Background(), &pb.UploadRequest{
			SessionId: req.SessionID,
			BlockId:   req.BlockSeq,
			RawData:   req.RawData,
		})
	})
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	res := rpc_res.(*pb.UploadReply)

	session.Percent += (1.0 / float32(session.BlockSize))

	err = leveldb.UpdataOne(session)
	if err != nil {
		util.Response.Error(c, nil, err.Error())
		return
	}
	// res.
	util.Response.Success(c, gin.H{"md5": res.Md5}, "success")
}
