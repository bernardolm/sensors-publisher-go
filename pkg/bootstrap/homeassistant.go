package bootstrap

import (
	"context"

	"github.com/bernardolm/sensors-publisher-go/pkg/contract"
)

type HomeassistantRepository struct {
	logger contract.Logger
	mqtt   contract.MqttClient
}

func (HomeassistantRepository) Do(context.Context) error { return nil }

func ProvideHomeassistantRepository(
	logger contract.Logger,
	mqtt contract.MqttClient,
) contract.HomeassistantRepository {
	obj := HomeassistantRepository{
		logger: logger,
		mqtt:   mqtt,
	}
	return &obj
}
