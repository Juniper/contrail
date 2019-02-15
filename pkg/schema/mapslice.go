package schema

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/logutil"
)

type mapSlice yaml.MapSlice

func (s mapSlice) get(key string) interface{} {
	for _, i := range s {
		k := i.Key.(string) //nolint: errcheck
		if k == key {
			return i.Value
		}
	}
	return nil
}

func (s mapSlice) keys() []string {
	result := []string{}
	for _, i := range s {
		k := i.Key.(string) //nolint: errcheck
		result = append(result, k)
	}
	return result
}

func (s mapSlice) getString(key string) string {
	i := s.get(key)
	result, _ := i.(string) //nolint: errcheck
	return result
}

func (s mapSlice) getBool(key string) bool {
	i := s.get(key)
	result, _ := i.(bool) //nolint: errcheck
	return result
}

func (s mapSlice) getMapSlice(key string) mapSlice {
	i := s.get(key)
	if i == nil {
		return nil
	}
	result := i.(yaml.MapSlice) //nolint: errcheck
	return mapSlice(result)
}

func (s mapSlice) getStringSlice(key string) []string {
	i := s.get(key)
	if i == nil {
		return nil
	}
	iResult := i.([]interface{}) //nolint: errcheck
	result := []string{}
	for _, a := range iResult {
		result = append(result, a.(string)) //nolint: errcheck
	}
	return result
}

var overridenTypes = map[string]struct{}{
	"types.json#/definitions/AccessType":       {},
	"types.json#/definitions/L4PortType":       {},
	"types.json#/definitions/IpAddressType":    {},
	"types.json#/definitions/MacAddressesType": {},
}

// JSONSchema creates JSONSchema using mapSlice data.
func (s mapSlice) JSONSchema() *JSONSchema {
	if s == nil {
		return nil
	}
	properties := s.getMapSlice("properties")
	schema := &JSONSchema{
		Title:           s.getString("title"),
		SQL:             s.getString("sql"),
		JSONTag:         s.getString("json_tag"),
		Default:         s.get("default"),
		Enum:            s.getStringSlice("enum"),
		Minimum:         s.get("minimum"),
		Maximum:         s.get("maximum"),
		Ref:             s.getString("$ref"),
		Permission:      s.getStringSlice("permission"),
		Operation:       s.getString("operation"),
		Type:            s.getString("type"),
		GoType:          s.getString("go_type"),
		ProtoType:       s.getString("proto_type"),
		Presence:        s.getString("presence"),
		Description:     s.getString("description"),
		Format:          s.getString("format"),
		Required:        s.getStringSlice("required"),
		Properties:      map[string]*JSONSchema{},
		PropertiesOrder: properties.keys(),
		CollectionType:  s.getString("collectionType"),
		MapKey:          s.getString("mapKey"),
	}
	if properties == nil {
		schema.Properties = nil
	}
	schema.OrderedProperties = []*JSONSchema{}
	for _, property := range properties {
		key := property.Key.(string) //nolint: errcheck
		if property.Value == nil {
			logutil.FatalWithStackTrace(errors.Errorf("property %s is null on %v", key, schema))
		}
		propertySchema := mapSlice(property.Value.(yaml.MapSlice)).JSONSchema() //nolint: errcheck

		// TODO: remove this workaround when schema is updated for zero-value required properties
		_, present := overridenTypes[propertySchema.Ref]

		if (present || propertySchema.Type == "boolean") &&
			(propertySchema.Presence == "required" || propertySchema.Presence == "true") {
			logrus.Warnf("property %s should be optional as it may have zero-value. Update schema.", key)
			logrus.Warnf("JSONSCHEMA: %v", propertySchema)
			propertySchema.Presence = "optional"
		}

		propertySchema.ID = key
		schema.Properties[key] = propertySchema
		schema.OrderedProperties = append(schema.OrderedProperties, propertySchema)
	}
	items := s.getMapSlice("items")
	schema.Items = items.JSONSchema()
	return schema
}
