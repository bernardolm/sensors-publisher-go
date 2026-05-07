package worker

import (
	"context"
	"time"
)

type taskFunc func(any) (any, error)

type worker struct {
	delta time.Duration
	flows [][]taskFunc
}

func New(_ context.Context, delta time.Duration) *worker {
	w := worker{
		delta: delta,
	}

	if w.delta == 0 {
		w.delta = 5 * 60 * time.Second // five minutes
	}

	return &w
}
