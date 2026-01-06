package logging

import (
	"os"

	"github.com/bernardolm/sensors-publisher-go/infrastructure/config"
	"github.com/sirupsen/logrus"
	"github.com/yukitsune/lokirus"
)

var Log = logrus.StandardLogger()

func Init() {
	var level logrus.Level = logrus.InfoLevel

	if ll := config.Get[string]("LOG_LEVEL"); ll != "" {
		var err error
		if level, err = logrus.ParseLevel(ll); err != nil {
			Log.WithError(err).Error("logging failed to set log level")
		}
	}

	Log.SetLevel(level)
	Log.SetOutput(os.Stdout)

	// hook, err := lSyslog.
	// 	NewSyslogHook(
	// 		"udp",
	// 		config.Get[string]("SYSLOG_HOST"),
	// 		syslog.LOG_NOTICE,
	// 		"")
	// if err == nil {
	// 	log.Hooks.Add(hook)
	// }

	lokiHost := config.Get[string]("LOKI_HOST")
	if lokiHost != "" {
		lokiOpts := lokirus.NewLokiHookOptions().
			WithFormatter(&logrus.JSONFormatter{}).
			WithStaticLabels(lokirus.Labels{
				"service": "sensors-publisher-go",
			})
		lokiHook := lokirus.NewLokiHookWithOpts(
			lokiHost, lokiOpts)
		Log.AddHook(lokiHook)
	}
}
