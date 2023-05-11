package api

import (
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Dial gRpc server
//
//	@param IpAddr
//	@param fc
//	@return interface{}
//	@return error
func Dial(IpAddr string, fc func(client pb.StorageClient) (interface{}, error)) (interface{}, error) {
	conn, err := grpc.Dial(IpAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Logger.Errorf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStorageClient(conn)
	return fc(client)

}
