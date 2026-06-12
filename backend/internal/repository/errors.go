package repository

import "errors"

var (
	ErrPasswordDoesNotMatch = errors.New("password does not match")
	ErrDuplicateEmail       = errors.New("email already exists")
	ErrDuplicateName        = errors.New("username already exists")
)
