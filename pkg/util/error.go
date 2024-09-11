package util

import "errors"

var (
	ErrBadStatus     = errors.New("bad status")
	ErrAlreadyExists = errors.New("already exists")
)
