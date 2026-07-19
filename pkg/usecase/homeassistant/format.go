package homeassistant

// import (
// 	"encoding/json"
// 	"fmt"

// 	"github.com/bernardolm/sensors-publisher-go/pkg/domain/model/measurement"
// 	"github.com/bernardolm/sensors-publisher-go/pkg/domain/model/sensor"
// 	"github.com/bernardolm/sensors-publisher-go/pkg/usecase/homeassistant/api"
// )

// // https://www.home-assistant.io/integrations/mqtt/#device-discovery-payload

// // format builds Home Assistant discovery and state messages from persisted data.
// func format(reading *measurement.Measurement, sensor *sensor.Sensor) ([][]byte, error) {
// 	if sensor == nil {
// 		return nil, fmt.Errorf("home assistant formatter: sensor is required")
// 	}

// 	stateTopic := fmt.Sprintf(stateTopicFormat, sensor.UniqueID)
// 	configTopic := fmt.Sprintf(configTopicFormat, sensor.UniqueID)
// 	discovery := api.Discovery{
// 		Sensor: api.Sensor{
// 			DeviceClass:       api.DeviceClass(sensor.Class),
// 			EntityPicture:     sensor.Picture,
// 			Icon:              sensor.Icon,
// 			Name:              sensor.Name,
// 			Platform:          api.SensorPlatform,
// 			StateClass:        api.MeasurementStateClass,
// 			StateTopic:        stateTopic,
// 			UniqueID:          sensor.UniqueID,
// 			UnitOfMeasurement: sensor.UnitOfMeasurement,
// 			ValueTemplate:     fmt.Sprintf("{{ value_json.%s }}", sensor.Class),
// 		},
// 		Device: api.Device{
// 			Identifiers:  []string{sensor.UniqueID},
// 			Manufacturer: sensor.Manufacturer,
// 			Model:        sensor.Model,
// 			Name:         sensor.Name,
// 		},
// 		Topic: configTopic,
// 	}
// 	configPayload, err := json.Marshal(discovery)
// 	if err != nil {
// 		return nil, err
// 	}

// 	state := api.State{
// 		"availability": "online",
// 		"t":            stateTopic,
// 		"time":         reading.CollectedAt,
// 		sensor.Class:   reading.Value,
// 	}

// 	statePayload, err := json.Marshal(state)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return [][]byte{configPayload, statePayload}, nil
// }
