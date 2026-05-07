package bootstrap

import (
	"context"

	"github.com/bernardolm/sensors-publisher-go/pkg/domain/formatter"
	fhass "github.com/bernardolm/sensors-publisher-go/pkg/domain/formatter/homeassistant"
	finfluxdb "github.com/bernardolm/sensors-publisher-go/pkg/domain/formatter/influxdb"
	fsqlite "github.com/bernardolm/sensors-publisher-go/pkg/domain/formatter/sqlite"
)

var (
	hassFormatter     formatter.Formatter
	influxdbFormatter formatter.Formatter
	sqliteFormatter   formatter.Formatter
)

func InitFormatters(ctx context.Context) error {
	var err error

	hassFormatter, err = fhass.New(ctx)
	if err != nil {
		return err
	}

	influxdbFormatter, err = finfluxdb.New(ctx)
	if err != nil {
		return err
	}

	sqliteFormatter, err = fsqlite.New(ctx)
	if err != nil {
		return err
	}

	return nil
}
