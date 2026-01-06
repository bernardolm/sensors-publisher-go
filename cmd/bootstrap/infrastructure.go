package bootstrap

import (
	"context"

	"github.com/bernardolm/sensors-publisher-go/infrastructure/config"
	iinfluxdb "github.com/bernardolm/sensors-publisher-go/infrastructure/influxdb"
	imqtt "github.com/bernardolm/sensors-publisher-go/infrastructure/mqtt"
	isqlite "github.com/bernardolm/sensors-publisher-go/infrastructure/sqlite"
)

var (
	influxClient *iinfluxdb.Client
	mqttClient   *imqtt.Client
	sqliteClient *isqlite.Client
)

func InitInfrastructures(ctx context.Context) error {
	var err error

	sqlitePath := config.Get[string]("SQLITE_PATH")
	sqliteClient, err = isqlite.New(sqlitePath)
	if err != nil {
		return err
	}

	mqttClient, err = imqtt.New(ctx)
	if err != nil {
		return err
	}

	influxClient, err = iinfluxdb.New(ctx)
	if err != nil {
		return err
	}

	return nil
}
