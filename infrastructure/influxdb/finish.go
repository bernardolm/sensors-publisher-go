package influxdb

import (
	"context"

	"github.com/bernardolm/sensors-publisher-go/infrastructure/logging"
)

func (c *Client) Finish(ctx context.Context) {
	logging.Log.Debug("influxdb: disconnecting")

	c.writer = nil
	c.client.Close()
	c.client = nil

	logging.Log.Info("influxdb: disconnected")
}
