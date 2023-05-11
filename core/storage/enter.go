package storage

import (
	"github.com/pipikai/yun/core/storage/svc"
)

// storage enter
func Run() {
	svc := svc.NewSVC()
	go svc.HTTPServer()
	svc.RpcServer()
}
