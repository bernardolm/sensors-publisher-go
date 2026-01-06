package mqtt

import (
	"context"

	imqtt "github.com/bernardolm/sensors-publisher-go/infrastructure/mqtt"
)

type mqtt struct {
	client *imqtt.Client
}

func New(_ context.Context, client *imqtt.Client) (*mqtt, error) {
	return &mqtt{
		client: client,
	}, nil
}
