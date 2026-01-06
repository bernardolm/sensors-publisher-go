package influxdb

import (
	"context"
	"fmt"

	"github.com/bernardolm/sensors-publisher-go/sensor"
)

type influxdb struct{}

func (a *influxdb) Build(s sensor.Sensor) (any, error) {
	value, err := s.Value()
	if err != nil {
		return nil, err
	}

	line := fmt.Sprintf("%s,entity_id=%s %s=%f",
		s.Class(),
		s.ID(),
		s.UnitOfMeasurement(),
		value)

	return line, nil
}

func New(ctx context.Context) (*influxdb, error) {
	f := influxdb{}
	return &f, nil
}
