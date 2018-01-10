// Package log facilitates creating of configured Logrus logger.
// Logger created with NewLogger() should be preferred over global Logger.
// Use log.Debug() and log.Info() forms of logging.
// Use log.WithField() and log.WithFields() methods with "dash-case" keys for additional log parameters.
package log

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Configuration for new logger instances.
var (
	minimalLevel logrus.Level
	writer       io.Writer = os.Stdout
)

const loggerKey = "logger"

// Configure configures global Logrus logger and sets configuration for new logger instances.
func Configure(level string) error {
	l, err := logrus.ParseLevel(level)
	if err != nil {
		return fmt.Errorf("parse log level: %s", err)
	}

	minimalLevel = l

	logrus.SetLevel(l)
	logrus.SetOutput(writer)
	return nil
}

// NewLogger creates configured logrus.Entry instance.
func NewLogger(loggerName string) *logrus.Entry {
	l := &logrus.Logger{
		Out:       writer,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     minimalLevel,
	}
	return l.WithField(loggerKey, loggerName)
}
