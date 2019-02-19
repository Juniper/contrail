package analytics

import (
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/collector"
)

const doNotSendCollectorKey = "no-collector"

type collectorLoggerHook struct {
	c collector.Collector
}

func (h *collectorLoggerHook) Fire(entry *logrus.Entry) error {
	if ignore, ok := entry.Data[doNotSendCollectorKey]; ok && ignore == true {
		return nil
	}

	h.c.Send(VNCAPIMessage(entry))
	return nil
}

func (h *collectorLoggerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// AddLoggerHook setup logrus logger to send entries to collector
func AddLoggerHook(c collector.Collector) {
	logrus.AddHook(&collectorLoggerHook{
		c: c,
	})
}

// ignoreAPIMessage add doNotSendCollectorKey key to logger. Used to avoid
// posible recursion during logging in collector
func ignoreAPIMessage() *logrus.Entry {
	return logrus.WithField(doNotSendCollectorKey, true)
}
