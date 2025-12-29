package influxdb

import (
	"context"
	"sync"
	"time"

	influxapi "github.com/influxdata/influxdb-client-go/api"
	log "github.com/sirupsen/logrus"
)

var (
	api          influxapi.WriteAPI
	lastErr      error
	lastErrAt    time.Time
	lastErrMutex sync.Mutex
)

func getAPI(ctx context.Context) influxapi.WriteAPI {
	if api == nil {
		api = getClient(ctx).WriteAPI("dummy-org", database)

		errorsCh := api.Errors()
		go func() {
			for err := range errorsCh {
				recordError(err)
				log.Errorf("influxdb: api write error - %s", err.Error())
			}
		}()
	}

	return api
}

func recordError(err error) {
	lastErrMutex.Lock()
	lastErr = err
	lastErrAt = time.Now()
	lastErrMutex.Unlock()
}

func lastWriteError(since time.Time) error {
	lastErrMutex.Lock()
	defer lastErrMutex.Unlock()

	if lastErr != nil && lastErrAt.After(since) {
		return lastErr
	}

	return nil
}
