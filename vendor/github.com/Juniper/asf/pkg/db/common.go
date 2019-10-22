package db

import (
	"encoding/json"

	"github.com/Juniper/asf/pkg/format"
	"github.com/Juniper/asf/pkg/services"
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
)

// Object is generic database model instance.
type Object interface {
	proto.Message
	ToMap() map[string]interface{}
}

// ParseFQName parses fqName string read from DB to string slice
func ParseFQName(fqNameStr string) ([]string, error) {
	var fqName []string
	err := json.Unmarshal([]byte(fqNameStr), &fqName)
	if err != nil {
		return nil, errors.Errorf("failed to parse fq name from string: %v", err)
	}
	return fqName, nil
}

func fqNameToString(fqName []string) (string, error) {
	fqNameStr, err := json.Marshal(fqName)
	if err != nil {
		return "", errors.Errorf("failed to parse fq name to string: %v", err)
	}
	return string(fqNameStr), nil
}

func makeInterfacePointerArray(length int) []interface{} {
	arr := make([]interface{}, length)
	for i := range arr {
		arr[i] = new(interface{})
	}
	return arr
}

// SingleObjectListSpec creates a list spec that gets one object with specified fields.
func SingleObjectListSpec(uuid string, fields []string) *services.ListSpec {
	return &services.ListSpec{
		Limit:  1,
		Detail: true,
		Shared: true,
		Fields: fields,
		Filters: []*services.Filter{
			{
				Key:    "uuid",
				Values: []string{uuid},
			},
		},
	}
}

// ResolveUUIDAndFQNameFromMap tries to extract uuid and fq_name fields from a dictionary.
func ResolveUUIDAndFQNameFromMap(m map[string]interface{}) (uuid string, fqName []string, err error) {
	uuid = format.InterfaceToString(m["to"])
	if uuid == "" {
		return "", nil, nil
	}
	fqNameStr := format.InterfaceToString(m["fq_name"])
	fqName, err = ParseFQName(fqNameStr)
	if err != nil {
		return "", nil, err
	}
	return uuid, fqName, nil
}
