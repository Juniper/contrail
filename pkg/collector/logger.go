package collector

import (
	"github.com/sirupsen/logrus"
)

const doNotSendCollectorKey = "no-collector"

type collectorLoggerHook struct {
	collector *Collector
}

func (h *collectorLoggerHook) Fire(entry *logrus.Entry) error {
	if ignore, ok := entry.Data[doNotSendCollectorKey]; ok && ignore == true {
		return nil
	}

	h.collector.APIMessage(entry)
	return nil
}

func (h *collectorLoggerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// AddLoggerHook setup logrus logger to send entries to collector
func AddLoggerHook(c *Collector) {
	logrus.AddHook(&collectorLoggerHook{
		collector: c,
	})
}

// ignoreAPIMessage add doNotSendCollectorKey key to logger. Used to avoid
// posible recursion during logging in collector
func ignoreAPIMessage() *logrus.Entry {
	return logrus.WithField(doNotSendCollectorKey, true)
}
