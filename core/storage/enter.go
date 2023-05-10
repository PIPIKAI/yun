package storage

import (
	"github.com/pipikai/yun/core/storage/svc"
)

func Run() {
	svc := svc.NewSVC()
	go svc.HTTPServer()
	svc.RpcServer()
}
