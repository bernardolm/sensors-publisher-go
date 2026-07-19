package measurement

import (
	"time"
)

// Measurement represents one measurement collected from a sensor.
type Measurement struct {
	ID       int64 `gorm:"primaryKey;autoIncrement"`
	IDSensor int64 `gorm:"column:id_sensor;not null;index"`
	// Sensor      *sensor.Sensor
	// `gorm:"foreignKey:IDSensor;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CollectedAt time.Time         `gorm:"column:collected_at;type:timestamp;not null;index"`
	Value       float64           `gorm:"column:value;not null"`
	Class       Class             `gorm:"column:class;not null"`
	Unit        UnitOfMeasurement `gorm:"column:unit_of_measurement;not null"`
}

// TableName returns the table name.
func (Measurement) TableName() string {
	return "measurement"
}
