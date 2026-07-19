package contract

import (
	"context"

	"gorm.io/gorm"

	"github.com/bernardolm/sensors-publisher-go/pkg/domain/model/measurement"
)

type HomeassistantRepository interface{ Do(context.Context) error }

type InfluxdbClient interface{ Do(context.Context) error }

type PublicationRepository interface{ Do(context.Context) error }

type Logger interface{ Do(context.Context) error }

type MeasurementRepository interface {
	Insert(context.Context, measurement.Measurement) error
}

type MqttClient interface {
	Publish(ctx context.Context, topic string, qos byte, retained bool, payload any) error
}

type PostgresClient interface{ Do(context.Context) error }

type ReplicationRepository interface{ Do(context.Context) error }

type Repository interface{ Do(context.Context) error }

type SensorRepository interface {
	Register(context.Context, SensorDevice) error
}

type SQLiteClient interface {
	DB(context.Context) *gorm.DB
}

type Worker interface{ Do(context.Context) error }
