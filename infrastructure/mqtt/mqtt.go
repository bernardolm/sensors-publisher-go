package mqtt

import (
	"context"

	"github.com/bernardolm/sensors-publisher-go/infrastructure/config"
	"github.com/bernardolm/sensors-publisher-go/infrastructure/logging"
	eclipsemqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/k0kubun/pp/v3"
)

type Client struct {
	client   eclipsemqtt.Client
	host     string
	id       string
	password string
	port     int
	store    eclipsemqtt.Store
	url      string
	username string
}

func NewWithStore(ctx context.Context, store eclipsemqtt.Store) (*Client, error) {
	c := Client{
		host:     config.Get[string]("MQTT_HOST"),
		id:       "sensors-publisher-go",
		password: config.Get[string]("MQTT_PASSWORD"),
		port:     config.Get[int]("MQTT_PORT"),
		store:    store,
		username: config.Get[string]("MQTT_USERNAME"),
	}

	if c.host == "" {
		logging.Log.Warnf("mqtt: no host configured")
		return nil, nil
	}

	c.config()

	pp.Println(c)

	if err := c.connect(ctx); err != nil {
		return nil, err
	}

	return &c, nil
}

func New(ctx context.Context) (*Client, error) {
	return NewWithStore(ctx, nil)
}
