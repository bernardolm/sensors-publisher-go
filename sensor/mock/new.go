package mock

import (
	"context"
	"fmt"
	"time"

	"github.com/bernardolm/sensors-publisher-go/formatter/homeassistant/api"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func New(_ context.Context) *mock {
	return &mock{
		// picture:           "https://www.pmi.org/-/media/pmi/microsites/disciplined-agile/the-design-patterns/crashtestdummy.jpg",
		caser:             cases.Title(language.BrazilianPortuguese),
		class:             string(api.AtmosphericPressureDeviceClass),
		icon:              "mdi: test-tube",
		id:                fmt.Sprintf("mock_%d", time.Now().Minute()),
		manufacturer:      "sensors-publisher-go",
		model:             "dummy",
		unitOfMeasurement: "hPa",
	}
}
