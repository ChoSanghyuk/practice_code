package go_error

import "errors"

var (
	Err1          = errors.New("CustomError1")
	Err2          = errors.New("CustomError2")
	ErrExistCheck = errors.New("ErrExistCheck")
)
