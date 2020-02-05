package cassandra

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/gocql/gocql"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	defaultCassandraVersion  = "3.4.4"
	defaultCassandraKeyspace = "config_db_uuid"

	exchangeName = "vnc_config.object-update"
)

// Config fields for cassandra
type Config struct {
	Host           string
	Port           int
	Timeout        time.Duration
	ConnectTimeout time.Duration
}

//GetConfig returns cassandra Config filled with data from config file.
func GetConfig() Config {
	return Config{
		Host:           viper.GetString("cassandra.host"),
		Port:           viper.GetInt("cassandra.port"),
		Timeout:        viper.GetDuration("cassandra.timeout"),
		ConnectTimeout: viper.GetDuration("cassandra.connect_timeout"),
	}
}

type contrailDBData struct {
	Cassandra *cassandra `json:"cassandra"`
}

type object map[string][]interface{}
type uuidTable map[string][]interface{}

type cassandra struct {
	ObjFQNameTable map[string]uuidTable `json:"obj_fq_name_table"`
	ObjUUIDTable   map[string]object    `json:"obj_uuid_table"`
}

type table map[string]map[string]interface{}

func (t table) get(uuid string) map[string]interface{} {
	r, ok := t[uuid]
	if !ok {
		r = map[string]interface{}{}
		t[uuid] = r
	}
	return r
}

func (t table) makeEventList() *services.EventList {
	events := &services.EventList{
		Events: []*services.Event{},
	}
	for uuid, data := range t {
		kind := data["type"].(string) //nolint: errcheck
		data["uuid"] = uuid
		event, err := services.NewEvent(services.EventOption{
			Kind: kind,
			Data: data,
			UUID: uuid,
		})

		if err != nil {
			logrus.Warnf("failed to create event - skipping: %v", err)
			continue
		}

		events.Events = append(events.Events, event)
	}
	return events
}

func (o object) Get(key string) interface{} {
	data, ok := o[key]
	if !ok {
		return nil
	}
	if data == nil {
		return nil
	}
	var response interface{}
	err := json.Unmarshal([]byte(data[0].(string)), &response) //nolint: errcheck
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
	return response
}

func (o object) GetString(key string) string {
	data := o.Get(key)
	response, _ := data.(string) //nolint: errcheck
	return response
}

func parseProperty(data map[string]interface{}, property string, value interface{}) {
	if strings.Contains(property, ":") {
		propertyList := strings.Split(property, ":")
		switch propertyList[0] {
		case "prop":
			data[propertyList[1]] = value
		case "propl", "propm":
			// TODO: preserve order in case of propl
			l, _ := data[propertyList[1]].([]interface{}) //nolint: errcheck
			data[propertyList[1]] = append(l, value)
		case "parent":
			data["parent_uuid"] = propertyList[2]
		case "ref":
			refProperty := propertyList[1] + "_refs"
			list, _ := data[refProperty].([]interface{}) //nolint: errcheck
			m, _ := value.(map[string]interface{})       //nolint: errcheck
			data[refProperty] = append(list, map[string]interface{}{
				"uuid": propertyList[2],
				"attr": m["attr"],
			})
		}
	} else {
		data[property] = value
	}
}

func (o object) convert(uuid string) map[string]interface{} {
	data := map[string]interface{}{
		"uuid": uuid,
	}
	for property := range o {
		value := o.Get(property)
		parseProperty(data, property, value)
	}
	return data
}

//DumpCassandra load all existing cassandra data.
func DumpCassandra(cfg Config) (*services.EventList, error) {
	logrus.Debug("Dumping data from cassandra")
	// connect to the cluster
	cluster := getCluster(cfg)
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	t := table{}
	// list all tweets
	iter := session.Query(`SELECT key, column1, value FROM obj_uuid_table`).Iter()
	var uuid, column, valueJSON string
	for iter.Scan(&uuid, &column, &valueJSON) {
		r := t.get(uuid)
		var value interface{}
		err = json.Unmarshal([]byte(valueJSON), &value)
		if err != nil {
			return nil, err
		}
		parseProperty(r, column, value)
	}
	if err = iter.Close(); err != nil {
		return nil, err
	}
	return t.makeEventList(), nil
}

//ReadCassandraDump reads json dump of cassandra.
func ReadCassandraDump(inFile string) (*services.EventList, error) {
	raw, err := ioutil.ReadFile(inFile)
	if err != nil {
		return nil, err
	}
	t := table{}
	var data contrailDBData
	err = json.Unmarshal(raw, &data)
	if err != nil {
		return nil, err
	}
	for uuid, object := range data.Cassandra.ObjUUIDTable {
		objectType := object.GetString("type")
		if objectType != "" {
			t[uuid] = object.convert(uuid)
		}
	}
	return t.makeEventList(), nil
}

// AmqpConfig groups config fields for AMQP
type AmqpConfig struct {
	host      string
	queueName string
}

func getQueueName() string {
	name, _ := os.Hostname() // nolint: errcheck
	return "contrail_process_" + name
}

//EventProducer send db update for processor.
//Event will be harvest from cassandra db and from amqp later on.
type EventProducer struct {
	Processor services.EventProcessor

	cassandraConfig Config
	amqpConfig      AmqpConfig

	log *logrus.Entry
}

//NewEventProducer makes event producer and couple it with processor given.
func NewEventProducer(
	processor services.EventProcessor,
) *EventProducer {
	cfg := GetConfig()
	return &EventProducer{
		cassandraConfig: cfg,
		amqpConfig: AmqpConfig{
			host:      viper.GetString("amqp.url"),
			queueName: getQueueName(),
		},
		Processor: processor,
		log:       logutil.NewLogger("cassandra-event-producer"),
	}
}

//WatchAMQP watches AMQP.
//nolint: gocyclo
func (p *EventProducer) WatchAMQP(ctx context.Context) error {
	p.log.Debug("Starting watch on AMQP")
	conn, err := amqp.Dial(p.amqpConfig.host)
	if err != nil {
		return err
	}
	defer conn.Close() // nolint: errcheck

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close() // nolint: errcheck

	q, err := ch.QueueDeclare(
		p.amqpConfig.queueName,
		false, // durable
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	err = ch.QueueBind(
		q.Name,       // queue name
		"",           // routing key
		exchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	for {
		select {
		case d := <-msgs:
			var data map[string]interface{}
			if err := json.Unmarshal(d.Body, &data); err != nil {
				p.log.WithError(err).WithField("data", string(d.Body)).Warn("Decoding failed - ignoring")
				continue
			}
			operation, _ := data["oper"].(string)               //nolint: errcheck
			kind, _ := data["type"].(string)                    //nolint: errcheck
			uuid, _ := data["uuid"].(string)                    //nolint: errcheck
			obj, _ := data["obj_dict"].(map[string]interface{}) //nolint: errcheck
			event, err := services.NewEvent(services.EventOption{
				Operation: operation,
				Data:      obj,
				Kind:      kind,
				UUID:      uuid,
			})
			if err != nil {
				p.log.WithField("data", data).Warnf("Failed to create event - ignoring: %v", err)
				continue
			}

			if _, err = p.Processor.Process(ctx, event); err != nil {
				p.log.WithError(err).WithFields(logrus.Fields{
					"context": ctx,
					"event":   event,
				}).Warn("Processing event failed - ignoring")
			}
		case <-ctx.Done():
			p.log.Debug("AQMP watcher cancelled by context")
			return nil
		}
	}
}

//Start starts event processor for cassandra.
func (p *EventProducer) Start(ctx context.Context) error {
	events, err := DumpCassandra(p.cassandraConfig)
	if err != nil {
		return err
	}

	if err = events.Sort(); err != nil {
		p.log.WithError(err).Warn("Sorting events failed - ignoring")
	}

	for _, e := range events.Events {
		if _, err = p.Processor.Process(ctx, e); err != nil {
			p.log.WithError(err).WithFields(logrus.Fields{
				"context": ctx,
				"event":   e,
			}).Warn("Processing event failed - ignoring")
		}
	}
	return p.WatchAMQP(ctx)
}

// EventProcessor writes events to cassandra and implements service.EventProcessor interface
type EventProcessor struct {
	config Config
}

// NewEventProcessor returns new cassandra.EventProcessor
func NewEventProcessor() *EventProcessor {
	cfg := GetConfig()
	return &EventProcessor{
		config: cfg}
}

func getCluster(cfg Config) *gocql.ClusterConfig {
	cluster := gocql.NewCluster(cfg.Host)
	if cfg.Port != 0 {
		cluster.Port = cfg.Port
	}
	cluster.Timeout = cfg.Timeout
	cluster.ConnectTimeout = cfg.ConnectTimeout
	cluster.Keyspace = defaultCassandraKeyspace
	cluster.CQLVersion = defaultCassandraVersion
	return cluster
}

// Process is a method needed to implement service.EventProcessor interface
func (p *EventProcessor) Process(ctx context.Context, event *services.Event) (*services.Event, error) {
	logrus.Debugf("Processing event %+v for cassandra", event)
	// connect to the cluster
	cluster := getCluster(p.config)
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	qry, err := getQuery(session, event)
	if err != nil {
		return nil, err
	}

	if err := qry.Exec(); err != nil {
		return nil, err
	}

	return event, nil
}

func getQuery(session *gocql.Session, event *services.Event) (qry *gocql.Query, err error) { // nolint: interfacer
	rsrc := event.GetResource()
	switch event.Operation() {
	case services.OperationCreate, services.OperationUpdate:
		rsrcJSON, err := json.Marshal(rsrc.ToMap())
		if err != nil {
			return nil, err
		}
		qry = session.Query(
			"INSERT INTO obj_uuid_table (key, column1, value) VALUES (?, ?, ?)",
			rsrc.GetUUID(),
			rsrc.Kind(),
			rsrcJSON,
		)
	case services.OperationDelete:
		qry = session.Query(
			"DELETE FROM obj_uuid_table WHERE key=? and column1=?",
			rsrc.GetUUID(),
			rsrc.Kind(),
		)
	}
	return qry, nil
}

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
	logrus.Debugf("Processing event %+v for amqp", event)

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
