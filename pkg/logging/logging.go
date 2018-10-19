package logging

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//SetLogLevel set global log level using viper configuraion.
func SetLogLevel() {
	logLevel := viper.GetString("log_level")
	switch logLevel {
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}
