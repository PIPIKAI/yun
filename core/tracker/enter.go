package tracker

import (
	"github.com/pipikai/yun/core/tracker/svc"
)

func Run() {

	tracker := svc.NewSvc()

	tracker.Server()
}
