package svc

import (
	"context"

	"github.com/pipikai/yun/pb"
)

func (s *Server) PreUpload(ctx context.Context, in *pb.PreUploadRequest) (*pb.PreUploadReply, error) {

	return &pb.PreUploadReply{}, nil
}
