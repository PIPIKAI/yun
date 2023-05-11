package errs

import (
	"errors"

	pkgerr "github.com/pkg/errors"
)

var (
	ErrObjectNotFound = errors.New("object not found")
	ErrNotFolder      = errors.New("not a folder")
	ErrNotFile        = errors.New("not a file")
)

func IsObjectNotFound(err error) bool {
	return errors.Is(pkgerr.Cause(err), ErrObjectNotFound)
}
