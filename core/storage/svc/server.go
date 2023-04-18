package svc

import (
	"context"
	"log"

	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/core/storage/api"
	"github.com/pipikai/yun/pb"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

func (s *Server) HeartBeat(ctx context.Context, in *pb.HeartBeatRequest) (reply *pb.HeartBeatReply, err error) {
	return
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

func (s *Server) HTTPServer(m cmux.CMux) {
	httpL := m.Match(cmux.HTTP1Fast())
	logger.Logger.Info(s.Config.DriverAddtion["rootpath"])

	s.G.Use(api.Proxy())
	if s.Config.DriverName == "Local" {
		s.G.Static("", s.Config.DriverAddtion["rootpath"])
	}
	s.G.RunListener(httpL)
}
