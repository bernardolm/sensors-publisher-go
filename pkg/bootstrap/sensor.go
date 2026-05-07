package bootstrap

import (
	"context"

	"github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/config"
	"github.com/bernardolm/sensors-publisher-go/pkg/domain/sensor"
	"github.com/bernardolm/sensors-publisher-go/pkg/domain/sensor/ds18b20"
	"github.com/bernardolm/sensors-publisher-go/pkg/domain/sensor/mock"
)

var (
	ds18b20Sensor sensor.Sensor
	mockSensor    sensor.Sensor
)

func InitSersors(ctx context.Context) error {
	if !config.Get[bool]("DEBUG") {
		ds, err := ds18b20.New(ctx)
		if err != nil {
			return err
		}
		ds18b20Sensor = ds[0]
	}

	mockSensor = mock.New(ctx)

	return nil
}
