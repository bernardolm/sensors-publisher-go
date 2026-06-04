package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

var Version string

// Load env vars and config file to get app config
func Load() {
	filenames := []string{
		"/etc/sensors-publisher-go/config.env",
		".env",
	}

	for _, filename := range filenames {
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			continue
		}

		if err := godotenv.Overload(filename); err != nil {
			log.
				WithError(err).
				WithField("filename", filename).
				Warn("config.Load: error loading env file")
		}

		return
	}
}

// Get value from env vars or config file
func Get[T any](key string) T {
	var out T
	var value any

	e := log.
		WithField("env_var", key).
		WithField("type", fmt.Sprintf("%T", out))

	rawValue := os.Getenv(key)

	if rawValue == "" {
		e.Debug("config.Get: value not found or not exist")

		return out
	}

	e = e.WithField("value", rawValue)

	e.Debug("config.Get: value recovered")

	switch any(out).(type) {
	case bool:
		value = cast.ToBool(rawValue)
	case float64:
		value = cast.ToFloat64(rawValue)
	case int:
		value = cast.ToInt(rawValue)
	case int64:
		value = cast.ToInt64(rawValue)
	case string:
		value = rawValue
	case time.Duration:
		value = cast.ToDuration(rawValue)
	default:
		value = *new(T)
		e.Error("config.Get: unknown type to cast")
	}

	return value.(T)
}
