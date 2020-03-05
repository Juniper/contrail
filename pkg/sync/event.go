package sync

import (
	"context"
)

// EventProducer based on RDBMS updates.
type EventProducer struct {
	Watcher watchCloser
}

// NewEventProducer makes EventProducer based RDBMS updates.
// Every EventProducer must have a unique id.
func NewEventProducer(id string, processor eventProcessor) (*EventProducer, error) {
	watcher, err := createWatcher(id, processor)
	if err != nil {
		return nil, err
	}
	return &EventProducer{
		Watcher: watcher,
	}, nil
}

// Start runs EventProducer.
func (e *EventProducer) Start(ctx context.Context) error {
	return e.Watcher.Watch(ctx)
}
