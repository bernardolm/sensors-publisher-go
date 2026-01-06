package influxdb

import (
	"context"
	"crypto/tls"
	"errors"

	"github.com/bernardolm/sensors-publisher-go/infrastructure/logging"
	influx "github.com/influxdata/influxdb-client-go"
)

func (c *Client) connect(_ context.Context) error {
	logging.Log.Debug("influxdb: trying to connect")

	opts := influx.DefaultOptions().
		SetTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})

	logging.Log.Debugf("influxdb: connecting to %q with %q", c.url, c.token)

	c.client = influx.NewClientWithOptions(c.url, c.token, opts)

	if c.client == nil {
		return errors.New("influxdb: couldn't create a client")
	}

	c.writer = c.client.WriteAPI("dummy-org", c.database)
	if c.writer == nil {
		return errors.New("influxdb: couldn't create a writer api")
	}

	return nil
}
