package bootstrap

import "github.com/bernardolm/sensors-publisher-go/pkg/contract"

func ProvideClientPostgres() contract.PostgresClient {
	logger := dummy{}
	return &logger
}
