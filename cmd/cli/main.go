package main

import (
	"context"
	"time"

	fhass "github.com/bernardolm/sensors-publisher-go/pkg/domain/formatter/homeassistant"
	pmqtt "github.com/bernardolm/sensors-publisher-go/pkg/domain/publisher/mqtt"
	"github.com/bernardolm/sensors-publisher-go/pkg/domain/sensor/ds18b20"
	"github.com/bernardolm/sensors-publisher-go/pkg/domain/sensor/mock"
	"github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/config"
	"github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/logging"
	imqtt "github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/mqtt"
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

	// queueWorkerDelta := config.Get[time.Duration]("QUEUE_WORKER_DELTA")
	// if queueWorkerDelta == 0 {
	// 	queueWorkerDelta = time.Duration(5*60) * time.Second
	// }

	publisherWorkerDelta := config.Get[time.Duration]("PUBLISHER_WORKER_DELTA")
	if publisherWorkerDelta == 0 {
		publisherWorkerDelta = time.Duration(5*60) * time.Second
	}

	for {
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
}
