package storage

import (
	"context"
	"encoding/json"
	"io"
	"log"

	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/storage/config"
	"github.com/pipikai/yun/core/storage/drivers"
	"github.com/pipikai/yun/core/storage/drivers/vo"
	"github.com/pipikai/yun/pb"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedStorageServer
	Driver vo.Driver
}

func (s *Server) InitDriver(config *config.StorageConfig) error {
	dr := drivers.GetDriver(config.DriverName)
	if dr == nil {
		logger.Logger.DPanicf("Driver %s Not Found", config.DriverName)
	}
	addtion, err := json.Marshal(config.DriverAddtion)
	if err != nil {
		return err
	}
	err = json.Unmarshal(addtion, dr.GetAddition())
	if err != nil {
		return err
	}

	logger.Logger.Infof("driver :%s  \nAddtion: %v", config.DriverName, dr.GetAddition())
	err = dr.Init(context.Background())
	if err != nil {
		logger.Logger.Error("driver :%s init err: %v", config.DriverName, err)

		return err
	}
	s.Driver = dr
	return nil
}
func (s *Server) HeartBeat(ctx context.Context, in *pb.HeartBeatRequest) (reply *pb.HeartBeatReply, err error) {
	// s.Driver.Link()
	return
}
func (s *Server) Upload(uploadServer pb.Storage_UploadServer) (err error) {
	// recide
	streamFile := vo.StreamFile{Content: make([]byte, 0)}
	for {
		req, err := uploadServer.Recv()
		if err != nil {
			if err == io.EOF { // 流结束会返回EOF错误
				logger.Logger.Debugf("rpc get size %v", len(streamFile.Content))

				link, err := s.Driver.Upload(context.Background(), streamFile)

				if err != nil {
					return err
				}
				link_bytes, err := util.Json.Marshal(link)
				if err != nil {
					return err
				}
				err = uploadServer.SendAndClose(&pb.UploadReply{
					Link: link_bytes,
				})
				if err != nil {
					return err
				}
				return nil
			}
			return err
		}
		streamFile.Name = req.File.FileName
		streamFile.Size = req.File.Size
		streamFile.Content = append(streamFile.Content, req.File.Content...)
	}

}
func (s *Server) Manage(ctx context.Context, in *pb.ManageRequest) (reply *pb.ManageReply, err error) {

	return
}

func (s *Server) RpcServer(m cmux.CMux) {

	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))

	gs := grpc.NewServer()
	pb.RegisterStorageServer(gs, s)
	logger.Logger.Infof("grpc server listening at %v", grpcL.Addr())
	if err := gs.Serve(grpcL); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
