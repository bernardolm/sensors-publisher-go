package lcd

import (
	"context"

	"github.com/bernardolm/sensors-publisher-go/infrastructure/logging"
)

func (a *lcd) Publish(_ context.Context, content any) error {
	if content == nil {
		return nil
	}

	logging.Log.
		WithField("content", string(content.([]byte))).
		WithField("publisher", "lcd").
		Debug("publisher: stubbing publish")

	return nil
}
