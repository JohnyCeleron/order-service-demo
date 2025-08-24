package db

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrExistsKey      = errors.New("exists key")
)
