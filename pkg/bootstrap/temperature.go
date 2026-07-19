package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bernardolm/sensors-publisher-go/pkg/contract"
	"github.com/bernardolm/sensors-publisher-go/pkg/domain/model/measurement"
)

type TemperatureCollectorWorker struct {
	logger     contract.Logger
	device     contract.SensorDevice
	repository contract.MeasurementRepository
}

func (o *TemperatureCollectorWorker) Collect(ctx context.Context) error {
	for {
		time.Sleep(time.Second)

		val, err := o.device.Value()
		if err != nil {
			return err
		}

		fmt.Printf("TemperatureCollector.Collect.ds18b20.Value.val: %v", val)

		f, ok := val.(float64)
		if !ok {
			return errors.New("invalid value")
		}

		m := measurement.Measurement{
			Value: f,
			Class: o.device.Class(),
			Unit:  o.device.UnitOfMeasurement(),
		}

		if err := o.repository.Insert(ctx, m); err != nil {
			return err
		}
	}
}

var _ contract.Sensor = &TemperatureCollectorWorker{}

func ProvideTemperatureCollectorWorker(
	logger contract.Logger,
	ds18b20 contract.SensorDevice,
	repository contract.MeasurementRepository,
) contract.Sensor {
	obj := TemperatureCollectorWorker{
		logger:     logger,
		device:     ds18b20,
		repository: repository,
	}
	return &obj
}
