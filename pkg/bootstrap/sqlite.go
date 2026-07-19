package bootstrap

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/bernardolm/sensors-publisher-go/pkg/contract"
	"github.com/bernardolm/sensors-publisher-go/pkg/domain/model/measurement"
	"github.com/bernardolm/sensors-publisher-go/pkg/domain/model/publication"
	"github.com/bernardolm/sensors-publisher-go/pkg/domain/model/sensor"
	"github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/config"
)

type sqliteClient struct {
	db *gorm.DB
}

func (o sqliteClient) DB(_ context.Context) *gorm.DB {
	return o.db
}

func ProvideClientSQLite() contract.SQLiteClient {
	path := config.Get[string]("SQLITE_PATH")
	var dsn string

	if path == ":memory:" {
		dsn = path
	} else {
		dir := filepath.Dir(path)
		if dir == "." || dir == "" {
			panic(0)
		}

		if err := os.MkdirAll(dir, 0o755); err != nil {
			panic(err)
		}

		separator := "?"
		if strings.Contains(path, "?") {
			separator = "&"
		}

		dsn = path + separator + "_texttotime=1&_inttotime=1&_time_format=sqlite"
	}

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		NowFunc:                func() time.Time { return time.Now().Local() },
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&sensor.Sensor{})
	db.AutoMigrate(&measurement.Measurement{})
	db.AutoMigrate(&publication.Publication{})

	return sqliteClient{db: db}
}
