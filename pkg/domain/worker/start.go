package worker

import (
	"time"

	"github.com/bernardolm/sensors-publisher-go/pkg/infrastructure/logging"
)

func (w *worker) Start() error {
	for {
		for _, flow := range w.flows {
			var content any
			var err error

		tasker:
			for _, task := range flow {
				content, err = task(content)
				if err != nil {
					logging.Log.Error(err)

					break tasker
				}
			}
		}

		time.Sleep(w.delta)
	}
}
