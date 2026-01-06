package ds18b20

// https://github.com/mertenats/open-home-automation/tree/master/ha_mqtt_sensor_dht22

import (
	"fmt"
	"strings"
	"time"
)

type ds18b20 struct {
	class             string
	icon              string
	id                string
	manufacturer      string
	model             string
	picture           string
	time              time.Time
	unitOfMeasurement string
}

func (s *ds18b20) Class() string             { return s.class }
func (s *ds18b20) Icon() string              { return s.icon }
func (s *ds18b20) ID() string                { return s.id }
func (s *ds18b20) Manufacturer() string      { return s.manufacturer }
func (s *ds18b20) Model() string             { return s.model }
func (s *ds18b20) Picture() string           { return s.picture }
func (s *ds18b20) Time() time.Time           { return s.time }
func (s *ds18b20) UnitOfMeasurement() string { return s.unitOfMeasurement }

func (s *ds18b20) Name() string {
	return strings.TrimSpace(fmt.Sprintf("%s %s %s", s.manufacturer, s.model, s.class))
}
