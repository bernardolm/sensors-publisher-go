package bootstrap

import (
	"context"

	"github.com/bernardolm/sensors-publisher-go/formatter"
	fhass "github.com/bernardolm/sensors-publisher-go/formatter/homeassistant"
	finfluxdb "github.com/bernardolm/sensors-publisher-go/formatter/influxdb"
	fsqlite "github.com/bernardolm/sensors-publisher-go/formatter/sqlite"
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
