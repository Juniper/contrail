package cache

import (
	"context"
	"time"

	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/services"
	log "github.com/sirupsen/logrus"
)

//EventProsessor processes cached events.
type EventProsessor struct {
	cache     *DB
	Processor services.EventProcessor
	Timeout   time.Duration
}

//NewEventProcessor makes an event processor for cache db.
func (cache *DB) NewEventProcessor(process services.EventProcessor) *EventProsessor {
	return &EventProsessor{
		cache:     cache,
		Processor: process,
		Timeout:   viper.GetDuration("cache.timeout"),
	}
}

//Start starts event processing.
func (p *EventProsessor) Start(ctx context.Context) {
	watcher := p.cache.AddWatcher(ctx, 0, p.Timeout)
	for {
		select {
		case e := <-watcher.ch:
			err := p.Processor.Process(e, p.Timeout)
			if err != nil {
				log.Warn(err)
			}
		case <-ctx.Done():
			log.Debugf("Process canceled by context")
			return
		}
	}
}
