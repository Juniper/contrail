package sync

import (
	"context"

	"github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/sirupsen/logrus"
)

//EventProducer based on RDBMS updates.
type EventProducer struct {
	watcher watchCloser
	log     *logrus.Entry
}

//NewEventProducer makes EventProducer based RDBMS updates.
func NewEventProducer(processor services.EventProcessor) (*EventProducer, error) {
	log := log.NewLogger("sync-event-producer")
	watcher, err := createWatcher(log, processor)
	if err != nil {
		return nil, err
	}
	return &EventProducer{
		log:     log,
		watcher: watcher,
	}, nil
}

// Start runs EventProducer.
func (e *EventProducer) Start(ctx context.Context) error {
	e.log.Info("Running Sync service")
	return e.watcher.Watch(ctx)
}

// Close closes EventProducer.
func (e *EventProducer) Close() {
	e.log.Info("Closing Sync service")
	e.watcher.Close()
}
