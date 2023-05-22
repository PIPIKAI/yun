package demo

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"testing"
)

var sr sync.RWMutex

func TestDownload(t *testing.T) {

	content := make([]byte, 26)
	for idx := 0; idx < 10; idx++ {
		BytesChan := make(chan []byte, 1)
		storageCtx, cancel := context.WithCancel(context.Background())
		defer cancel()
		// 当第一个协程完成传输，调用cancel
		for storage := 0; storage < 5; storage++ {
			go func(ctx context.Context, storage int) {
				res := []byte(strconv.Itoa(idx))
				select {
				case <-ctx.Done():
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

func TestFile(t *testing.T) {
	content, _ := os.ReadFile("Map_.png")
	fmt.Println(len(content))
}
