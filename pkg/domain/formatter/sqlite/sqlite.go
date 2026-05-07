package sqlite

import "context"

type sqlite struct{}

func New(_ context.Context) (*sqlite, error) {
	return &sqlite{}, nil
}
