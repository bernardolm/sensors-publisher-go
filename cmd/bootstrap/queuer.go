package bootstrap

import (
	"context"

	"github.com/bernardolm/sensors-publisher-go/queuer"
	qmqtt "github.com/bernardolm/sensors-publisher-go/queuer/mqtt"
	qsqlite "github.com/bernardolm/sensors-publisher-go/queuer/sqlite"
)

var (
	mqttQueuer   queuer.Queuer
	sqliteQueuer queuer.Queuer
)

func InitQueuers(ctx context.Context) error {
	var err error

	mqttQueuer, err = qmqtt.New(ctx)
	if err != nil {
		return err
	}

	sqliteQueuer, err = qsqlite.New(ctx)
	if err != nil {
		return err
	}

	return nil
}
