package errs

import "errors"

var (
	ErrEmptyUsername      = errors.New("username is empty")
	ErrEmptyPassword      = errors.New("password is empty")
	ErrWrongPassword      = errors.New("password is incorrect")
	ErrDeleteAdminOrGuest = errors.New("cannot delete admin or guest")
)
