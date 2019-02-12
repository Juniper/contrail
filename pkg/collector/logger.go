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

// AddCollectorHook setup logrus logger to send entries to collector
func AddCollectorHook(c *Collector) {
	logrus.AddHook(&collectorLoggerHook{
		collector: c,
	})
}

// NotCollectorEntry return entry with field marking this message not to send
// to collector
func NotCollectorEntry(entry *logrus.Entry) *logrus.Entry {
	return entry.WithField(doNotSendCollectorKey, true)
}

// NotCollectorEntry return entry with field marking this message not to send
// to collector
func NotCollectorLogger() *logrus.Entry {
	return logrus.WithField(doNotSendCollectorKey, true)
}
