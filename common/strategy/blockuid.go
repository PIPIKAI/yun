// gen blockid fileid strategy
package strategy

import (
	"strconv"
)

// GenBlockUid
//
//	@param Md5
//	@param blocksize
//	@return string
func GenFileHash(Md5 string, blocksize int64) string {
	return Md5 + "+" + strconv.Itoa(int(blocksize))
}
