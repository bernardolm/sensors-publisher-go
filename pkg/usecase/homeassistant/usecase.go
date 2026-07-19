package homeassistant

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"time"

// 	measurementrepository "github.com/bernardolm/sensors-publisher-go/pkg/adapter/repository/measurement"
// 	publicationrepository "github.com/bernardolm/sensors-publisher-go/pkg/adapter/repository/publication"
// 	sensorrepository "github.com/bernardolm/sensors-publisher-go/pkg/adapter/repository/sensor"
// 	"github.com/bernardolm/sensors-publisher-go/pkg/domain/model/publication"
// 	"github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/logger"
// 	mqttinfrastructure "github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/messaging/mqtt"
// 	"github.com/bernardolm/sensors-publisher-go/pkg/usecase/worker"
// )

// // UseCase publishes locally stored readings to MQTT for Home Assistant.
// type UseCase struct {
// 	batchSize    int
// 	client       *mqttinfrastructure.Client
// 	measurements *measurementrepository.MeasurementRepository
// 	publications *publicationrepository.Repository
// 	sensors      *sensorrepository.SensorRepository
// 	worker       *worker.Worker
// }

// // New creates the Home Assistant MQTT publication use case.
// func New(
// 	_ context.Context,
// 	measurements *measurementrepository.MeasurementRepository,
// 	publications *publicationrepository.Repository,
// 	sensors *sensorrepository.SensorRepository,
// 	client *mqttinfrastructure.Client,
// 	interval time.Duration,
// 	batchSize int,
// ) (*UseCase, error) {
// 	if measurements == nil {
// 		return nil, fmt.Errorf("home assistant: measurement repository is required")
// 	}
// 	if publications == nil {
// 		return nil, fmt.Errorf("home assistant: publication repository is required")
// 	}
// 	if sensors == nil {
// 		return nil, fmt.Errorf("home assistant: sensor repository is required")
// 	}
// 	if client == nil {
// 		return nil, fmt.Errorf("home assistant: MQTT client is required")
// 	}

// 	if batchSize <= 0 {
// 		batchSize = 100
// 	}
// 	useCase := &UseCase{
// 		batchSize:    batchSize,
// 		client:       client,
// 		measurements: measurements,
// 		publications: publications,
// 		sensors:      sensors,
// 	}
// 	worker, err := worker.New(interval, useCase.run)
// 	if err != nil {
// 		return nil, err
// 	}

// 	useCase.worker = worker

// 	return useCase, nil
// }

// // Interval returns the MQTT publication schedule.
// func (u *UseCase) Interval() time.Duration {
// 	return u.worker.Interval()
// }

// // Start starts the MQTT worker and returns immediately.
// func (u *UseCase) Start(ctx context.Context) {
// 	u.worker.Start(ctx)
// }

// // run executes one MQTT publication iteration.
// func (u *UseCase) run(ctx context.Context) {
// 	if err := u.Execute(ctx); err != nil {
// 		logger.Log.WithError(err).Error("home assistant: publication failed")
// 	}
// }

// // Execute publishes one batch of readings pending for MQTT.
// func (u *UseCase) Execute(ctx context.Context) error {
// 	readings, err := u.measurements.ListPending(ctx, publication.DestinationMQTT, u.batchSize)
// 	if err != nil {
// 		return fmt.Errorf("home assistant: list pending measurements: %w", err)
// 	}

// 	var publicationErrors error
// 	for _, reading := range readings {
// 		if reading.Sensor == nil {
// 			publicationErrors = errors.Join(
// 				publicationErrors,
// 				fmt.Errorf("home assistant: sensor for measurement %d is missing", reading.ID),
// 			)

// 			continue
// 		}
// 		sensor, err := u.sensors.Get(ctx, reading.Sensor.UniqueID)
// 		if err != nil {
// 			publicationErrors = errors.Join(
// 				publicationErrors,
// 				fmt.Errorf("home assistant: get sensor %s: %w", reading.Sensor.UniqueID, err),
// 			)

// 			continue
// 		}
// 		if err := u.publish(ctx, reading, sensor); err != nil {
// 			publicationErrors = errors.Join(publicationErrors, err)

// 			continue
// 		}
// 		if err := u.publications.Save(ctx, &publication.Publication{
// 			MeasurementID: reading.ID,
// 			Destination:   publication.DestinationMQTT,
// 		}); err != nil {
// 			publicationErrors = errors.Join(publicationErrors, err)
// 		}
// 	}

// 	return publicationErrors
// }
