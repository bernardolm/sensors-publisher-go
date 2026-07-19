package bootstrap

import "github.com/bernardolm/sensors-publisher-go/pkg/contract"

func ProvideClientInfluxdb() contract.InfluxdbClient {
	logger := dummy{}
	return &logger
}
