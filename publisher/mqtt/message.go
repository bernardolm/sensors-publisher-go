package mqtt

import "time"

type message struct {
	Qos    byte      `json:"qos"`
	Time   time.Time `json:"time"`
	Topic1 string    `json:"t"`
	Topic2 string    `json:"topic"`
}

func (t *message) Topic() string {
	if t.Topic1 != "" {
		return t.Topic1
	} else if t.Topic2 != "" {
		return t.Topic2
	}

	return "message: empty topic"
}
