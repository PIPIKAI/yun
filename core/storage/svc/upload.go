package svc

import (
	"context"

	"github.com/pipikai/yun/pb"
)

// Upload
//
//	@receiver s
//	@param ctx
//	@param in
//	@return *pb.UploadReply
//	@return error
func (s *Server) Upload(ctx context.Context, in *pb.UploadRequest) (*pb.UploadReply, error) {

	reply, err := s.Driver.Upload(ctx, in)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
