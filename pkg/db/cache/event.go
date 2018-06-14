package cache

import (
	"context"

	"github.com/Juniper/contrail/pkg/services"
	log "github.com/sirupsen/logrus"
)

//EventProducer processes cached events.
type EventProducer struct {
	cache     *DB
	Processor services.EventProcessor
}

//NewEventProducer makes an event processor for cache db.
func (cache *DB) NewEventProducer(process services.EventProcessor) *EventProducer {
	return &EventProducer{
		cache:     cache,
		Processor: process,
	}
}

//Start starts event processing.
func (p *EventProducer) Start(ctx context.Context) {
	//adding watcher with verison 0 always success.
	watcher, _ := p.cache.AddWatcher(ctx, 0) // nolint: noerror
	for {
		select {
		case e := <-watcher.ch:
			_, err := p.Processor.Process(ctx, e)
			if err != nil {
				log.Warn(err)
			}
		case <-ctx.Done():
			log.Debugf("Process canceled by context")
			return
		}
	}
}
