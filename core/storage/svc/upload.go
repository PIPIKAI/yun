package svc

import (
	"context"

	"github.com/pipikai/yun/pb"
)

func (s *Server) Upload(ctx context.Context, in *pb.UploadRequest) (*pb.UploadReply, error) {
	return &pb.UploadReply{}, nil
}
