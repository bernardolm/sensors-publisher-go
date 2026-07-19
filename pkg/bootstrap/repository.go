package bootstrap

import "github.com/bernardolm/sensors-publisher-go/pkg/contract"

func ProvideMyRepository() contract.Repository {
	repository := dummy{}
	return &repository
}
