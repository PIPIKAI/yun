package tracker

import (
	"context"
)

type svc struct {
	ctx context.Context
}

func NewSvc() *svc {
	return &svc{
		ctx: context.Background(),
	}
}
