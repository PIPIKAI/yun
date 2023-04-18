package storage

import (
	"net"

	"github.com/pipikai/yun/core/storage/svc"
	"github.com/soheilhy/cmux"
)

func Run() {
	svc := svc.NewSVC()
	l, err := net.Listen("tcp", svc.Config.ListenOn)
	if err != nil {
		panic(err)
	}
	m := cmux.New(l)

	go svc.RpcServer(m)
	go svc.HTTPServer(m)
	m.Serve()
}
