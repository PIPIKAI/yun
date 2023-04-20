package main

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/common/models"
)

func TestUpdataOne(t *testing.T) {
	err := leveldb.UpdataOne(models.FileInfo{
		ID:      "testid",
		Storage: "testStorage",
		FileMeta: models.FileMeta{
			Md5:         "testMd5",
			Size:        123456,
			CreatedTime: time.Now(),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetOne(t *testing.T) {
	v, err := leveldb.GetOne[models.FileInfo]("testid")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}

func TestGetAll(t *testing.T) {
	v, err := leveldb.GetAll[models.FileInfo]()
	if err != nil {
		t.Fatal(err)
	}
	res, _ := json.Marshal(v)
	t.Log(string(res))
}

func TestDeleteOne(t *testing.T) {
	err := leveldb.DeleteOne[models.FileInfo]("testid")
	if err != nil {
		t.Fatal(err)
	}
}
