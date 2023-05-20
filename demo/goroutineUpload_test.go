package demo

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
)

// Retry 重试指定次数的函数调用
func Retry(attempts int, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		if err = fn(); err == nil {
			return nil
		}
	}
	return fmt.Errorf("after %d attempts, last error: %s", attempts, err.Error())
}

func MockUpload() error {
	// time.Sleep(time.Millisecond * 500)
	if rand.Float64() <= 0.5 {
		return errors.New("MockUpload  wrong")
	}
	return nil
}

func TestGroutineUpload(t *testing.T) {

	// var wg sync.WaitGroup
	// 启动5个协程执行函数，并在发生错误时发送错误信息到通道中

	Errchan := make(chan error, 5)
	for i := 0; i < 5; i++ {
		go func(id int) {
			err := Retry(2, MockUpload)
			Errchan <- err
		}(i)
	}
	for i := 0; i < 5; i++ {
		res := <-Errchan
		fmt.Println(res)

		// if res != nil {
		// 	fmt.Println(res)
		// }
	}
}
