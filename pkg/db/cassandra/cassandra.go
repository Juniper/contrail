package cassandra

import (
	"io/ioutil"
	"strings"
	"time"

	"encoding/json"

	"github.com/Juniper/contrail/pkg/services"
	"github.com/gocql/gocql"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"golang.org/x/net/context"
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
			services.InterfaceToEvent(map[string]interface{}{
				"kind": kind,
				"data": data,
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
		case "propm":
			m, _ := data[propertyList[1]].(map[string]interface{})
			if m == nil {
				m = map[string]interface{}{}
				data[propertyList[1]] = m
			}
			mValue, _ := value.(map[string]interface{})
			m[propertyList[2]] = mValue["value"]
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
func DumpCassandra(host string) (*services.EventList, error) {
	log.Debug("Dumping data from cassandra")
	// connect to the cluster
	cluster := gocql.NewCluster(host)
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

//EventProcessor send db update for processor.
type EventProcessor struct {
	cassandraHost string
	rabbitMQHost  string
	queueName     string
	Processor     services.EventProcessor
	Timeout       time.Duration
}

//NewEventProcessor makes new event processor for cassandra.
func NewEventProcessor(processor services.EventProcessor, queueName, cassandraHost, rabbitMQHost string) *EventProcessor {
	return &EventProcessor{
		cassandraHost: cassandraHost,
		queueName:     queueName,
		rabbitMQHost:  rabbitMQHost,
		Processor:     processor,
		Timeout:       10 * time.Second,
	}
}

//WatchAMQP watches AMQP.
func (p *EventProcessor) WatchAMQP(ctx context.Context) error {
	log.Debug("starting watch amqp")
	conn, err := amqp.Dial(p.rabbitMQHost)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

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
			var obj interface{}
			err := json.Unmarshal(d.Body, &data)
			if err != nil {
				log.Warnf("decode error for %s", string(d.Body))
				continue
			}
			operation := data["oper"]
			kind := data["type"]
			uuid := data["uuid"]
			obj = data["obj_dict"]
			e := services.InterfaceToEvent(
				map[string]interface{}{
					"operation": operation,
					"data":      obj,
					"kind":      kind,
					"uuid":      uuid,
				},
			)
			if e == nil {
				log.Warn("invalid event %v", data)
				continue
			}
			p.Processor.Process(e, p.Timeout)
		case <-ctx.Done():
			log.Debug("AQMP watcher canncelled by context")
			return nil
		}
	}
}

//Start starts event processor for cassandra.
func (p *EventProcessor) Start(ctx context.Context) error {
	events, err := DumpCassandra(p.cassandraHost)
	if err != nil {
		return err
	}
	events.Sort()
	for _, e := range events.Events {
		p.Processor.Process(e, p.Timeout)
	}
	return p.WatchAMQP(ctx)
}
