package svc

import (
	"context"

	"github.com/pipikai/yun/pb"
)

func (s *Server) Download(ctx context.Context, in *pb.DownloadRequest) (*pb.DownloadReply, error) {
	reply, err := s.Driver.Download(ctx, in)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
