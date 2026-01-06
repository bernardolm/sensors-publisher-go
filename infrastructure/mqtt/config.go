package mqtt

import (
	"fmt"
	// "github.com/bernardolm/sensors-publisher-go/infrastructure/config"
	// "github.com/bernardolm/sensors-publisher-go/infrastructure/logging"
	// eclipsemqtt "github.com/eclipse/paho.mqtt.golang"
)

func (c *Client) config() {
	// if config.Get[bool]("DEBUG") {
	// 	eclipsemqtt.DEBUG = logging.Log
	// }

	if c.port == 0 {
		c.port = 1883
	}

	c.url = fmt.Sprintf("tcp://%s:%d", c.host, c.port)
}
