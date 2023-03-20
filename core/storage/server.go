package storage

import (
	"context"
	"log"

	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/core/storage/drivers"
	"github.com/pipikai/yun/pb"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedStorageServer
	Driver drivers.Driver
}

func (s *Server) HeartBeat(ctx context.Context, in *pb.HeartBeatRequest) (reply *pb.HeartBeatReply, err error) {
	s.Driver.Link()
	return
}
func (s *Server) Upload(ctx context.Context, in *pb.UploadRequest) (reply *pb.UploadReply, err error) {
	return &pb.UploadReply{
		Url:   in.File.FileName,
		Token: in.File.FileName,
	}, nil
}
func (s *Server) Manage(ctx context.Context, in *pb.ManageRequest) (reply *pb.ManageReply, err error) {

	return
}

func (s *Server) RpcServer(m cmux.CMux) {

	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))

	gs := grpc.NewServer()
	pb.RegisterStorageServer(gs, &Server{})
	logger.Logger.Infof("grpc server listening at %v", grpcL.Addr())
	if err := gs.Serve(grpcL); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
