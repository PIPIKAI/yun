package svc

// PreUpload
//
//	@receiver s
//	@param ctx
//	@param in
//	@return *pb.PreUploadReply
//	@return error
// func (s *Server) PreUpload(ctx context.Context, in *pb.PreUploadRequest) (*pb.PreUploadReply, error) {

// 	v, err := leveldb.GetOne[models.File](in.Filemata.Md5)
// 	file_meta := models.FileMeta{
// 		Size:    in.Filemata.Size,
// 		ModTime: in.Filemata.ModTime,
// 		Md5:     in.Filemata.Md5,
// 	}
// 	// logger.Logger.Debug("File info", v)
// 	if err == nil {
// 		if v.GetStatus() &&
// 			len(v.BlockMd5) == len(in.BlockMd5) &&
// 			v.BlockMd5[0] == in.BlockMd5[0] {
// 			return &pb.PreUploadReply{
// 				Code: 2,
// 			}, nil
// 		} else {
// 			return &pb.PreUploadReply{
// 				Code:        1,
// 				Blockstatus: v.BlockStatus,
// 			}, nil
// 		}
// 	}

// 	reply, err := s.Driver.PreUpload(ctx, in)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// logger.Logger.Debug("Driver reply ", reply)

// 	err = leveldb.UpdataOne(models.File{
// 		ID:          in.Filemata.Md5,
// 		FileMeta:    file_meta,
// 		Name:        in.Filemata.Name,
// 		Status:      false,
// 		BlockMd5:    in.BlockMd5,
// 		BlockStatus: reply.Blockstatus,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return reply, nil
// }
