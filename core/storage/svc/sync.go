package svc

import (
	"context"
	"sync"

	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/common/models"
	"github.com/pipikai/yun/common/rpc"
	"github.com/pipikai/yun/pb"
)

var syncLock sync.RWMutex

func (s *Server) BeginSync(sessions *models.SyncReport) {
	for _, md5block := range sessions.BlockMd5 {
		downloadRes, err := s.Driver.Download(context.Background(), &pb.DownloadRequest{
			Fid: sessions.FID,
			Md5: md5block,
		})
		if err != nil {
			logger.Logger.Error("BeginSync Err :", err)
			return
		}

		for idx, detail := range sessions.SyncDetails {
			_, err := rpc.Dial(detail.ServerAddr, func(client pb.StorageClient) (interface{}, error) {
				return client.Upload(context.Background(), &pb.UploadRequest{
					Fid:     sessions.FID,
					Md5:     md5block,
					RawData: downloadRes.Content,
				})
			})
			if err != nil {
				sessions.SyncDetails[idx].Status = err.Error()
			} else {
				sessions.SyncDetails[idx].Status = "正在上传"
				sessions.SyncDetails[idx].Percent++
			}
			syncLock.Lock()
			leveldb.UpdataOne(sessions)
			ReportSyncQueue = append(ReportSyncQueue, *sessions)
			syncLock.Unlock()
		}
	}

}
func (s *Server) Sync(ctx context.Context, in *pb.SyncRequest) (*pb.SyncReply, error) {
	status := make([]models.SyncDetail, 0)
	for _, reqServer := range in.Target {
		status = append(status, models.SyncDetail{
			Status:     "正在上传",
			Percent:    0,
			ServerAddr: reqServer,
		})
	}
	syncReport := &models.SyncReport{
		FID:         in.Fid,
		SessionID:   in.SessionId,
		SyncDetails: status,
		BlockMd5:    in.Md5S,
	}
	err := leveldb.UpdataOne(syncReport)
	go s.BeginSync(syncReport)
	return &pb.SyncReply{Msg: "ok"}, err
}
