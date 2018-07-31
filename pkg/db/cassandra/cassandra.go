package cassandra

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gocql/gocql"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	defaultCassandraVersion  = "3.4.4"
	defaultCassandraKeyspace = "config_db_uuid"

	exchangeName = "vnc_config.object-update"
)

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
		events.Events = append(events.Events,
			services.NewEvent(&services.EventOption{
				Kind: kind,
				Data: data,
				UUID: uuid,
			}))
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
		logrus.Fatal(err)
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
		case "propm":
			m, _ := data[propertyList[1]].(map[string]interface{})
			if m == nil {
				m = map[string]interface{}{}
				data[propertyList[1]] = m
			}
			CassandraToMapProperty(m, propertyList[1], value)
		case "propl":
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

// CassandraToMapProperty adds a single value from Cassandra to a property of type "map"
func CassandraToMapProperty(propMap map[string]interface{}, propName string, value interface{}) {
	// TODO (Kamil): temporary solution; this code should be generated from schema
	if propName == "virtual_machine_interface_bindings" {
		kvps, _ := propMap["key_value_pair"].([]interface{})
		kvps = append(kvps, value)
		propMap["key_value_pair"] = kvps
	}
}

// MapPropertyToCassandra turns a property of type "map" to a list of (property, value) pairs for Cassandra
func MapPropertyToCassandra(propMap map[string]interface{}, propName string) []struct {
	property string
	value    interface{}
} {
	// TODO (Kamil): temporary solution; this code should be generated from schema
	var result []struct {
		property string
		value    interface{}
	}

	if propName == "virtual_machine_interface_bindings" {
		kvps, _ := propMap["key_value_pair"].([]interface{})
		for _, kvp := range kvps {
			key := kvp.(map[string]interface{})["key"]
			result = append(result, struct {
				property string
				value    interface{}
			}{
				property: fmt.Sprintf("propm:%v:%v", propName, key),
				value:    kvp,
			})
		}
	}
	return result
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
func DumpCassandra(host string, port int, timeout time.Duration) (*services.EventList, error) {
	logrus.Debug("Dumping data from cassandra")
	// connect to the cluster
	cluster := gocql.NewCluster(host)
	if port != 0 {
		cluster.Port = port
	}
	cluster.Timeout = timeout
	cluster.Keyspace = defaultCassandraKeyspace
	cluster.CQLVersion = defaultCassandraVersion
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

//EventProducer send db update for processor.
type EventProducer struct {
	Processor services.EventProcessor

	cassandraHost    string
	cassandraPort    int
	cassandraTimeout time.Duration
	rabbitMQHost     string
	queueName        string

	log *logrus.Entry
}

//NewEventProducer makes new event processor for cassandra.
func NewEventProducer(
	processor services.EventProcessor,
	queueName string,
	cassandraHost string,
	cassandraPort int,
	cassandraTimeout time.Duration,
	rabbitMQHost string,
) *EventProducer {
	return &EventProducer{
		cassandraHost:    cassandraHost,
		cassandraPort:    cassandraPort,
		cassandraTimeout: cassandraTimeout,
		queueName:        queueName,
		rabbitMQHost:     rabbitMQHost,
		Processor:        processor,
		log:              pkglog.NewLogger("cassandra-event-producer"),
	}
}

//WatchAMQP watches AMQP.
//nolint: gocyclo
func (p *EventProducer) WatchAMQP(ctx context.Context) error {
	p.log.Debug("Starting watch on AMQP")
	conn, err := amqp.Dial(p.rabbitMQHost)
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
		p.queueName,
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
			err := json.Unmarshal(d.Body, &data)
			if err != nil {
				p.log.WithError(err).WithField("data", string(d.Body)).Warn("Decoding failed - ignoring")
				continue
			}
			operation, _ := data["oper"].(string)
			kind, _ := data["type"].(string)
			uuid, _ := data["uuid"].(string)
			obj, _ := data["obj_dict"].(map[string]interface{})
			e := services.NewEvent(&services.EventOption{
				Operation: operation,
				Data:      obj,
				Kind:      kind,
				UUID:      uuid,
			})
			if e == nil {
				p.log.WithField("data", data).Warn("Got invalid event - ignoring")
				continue
			}

			if _, err = p.Processor.Process(ctx, e); err != nil {
				p.log.WithError(err).WithFields(logrus.Fields{
					"context": ctx,
					"event":   e,
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
	events, err := DumpCassandra(p.cassandraHost, p.cassandraPort, p.cassandraTimeout)
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
