package storage

import (
	"github.com/soheilhy/cmux"
	"github.com/spf13/viper"
)

func (s *storage) StartHTTP(m cmux.CMux) {
	httpL := m.Match(cmux.HTTP1Fast())
	if viper.GetString("DriverName") == "Local" {
		s.g.Static("", viper.GetString("DriverAddtion.dir"))
	}
	s.g.RunListener(httpL)
}
