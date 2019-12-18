package contrailutil

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/cobra"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"

	yaml "gopkg.in/yaml.v2"
)

func init() {
	ContrailUtil.AddCommand(checktCmd)
}

var checktCmd = &cobra.Command{
	Use:   "check",
	Short: "check data schema",
	Long:  `This command checks data schema against the correct one`,
	Run: func(cmd *cobra.Command, args []string) {
		err := check("/Users/dji/go/src/github.com/Juniper/contrail/dump.yaml")
		if err != nil {
			logutil.FatalWithStackTrace(err)
		}
	},
}

func check(inputFile string) error {
	var data map[string][]map[string]interface{}

	bodyBuff, err := fileutil.GetContent(inputFile)
	if err != nil {
		logutil.FatalWithStackTrace(err)
		return errors.Wrap(err, "Load content of input file failed")
	}

	err = yaml.Unmarshal(bodyBuff, &data)
	if err != nil {
		logutil.FatalWithStackTrace(err)
		return errors.Wrap(err, "YAML unmarshaling of file content failed")
	}

	for _, request := range data["resources"] {
		inputFieldNames := make(map[string]int)
		for key := range request["data"].(map[interface{}]interface{}) {
			inputFieldNames[key.(string)] = 1
		}

		event, err := services.NewEvent(services.EventOption{
			Operation: request["operation"].(string),
			Kind:      request["kind"].(string),
		})
		if err != nil {
			logutil.FatalWithStackTrace(err)
			return errors.Wrap(err, "Creation of new services.Event failed")
		}
		// TODO: get the type of field
		val := reflect.TypeOf(event.Request).Elem()
		field := val.Field(0)
		val = field.Type.Elem()
		field = val.Field(0)
		val = field.Type.Elem()
		standardFieldNames := make(map[string]struct{})
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			yamlDescription, _ := field.Tag.Lookup("yaml")
			yamlFieldName := strings.Split(yamlDescription, ",")[0]
			standardFieldNames[yamlFieldName] = struct{}{}
		}

		for fieldName := range inputFieldNames {
			_, ok := standardFieldNames[fieldName]
			if !ok {
				fmt.Println("Non-existent field name:", fieldName)
			}
		}
	}
	return nil
}
