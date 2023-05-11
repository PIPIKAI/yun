// package errs define
package errs

import (
	"errors"
)

var (
	ErrNotImplement = errors.New("not implement")
	ErrNotSupport   = errors.New("not support")
	ErrRelativePath = errors.New("access using relative path is not allowed")

	ErrMoveBetweenTwoStorages = errors.New("can't move files between two storages, try to copy")
	ErrUploadNotSupported     = errors.New("upload not supported")

	ErrMetaNotFound = errors.New("meta not found")
	ErrEmptyToken   = errors.New("empty token")
)
