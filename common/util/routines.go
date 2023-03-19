package util

import (
	"context"
	"time"

	"github.com/PIPIKAI/gomdfs/common/errs"
)

func WithTimeoutRoutine(td time.Duration, args interface{}, fc func(interface{}) (interface{}, error)) (interface{}, error) {
	ctx, cancer := context.WithTimeout(context.Background(), td)
	defer cancer()
	c_result := make(chan interface{})
	c_err := make(chan error)
	go func() {
		res, err := fc(args)
		if err != nil {
			c_err <- err
		} else {
			c_result <- res
		}
	}()
	select {
	case <-ctx.Done():
		return nil, errs.TimeOut
	case e := <-c_err:
		return nil, e
	case res := <-c_result:
		return res, nil
	}
}
