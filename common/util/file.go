package util

import (
	"os"
)

func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CreateTempFile(data []byte) (*os.File, error) {

	f, err := os.Create("temp/_tmp/" + GetUnixString())
	if err != nil {
		return nil, err
	}
	_, err = f.Read(data)
	return f, err
}
