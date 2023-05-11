package util

import (
	"strconv"
	"time"
)

// GetUnixString
//
//	@return string
func GetUnixString() string {
	return strconv.Itoa(int(time.Now().Unix()))
}

// GetYMDStylePath
//
//	@return string
func GetYMDStylePath() string {
	return time.Now().Format("2006") + "/" + time.Now().Format("01") + "/" + time.Now().Format("02")
}
