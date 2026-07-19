package bootstrap

import (
	"context"

	"github.com/bernardolm/sensors-publisher-go/pkg/contract"
	"github.com/bernardolm/sensors-publisher-go/pkg/domain/model/measurement"
)

type MeasurementRepository struct {
	client contract.SQLiteClient
}

func (o *MeasurementRepository) Insert(ctx context.Context, m measurement.Measurement) error {
	o.client.DB(ctx).Exec("foobar")
	return nil
}

func ProvideMeasurementRepository(
	client contract.SQLiteClient,
) contract.MeasurementRepository {
	obj := MeasurementRepository{
		client: client,
	}
	return &obj
}

// --------------

type MeasurementWorker struct {
	logger                contract.Logger
	sensor                contract.Sensor
	measurementRepository contract.MeasurementRepository
}

func (o *MeasurementWorker) Do(ctx context.Context) error {
	return nil
}

var _ contract.Worker = &MeasurementWorker{}

func ProvideMeasurementWorker(
	logger contract.Logger,
	sensor contract.Sensor,
	measurementRepository contract.MeasurementRepository,
) contract.Worker {
	obj := MeasurementWorker{
		logger:                logger,
		sensor:                sensor,
		measurementRepository: measurementRepository,
	}
	return &obj
}

// --------------
