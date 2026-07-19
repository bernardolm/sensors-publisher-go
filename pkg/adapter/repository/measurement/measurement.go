package measurement

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	modelmeasurement "github.com/bernardolm/sensors-publisher-go/pkg/domain/model/measurement"
	"github.com/bernardolm/sensors-publisher-go/pkg/domain/model/publication"
)

// MeasurementRepository implements the shared GORM persistence behavior for measurements.
type MeasurementRepository struct {
	db     *gorm.DB
	driver string
}

// NewRepository creates a measurement repository over a prepared database.
func NewRepository(db *gorm.DB, driver string) *MeasurementRepository {
	return &MeasurementRepository{db: db, driver: driver}
}

// Save persists a measurement.
func (r *MeasurementRepository) Save(ctx context.Context, val *modelmeasurement.Measurement) error {
	val.CollectedAt = val.CollectedAt.Local()

	if err := r.db.WithContext(ctx).Create(val).Error; err != nil {
		return fmt.Errorf("measurement repository: save %s measurement: %w", r.driver, err)
	}

	return nil
}

// ListPending returns readings without a successful delivery to one destination.
func (r *MeasurementRepository) ListPending(
	ctx context.Context,
	destination publication.Destination,
	limit int,
) ([]*modelmeasurement.Measurement, error) {
	if limit <= 0 {
		limit = 100
	}

	successfulDelivery := r.db.
		Model(&publication.Publication{}).
		Select("1").
		Where("publication.id_measurement = measurement.id").
		Where("publication.destination = ?", destination)

	var obj []*modelmeasurement.Measurement
	if err := r.db.WithContext(ctx).
		Where("NOT EXISTS (?)", successfulDelivery).
		Preload("Sensor").
		Order("collected_at ASC").
		Order("id ASC").
		Limit(limit).
		Find(&obj).Error; err != nil {
		return nil, fmt.Errorf("measurement repository: list pending %s measurements: %w", r.driver, err)
	}

	return obj, nil
}
