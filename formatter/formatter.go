package formatter

import (
	"github.com/bernardolm/sensors-publisher-go/message"
	"github.com/bernardolm/sensors-publisher-go/sensor"
)

type Formatter interface {
	Build(s sensor.Sensor) ([]message.Message, error)
}
