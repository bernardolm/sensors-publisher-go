package bootstrap

import (
	"context"

	"github.com/bernardolm/sensors-publisher-go/pkg/contract"
)

type SensorRepository struct {
	client contract.SQLiteClient
}

func (o *SensorRepository) Register(ctx context.Context, sensor contract.SensorDevice) error {
	o.client.DB(ctx).Exec("foobar")
	return nil
}

func ProvideSensorRepository(
	client contract.SQLiteClient,
) contract.SensorRepository {
	obj := SensorRepository{
		client: client,
	}
	return &obj
}

// --------------
