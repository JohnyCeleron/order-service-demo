package order

import "errors"

var (
	ErrOrderNotFound = errors.New("order not found")
	ErrOrderExists   = errors.New("order exists")
)
