package log

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestConfigureLoggingFailsWhenInvalidLevelGiven(t *testing.T) {
	tests := []struct {
		name  string
		level string
	}{
		{"Empty", ""},
		{"Invalid", "invalid"},
		{"TrailingWhitespace", "warn "},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := Configure(test.level)

			assert.Error(t, err)
		})
	}
}

func TestConfigureLoggingSetsMinimalLevelPackageVariable(t *testing.T) {
	tests := []struct {
		name  string
		level string
		out   logrus.Level
	}{
		{"Debug", "debug", logrus.DebugLevel},
		{"Info", "info", logrus.InfoLevel},
		{"Warn", "warn", logrus.WarnLevel},
		{"Error", "error", logrus.ErrorLevel},
		{"MixedCase", "InFo", logrus.InfoLevel},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := Configure(test.level)

			assert.Nil(t, err)
			assert.Equal(t, test.out, minimalLevel)
		})
	}
}

func TestConfigureLoggingConfiguresGlobalLogger(t *testing.T) {
	globalLogger := logrus.StandardLogger()

	err := Configure("warn")

	assert.Nil(t, err)
	assert.Equal(t, os.Stdout, globalLogger.Out)
	assert.Equal(t, &logrus.TextFormatter{}, globalLogger.Formatter)
	assert.Equal(t, logrus.WarnLevel, globalLogger.Level)
}

func TestNewLoggerContainsLoggerName(t *testing.T) {
	l := NewLogger("test-logger")

	assert.Equal(t, "test-logger", l.Data[loggerKey])
}
