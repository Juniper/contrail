package collector

import (
	"github.com/sirupsen/logrus"
)

type collectorLoggerHook struct {
	collector *Collector
}

func (h *collectorLoggerHook) Fire(entry *logrus.Entry) error {
	h.collector.APIMessage(entry)
	return nil
}

func (h *collectorLoggerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// AddCollectorHook setup logrus logger to send entries to collector
func AddCollectorHook(c *Collector) {
	logrus.AddHook(&collectorLoggerHook{
		collector: c,
	})
}
