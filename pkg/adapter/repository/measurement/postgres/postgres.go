package postgres

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/bernardolm/sensors-publisher-go/pkg/domain/model/measurement"
	postgresinfrastructure "github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/database/postgres"
)

// Repository persists measurements in PostgreSQL.
type Repository struct {
	db *gorm.DB
}

// New creates a PostgreSQL measurement repository.
func New(_ context.Context, client *postgresinfrastructure.Client) (*Repository, error) {
	if client == nil || client.DB() == nil {
		return nil, fmt.Errorf("measurement repository: PostgreSQL client is required")
	}

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: client.DB()}), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
	})
	if err != nil {
		return nil, fmt.Errorf("measurement repository: open PostgreSQL GORM connection: %w", err)
	}

	return &Repository{db: db}, nil
}

// Save persists one measurement idempotently by its local ID.
func (r *Repository) Save(ctx context.Context, item *measurement.Measurement) error {
	item.CollectedAt = item.CollectedAt.Local()

	if err := r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoNothing: true,
		}).
		Omit("Sensor").
		Create(item).Error; err != nil {
		return fmt.Errorf("measurement repository: save PostgreSQL measurement: %w", err)
	}

	return nil
}
