package svc

// Merge
//
//	@receiver s
//	@param ctx
//	@param in
//	@return *pb.MergeReply
//	@return error
// func (s *Server) Merge(ctx context.Context, in *pb.MergeRequest) (*pb.MergeReply, error) {

// 	file, err := leveldb.GetOne[models.File](in.Md5)
// 	if err != nil {
// 		return nil, err
// 	}

// 	reply, err := s.Driver.CreateFile(ctx, file)
// 	if err != nil {
// 		return nil, err
// 	}

// 	file.Status = true
// 	file.CreatedTime = time.Now().Unix()
// 	file.Path = reply.Path
// 	leveldb.UpdataOne(file)
// 	return reply, nil
// }
