package influxdb

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

func Start(ctx context.Context) {
	log.Info("influxdb: starting client")
	_ = getAPI(ctx)
}

func Send(ctx context.Context, _ string, payload []byte) error {
	log.Debug("influxdb: publishing")

	startedAt := time.Now()
	getAPI(ctx).WriteRecord(string(payload))
	if err := lastWriteError(startedAt); err != nil {
		return err
	}

	log.
		WithField("payload", fmt.Sprintf("%s", payload)).
		Info("influxdb: sent")

	return nil
}

func Finish(ctx context.Context) {
	log.Debug("influxdb: disconnecting")

	api = nil
	getClient(ctx).Close()
	client = nil

	log.Info("influxdb: disconnected")
}
