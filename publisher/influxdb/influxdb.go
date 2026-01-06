package influxdb

import (
	"context"

	iinfluxdb "github.com/bernardolm/sensors-publisher-go/infrastructure/influxdb"
)

type influxdb struct {
	client *iinfluxdb.Client
}

func New(_ context.Context, client *iinfluxdb.Client) (*influxdb, error) {
	return &influxdb{
		client: client,
	}, nil
}
