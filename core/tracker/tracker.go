package tracker

import (
	"crypto/md5"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/consts"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/models"
	"github.com/pipikai/yun/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (t *tracker) Dial(IpAddr string, fc func(client pb.StorageClient) (interface{}, error)) (interface{}, error) {
	conn, err := grpc.Dial(IpAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Logger.Errorf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStorageClient(conn)
	return fc(client)

}
func (t *tracker) SelectMaxCapGroup() (*models.Group, error) {
	var res models.Group
	ldb, err := leveldb.NewLDB(consts.Group_Storage_DB)
	if err != nil {
		return nil, err
	}
	groups, err := ldb.GetAllGroups()

	if err != nil {
		return nil, err
	}
	for _, v := range groups {
		if res.Cap <= v.Cap && v.Status == "work" {
			res = v
		}
	}
	return &res, nil
}

func (t *tracker) SelectStorage(c *gin.Context, group models.Group) (*models.Storage, error) {
	storages := group.GetValidStorages()
	if storages == nil {
		return nil, errors.New("no avalid storage server")
	}
	// calculate ip hash , find the storage server
	signByte := []byte(c.ClientIP())
	hash := md5.New()
	hash.Write(signByte)
	md5Hex := hash.Sum(nil)
	hashIndex := int(md5Hex[len(md5Hex)-1]) % len(storages)
	vsm := storages[hashIndex]
	return &vsm, nil
}
