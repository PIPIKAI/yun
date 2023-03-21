package tracker

import (
	"github.com/gin-gonic/gin"
	"github.com/pipikai/yun/common/consts"
	"github.com/pipikai/yun/common/schedule"
	"github.com/pipikai/yun/core/tracker/config"
)

func Run() {

	tracker := &tracker{
		Svc:    NewSvc(),
		g:      gin.Default(),
		config: config.NewTrackerConfig(),
	}

	schedule.StartCronTask(consts.FreshSchedule, tracker.FreshStorage)

	tracker.Server()
}

type tracker struct {
	Svc    *svc
	g      *gin.Engine
	config *config.TrackerConfig
}
