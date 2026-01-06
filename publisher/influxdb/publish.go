package influxdb

import (
	"context"
	"fmt"

	"github.com/bernardolm/sensors-publisher-go/infrastructure/logging"
)

func (a *influxdb) Publish(ctx context.Context, content any) error {
	if content == nil {
		return nil
	}

	payload, ok := content.([]byte)
	if !ok {
		return fmt.Errorf("publisher: influxdb payload type %T", content)
	}

	logging.Log.
		WithField("message", string(payload)).
		WithField("publisher", "influxdb").
		Debug("publisher: trying to publish")

	if err := a.client.Send(ctx, payload); err != nil {
		return err
	}

	return nil
}
