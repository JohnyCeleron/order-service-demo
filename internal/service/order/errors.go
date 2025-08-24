package order

import "errors"

var (
	ErrOrderNotFound = errors.New("Order not found")
	ErrOrderExists   = errors.New("Order exists")
)
