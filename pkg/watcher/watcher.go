package watcher

import (
	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
)

type abstractCanal interface {
	Run() error
	Close()
}

type binlogWatcher struct {
	canal abstractCanal
	log   *logrus.Entry
}

func newBinlogWatcher(c abstractCanal) *binlogWatcher {
	if c == nil {
		c = &noopCanal{}
	}

	return &binlogWatcher{
		canal: c,
		log:   pkglog.NewLogger("binlog-watcher"),
	}
}

func (w *binlogWatcher) watch() error {
	w.log.Debug("Watching events on MySQL binlog")
	return w.canal.Run()
}

func (w *binlogWatcher) close() {
	w.log.Debug("Stopping watching events on MySQL binlog")
	w.canal.Close()
}

type noopCanal struct{}

func (c *noopCanal) Run() error { return nil }

func (c *noopCanal) Close() {}
