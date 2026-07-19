package bootstrap

import "github.com/bernardolm/sensors-publisher-go/pkg/contract"

func ProvideLogger() contract.Logger {
	logger := dummy{}
	return &logger
}
