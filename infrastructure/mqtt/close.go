package mqtt

import (
	"context"

	"github.com/bernardolm/sensors-publisher-go/infrastructure/logging"
)

func (c *Client) Close(_ context.Context) {
	logging.Log.Debug("mqtt: disconnecting")
	c.client.Disconnect(500)
	logging.Log.Info("mqtt: disconnected")
}
