package contract

import (
	"context"
	"time"

	"github.com/bernardolm/sensors-publisher-go/pkg/domain/model/measurement"
)

type SensorDevice interface {
	Class() measurement.Class
	Icon() string
	ID() string
	Manufacturer() string
	Model() string
	Name() string
	Picture() string
	Time() time.Time
	UnitOfMeasurement() measurement.UnitOfMeasurement
	Value() (any, error)
}

type Sensor interface {
	Collect(context.Context) error
}
