package bootstrap

import (
	"context"
	"errors"
	"time"

	eclipsemqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/bernardolm/sensors-publisher-go/pkg/contract"
	"github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/config"
	"github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/logger"
)

type MqttClient struct {
	client eclipsemqtt.Client
}

func (c MqttClient) Publish(ctx context.Context, topic string, qos byte, retained bool, payload any) error {
	if c.client == nil {
		return errors.New("mqtt: client not initialized")
	}

	if !c.client.IsConnected() {
		return errors.New("mqtt: not connected")
	}

	token := c.client.Publish(topic, qos, retained, payload)
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return errors.New("mqtt: publish timed out")
	case <-token.Done():
	}

	if err := token.Error(); err != nil {
		return err
	}

	return nil
}

func ProvideMqttClient() contract.MqttClient {
	host := config.Get[string]("MQTT_HOST")
	password := config.Get[string]("MQTT_PASSWORD")
	username := config.Get[string]("MQTT_USERNAME")

	url := "tcp://" + host + ":1883"

	opts := eclipsemqtt.NewClientOptions().
		AddBroker(url).
		SetAutoReconnect(true).
		SetClientID("sensors-publisher-go").
		SetConnectRetry(true).
		SetConnectTimeout(3 * time.Second).
		SetOrderMatters(true).
		SetPassword(password).
		SetUsername(username)

	client := eclipsemqtt.NewClient(opts)
	token := client.Connect()
	timer := time.NewTimer(3 * time.Second)
	defer timer.Stop()

	ctx := context.TODO()

	select {
	case <-ctx.Done():
		panic(ctx.Err())
	case <-timer.C:
		panic(errors.New("mqtt: initial connection timed out"))
	case <-token.Done():
	}

	if err := token.Error(); err != nil {
		logger.Log.
			WithError(err).
			Error("mqtt: failed to connect")

		panic(err)
	}

	obj := MqttClient{
		client: client,
	}

	return &obj
}
