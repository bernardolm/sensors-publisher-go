package influxdb

import (
	"context"
	"fmt"

	"github.com/bernardolm/sensors-publisher-go/infrastructure/config"
	"github.com/bernardolm/sensors-publisher-go/infrastructure/logging"
	influx "github.com/influxdata/influxdb-client-go"
	influxapi "github.com/influxdata/influxdb-client-go/api"
)

type Client struct {
	client   influx.Client
	database string
	password string
	token    string
	url      string
	username string
	writer   influxapi.WriteAPI
}

func New(ctx context.Context) (*Client, error) {
	c := Client{
		database: config.Get[string]("INFLUX_DATABASE"),
		password: config.Get[string]("INFLUX_PASSWORD"),
		url:      config.Get[string]("INFLUX_URL"),
		username: config.Get[string]("INFLUX_USERNAME"),
	}

	if c.url == "" {
		logging.Log.Warnf("influxdb: no host configured")
		return nil, nil
	}

	c.token = fmt.Sprintf("%s:%s", c.username, c.password)

	if c.client == nil || c.writer == nil {
		if err := c.connect(ctx); err != nil {
			return nil, err
		}
	}

	return &c, nil
}
