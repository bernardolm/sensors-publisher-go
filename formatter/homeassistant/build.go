package hass

import (
	"encoding/json"
	"fmt"

	"github.com/bernardolm/sensors-publisher-go/formatter/homeassistant/api"
	"github.com/bernardolm/sensors-publisher-go/infrastructure/config"
	"github.com/bernardolm/sensors-publisher-go/sensor"
	"github.com/k0kubun/pp/v3"
)

// https://www.home-assistant.io/integrations/mqtt/#device-discovery-payload

func (f *hass) Build(s sensor.Sensor) (any, error) {
	result := []any{}

	v, err := s.Value()
	if err != nil {
		return nil, err
	}

	val, ok := v.(float64)
	if !ok {
		return nil, fmt.Errorf("home assistant: sensor %v couldn't return value", s.ID())
	}

	stateTopic := fmt.Sprintf(stateTopicFormat, s.ID())

	if !f.deviceConfigSent {
		p := api.Discovery{
			Device: api.Device{
				Identifiers:  []string{s.ID(), s.Name()},
				Manufacturer: s.Manufacturer(),
				Model:        s.Model(),
				Name:         s.Name(),
			},

			Entity: api.Entity{
				Icon:              s.Icon(),
				UnitOfMeasurement: s.UnitOfMeasurement(),
			},

			Sensor: api.Sensor{
				Availability: []api.Availability{{
					PayloadAvailable:    "online",
					PayloadNotAvailable: "offline",
					Topic:               stateTopic,
					ValueTemplate:       "{{ value_json.availability }}",
				}},
				DeviceClass:               api.DeviceClass(s.Class()),
				ForceUpdate:               true,
				Platform:                  api.SensorPlatform,
				Qos:                       2,
				StateClass:                api.MeasurementStateClass,
				StateTopic:                fmt.Sprintf(stateTopicFormat, s.ID()),
				SuggestedDisplayPrecision: 1,
				UniqueID:                  s.ID(),
				ValueTemplate:             "{{ value_json." + s.Class() + " }}", // NOTE: fail when enum is unknown
			},

			Origin: api.Origin{
				Name:       "sensors-publisher-go",
				SupportURL: "https://github.com/bernardolm/sensors-publisher-go",
			},

			Topic: fmt.Sprintf(configTopicFormat, s.ID()),
		}

		if config.Get[bool]("DEBUG") {
			pp.Println(p)
		}

		b, err := json.Marshal(p)
		if err != nil {
			return nil, err
		}

		result = append(result, b)
	}

	p := api.State{
		"availability": "online",
		"t":            fmt.Sprintf(stateTopicFormat, s.ID()),
		"time":         s.Time(),
		s.Class():      val,
	}

	if config.Get[bool]("DEBUG") {
		pp.Println(p)
	}

	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	result = append(result, b)

	return result, nil
}
