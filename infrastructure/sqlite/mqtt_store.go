package sqlite

import (
	eclipsemqtt "github.com/eclipse/paho.mqtt.golang"
)

func (c *Client) AsMqttStore() (eclipsemqtt.Store, error) {
	return nil, nil
}
