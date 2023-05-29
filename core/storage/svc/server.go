package svc

import (
	"log"
	"net"

	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/core/storage/api"
	"github.com/pipikai/yun/pb"
	"google.golang.org/grpc"
)

// // HeartBeat
// //  @receiver s
// //  @param ctx
// //  @param in
// //  @return reply
// //  @return err
// func (s *Server) HeartBeat(ctx context.Context, in *pb.HeartBeatRequest) (reply *pb.HeartBeatReply, err error) {
// 	return
// }

// // Manage
// //  @receiver s
// //  @param ctx
// //  @param in
// //  @return reply
// //  @return err
// func (s *Server) Manage(ctx context.Context, in *pb.ManageRequest) (reply *pb.ManageReply, err error) {

// 	return
// }

// RpcServer
//
//	@receiver s
func (s *Server) RpcServer() {
	// grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))

	lis, err := net.Listen("tcp", s.Config.RpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gs := grpc.NewServer(
		grpc.MaxRecvMsgSize(20*1024*1024),
		grpc.MaxSendMsgSize(20*1024*1024),
	)
	pb.RegisterStorageServer(gs, s)
	logger.Logger.Infof("grpc server listening at %v", lis.Addr())
	if err := gs.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// HTTPServer
//
//	@receiver s
func (s *Server) HTTPServer() {
	s.G.Use(api.Proxy())
	if s.Config.DriverName == "Local" {
		s.G.Static("", s.Config.DriverAddtion["rootpath"])
	}

	if err := s.G.Run(s.Config.HttpPort); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
