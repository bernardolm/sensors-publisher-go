package mqtt

import "context"

type mqtt struct{}

func New(_ context.Context) (*mqtt, error) {
	return &mqtt{}, nil
}
