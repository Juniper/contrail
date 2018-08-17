package etcd

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/services"
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
	e, err := DialByConfig()
	p.client = NewClient(e)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// HandleMessage handles message received from etcd pubsub.
func (p *EventProducer) HandleMessage(
	ctx context.Context, index int64, oper int32, key string, newValue []byte,
) {
	log.Debugf("Index: %d, oper: %d, Got Message %s: %s",
		index, oper, key, newValue)

	event, err := ParseEvent(oper, key, newValue)
	if err != nil {
		log.WithError(err).Error("Failed to parse event")
		return
	}

	_, err = p.Processor.Process(ctx, event)
	if err != nil {
		log.WithError(err).Error("Failed to process event")
	}
}

// ParseEvent returns an Event corresponding to a change in ETCD.
func ParseEvent(oper int32, key string, newValue []byte) (*services.Event, error) {

	//TODO(nati) use sync.Codec

	kind, uuid, err := parseKey(key)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse kind and UUID from etcd key: %s", key)
	}

	operation, err := parseOperation(oper)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse ETCD operation")
	}

	var data map[string]interface{}
	if operation == services.OperationCreate {
		err = json.Unmarshal(newValue, &data)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to decode %s", string(newValue))
		}
	}

	event, err := services.NewEvent(&services.EventOption{
		UUID:      uuid,
		Kind:      kind,
		Operation: operation,
		Data:      data,
	})
	if err != nil {
		return nil, errors.Errorf("failed to create event from data: %v, reason: %v", data, err)
	}
	return event, nil
}

func parseKey(key string) (kind string, uuid string, err error) {
	subkeys := strings.Split(key, "/")

	if len(subkeys) < 4 {
		return "", "", errors.New("Key has too few fields")
	}
	kind = subkeys[2]
	uuid = subkeys[3]
	return kind, uuid, nil
}

func parseOperation(etcdOperation int32) (string, error) {
	switch etcdOperation {
	case int32(mvccpb.PUT):
		return services.OperationCreate, nil
	case int32(mvccpb.DELETE):
		return services.OperationDelete, nil
	default:
		return "", errors.Errorf("unsupported ETCD operation: %v", etcdOperation)
	}
}

//Start watch etcd.
func (p *EventProducer) Start(ctx context.Context) error {
	eventChan := p.client.WatchRecursive(ctx, "/"+p.WatchPath, int64(0))
	log.Debug("Starting handle loop")
	for {
		select {
		case <-ctx.Done():
			return nil
		case e, ok := <-eventChan:
			if !ok {
				return errors.New("event channel unsuspectingly closed")
			}
			p.HandleMessage(ctx, e.Revision, e.Type, e.Key, e.Value)
		}
	}
}
