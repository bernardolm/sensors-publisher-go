package homeassistant

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"

// 	"github.com/bernardolm/sensors-publisher-go/pkg/domain/model/measurement"
// 	"github.com/bernardolm/sensors-publisher-go/pkg/domain/model/sensor"
// )

// // publish formats and sends one weather measurement to MQTT.
// func (u *UseCase) publish(ctx context.Context, reading *measurement.Measurement, sensor *sensor.Sensor) error {
// 	content, err := format(reading, sensor)
// 	if err != nil {
// 		return fmt.Errorf("home assistant: format MQTT payload: %w", err)
// 	}

// 	return u.publishContent(ctx, content)
// }

// // publishContent sends all formatted MQTT messages in order.
// func (u *UseCase) publishContent(ctx context.Context, content [][]byte) error {
// 	for _, payload := range content {
// 		m := message{Qos: 2}
// 		if err := json.Unmarshal(payload, &m); err != nil {
// 			return err
// 		}

// 		if err := u.client.Publish(
// 			ctx,
// 			m.Topic(),
// 			m.Qos,
// 			true,
// 			payload); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
