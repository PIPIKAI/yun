package local

import "github.com/pipikai/yun/core/storage/drivers/vo"

func (l *Local) GetAddition() vo.Addition {
	return &l.Addition
}

func New() *Local {
	return &Local{}
}
