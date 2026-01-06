package lcd

import (
	"context"
)

type lcd struct{}

func New(_ context.Context) (*lcd, error) {
	return &lcd{}, nil
}
