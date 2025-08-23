package app

import "sync"

type Closer struct {
	mu    sync.Mutex
	funcs []func() error
}
