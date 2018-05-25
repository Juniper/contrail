package schema

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

type mapSlice yaml.MapSlice

func (s mapSlice) get(key string) interface{} {
	for _, i := range s {
		k := i.Key.(string)
		if k == key {
			return i.Value
		}
	}
	return nil
}

func (s mapSlice) keys() []string {
	result := []string{}
	for _, i := range s {
		k := i.Key.(string)
		result = append(result, k)
	}
	return result
}

func (s mapSlice) getString(key string) string {
	i := s.get(key)
	result, _ := i.(string)
	return result
}

func (s mapSlice) getInt(key string) int {
	i := s.get(key)
	result, _ := i.(int)
	return result
}

func (s mapSlice) getMapSlice(key string) mapSlice {
	i := s.get(key)
	if i == nil {
		return nil
	}
	result := i.(yaml.MapSlice)
	return mapSlice(result)
}

func (s mapSlice) getStringSlice(key string) []string {
	i := s.get(key)
	if i == nil {
		return nil
	}
	iResult := i.([]interface{})
	result := []string{}
	for _, a := range iResult {
		result = append(result, a.(string))
	}
	return result
}

//Copy copies a json schema
func (s mapSlice) JSONSchema() *JSONSchema {
	if s == nil {
		return nil
	}
	properties := s.getMapSlice("properties")
	schema := &JSONSchema{
		Title:           s.getString("title"),
		SQL:             s.getString("sql"),
		Default:         s.get("default"),
		Enum:            s.getStringSlice("enum"),
		Minimum:         s.get("minimum"),
		Maximum:         s.get("maximum"),
		Ref:             s.getString("$ref"),
		Permission:      s.getStringSlice("permission"),
		Operation:       s.getString("operation"),
		Type:            s.getString("type"),
		Presence:        s.getString("presence"),
		Description:     s.getString("description"),
		Format:          s.getString("format"),
		Required:        s.getStringSlice("required"),
		Properties:      map[string]*JSONSchema{},
		PropertiesOrder: properties.keys(),
	}
	if properties == nil {
		schema.Properties = nil
	}
	schema.OrderedProperties = []*JSONSchema{}
	for _, property := range properties {
		key := property.Key.(string)
		if property.Value == nil {
			log.Fatal(fmt.Sprintf("Property is null on key %s", key))
		}
		propertySchema := mapSlice(property.Value.(yaml.MapSlice)).JSONSchema()
		propertySchema.ID = key
		schema.Properties[key] = propertySchema
		schema.OrderedProperties = append(schema.OrderedProperties, propertySchema)
	}
	items := s.getMapSlice("items")
	schema.Items = items.JSONSchema()
	return schema
}

//Reference convert a mapslice for reference
func (s mapSlice) Reference() *Reference {
	if s == nil {
		return nil
	}
	reference := &Reference{
		Description: s.getString("description"),
		Operations:  s.getString("operations"),
		Presence:    s.getString("presence"),
		Ref:         s.getString("$ref"),
		AttrSlice:   yaml.MapSlice(s.getMapSlice("attr")),
	}
	return reference
}
