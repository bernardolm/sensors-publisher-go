package main

import (
	"context"
	"time"

	fhass "github.com/bernardolm/sensors-publisher-go/formatter/homeassistant"
	"github.com/bernardolm/sensors-publisher-go/infrastructure/config"
	"github.com/bernardolm/sensors-publisher-go/infrastructure/logging"
	imqtt "github.com/bernardolm/sensors-publisher-go/infrastructure/mqtt"
	pmqtt "github.com/bernardolm/sensors-publisher-go/publisher/mqtt"
	"github.com/bernardolm/sensors-publisher-go/sensor/ds18b20"
	"github.com/bernardolm/sensors-publisher-go/sensor/mock"
)

func main() {
	logging.Log.Info("cmd.console: starting")

	config.Load()

	logging.Init()

	ctx := context.Background()
	ctx, ctxCancelFunc := context.WithCancel(ctx)
	defer ctxCancelFunc()

	//
	//
	// bootstrap - begin
	//
	//

	mockSensor := mock.New(ctx)

	ds18b20Sensor, err := ds18b20.New(ctx)
	if err != nil {
		if !config.Get[bool]("DEBUG") {
			panic(err)
		}
	}

	hassFormatter, err := fhass.New(ctx)
	if err != nil {
		panic(err)
	}

	mqttClient, err := imqtt.New(ctx)
	if err != nil {
		panic(err)
	}

	mqttPublisher, err := pmqtt.New(ctx, mqttClient)
	if err != nil {
		panic(err)
	}

	//
	//
	// bootstrap - end
	//
	//

	queueWorkerDelta := config.Get[time.Duration]("QUEUE_WORKER_DELTA")
	if queueWorkerDelta == 0 {
		queueWorkerDelta = 5 * 60 * time.Second
	}

	publisherWorkerDelta := config.Get[time.Duration]("PUBLISHER_WORKER_DELTA")
	if publisherWorkerDelta == 0 {
		publisherWorkerDelta = 5 * 60 * time.Second
	}

	for true {
		if config.Get[bool]("DEBUG") {
			mockContent, err := hassFormatter.Build(mockSensor)
			if err != nil {
				panic(err)
			}

			if err := mqttPublisher.Publish(ctx, mockContent); err != nil {
				panic(err)
			}
		}

		for _, s := range ds18b20Sensor {
			content, err := hassFormatter.Build(s)
			if err != nil {
				panic(err)
			}

			if err := mqttPublisher.Publish(ctx, content); err != nil {
				panic(err)
			}
		}

		time.Sleep(publisherWorkerDelta)
	}

	logging.Log.Info("cmd.console: graceful shutdown complete")
}
