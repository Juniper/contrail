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
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"

	pkglog "github.com/Juniper/contrail/pkg/log"
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
		kind := data["type"].(string)
		data["uuid"] = uuid
		event, err := services.NewEvent(&services.EventOption{
			Kind: kind,
			Data: data,
			UUID: uuid,
		})

		if err != nil {
			log.Warnf("failed to create event - skipping: %v", err)
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
	err := json.Unmarshal([]byte(data[0].(string)), &response)
	if err != nil {
		log.Fatal(err)
	}
	return response
}

func (o object) GetString(key string) string {
	data := o.Get(key)
	response, _ := data.(string)
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
			l, _ := data[propertyList[1]].([]interface{})
			data[propertyList[1]] = append(l, value)
		case "parent":
			data["parent_uuid"] = propertyList[2]
		case "ref":
			refProperty := propertyList[1] + "_refs"
			list, _ := data[refProperty].([]interface{})
			m, _ := value.(map[string]interface{})
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
	log.Debug("Dumping data from cassandra")
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
	name, _ := os.Hostname() // nolint: noerror
	return "contrail_process_" + name
}

//EventProducer send db update for processor.
//Event will be harvest from cassandra db and from amqp later on.
type EventProducer struct {
	Processor services.EventProcessor

	cassandraConfig Config
	amqpConfig      AmqpConfig

	log *log.Entry
}

//NewEventProducer makes event producer and couple it with processor given.
func NewEventProducer(
	processor services.EventProcessor,
) *EventProducer {
	cfg := GetConfig()
	return &EventProducer{
		cassandraConfig: cfg,
		amqpConfig: AmqpConfig{
			host:      viper.GetString("cache.cassandra.amqp"),
			queueName: getQueueName(),
		},
		Processor: processor,
		log:       pkglog.NewLogger("cassandra-event-producer"),
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
			operation, _ := data["oper"].(string)
			kind, _ := data["type"].(string)
			uuid, _ := data["uuid"].(string)
			obj, _ := data["obj_dict"].(map[string]interface{})
			event, err := services.NewEvent(&services.EventOption{
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
				p.log.WithError(err).WithFields(log.Fields{
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
			p.log.WithError(err).WithFields(log.Fields{
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
		config: cfg,
	}
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
	log.Debugf("Processing event %+v for cassandra", event)
	// connect to the cluster
	cluster := getCluster(p.config)
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	if err = handleEvent(session, event); err != nil {
		return nil, err
	}

	return event, nil
}

const selectResource = "SELECT key, column1, value FROM obj_uuid_table WHERE key = ?"
const insertQuery = "INSERT INTO obj_uuid_table (key, column1, value) VALUES (?, ?, ?)"
const deleteRowByColumn1 = "DELETE FROM obj_uuid_table WHERE key=? and column1=?"
const deleteResource = "DELETE FROM obj_uuid_table WHERE key=?"
const insertIntoFQNameTable = "INSERT INTO obj_fq_name_table (key, column1, value) VALUES (?, ?, ?)"
const deleteFromFQNameTable = "DELETE FROM obj_fq_name_table WHERE key=? and column1=?"

func handleEvent(session *gocql.Session, event *services.Event) error { // nolint: interfacer
	rsrc := event.GetResource()
	switch event.Operation() {
	case services.OperationCreate, services.OperationUpdate:
		cassandraMap, err := resourceToCassandraMap(rsrc)
		if err != nil {
			return err
		}

		// select the object's children from cassandra and update our object map
		iter := session.Query(selectResource, rsrc.GetUUID()).Iter()
		var uuid, column1, value string
		for iter.Scan(&uuid, &column1, &value) {
			if strings.HasPrefix(column1, "children") {
				cassandraMap[column1] = value
			}
			fmt.Println(uuid, column1, value)
		}

		batch := gocql.NewBatch(gocql.LoggedBatch)

		// delete the old object from cassandra
		withTimestamp(batch, deleteResource, rsrc.GetUUID())

		// insert the new version
		for column1, value := range cassandraMap {
			withTimestamp(batch, insertQuery, rsrc.GetUUID(), column1, value)
		}

		// if the object has a parent, update the parent's children
		if parentUUID := rsrc.GetParentUUID(); parentUUID != "" {
			updateChild(batch, rsrc)
		}

		updateFQName(batch, rsrc)
		if err = session.ExecuteBatch(batch); err != nil {
			return err
		}
	case services.OperationDelete:
		batch := gocql.NewBatch(gocql.LoggedBatch)
		deleteChild(batch, rsrc)
		withTimestamp(batch, deleteResource, rsrc.GetUUID())
		if err := session.ExecuteBatch(batch); err != nil {
			return err
		}
	}
	return nil
}

func withTimestamp(b *gocql.Batch, stmt string, args ...interface{}) {
	strs := strings.Split(stmt, " WHERE")
	stmt = strings.Join([]string{strs[0], " USING TIMESTAMP ", fmt.Sprint(time.Now().UnixNano())}, "")
	if len(strs) > 1 {
		stmt += " WHERE" + strs[1]
	}
	log.Warn(stmt)
	b.Query(stmt, args...)
}

func childColumn1(r services.Resource) string {
	childType := strings.Replace(r.Kind(), "_", "-", -1)
	return fmt.Sprintf("children:%s:%s", childType, r.GetUUID())
}

func getFQNameToUUIDColumn1(r services.Resource) string {
	return strings.Join(append(r.GetFQName(), r.GetUUID()), ":")
}

func updateFQName(b *gocql.Batch, r services.Resource) {
	withTimestamp(b, deleteFromFQNameTable, r.Kind(), getFQNameToUUIDColumn1(r))
	withTimestamp(b, insertIntoFQNameTable, r.Kind(), getFQNameToUUIDColumn1(r), "null")
}

func deleteChild(b *gocql.Batch, r services.Resource) {
	withTimestamp(b, deleteRowByColumn1, r.GetParentUUID(), childColumn1(r))
}

func updateChild(b *gocql.Batch, r services.Resource) {
	withTimestamp(b, deleteRowByColumn1, r.GetParentUUID(), childColumn1(r))
	withTimestamp(b, insertQuery, r.GetParentUUID(), childColumn1(r), "null")
}

//AmqpEventProcessor implements EventProcessor
type AmqpEventProcessor struct {
	config AmqpConfig
}

//NewAmqpEventProcessor return Amqp event processor
func NewAmqpEventProcessor() *AmqpEventProcessor {
	return &AmqpEventProcessor{
		config: AmqpConfig{
			host:      viper.GetString("cache.cassandra.amqp"),
			queueName: getQueueName(),
		},
	}
}

//AmqpMessage type
type AmqpMessage struct {
	RequestID string   `json:"request_id"`
	Oper      string   `json:"oper"`
	Type      string   `json:"type"`
	UUID      string   `json:"uuid"`
	FqName    []string `json:"fq_name"`
	Data      []byte   `json:"obj_dict"`
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
		Data:      data,
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
