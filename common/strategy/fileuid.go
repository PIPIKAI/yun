package strategy

import (
	"time"
)

// GenFileUid
//
//	@param md5
//	@param modTime
//	@return string
func GenFileUid(md5 string, modTime time.Time) string {
	return md5 + modTime.GoString()
}
