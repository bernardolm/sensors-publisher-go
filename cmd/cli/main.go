package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/dig"

	"github.com/bernardolm/sensors-publisher-go/pkg/bootstrap"
	"github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/config"
	"github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/logger"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	config.Load()
	logger.Init()
	container := dig.New()

	container.Provide(bootstrap.ProvideLogger)

	container.Provide(bootstrap.ProvideClientSQLite)
	// container.Provide(bootstrap.ProvideClientPostgres)

	// container.Provide(bootstrap.ProvideDs18b20Sensor)

	// container.Provide(bootstrap.ProvideTemperatureSensor)

	container.Provide(bootstrap.ProvideMeasurementRepository)
	container.Provide(bootstrap.ProvideSensorRepository)
	// container.Provide(bootstrap.ProvidePublicationRepository)

	container.Provide(bootstrap.ProvideTemperatureCollectorWorker)
	// container.Provide(bootstrap.ProvideMeasurementWorker)
	// container.Provide(bootstrap.ProvideReplicationWorker)
	// container.Provide(bootstrap.ProvidePublicationWorker)

	workerProcess := func(srv *bootstrap.TemperatureCollectorWorker) {
		srv.Collect(ctx)
	}

	if err := container.Invoke(workerProcess); err != nil {
		panic(err)
	}
}
