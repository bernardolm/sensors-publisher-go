package sensor

import (
	"time"

	"github.com/bernardolm/sensors-publisher-go/pkg/domain/model/measurement"
)

// Sensor represents persisted metadata and the measurement class of one physical sensor.
type Sensor struct {
	ID                int64  `gorm:"primaryKey;autoIncrement"`
	UniqueID          string `gorm:"column:unique_id;uniqueIndex;not null"`
	Class             measurement.Class
	Icon              string
	Manufacturer      string
	Model             string
	Name              string `gorm:"not null"`
	Picture           string
	RegisteredAt      time.Time  `gorm:"column:registered_at;type:timestamp;not null;autoCreateTime"`
	UpdatedAt         *time.Time `gorm:"column:updated_at;type:timestamp;autoUpdateTime"`
	UnitOfMeasurement measurement.UnitOfMeasurement
}

// TableName returns the sensor table name.
func (Sensor) TableName() string {
	return "sensor"
}
