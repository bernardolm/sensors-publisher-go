package formatter

import (
	"github.com/bernardolm/sensors-publisher-go/pkg/domain/sensor"
)

type Formatter interface {
	Build(s sensor.Sensor) (any, error)
}
