package local

import "github.com/pipikai/yun/core/storage/drivers/vo"

// GetAddition
//
//	@receiver l
//	@return vo.Addition
func (l *Local) GetAddition() vo.Addition {
	return &l.Addition
}

// New
//
//	@return *Local
func New() *Local {
	return &Local{}
}
