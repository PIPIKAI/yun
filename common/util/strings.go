package util

import (
	"strconv"
	"time"
)

func GetUnixString() string {
	return strconv.Itoa(int(time.Now().Unix()))
}

func GetYMDStylePath() string {
	return time.Now().Format("2006") + "/" + time.Now().Format("01") + "/" + time.Now().Format("02")
}
