package main

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/pipikai/yun/common/leveldb"
	"github.com/pipikai/yun/core/tracker/models"
)

func TestUpdataOne(t *testing.T) {
	err := leveldb.UpdataOne(models.File{
		ID: "testid",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetOne(t *testing.T) {
	v, err := leveldb.GetOne[models.File]("testid")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}

func TestGetAll(t *testing.T) {
	v, err := leveldb.GetAll[models.File]()
	if err != nil {
		t.Fatal(err)
	}
	res, _ := json.Marshal(v)
	t.Log(string(res))
}

func TestDeleteOne(t *testing.T) {
	err := leveldb.DeleteOne[models.File]("testid")
	if err != nil {
		t.Fatal(err)
	}
}

func TestTimeSub(t *testing.T) {
	updatatime := time.Unix(int64(1683257910), 0)
	nowtime := time.Now()
	subT := nowtime.Sub(updatatime)
	t.Log(subT.Minutes())
}
