package storage

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/core/storage/config"
	"github.com/soheilhy/cmux"
)

type storage struct {
	Config config.StorageConfig
	Grpc   *Server
	status string
	g      *gin.Engine
}

func (s *storage) FreshDriver() {

	err := s.Grpc.InitDriver(&s.Config)
	if err != nil {
		s.status = err.Error()
	}
	logger.Logger.Info("FreshDriver status:", s.status)
}
func Run() {
	config := *config.NewStorageConfig()
	l, err := net.Listen("tcp", config.ListenOn)
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	m := cmux.New(l)
	s := &storage{
		Grpc:   &Server{},
		Config: config,
		g:      gin.Default(),
	}

	if err := s.Grpc.InitDriver(&s.Config); err != nil {
		s.status = err.Error()
	}
	s.status = "work"
	// schedule.StartCronTask(consts.ReportSchedule, s.ReportStatus)
	// schedule.StartCronTask(consts.StorageFreshSchedule, s.FreshDriver)
	go s.Grpc.RpcServer(m)
	go s.StartHTTP(m)

	m.Serve()
}
