package contrailutil

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var oldJSONFile string

type contrailDBData struct {
	Cassandra *cassandra `json:"cassandra"`
}

type object map[string][]interface{}
type uuidTable map[string][]interface{}

type cassandra struct {
	ObjFQNameTable map[string]uuidTable `json:"obj_fq_name_table"`
	ObjUUIDTable   map[string]object    `json:"obj_uuid_table"`
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
}

func (o object) convert(uuid string) map[string]interface{} {
	data := map[string]interface{}{
		"uuid": uuid,
	}
	for property := range o {
		value := o.Get(property)
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
				data["parent_uuid"] = propertyList[1]
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
	return data
}

func importData() {
	raw, err := ioutil.ReadFile(oldJSONFile)
	if err != nil {
		log.Fatal(err)
	}
	etcd := map[string]interface{}{}
	var data contrailDBData
	err = json.Unmarshal(raw, &data)
	if err != nil {
		log.Fatal(err)
	}
	prefix := "/contrail"
	for uuid, object := range data.Cassandra.ObjUUIDTable {
		objectType := object.GetString("type")
		if objectType != "" {
			key := filepath.Join(prefix, objectType, uuid)
			etcd[key] = object.convert(uuid)
		}
	}
	output, err := json.Marshal(etcd)
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
		importData()
	},
}
