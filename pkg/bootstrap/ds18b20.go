package bootstrap

import (
	"context"
	"errors"

	"github.com/bernardolm/sensors-publisher-go/pkg/contract"
	"github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/sensor/ds18b20"
)

func ProvideDs18b20SensorDevice(
	logger contract.Logger,
	repository contract.SensorRepository,
) contract.SensorDevice {
	obj, err := ds18b20.New(context.Background())
	if err != nil {
		panic(err)
	}
	if len(obj) == 0 {
		panic(errors.New("no ds18b20 devices found"))
	}

	if err := repository.Register(context.Background(), obj[0]); err != nil {
		panic(err)
	}

	return obj[0]
}
