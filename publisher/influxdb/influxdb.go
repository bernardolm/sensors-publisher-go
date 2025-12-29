package influxdb

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	influxdbclient "github.com/bernardolm/sensors-publisher-go/infrastructure/influxdb"
	sqlitequeue "github.com/bernardolm/sensors-publisher-go/infrastructure/sqlite"
)

const influxQueueTarget = "influxdb"

type influxdb struct {
	queue *sqlitequeue.Queue
}

func (a *influxdb) Publish(ctx context.Context, topic string, message interface{}) error {
	if a.queue != nil {
		if sent, err := a.queue.Flush(ctx, influxQueueTarget, func(ctx context.Context, topic string, payload []byte) error {
			return influxdbclient.Send(ctx, topic, payload)
		}); err != nil {
			log.WithError(err).Warn("publisher: influxdb queue flush failed")
		} else if sent > 0 {
			log.WithField("count", sent).Info("publisher: influxdb queue flushed")
		}
	}

	if message == nil {
		return nil
	}

	payload, ok := message.([]byte)
	if !ok {
		return fmt.Errorf("publisher: influxdb payload type %T", message)
	}

	log.WithField("message", string(payload)).
		WithField("publisher", "influxdb").
		Debug("publisher: trying to publish")

	if err := influxdbclient.Send(ctx, topic, payload); err != nil {
		log.WithError(err).Warn("publisher: influxdb send failed, queueing")
		if a.queue != nil {
			if err := a.queue.Enqueue(ctx, influxQueueTarget, topic, payload); err != nil {
				return err
			}
			return nil
		}
		return err
	}

	return nil
}

func New(queue *sqlitequeue.Queue) *influxdb {
	return &influxdb{queue: queue}
}
