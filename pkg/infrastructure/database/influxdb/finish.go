package influxdb

import (
	"context"

	"github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/logger"
)

func (c *Client) Finish(ctx context.Context) {
	logger.Log.Debug("influxdb: disconnecting")

	c.writer = nil
	c.client.Close()
	c.client = nil

	logger.Log.Info("influxdb: disconnected")
}
