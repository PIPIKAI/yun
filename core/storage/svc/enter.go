package svc

import (
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/logger"
	"github.com/pipikai/yun/common/schedule"
	"github.com/pipikai/yun/core/storage/config"
	"github.com/pipikai/yun/core/storage/drivers"
	"github.com/pipikai/yun/core/storage/drivers/vo"
	"github.com/pipikai/yun/pb"
)

// Server
type Server struct {
	pb.UnimplementedStorageServer
	Config   *config.StorageConfig
	Schedule schedule.ScheduleManage
	Status   string
	G        *gin.Engine
	Driver   vo.Driver
}

func (s *Server) InitDriver() error {
	dr := drivers.GetDriver(s.Config.DriverName)
	if dr == nil {
		logger.Logger.DPanicf("Driver %s Not Found", s.Config.DriverName)
	}
	addtion, err := json.Marshal(s.Config.DriverAddtion)
	if err != nil {
		return err
	}
	err = json.Unmarshal(addtion, dr.GetAddition())
	if err != nil {
		return err
	}

	err = dr.Init(context.Background())
	if err != nil {
		logger.Logger.Error("driver :%s init err: %v", s.Config.DriverName, err)

		return err
	}
	s.Driver = dr
	logger.Logger.Infof("driver %s init done ! \nAddtion: %v", s.Config.DriverName, dr.GetAddition())

	return nil
}

// NewSVC
//
//	@return *Server
func NewSVC() *Server {
	s := &Server{
		Config:   config.NewStorageConfig(),
		Schedule: *schedule.NewScheduleManage(),
		Status:   "staring",
		G:        gin.Default(),
	}
	err := s.InitDriver()

	if err != nil {
		s.Status = err.Error()
		panic(err)
	}
	s.Status = "work"

	s.Schedule.Add(ReportSchedule, s.ReportStatus)
	s.Schedule.StartAll()

	return s
}
