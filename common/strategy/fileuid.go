package strategy

import (
	"strconv"
	"sync"
	"time"
)

type fid struct {
	count uint64
	sync.Mutex
}

var Fid fid

func DecodeFid(s string) string {
	k, _ := strconv.Atoi(s)
	seq := strconv.Itoa(int(k % (1 << 32)))
	return time.Unix(int64(k>>32), 1).Format("15:04:05") + "_" + seq
}
func (f *fid) GenFID() string {
	f.Lock()
	defer f.Unlock()
	now := time.Now().Unix()
	f.count++
	id := uint64(now<<32) | f.count
	return strconv.Itoa(int(id))
}
