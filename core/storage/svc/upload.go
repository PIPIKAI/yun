package svc

import (
	"context"

	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/core/storage/models"
	"github.com/pipikai/yun/pb"
)

func (s *Server) Upload(ctx context.Context, in *pb.UploadRequest) (*pb.UploadReply, error) {
	fileinfo, err := leveldb.GetOne[models.File](in.FileId)
	if err != nil {
		return nil, err
	}
	if fileinfo.BlockStatus[in.BlockId] {
		return &pb.UploadReply{
			Md5: fileinfo.BlockMd5[in.BlockId],
		}, nil
	}
	reply, err := s.Driver.Upload(ctx, in)
	if err != nil {
		return nil, err
	}

	fileinfo.BlockStatus[in.BlockId] = true
	err = leveldb.UpdataOne(fileinfo)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
