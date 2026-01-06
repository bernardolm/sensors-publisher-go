package ds18b20

import (
	"fmt"
	"time"

	"github.com/bernardolm/sensors-publisher-go/infrastructure/logging"
	yryzds18b20 "github.com/yryz/ds18b20"
)

func (s *ds18b20) Value() (any, error) {
	// var value *float64

	if len(s.id) == 0 {
		return nil, fmt.Errorf("ds18b20 sensor: none devices found")
	}

	t, err := yryzds18b20.Temperature(s.id)
	if err != nil {
		return nil, err
	}

	s.time = time.Now()

	logging.Log.
		WithField("id", s.id).
		WithField("name", s.model).
		WithField("value", t).
		Debug("ds18b20 sensor: retrieved value")

	return t, nil
}
