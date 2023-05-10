package svc

import (
	"context"
	"time"

	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/core/storage/models"
	"github.com/pipikai/yun/pb"
)

func (s *Server) Merge(ctx context.Context, in *pb.MergeRequest) (*pb.MergeReply, error) {

	file, err := leveldb.GetOne[models.File](in.Md5)
	if err != nil {
		return nil, err
	}

	reply, err := s.Driver.CreateFile(ctx, file)
	if err != nil {
		return nil, err
	}

	file.Status = true
	file.CreatedTime = time.Now().Unix()
	file.Path = reply.Path
	leveldb.UpdataOne(file)
	return reply, nil
}
