package strategy

import (
	"time"
)

func GenFileUid(md5 string, modTime time.Time) string {
	return md5 + modTime.GoString()
}
