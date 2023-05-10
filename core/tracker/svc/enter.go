package svc

import (
	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/middleware"
	"github.com/pipikai/yun/common/schedule"
	"github.com/pipikai/yun/core/tracker/api/manage"
	"github.com/pipikai/yun/core/tracker/config"
	"github.com/pipikai/yun/core/tracker/crons"
)

type Svc struct {
	g         *gin.Engine
	config    *config.TrackerConfig
	schedules schedule.ScheduleManage
}

func NewSvc() *Svc {
	svc := &Svc{
		g:         gin.Default(),
		config:    config.NewTrackerConfig(),
		schedules: *schedule.NewScheduleManage(),
	}
	svc.g.Use(middleware.Cors())
	svc.g = Router(svc.g)
	manage.ManageRouter(svc.g)
	return svc
}

func (s *Svc) Server() {
	s.schedules.Add(crons.FreshStorageSpec, crons.UpdateStorageStatus)
	s.schedules.Add(crons.CheckSessionStatusSpec, crons.CheckSessionStatus)
	s.schedules.StartAll()
	s.g.Run(s.config.ListenOn)
}
