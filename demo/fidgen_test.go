package demo

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

type Fid struct {
	count uint64
	sync.Mutex
}

var rw sync.Mutex

func fs(s string) string {
	k, _ := strconv.Atoi(s)
	seq := strconv.Itoa(int(k % (1 << 32)))
	return time.Unix(int64(k>>32), 1).Format("15:04:05") + "_" + seq
}
func (f *Fid) GenID() string {
	f.Lock()
	defer f.Unlock()
	now := time.Now().Unix()
	f.count++
	id := uint64(now<<32) | f.count
	return strconv.Itoa(int(id))
}

func TestNewFid(t *testing.T) {
	var wg sync.WaitGroup
	var F Fid

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id := F.GenID()
			fmt.Println(fs(id))

		}()
	}
	wg.Wait()

	fmt.Println("###")
	time.Sleep(time.Second)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id := F.GenID()
			fmt.Println(id)
			fmt.Println(fs(id))
		}()
	}
	wg.Wait()
}
