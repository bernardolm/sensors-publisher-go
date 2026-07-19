package publication

import (
	"time"

	"github.com/bernardolm/sensors-publisher-go/pkg/domain/model/measurement"
)

// Publication records a successful delivery to one destination.
type Publication struct {
	ID            int64                    `gorm:"primaryKey;autoIncrement"`
	MeasurementID int64                    `gorm:"column:id_measurement;not null;uniqueIndex:idx_publication_measurement_destination"`
	Measurement   *measurement.Measurement `gorm:"foreignKey:measurementID;constraint:OnDelete:CASCADE"`
	SentAt        time.Time                `gorm:"column:sent_at;type:timestamp;not null;autoCreateTime"`
	Destination   Destination              `gorm:"column:destination;not null;uniqueIndex:idx_publication_measurement_destination"`
}

// TableName returns the publication table name.
func (Publication) TableName() string {
	return "publication"
}
