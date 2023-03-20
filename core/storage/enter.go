package storage

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/consts"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/common/schedule"
	"github.com/pipikai/yun/common/util"
	"github.com/pipikai/yun/core/storage/config"
	"github.com/pipikai/yun/models"
	"github.com/soheilhy/cmux"
	"github.com/spf13/viper"
)

type storage struct {
	storageConfig config.StorageConfig
	grpc          Server
	status        string
	g             *gin.Engine
}

func (s *storage) InitDriver() {
	s.grpc.Driver = nil
	s.status = "work"
}

func (s *storage) ReportStatus() {
	var data models.ServerReport

	data.Group = s.storageConfig.Group
	data.IpAddr = s.storageConfig.ServerAddr
	data.Status = s.status
	data.Driver = s.storageConfig.DriverName

	logger.Logger.Info(data)
	for _, ip := range s.storageConfig.Trackers {
		if res, err := util.PostJSON(ip+"/report-status", data, nil, consts.TimeOut); err != nil {
			logger.Logger.Errorf("Report Err :%s \n err : %v", string(res), err)
		}
	}

}

func (s *storage) FreshDriver() {
	logger.Logger.Info("FreshDriver")
}
func Run() {
	l, err := net.Listen("tcp", viper.GetString("ListenOn"))
	if err != nil {
		logger.Logger.Error(err)
		return
	}
	m := cmux.New(l)
	s := &storage{
		grpc:          Server{},
		storageConfig: *config.NewStorageConfig(),
		g:             gin.Default(),
	}

	s.InitDriver()

	schedule.StartCronTask(consts.ReportSchedule, s.ReportStatus)
	schedule.StartCronTask(consts.FreshSchedule, s.FreshDriver)
	go s.grpc.RpcServer(m)
	go s.StartHTTP(m)

	m.Serve()
}
