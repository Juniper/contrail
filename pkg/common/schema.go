package common

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

var sqlTypeMap = map[string]string{
	"object":  "json",
	"integer": "int",
	"array":   "json",
	"boolean": "bool",
	"number":  "float",
	"string":  "varchar(255)",
}

var sqlBindMap = map[string]string{
	"object":  "json",
	"integer": "int",
	"array":   "json",
	"boolean": "bool",
	"number":  "float",
	"string":  "string",
}

//API object has schemas and types for API definition.
type API struct {
	Schemas     []*Schema              `yaml:"schemas" json:"schemas,omitempty"`
	Definitions []*Schema              `yaml:"-" json:"-"`
	Types       map[string]*JSONSchema `yaml:"-" json:"-"`
}

//ColumnConfig is for database configuration.
type ColumnConfig struct {
	Path         string
	Type         string
	GoType       string
	Bind         string
	Column       string
	ParentColumn []string
	Name         string
	GoPremitive  bool
}

//ColumnConfigs is for list of columns
type ColumnConfigs []*ColumnConfig

func (c ColumnConfigs) Len() int {
	return len(c)
}
func (c ColumnConfigs) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c ColumnConfigs) Less(i, j int) bool {
	return strings.Compare(
		strings.Join(c[i].ParentColumn, "")+c[i].Column,
		strings.Join(c[j].ParentColumn, "")+c[j].Column,
	) > 0
}

func (c ColumnConfigs) shortenColumn() {
	sort.Sort(c)
	if len(c) < 2 {
		return
	}
	for i := 0; i < len(c)-1; i++ {
		for j := i + 1; j < len(c); j++ {
			if c[j].Column == c[i].Column {
				c[j].Column = c[j].ParentColumn[len(c[j].ParentColumn)-1] + "_" + c[j].Column
				c[j].ParentColumn = c[j].ParentColumn[:len(c[j].ParentColumn)-1]
			}
		}
	}
}

//Schema represents a data model
type Schema struct {
	FileName         string                   `yaml:"-" json:"-"`
	ID               string                   `yaml:"id" json:"id,omitempty"`
	Plural           string                   `yaml:"plural" json:"plural,omitempty"`
	Type             string                   `yaml:"type" json:"type,omitempty"`
	Title            string                   `yaml:"title" json:"title,omitempty"`
	Description      string                   `yaml:"description" json:"description,omitempty"`
	Parents          map[string]*Reference    `yaml:"parents" json:"parents,omitempty"`
	References       map[string]*Reference    `yaml:"references" json:"references,omitempty"`
	Prefix           string                   `yaml:"prefix" json:"prefix,omitempty"`
	JSONSchema       *JSONSchema              `yaml:"-" json:"schema,omitempty"`
	JSONSchemaSlice  yaml.MapSlice            `yaml:"schema" json:"-"`
	Definitions      map[string]*JSONSchema   `yaml:"-" json:"-"`
	DefinitionsSlice map[string]yaml.MapSlice `yaml:"definitions" json:"-"`
	Extends          []string                 `yaml:"extends" json:"extends,omitempty"`
	Columns          ColumnConfigs            `yaml:"-" json:"-"`
	Path             string                   `yaml:"-" json:"-"`
	PluralPath       string                   `yaml:"-" json:"-"`
	Children         []*Schema                `yaml:"-" json:"-"`
}

//JSONSchema is a standard JSONSchema representation plus data for code generation.
type JSONSchema struct {
	ID                string                 `yaml:"-" json:"-"`
	Index             int                    `yaml:"-" json:"-"`
	Title             string                 `yaml:"title" json:"title,omitempty"`
	Description       string                 `yaml:"description" json:"description,omitempty"`
	SQL               string                 `yaml:"sql" json:"-"`
	Default           interface{}            `yaml:"default" json:"default,omitempty"`
	Operation         string                 `yaml:"operation" json:"-"`
	Presence          string                 `yaml:"presence" json:"-"`
	Type              string                 `yaml:"type" json:"type,omitempty"`
	Permission        []string               `yaml:"permission" json:"permission,omitempty"`
	Properties        map[string]*JSONSchema `yaml:"properties" json:"properties,omitempty"`
	PropertiesOrder   []string               `yaml:"-" json:"propertiesOrder,omitempty"`
	OrderedProperties []*JSONSchema          `yaml:"-" json:"-"`
	Enum              []string               `yaml:"enum" json:"enum,omitempty"`
	Minimum           interface{}            `yaml:"minimum" json:"minimum,omitempty"`
	Maximum           interface{}            `yaml:"maximum" json:"maximum,omitempty"`
	Ref               string                 `yaml:"$ref" json:"-"`
	CollectionType    string                 `yaml:"-" json:"-"`
	Items             *JSONSchema            `yaml:"items" json:"items,omitempty"`
	GoName            string                 `yaml:"-" json:"-"`
	GoType            string                 `yaml:"-" json:"-"`
	ProtoType         string                 `yaml:"-" json:"-"`
	Required          []string               `yaml:"required" json:"-"`
	GoPremitive       bool                   `yaml:"-" json:"-"`
}

//String makes string format for json schema
func (s *JSONSchema) String() string {
	data, _ := json.Marshal(s) // nolint: gas
	return string(data)
}

//Reference object represents many to many relationships between resources.
type Reference struct {
	GoName      string        `yaml:"-" json:"-"`
	Description string        `yaml:"description" json:"description,omitempty"`
	Operation   string        `yaml:"operation" json:"operation,omitempty"`
	Presence    string        `yaml:"presence" json:"presence,omitempty"`
	RefType     string        `yaml:"-" json:"-"`
	Columns     ColumnConfigs `yaml:"-" json:"-"`
	Attr        *JSONSchema   `yaml:"-" json:"attr"`
	AttrSlice   yaml.MapSlice `yaml:"attr" json:"-"`
	LinkTo      *Schema       `yaml:"-" json:"-"`
	Ref         string        `yaml:"$ref" json:"$ref,omitempty"`
}

func parseRef(ref string) (string, string) {
	if ref == "" {
		return "", ""
	}
	refs := strings.Split(ref, "#")
	types := strings.Split(ref, "/")
	return refs[0], types[len(types)-1]
}

func (s *JSONSchema) getRefType() string {
	_, goType := parseRef(s.Ref)
	return goType
}

//Copy copies a json schema
func (s *JSONSchema) Copy() *JSONSchema {
	copied := &JSONSchema{
		ID:         s.ID,
		Title:      s.Title,
		SQL:        s.SQL,
		Default:    s.Default,
		Enum:       s.Enum,
		Minimum:    s.Minimum,
		Maximum:    s.Maximum,
		Ref:        s.Ref,
		Permission: s.Permission,
		Operation:  s.Operation,
		Type:       s.Type,
		Presence:   s.Presence,
		Required:   s.Required,
		Properties: map[string]*JSONSchema{},
	}
	for name, property := range s.Properties {
		copied.Properties[name] = property.Copy()
	}
	if s.Items != nil {
		copied.Items = s.Items.Copy()
	}
	return copied
}

//Update merges two JSONSchema
func (s *JSONSchema) Update(s2 *JSONSchema) {
	if s2 == nil {
		return
	}
	if s.Title == "" {
		s.Title = s2.Title
	}
	if s.Description == "" {
		s.Description = s2.Description
	}
	if s.SQL == "" {
		s.SQL = s2.SQL
	}
	if s.Default == nil {
		s.Default = s2.Default
	}
	if s.Operation == "" {
		s.Operation = s2.Operation
	}
	if s.Presence == "" {
		s.Presence = s2.Presence
	}

	if s.Properties == nil {
		s.Properties = map[string]*JSONSchema{}
	}
	for name, property := range s2.Properties {
		if _, ok := s.Properties[name]; !ok {
			s.Properties[name] = property.Copy()
		}
	}
	s.Required = append(s2.Required, s.Required...)
	s.PropertiesOrder = append(s2.PropertiesOrder, s.PropertiesOrder...)
	s.OrderedProperties = []*JSONSchema{}
	for _, id := range s.PropertiesOrder {
		s.OrderedProperties = append(s.OrderedProperties, s.Properties[id])
	}
	if s.Type == "" {
		s.Type = s2.Type
	}
	if s.Enum == nil {
		s.Enum = s2.Enum
	}
	if s.Minimum == nil {
		s.Minimum = s2.Minimum
	}
	if s.Maximum == nil {
		s.Maximum = s2.Maximum
	}
}

//Walk apply one function for json schema recursively.
func (s *JSONSchema) Walk(do func(s2 *JSONSchema) error) error {
	if s == nil {
		return nil
	}
	err := do(s)
	if err != nil {
		return err
	}
	if s.Properties == nil {
		return nil
	}
	for _, property := range s.Properties {
		err = property.Walk(do)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *JSONSchema) resolveSQL(parentColumn []string, columnName string, goPath string, columns *ColumnConfigs) error {
	if s == nil {
		return nil
	}
	if len(s.Properties) == 0 || s.CollectionType != "" || s.Type == "array" {
		if s.SQL == "" {
			s.SQL = sqlTypeMap[s.Type]
		}
		bind := ""
		if s.GoType != "" {
			bind = sqlBindMap[s.Type]
		}

		*columns = append(*columns, &ColumnConfig{
			Path:         goPath,
			GoType:       s.GoType,
			Bind:         bind,
			Type:         s.SQL,
			Column:       columnName,
			ParentColumn: parentColumn,
			GoPremitive:  s.GoPremitive,
			Name:         columnName,
		})
		return nil
	}
	for name, property := range s.Properties {
		nextParentColumn := make([]string, len(parentColumn))
		copy(nextParentColumn, parentColumn)
		nextParentColumn = append(nextParentColumn, columnName)
		err := property.resolveSQL(nextParentColumn, name, goPath+"."+property.GoName, columns)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *JSONSchema) resolveGoName(name string) error {
	if s == nil {
		return nil
	}
	s.GoName = SnakeToCamel(name)
	goType := s.getRefType()

	protoType := ""
	if goType == "" {
		s.GoPremitive = true
		switch s.Type {
		case "integer":
			goType = "int"
		case "number":
			goType = "float64"
		case "string":
			goType = "string"
		case "boolean":
			goType = "bool"
		case "object":
			if s.Properties == nil {
				goType = "map[string]interface{}"
			}
		case "array":
			err := s.Items.resolveGoName(name)
			if err != nil {
				return err
			}
			if s.Items == nil {
				goType = "[]interface{}"
			} else if s.Items.Type == "integer" || s.Items.Type == "number" || s.Items.Type == "boolean" ||
				s.Items.Type == "string" {
				goType = "[]" + s.Items.GoType
			} else {
				goType = "[]*" + s.Items.GoType
			}
		}
	}
	switch s.Type {
	case "integer":
		protoType = "int64"
	case "number":
		protoType = "float"
	case "string":
		protoType = "string"
	case "boolean":
		protoType = "bool"
	case "object":
		if goType != "" {
			protoType = goType
		}
		if s.Properties == nil {
			protoType = "bytes"
		}
	case "array":
		if s.Items == nil {
			protoType = "repeated bytes"
		} else {
			protoType = "repeated " + s.Items.ProtoType
		}
	}

	// if s.CollectionType == "list" {
	// 	goType = "[]string"
	// }
	// if s.CollectionType == "map" {
	// 	goType = "map[string]string"
	// }
	s.GoType = goType
	s.ProtoType = protoType
	for name, property := range s.Properties {
		err := property.resolveGoName(name)
		if err != nil {
			return err
		}
	}
	return nil
}

func (api *API) definitionByFileName(fileName string) *Schema {
	for _, s := range api.Definitions {
		if s.FileName == fileName {
			return s
		}
	}
	return nil
}

func (api *API) schemaByID(id string) *Schema {
	for _, s := range api.Schemas {
		if s.ID == id {
			return s
		}
	}
	return nil
}

func (api *API) loadType(schemaFile, typeName string) (*JSONSchema, error) {
	if definition, ok := api.Types[typeName]; ok {
		return definition, nil
	}
	definitions := api.definitionByFileName(schemaFile)
	if definitions == nil {
		return nil, fmt.Errorf("Can't find file for %s", schemaFile)
	}
	definition, ok := definitions.Definitions[typeName]
	if !ok {
		return nil, fmt.Errorf("%s isn't defined in %s", typeName, schemaFile)
	}
	err := definition.Walk(api.resolveRef)
	if err != nil {
		return nil, err
	}
	err = definition.resolveGoName("")
	if err != nil {
		return nil, err
	}
	api.Types[typeName] = definition
	return definition, nil
}

func (api *API) resolveRef(schema *JSONSchema) error {
	if schema == nil {
		return nil
	}
	if schema.Type == "array" {
		err := api.resolveRef(schema.Items)
		if err != nil {
			return err
		}
	}
	if schema.Ref == "" {
		return nil
	}
	definition, err := api.loadType(parseRef(schema.Ref))
	if err != nil {
		return err
	}
	schema.Update(definition)
	return nil
}

func (api *API) resolveAllRef() error {
	for _, s := range api.Schemas {
		err := s.JSONSchema.Walk(api.resolveRef)
		if err != nil {
			return err
		}
		if s.Type == "abstract" {
			continue
		}
		for parent := range s.Parents {
			parentSchema := api.schemaByID(parent)
			if parentSchema == nil {
				return fmt.Errorf("Parent schema %s not found", parent)
			}
			parentSchema.Children = append(parentSchema.Children, s)
		}
	}
	return nil
}

func (api *API) resolveAllSQL() error {
	for _, s := range api.Schemas {
		err := s.JSONSchema.resolveSQL([]string{}, "", "", &s.Columns)
		if err != nil {
			return err
		}
		s.Columns.shortenColumn()
	}
	return nil
}

func (api *API) resolveRelation(linkTo string, reference *Reference) error {
	reference.GoName = SnakeToCamel(linkTo)
	reference.Attr = mapSlice(reference.AttrSlice).JSONSchema()
	linkToSchema := api.schemaByID(linkTo)
	if linkToSchema == nil {
		return fmt.Errorf("Can't find linked schema %s", linkTo)
	}
	reference.LinkTo = linkToSchema
	ref := reference.Ref
	if ref == "" {
		return nil
	}
	file, jsonType := parseRef(ref)
	reference.RefType = jsonType
	definition, err := api.loadType(file, jsonType)
	if err != nil {
		return err
	}
	reference.Attr = definition
	err = definition.resolveGoName("")
	if err != nil {
		return err
	}
	err = definition.resolveSQL([]string{}, "", "", &reference.Columns)
	if err != nil {
		return err
	}
	return nil
}

func (api *API) resolveAllRelation() error {
	for _, s := range api.Schemas {
		for linkTo, reference := range s.References {
			if err := api.resolveRelation(linkTo, reference); err != nil {
				return err
			}
		}
		for linkTo, reference := range s.Parents {
			if err := api.resolveRelation(linkTo, reference); err != nil {
				return err
			}
		}
	}
	return nil
}

func (api *API) resolveIndex() error {
	for _, s := range api.Schemas {
		for index, property := range s.JSONSchema.OrderedProperties {
			property.Index = index + 1
		}
	}
	for _, t := range api.Types {
		for index, property := range t.OrderedProperties {
			property.Index = index + 1
		}
	}
	return nil
}

func (api *API) resolveAllGoName() error {
	for _, s := range api.Schemas {
		err := s.JSONSchema.resolveGoName(s.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (api *API) resolveExtend() error {
	for _, s := range api.Schemas {
		for _, baseSchemaID := range s.Extends {
			baseSchema := api.schemaByID(baseSchemaID)
			if baseSchema == nil {
				continue
			}
			s.JSONSchema.Update(baseSchema.JSONSchema)
		}
	}
	return nil
}

//MakeAPI load directory and generate API definitions.
func MakeAPI(dir string) (*API, error) {
	api := &API{
		Schemas:     []*Schema{},
		Definitions: []*Schema{},
		Types:       map[string]*JSONSchema{},
	}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}
		var schema Schema
		err = LoadFile(path, &schema)
		if err != nil {
			return nil
		}
		schema.FileName = strings.Replace(filepath.Base(path), ".yml", ".json", 1)
		if &schema == nil {
			return nil
		}
		schema.JSONSchema = mapSlice(schema.JSONSchemaSlice).JSONSchema()
		schema.Definitions = map[string]*JSONSchema{}
		for key, definitionSlice := range schema.DefinitionsSlice {
			schema.Definitions[key] = mapSlice(definitionSlice).JSONSchema()
		}
		schema.Path = strings.Replace(schema.ID, "_", "-", -1)
		schema.PluralPath = strings.Replace(schema.Plural, "_", "-", -1)
		if schema.ID != "" {
			api.Schemas = append(api.Schemas, &schema)
		}
		api.Definitions = append(api.Definitions, &schema)
		return nil
	})
	err = api.resolveAllRef()
	if err != nil {
		return nil, err
	}
	err = api.resolveExtend()
	if err != nil {
		return nil, err
	}
	err = api.resolveAllGoName()
	if err != nil {
		return nil, err
	}
	err = api.resolveAllSQL()
	if err != nil {
		return nil, err
	}
	err = api.resolveAllRelation()
	if err != nil {
		return nil, err
	}
	err = api.resolveIndex()
	if err != nil {
		return nil, err
	}
	return api, err
}
