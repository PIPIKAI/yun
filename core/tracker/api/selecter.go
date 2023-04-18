package api

import (
	"crypto/md5"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/core/tracker/models"
)

func SelectMaxCapGroup() (*models.Group, error) {
	var res models.Group

	groups, err := leveldb.GetAll[models.Group]()

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

func SelectStorage(c *gin.Context, group models.Group) (*models.Storage, error) {
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
