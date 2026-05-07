package hass

import "context"

const (
	configTopicFormat string = "homeassistant/sensor/%s/config"
	stateTopicFormat  string = "sensors-publisher-go/sensor/%s"
)

type hass struct {
	deviceConfigSent bool
}

func New(ctx context.Context) (*hass, error) {
	return &hass{}, nil
}
