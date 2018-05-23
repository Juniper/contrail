package contrailutil

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Juniper/contrail/pkg/db"

	"encoding/json"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/gocql/gocql"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	defaultCassandraVersion  = "3.4.4"
	defaultCassandraKeyspace = "config_db_uuid"
	yamlType                 = "yaml"
	cassandraType            = "cassandra"
	cassandraDumpType        = "cassandra_dump"
	rdbmsType                = "rdbms"
)

func init() {
	ContrailUtil.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&inType, "intype", "", "",
		"input type: cassandra,cassandra_dump and yaml are supported")
	convertCmd.Flags().StringVarP(&inFile, "in", "i", "", "Input file or Cassandra host")
	convertCmd.Flags().StringVarP(&outType, "outtype", "", "",
		"output type: rdbms and yaml are supported")
	convertCmd.Flags().StringVarP(&outFile, "out", "o", "", "Output file")
}

var inType, inFile string
var outType, outFile string

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

func (t table) makeSyncResources() *services.ResourceList {
	request := &services.ResourceList{
		Resources: []*services.ResourceEvent{},
	}
	for uuid, data := range t {
		kind := data["type"].(string)
		data["uuid"] = uuid
		request.Resources = append(request.Resources, &services.ResourceEvent{
			Kind: kind,
			Data: data,
		})
	}
	return request
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

func readCassandra() (*services.ResourceList, error) {
	// connect to the cluster
	cluster := gocql.NewCluster(inFile)
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
	return t.makeSyncResources(), nil
}

func readCassandraDump() (*services.ResourceList, error) {
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
	return t.makeSyncResources(), nil
}

func readYAML() (*services.ResourceList, error) {
	var resources services.ResourceList
	err := common.LoadFile(inFile, &resources)
	resources.YAMLCompat()
	return &resources, err

}

func writeYAML(resources *services.ResourceList) error {
	return common.SaveFile(outFile, resources)
}

func writeRDBMS(resources *services.ResourceList) error {
	dbService, err := db.NewServiceFromConfig()
	if err != nil {
		return errors.Wrap(err, "failed to connect DB")
	}
	ctx := context.Background()
	err = db.DoInTransaction(ctx,
		dbService.DB(),
		func(ctx context.Context) error {
			_, err = resources.Process(ctx, dbService)
			return err
		})
	return err
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert data format",
	Long:  `This command converts data formats from one to another`,
	Run: func(cmd *cobra.Command, args []string) {
		var resources *services.ResourceList
		var err error
		switch inType {
		case cassandraType:
			resources, err = readCassandra()
		case cassandraDumpType:
			resources, err = readCassandraDump()
		case yamlType:
			resources, err = readYAML()
		default:
			err = fmt.Errorf("Unsupported input type %s", inType)
		}
		if err != nil {
			log.Fatal(err)
		}

		err = resources.Sort()
		if err != nil {
			log.Fatal(err)
		}

		switch outType {
		case rdbmsType:
			err = writeRDBMS(resources)
		case yamlType:
			err = writeYAML(resources)
		default:
			err = fmt.Errorf("Unsupported input type %s", inType)
			log.Fatal("Unsupported output type")
		}
		if err != nil {
			log.Fatal(err)
		}
	},
}
