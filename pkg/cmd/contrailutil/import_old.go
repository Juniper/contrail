package contrailutil

import (
	"fmt"
	"io/ioutil"
	"strings"

	"encoding/json"

	"github.com/Juniper/contrail/pkg/services"
	"github.com/gocql/gocql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

var oldJSONFile, host, keyspace string

type contrailDBData struct {
	Cassandra *cassandra `json:"cassandra"`
}

type object map[string][]interface{}
type uuidTable map[string][]interface{}

type cassandra struct {
	ObjFQNameTable map[string]uuidTable `json:"obj_fq_name_table"`
	ObjUUIDTable   map[string]object    `json:"obj_uuid_table"`
}

type database map[string]table
type table map[string]resource
type resource map[string]interface{}

func (t table) get(uuid string) resource {
	r, ok := t[uuid]
	if !ok {
		r = resource{}
		t[uuid] = r
	}
	return r
}

func (t table) makeSyncResources() *services.RESTSyncRequest {
	request := &services.RESTSyncRequest{
		Resources: []*services.RESTResource{},
	}
	for uuid, data := range t {
		kind := data["type"].(string)
		data["uuid"] = uuid
		request.Resources = append(request.Resources, &services.RESTResource{
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

func init() {
	ContrailUtil.AddCommand(importCmd)
	importCmd.Flags().StringVarP(&oldJSONFile, "input", "i", "", "Contrail old version dump file")
	importCmd.Flags().StringVarP(&host, "host", "", "localhost", "Cassandra host")
	importCmd.Flags().StringVarP(&keyspace, "keyspace", "k", "config_db_uuid", "Cassandra keyspace")
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

func importCassandraData() {
	// connect to the cluster
	cluster := gocql.NewCluster(host)
	cluster.Keyspace = keyspace
	cluster.CQLVersion = "3.4.4"
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	t := table{}
	// list all tweets
	iter := session.Query(`SELECT key, column1, value FROM obj_uuid_table`).Iter()
	var uuid, column, valueJSON string
	for iter.Scan(&uuid, &column, &valueJSON) {
		r := t.get(uuid)
		var value interface{}
		err := json.Unmarshal([]byte(valueJSON), &value)
		if err != nil {
			log.Warn(err)
		}
		parseProperty(r, column, value)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	output, err := yaml.Marshal(t.makeSyncResources())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}

func importJSONData() {
	raw, err := ioutil.ReadFile(oldJSONFile)
	if err != nil {
		log.Fatal(err)
	}
	t := table{}
	var data contrailDBData
	err = json.Unmarshal(raw, &data)
	if err != nil {
		log.Fatal(err)
	}
	for uuid, object := range data.Cassandra.ObjUUIDTable {
		objectType := object.GetString("type")
		if objectType != "" {
			t[uuid] = object.convert(uuid)
		}
	}
	output, err := yaml.Marshal(t.makeSyncResources())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))
}

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import old version server dump",
	Long:  `This command imports json data by cassandra_to_json.rb`,
	Run: func(cmd *cobra.Command, args []string) {
		if oldJSONFile != "" {
			importJSONData()
		} else {
			importCassandraData()
		}
	},
}
