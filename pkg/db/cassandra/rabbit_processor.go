package cassandra

import (
	"context"
	"encoding/json"
	"fmt"

	uuid "github.com/satori/go.uuid"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"

	"github.com/Juniper/contrail/pkg/services"
)

//AmqpEventProcessor implements EventProcessor
type AmqpEventProcessor struct {
	config AmqpConfig
}

//NewAmqpEventProcessor return Amqp event processor
func NewAmqpEventProcessor() *AmqpEventProcessor {
	return &AmqpEventProcessor{
		config: AmqpConfig{
			host:      viper.GetString("amqp.url"),
			queueName: getQueueName(),
		},
	}
}

//AmqpMessage type
type AmqpMessage struct {
	RequestID string          `json:"request_id"`
	Oper      string          `json:"oper"`
	Type      string          `json:"type"`
	UUID      string          `json:"uuid"`
	FqName    []string        `json:"fq_name"`
	Data      json.RawMessage `json:"obj_dict"`
}

//Process sends msg to amqp exchange
func (p *AmqpEventProcessor) Process(ctx context.Context, event *services.Event) (*services.Event, error) {
	log.Debugf("Processing event %+v for amqp", event)

	conn, err := amqp.Dial(p.config.host)
	if err != nil {
		return nil, err
	}
	defer conn.Close() // nolint: errcheck

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	defer ch.Close() // nolint: errcheck

	rsrc := event.GetResource()
	data, err := json.Marshal(rsrc.ToMap())
	if err != nil {
		return nil, err
	}

	msg := AmqpMessage{
		RequestID: fmt.Sprintf("req-%s", uuid.NewV4().String()),
		Oper:      event.Operation(),
		Type:      rsrc.Kind(),
		UUID:      rsrc.GetUUID(),
		FqName:    rsrc.GetFQName(),
		Data:      json.RawMessage(data),
	}

	msgJSON, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	err = ch.Publish(
		exchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msgJSON,
		},
	)
	if err != nil {
		return nil, err
	}
	return event, nil
}
