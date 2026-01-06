package stdout

import "context"

type stdout struct{}

func New(_ context.Context) (*stdout, error) {
	return &stdout{}, nil
}
