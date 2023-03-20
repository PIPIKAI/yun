package util

import (
	"strconv"
	"time"
)

func GetUnixString() string {
	return strconv.Itoa(int(time.Now().Unix()))
}
