package strategy

import (
	"strconv"
)

func GenBlockUid(Md5 string, blocksize int64) string {
	return Md5 + strconv.Itoa(int(blocksize))
}
