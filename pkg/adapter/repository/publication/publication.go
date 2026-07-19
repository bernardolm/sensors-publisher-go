package publication

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	model "github.com/bernardolm/sensors-publisher-go/pkg/domain/model/publication"
)

// Repository persists successful publications.
type Repository struct {
	db     *gorm.DB
	driver string
}

// NewRepository creates a publication repository over a prepared database.
func NewRepository(db *gorm.DB, driver string) *Repository {
	return &Repository{db: db, driver: driver}
}

// Save persists one successful publication.
func (r *Repository) Save(ctx context.Context, item *model.Publication) error {
	if !item.SentAt.IsZero() {
		item.SentAt = item.SentAt.Local()
	}
	if err := r.db.WithContext(ctx).Create(item).Error; err != nil {
		return fmt.Errorf("publication repository: save %s publication: %w", r.driver, err)
	}

	return nil
}
