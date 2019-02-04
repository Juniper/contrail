// Package log facilitates creating of configured Logrus logger.
// Logger created with NewLogger() should be preferred over global Logger.
// Use log.Debug() and log.Info() forms of logging.
// Use log.WithField() and log.WithFields() methods with "dash-case" keys for additional log parameters.
package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// Configuration for new logger instances.
var (
	minimalLevel           = logrus.DebugLevel
	dirMode                = 0755
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
	return newLogger(loggerName, writer)
}

// NewFileLogger creates configured logrus.Entry instance.
func NewFileLogger(loggerName string, filename string) *logrus.Entry {

	// create the log directory if it doesnt exist
	err := os.MkdirAll(filepath.Dir(filename), os.ModeDir)
	if err != nil {
		logrus.Error(err)
		logrus.Info("fall back to stdout writer")
		return newLogger(loggerName, writer)
	}

	// Create the log file if doesn't exist.
	// append if already exists.
	w, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		logrus.Error(err)
		logrus.Info("fall back to stdout writer")
		return newLogger(loggerName, writer)
	}
	return newLogger(loggerName, w)
}

// NewLogger creates configured logrus.Entry instance.
func newLogger(loggerName string, writer io.Writer) *logrus.Entry {
	l := &logrus.Logger{
		Out:       writer,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     minimalLevel,
	}
	return l.WithField(loggerKey, loggerName)
}

// SetLogLevel configure any level for any logger
func SetLogLevel(log *logrus.Entry, level string) {
	switch level {
	case "panic":
		log.Logger.SetLevel(logrus.PanicLevel)
	case "fatal":
		log.Logger.SetLevel(logrus.FatalLevel)
	case "error":
		log.Logger.SetLevel(logrus.ErrorLevel)
	case "warn":
		log.Logger.SetLevel(logrus.WarnLevel)
	case "debug":
		log.Logger.SetLevel(logrus.DebugLevel)
	default:
		log.Logger.SetLevel(logrus.InfoLevel)
	}
}
