package bootstrap

import (
	"context"

	"github.com/bernardolm/sensors-publisher-go/pkg/domain/publisher"
	pinfluxdb "github.com/bernardolm/sensors-publisher-go/pkg/domain/publisher/influxdb"
	plcd "github.com/bernardolm/sensors-publisher-go/pkg/domain/publisher/lcd"
	pmqtt "github.com/bernardolm/sensors-publisher-go/pkg/domain/publisher/mqtt"
	pstdout "github.com/bernardolm/sensors-publisher-go/pkg/domain/publisher/stdout"
)

var (
	influxdbPublisher publisher.Publisher
	lcdPublisher      publisher.Publisher
	mqttPublisher     publisher.Publisher
	stdoutPublisher   publisher.Publisher
)

func InitPublishers(ctx context.Context) error {
	var err error

	mqttPublisher, err = pmqtt.New(ctx, mqttClient)
	if err != nil {
		return err
	}

	influxdbPublisher, err = pinfluxdb.New(ctx, influxClient)
	if err != nil {
		return err
	}

	stdoutPublisher, err = pstdout.New(ctx)
	if err != nil {
		return err
	}

	lcdPublisher, err = plcd.New(ctx)
	if err != nil {
		return err
	}

	return nil
}
