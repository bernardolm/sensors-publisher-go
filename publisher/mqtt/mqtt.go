package mqtt

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	mqttclient "github.com/bernardolm/sensors-publisher-go/infrastructure/mqtt"
	sqlitequeue "github.com/bernardolm/sensors-publisher-go/infrastructure/sqlite"
)

const mqttQueueTarget = "mqtt"

type mqtt struct {
	queue *sqlitequeue.Queue
}

func (a *mqtt) Publish(ctx context.Context, topic string, message interface{}) error {
	if a.queue != nil && mqttclient.IsConnected() {
		if sent, err := a.queue.Flush(ctx, mqttQueueTarget, func(ctx context.Context, topic string, payload []byte) error {
			return mqttclient.Send(ctx, topic, payload)
		}); err != nil {
			log.WithError(err).Warn("publisher: mqtt queue flush failed")
		} else if sent > 0 {
			log.WithField("count", sent).Info("publisher: mqtt queue flushed")
		}
	}

	if message == nil {
		return nil
	}

	payload, ok := message.([]byte)
	if !ok {
		return fmt.Errorf("publisher: mqtt payload type %T", message)
	}

	log.WithField("topic", topic).
		WithField("message", string(payload)).
		WithField("publisher", "mqtt").
		Debug("publisher: trying to publish")

	if err := mqttclient.Send(ctx, topic, payload); err != nil {
		log.WithError(err).
			WithField("topic", topic).
			Warn("publisher: mqtt send failed, queueing")
		if a.queue != nil {
			if err := a.queue.Enqueue(ctx, mqttQueueTarget, topic, payload); err != nil {
				return err
			}
			return nil
		}
		return err
	}

	return nil
}

func New(queue *sqlitequeue.Queue) *mqtt {
	return &mqtt{queue: queue}
}
