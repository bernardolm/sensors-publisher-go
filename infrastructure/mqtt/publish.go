package mqtt

import (
	"errors"
	"fmt"
	"time"

	"github.com/bernardolm/sensors-publisher-go/infrastructure/logging"
)

func (c *Client) Publish(topic string, qos byte, retained bool, payload any) error {
	e := logging.Log.
		WithField("payload", fmt.Sprintf("%s", payload)).
		WithField("qos", qos).
		WithField("retained", retained).
		WithField("topic", topic)

	e.Debug("mqtt: starting publish")

	if c.client == nil {
		return errors.New("mqtt: client not initialized")
	}

	if !c.client.IsConnected() {
		return errors.New("mqtt: not connected")
	}

	token := c.client.Publish(topic, qos, retained, payload)
	if token.Error() != nil {
		return token.Error()
	}

	_ = token.WaitTimeout(5 * time.Second) // Can also use '<-token.Done()' in releases > 1.2.0
	if token.Error() != nil {
		e.
			WithError(token.Error()).
			Error("mqtt: failed to publish")
		return token.Error()
	}

	e.Info("mqtt: published")

	return nil
}
