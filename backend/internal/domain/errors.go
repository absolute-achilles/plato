package domain

import "errors"

// Sentinel Errors (Great for simple states)
var (
	ErrNotFound     = errors.New("record not found")
	ErrDuplicate    = errors.New("record already exists")
	ErrInvalidInput = errors.New("invalid input provided")
)
