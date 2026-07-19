package influxdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/logger"
)

func (c *Client) Send(ctx context.Context, payload []byte) error {
	logger.Log.Debug("influxdb: publishing")
	if c == nil || c.writer == nil {
		return errors.New("influxdb: client not configured")
	}

	c.writer.WriteRecord(string(payload))

	ch := c.writer.Errors()

	for err := range ch {
		if err != nil {
			return errors.Join(err, fmt.Errorf("influxdb: can't sent"))
		}
	}

	logger.Log.
		WithField("payload", string(payload)).
		Info("influxdb: sent")

	return nil
}
