package mqtt

import (
	"sync"

	"github.com/bernardolm/sensors-publisher-go/infrastructure/logging"
	eclipsemqtt "github.com/eclipse/paho.mqtt.golang"
)

var sorh = sync.Once{}

var reconnecthandler eclipsemqtt.ReconnectHandler = func(c eclipsemqtt.Client, options *eclipsemqtt.ClientOptions) {
	sorh.Do(func() {
		logging.Log.Warn("mqtt: client reconnected")
	})
}

var socl = sync.Once{}

var connectionLostHandler eclipsemqtt.ConnectionLostHandler = func(c eclipsemqtt.Client, err error) {
	socl.Do(func() {
		logging.Log.WithError(err).Error("mqtt: connection lost")
	})
}

var soch = sync.Once{}

var onConnectHandler eclipsemqtt.OnConnectHandler = func(c eclipsemqtt.Client) {
	soch.Do(func() {
		logging.Log.Info("mqtt: ready (connected now)")
	})
}
