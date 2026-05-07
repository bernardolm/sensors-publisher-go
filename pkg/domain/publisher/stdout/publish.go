package stdout

import (
	"context"

	"github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/config"
	"github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/logging"
)

func (a *stdout) Publish(_ context.Context, content any) error {
	if !config.Get[bool]("DEBUG") || content == nil {
		return nil
	}

	logging.Log.
		WithField("content", string(content.([]byte))).
		WithField("publisher", "stdout").
		Debug("publisher: trying to publish")

	return nil
}
