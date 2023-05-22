package demo

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestDownload(t *testing.T) {
	content := make([]byte, 16)
	for idx := 0; idx < 16; idx++ {
		BytesChan := make(chan []byte, 1)
		storageCtx, cancel := context.WithCancel(context.Background())
		defer cancel()
		// 当第一个协程完成传输，调用cancel
		for storage := 0; storage < 5; storage++ {
			go func(ctx context.Context, storage int) {
				res := []byte(strconv.Itoa(storage))
				select {
				case <-ctx.Done():
					return
				case _, ok := <-BytesChan:
					if ok {
						BytesChan <- res
					}
					return
				default:
					BytesChan <- res
					return
				}

			}(storageCtx, storage)
		}
		res := <-BytesChan
		content = append(content, res...)
	}

	fmt.Println(string(content))
}
