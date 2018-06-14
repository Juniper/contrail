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

//EventProcessor watches etcd and call event processor.
type EventProcessor struct {
	client    *Client
	Processor services.EventProcessor
	WatchPath string
	Timeout   time.Duration
}

//NewEventProcessor makes a event processor.
func NewEventProcessor(processor services.EventProcessor) (p *EventProcessor, err error) {
	p = &EventProcessor{
		Processor: processor,
		WatchPath: viper.GetString("etcd.path"),
		Timeout:   10 * time.Second,
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
func (p *EventProcessor) HandleMessage(
	ctx context.Context, index int64, oper int32, key string, newValue []byte,
) {
	log.Debug("Index: %d, oper: %d, Got Message %s: %s",
		index, oper, key, newValue)
	var data interface{}

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
	e := services.InterfaceToEvent(
		map[string]interface{}{
			"operation": operation,
			"data":      data,
			"kind":      kind,
			"uuid":      uuid,
		},
	)
	if e == nil {
		log.Warn("invalid event %v", data)
		return
	}
	p.Processor.Process(e, p.Timeout)
}

//Start watch etcd.
func (p *EventProcessor) Start(ctx context.Context) error {
	p.client.WatchRecursive(ctx, "/"+p.WatchPath, p.HandleMessage)
	return nil
}
