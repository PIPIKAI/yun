package util

import (
	"os"
)

// Exists
//
//	@param path
//	@return bool
func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// CreateTempFile
//
//	@param data
//	@return *os.File
//	@return error
func CreateTempFile(data []byte) (*os.File, error) {

	f, err := os.Create("temp/_tmp/" + GetUnixString())
	if err != nil {
		return nil, err
	}
	_, err = f.Read(data)
	return f, err
}
