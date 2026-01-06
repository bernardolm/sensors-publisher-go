package influxdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/bernardolm/sensors-publisher-go/infrastructure/logging"
)

func (c *Client) Send(ctx context.Context, payload []byte) error {
	logging.Log.Debug("influxdb: publishing")

	c.writer.WriteRecord(string(payload))

	ch := c.writer.Errors()

	for err := range ch {
		if err != nil {
			return errors.Join(err, fmt.Errorf("influxdb: can't sent"))
		}
	}

	logging.Log.
		WithField("payload", fmt.Sprintf("%s", payload)).
		Info("influxdb: sent")

	return nil
}
