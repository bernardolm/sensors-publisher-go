package sensor

import "time"

type Sensor interface {
	Class() string
	Icon() string
	ID() string
	Manufacturer() string
	Model() string
	Name() string
	Picture() string
	Time() time.Time
	UnitOfMeasurement() string
	Value() (any, error)
}
