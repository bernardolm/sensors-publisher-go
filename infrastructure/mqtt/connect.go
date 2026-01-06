package mqtt

import (
	"context"
	"time"

	"github.com/bernardolm/sensors-publisher-go/infrastructure/logging"
	eclipsemqtt "github.com/eclipse/paho.mqtt.golang"
)

func (c *Client) connect(_ context.Context) error {
	logging.Log.Debug("mqtt: trying to connect")

	opts := eclipsemqtt.NewClientOptions().
		AddBroker(c.url).
		SetAutoReconnect(true).
		SetClientID(c.id).
		SetConnectionLostHandler(connectionLostHandler).
		SetConnectRetry(true).
		SetConnectTimeout(3 * time.Second).
		SetOnConnectHandler(onConnectHandler).
		SetOrderMatters(true).
		SetPassword(c.password).
		SetReconnectingHandler(reconnecthandler).
		SetUsername(c.username)

	if c.store != nil {
		opts.SetStore(c.store)
	}

	c.client = eclipsemqtt.NewClient(opts)

	token := c.client.Connect()

	if token.Error() != nil {
		return token.Error()
	}

	_ = token.Wait() // Can also use '<-token.Done()' in releases > 1.2.0
	if token.Error() != nil {
		logging.Log.
			WithError(token.Error()).
			Error("mqtt: failed to connect")
		return token.Error()
	}

	logging.Log.Info("mqtt: connected")

	return nil
}
