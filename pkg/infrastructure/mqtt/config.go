package mqtt

import (
	"fmt"
	"net"
	// "github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/config"
	// "github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/logging"
	// eclipsemqtt "github.com/eclipse/paho.mqtt.golang"
)

func (c *Client) config() {
	// if config.Get[bool]("DEBUG") {
	// 	eclipsemqtt.DEBUG = logging.Log
	// }

	if c.port == 0 {
		c.port = 1883
	}

	c.url = "tcp://" + net.JoinHostPort(c.host, fmt.Sprint(rune(c.port)))
}
