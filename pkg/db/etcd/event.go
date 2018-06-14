package etcd

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/Juniper/contrail/pkg/services"
	"github.com/coreos/etcd/mvcc/mvccpb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//EventProducer watches etcd and call event processor.
type EventProducer struct {
	client    *Client
	Processor services.EventProcessor
	WatchPath string
	Timeout   time.Duration
}

//NewEventProducer makes a event processor.
func NewEventProducer(processor services.EventProcessor) (p *EventProducer, err error) {
	p = &EventProducer{
		Processor: processor,
		WatchPath: viper.GetString("etcd.path"),
		Timeout:   viper.GetDuration("cache.timeout"),
	}
	p.client, err = DialByConfig()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func parseKey(key string) (string, string) {
	subkeys := strings.Split(key, "/")
	return subkeys[2], subkeys[3]
}

// HandleMessage handles message received from etcd pubsub.
func (p *EventProducer) HandleMessage(
	ctx context.Context, index int64, oper int32, key string, newValue []byte,
) {
	log.Debug("Index: %d, oper: %d, Got Message %s: %s",
		index, oper, key, newValue)
	var data map[string]interface{}

	//TODO(nati) use sync.Codec

	var operation string
	kind, uuid := parseKey(key)

	switch oper {
	case int32(mvccpb.PUT):
		operation = services.OperationCreate
		err := json.Unmarshal(newValue, &data)
		if err != nil {
			log.Warn("decode error for %s", string(newValue))
			return
		}
	case int32(mvccpb.DELETE):
		operation = services.OperationDelete
	}
	e := services.NewEvent(&services.EventOption{
		UUID:      uuid,
		Kind:      kind,
		Operation: operation,
		Data:      data,
	})
	if e == nil {
		log.Warn("invalid event %v", data)
		return
	}
	p.Processor.Process(ctx, e)
}

//Start watch etcd.
func (p *EventProducer) Start(ctx context.Context) error {
	p.client.WatchRecursive(ctx, "/"+p.WatchPath, p.HandleMessage)
	return nil
}
