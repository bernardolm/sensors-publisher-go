package postgres

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	sensorrepository "github.com/bernardolm/sensors-publisher-go/pkg/adapter/repository/sensor"
	postgresinfrastructure "github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/database/postgres"
)

// New creates a PostgreSQL sensor repository over a prepared connection.
func New(_ context.Context, client *postgresinfrastructure.Client) (*sensorrepository.SensorRepository, error) {
	if client == nil || client.DB() == nil {
		return nil, fmt.Errorf("sensor repository: PostgreSQL client is required")
	}

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: client.DB()}), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
	})
	if err != nil {
		return nil, fmt.Errorf("sensor repository: open PostgreSQL GORM connection: %w", err)
	}

	return sensorrepository.NewRepository(db, "PostgreSQL"), nil
}
