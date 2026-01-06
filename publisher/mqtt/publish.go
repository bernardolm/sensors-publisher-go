package mqtt

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/bernardolm/sensors-publisher-go/infrastructure/logging"
	"github.com/k0kubun/pp/v3"
	"github.com/tidwall/pretty"
)

func (a *mqtt) Publish(ctx context.Context, content any) error {
	if content == nil {
		return nil
	}

	e := logging.Log.WithField("publisher", "mqtt")

	payloads, ok := content.([]any)
	if !ok {
		return fmt.Errorf("publisher: mqtt payload in unkown type %T", content)
	}

	for _, p := range payloads {
		b, ok := p.([]byte)
		if !ok {
			return fmt.Errorf("publisher: mqtt payload in unkown type %T", content)
		}

		e = e.
			WithField("payload", string(b))

		fmt.Println([]string{"\n\n", string(pretty.Pretty(b)), "\n\n"})

		m := message{Qos: 2}
		if err := json.Unmarshal(b, &m); err != nil {
			return err
		}

		pp.Println(m)

		e = e.
			WithField("qos", m.Qos).
			WithField("topic", m.Topic())

		e.Debug("publisher: trying to publish")

		if err := a.client.Publish(
			m.Topic(),
			m.Qos,
			true,
			b); err != nil {
			return err
		}

		e.Debug("publisher: published")
	}

	return nil
}
