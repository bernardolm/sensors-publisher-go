package ds18b20

import (
	"context"
	"strings"

	yryzds18b20 "github.com/yryz/ds18b20"

	"github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/logger"
)

func New(_ context.Context) ([]*ds18b20, error) {
	sensorIDs, err := yryzds18b20.Sensors()
	if err != nil {
		return nil, err
	}

	if len(sensorIDs) == 0 {
		logger.Log.Debug("ds18b20 sensor: ds18b20 not found")

		return nil, nil
	}

	logger.Log.
		WithField("sensors",
			strings.Join(sensorIDs, ",")).
		Debugf("ds18b20 sensor: %d found", len(sensorIDs))

	sensors := []*ds18b20{}

	for _, id := range sensorIDs {
		sensors = append(sensors, &ds18b20{
			// picture:           "https://cdn.awsli.com.br/2500x2500/468/468162/produto/19414360586929efad.jpg",
			class:             "temperature",
			icon:              "mdi: home-thermometer-outline",
			id:                id,
			manufacturer:      "unknown",
			model:             "ds18b20",
			unitOfMeasurement: "°C",
		})
	}

	return sensors, nil
}
