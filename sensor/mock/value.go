package mock

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/bernardolm/sensors-publisher-go/infrastructure/logging"
)

func (s *mock) Value() (any, error) {
	r := rand.Float64()
	r = r * 1.234 * float64(time.Now().Second())

	// fmt.Printf("\n\033[0;32m************************** r **************************\033[0m\n")
	// pp.Println(r)
	// fmt.Printf("\n\n")
	// // remember import github.com/k0kubun/pp/v3

	value := aws.Float64(r)

	// fmt.Printf("\n\033[0;32m************************** value **************************\033[0m\n")
	// pp.Println(value)
	// fmt.Printf("\n\n")
	// // remember import github.com/k0kubun/pp/v3

	if value == nil {
		return nil, fmt.Errorf("mock sensor: get value failed")
	}

	s.time = time.Now()

	logging.Log.
		WithField("name", "mock").
		WithField("value", *value).
		Debug("mock sensor: retrieved value")

	return *value, nil
}
