package storage

import (
	"context"
	"encoding/json"
	"log"

	"github.com/pipikai/yun/common/logger"
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
	addtion, err := json.Marshal(config.DriverAddtion)
	if err != nil {
		return err
	}
	err = json.Unmarshal(addtion, dr.GetAddition())
	if err != nil {
		return err
	}
	err = dr.Init(context.Background())
	if err != nil {
		return err
	}
	s.Driver = dr
	return nil
}
func (s *Server) HeartBeat(ctx context.Context, in *pb.HeartBeatRequest) (reply *pb.HeartBeatReply, err error) {
	// s.Driver.Link()
	return
}
func (s *Server) Upload(ctx context.Context, in *pb.UploadRequest) (reply *pb.UploadReply, err error) {
	link, err := s.Driver.Upload(ctx, vo.StreamFile{
		Name:    in.File.FileName,
		Size:    in.File.Size,
		Content: in.File.Content,
	})
	if err != nil {
		return nil, err
	}
	logger.Logger.Info(link.GetPath())
	return &pb.UploadReply{
		Url:   link.GetPath(),
		Token: in.File.FileName,
	}, nil
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
