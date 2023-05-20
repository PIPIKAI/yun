package util

import (
	"crypto/md5"
	"encoding/hex"
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

func Md5hex(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)
	hashStr := hex.EncodeToString(hashBytes)
	return hashStr
}
