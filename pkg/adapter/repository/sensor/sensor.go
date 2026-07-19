package sensor

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	model "github.com/bernardolm/sensors-publisher-go/pkg/domain/model/sensor"
)

// SensorRepository persists sensor metadata.
type SensorRepository struct {
	db     *gorm.DB
	driver string
}

// NewRepository creates a sensor repository over a prepared database.
func NewRepository(db *gorm.DB, driver string) *SensorRepository {
	return &SensorRepository{db: db, driver: driver}
}

// Register creates or updates sensor metadata by its unique ID.
func (r *SensorRepository) Register(ctx context.Context, item *model.Sensor) (*model.Sensor, error) {
	if !item.RegisteredAt.IsZero() {
		item.RegisteredAt = item.RegisteredAt.Local()
	}
	if item.UpdatedAt != nil && !item.UpdatedAt.IsZero() {
		updatedAt := item.UpdatedAt.Local()
		item.UpdatedAt = &updatedAt
	}

	query := r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "unique_id"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"class",
				"icon",
				"manufacturer",
				"model",
				"name",
				"picture",
				"updated_at",
				"unit_of_measurement",
			}),
		})
	if r.driver == "PostgreSQL" {
		query = query.Omit("ID")
	}
	if err := query.Create(item).Error; err != nil {
		return nil, fmt.Errorf("sensor repository: save %s sensor: %w", r.driver, err)
	}

	return r.Get(ctx, item.UniqueID)
}

// Get returns persisted metadata for one sensor unique ID.
func (r *SensorRepository) Get(ctx context.Context, uniqueID string) (*model.Sensor, error) {
	item := &model.Sensor{}
	if err := r.db.WithContext(ctx).First(item, "unique_id = ?", uniqueID).Error; err != nil {
		return nil, fmt.Errorf("sensor repository: get %s sensor %s: %w", r.driver, uniqueID, err)
	}

	return item, nil
}
