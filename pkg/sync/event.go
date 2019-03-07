package sync

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/services"
)

// EventProducer based on RDBMS updates.
type EventProducer struct {
	watcher watchCloser
	log     *logrus.Entry
}

// NewEventProducer makes EventProducer based RDBMS updates.
// Every EventProducer must have unique id.
func NewEventProducer(id string, processor services.EventProcessor) (*EventProducer, error) {
	watcher, err := createWatcher(id, processor)
	if err != nil {
		return nil, err
	}
	return &EventProducer{
		log:     logutil.NewLogger("sync-event-producer"),
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
