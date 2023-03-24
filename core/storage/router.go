package storage

import (
	"github.com/pipikai/yun/common/logger"
	"github.com/soheilhy/cmux"
)

func (s *storage) StartHTTP(m cmux.CMux) {
	httpL := m.Match(cmux.HTTP1Fast())
	logger.Logger.Info(s.Config.DriverAddtion["rootpath"])

	s.g.Use(s.Redirect())
	if s.Config.DriverName == "Local" {
		s.g.Static("", s.Config.DriverAddtion["rootpath"])
	}
	s.g.RunListener(httpL)
}
