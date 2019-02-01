// Package log facilitates creating of configured Logrus logger.
// Logger created with NewLogger() should be preferred over global Logger.
// Use log.Debug() and log.Info() forms of logging.
// Use log.WithField() and log.WithFields() methods with "dash-case" keys for additional log parameters.
package log

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	shellwords "github.com/mattn/go-shellwords"
	"github.com/sirupsen/logrus"
	"github.com/yudai/gotty/backend/localcommand"
	"github.com/yudai/gotty/server"
	"github.com/yudai/gotty/utils"
)

// Configuration for new logger instances.
var (
	minimalLevel           = logrus.DebugLevel
	writer       io.Writer = os.Stdout
)

const (
	loggerKey        = "logger"
	streamServerPort = "9011"
)

// Configure configures global Logrus logger and sets configuration for new logger instances.
func Configure(level string) error {
	if level == "" {
		level = "debug"
	}

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
// TODO(Daniel): remove that and use Configure() only
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

// StreamServer represents log streaming server
type StreamServer struct {
	terminalCommand string
	listenPort      string
	shutDown        context.CancelFunc
}

// NewStreamServer creates a log streaming server
func NewStreamServer(filename string) *StreamServer {
	c := fmt.Sprintf("tail -f %s", filename)
	l := &StreamServer{
		terminalCommand: c,
		listenPort:      streamServerPort,
	}
	return l
}

// Serve starts the log server
func (l *StreamServer) Serve() {
	// fill server options
	appOptions := &server.Options{}
	if err := utils.ApplyDefaultValues(appOptions); err != nil {
		logrus.Fatal(err)

	}
	appOptions.Port = l.listenPort

	// fill command factory backend options
	backendOptions := &localcommand.Options{}
	if err := utils.ApplyDefaultValues(backendOptions); err != nil {
		logrus.Fatal(err)
	}
	args, err := shellwords.Parse(l.terminalCommand)
	if err != nil {
		logrus.Fatal(err)
	}

	// create command factory
	factory, err := localcommand.NewFactory(args[0], args[1:], backendOptions)
	if err != nil {
		logrus.Fatal(err)
	}

	// create server
	srv, err := server.New(factory, appOptions)
	if err != nil {
		logrus.Fatal(err)
	}

	// run server
	ctx, cancel := context.WithCancel(context.Background())
	gCtx, gCancel := context.WithCancel(context.Background())
	l.shutDown = gCancel
	go func() {
		err := srv.Run(ctx, server.WithGracefullContext(gCtx))
		if err != nil {
			cancel()
		}
	}()
}

// Close stops serving log server
func (l *StreamServer) Close() {
	l.shutDown()
}

// FatalWithStackTrace logs error with extended format and calls os.Exit(1)
// If given error is constructed with pkg/errors library, stack trace is printed.
// See: https://godoc.org/github.com/pkg/errors#hdr-Formatted_printing_of_errors
func FatalWithStackTrace(err error) {
	log.Fatalf("%+v", err)
}
