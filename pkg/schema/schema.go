package schema

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/common"
)

//Version is version for schema format.
var Version = "1.0"

const (
	schemaStartIndex    = 10
	schemaIndexOffset   = 3
	propertyIndexOffset = 1000
	stringType          = "string"
)

// Available type values.
const (
	AbstractType = "abstract"
	ObjectType   = "object"
	IntegerType  = "integer"
	ArrayType    = "array"
	BooleanType  = "boolean"
	NumberType   = "number"
	StringType   = stringType
	Base64Type   = "base64"
)

// Available Go type values.
const (
	IntGoType   = "int64"
	FloatGoType = "float64"
)

// Available Proto type values
const (
	IntProtoType   = IntGoType
	FloatProtoType = "float"
)

const (
	maxColumnLen = 55
	//RefPrefix is table column name prefix for reference
	RefPrefix = "ref"
	//ParentPrefix is table column name prefix for parent
	ParentPrefix = "parent"
	configRoot   = "config_root"
	optional     = "optional"
)

var sqlTypeMap = map[string]string{
	ObjectType:  "json",
	IntegerType: "bigint",
	ArrayType:   "json",
	BooleanType: "bool",
	NumberType:  "float",
	StringType:  "varchar(255)",
	Base64Type:  "varchar(255)",
}

var sqlBindMap = map[string]string{
	ObjectType:  "json",
	IntegerType: "int",
	ArrayType:   "json",
	BooleanType: "bool",
	NumberType:  "float",
	StringType:  stringType,
	Base64Type:  stringType,
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
	GetPath      string
	UpdatePath   string
	Bind         string
	Column       string
	ParentColumn []string
	Name         string
	JSONSchema   *JSONSchema
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
	FileName              string                    `yaml:"-" json:"-"`
	ID                    string                    `yaml:"id" json:"id,omitempty"`
	Plural                string                    `yaml:"plural" json:"plural,omitempty"`
	Type                  string                    `yaml:"type" json:"type,omitempty"`
	Title                 string                    `yaml:"title" json:"title,omitempty"`
	Description           string                    `yaml:"description" json:"description,omitempty"`
	Parents               map[string]*Reference     `yaml:"-" json:"parents,omitempty"`
	ParentsSlice          yaml.MapSlice             `yaml:"parents" json:"-"`
	References            map[string]*Reference     `yaml:"-" json:"references,omitempty"`
	BackReferences        map[string]*BackReference `yaml:"-" json:"back_references,omitempty"`
	ReferencesSlice       yaml.MapSlice             `yaml:"references" json:"-"`
	Prefix                string                    `yaml:"prefix" json:"prefix,omitempty"`
	JSONSchema            *JSONSchema               `yaml:"-" json:"schema,omitempty"`
	JSONSchemaSlice       yaml.MapSlice             `yaml:"schema" json:"-"`
	Definitions           map[string]*JSONSchema    `yaml:"-" json:"-"`
	DefinitionsSlice      map[string]yaml.MapSlice  `yaml:"definitions" json:"-"`
	Extends               []string                  `yaml:"extends" json:"extends,omitempty"`
	Columns               ColumnConfigs             `yaml:"-" json:"-"`
	TypeName              string                    `yaml:"-" json:"-"`
	Path                  string                    `yaml:"-" json:"-"`
	PluralPath            string                    `yaml:"-" json:"-"`
	Children              []*BackReference          `yaml:"-" json:"-"`
	Index                 int                       `yaml:"-" json:"-"`
	ParentOptional        bool                      `yaml:"-" json:"-"`
	IsConfigRootInParents bool                      `yaml:"-" json:"-"`
	HasParents            bool                      `yaml:"-" json:"-"`
	DefaultParent         *Reference                `yaml:"-" json:"-"`
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
	Presence          string                 `yaml:"presence" json:"presence,omitempty"`
	Type              string                 `yaml:"type" json:"type,omitempty"`
	Permission        []string               `yaml:"permission" json:"permission,omitempty"`
	Properties        map[string]*JSONSchema `yaml:"properties" json:"properties,omitempty"`
	PropertiesOrder   []string               `yaml:"-" json:"propertiesOrder,omitempty"`
	OrderedProperties []*JSONSchema          `yaml:"-" json:"-"`
	Enum              []string               `yaml:"enum" json:"enum,omitempty"`
	Minimum           interface{}            `yaml:"minimum" json:"minimum,omitempty"`
	Maximum           interface{}            `yaml:"maximum" json:"maximum,omitempty"`
	Ref               string                 `yaml:"$ref" json:"-"`
	Items             *JSONSchema            `yaml:"items" json:"items,omitempty"`
	GoName            string                 `yaml:"-" json:"-"`
	GoType            string                 `yaml:"-" json:"-"`
	ProtoType         string                 `yaml:"-" json:"-"`
	Required          []string               `yaml:"required" json:"-"`
	GoPremitive       bool                   `yaml:"-" json:"-"`
	Format            string                 `yaml:"format" json:"format,omitempty"`

	// Properties relevant for collection types (with CollectionType == "map" or "list"):
	CollectionType string `yaml:"collectionType" json:"collectionType,omitempty"`
	// MapKey is name of MapKeyProperty.
	MapKey         string      `yaml:"mapKey" json:"mapKey,omitempty"`
	MapKeyProperty *JSONSchema `yaml:"mapKeyProperty" json:"mapKeyProperty,omitempty"`
}

//String makes string format for json schema.
func (s *JSONSchema) String() string {
	data, err := json.Marshal(s)
	if err != nil {
		log.WithError(err).Debug("Could not stringify JSONSchema")
	}
	return string(data)
}

// IsInt returns true if schema is of int type.
func (s *JSONSchema) IsInt() bool {
	return s.GoType == IntGoType
}

// IsFloat returns true if schema is of float type.
func (s *JSONSchema) IsFloat() bool {
	return s.GoType == FloatGoType
}

// HasNumberFields returns true if JSONSchema has any number fields (int or float).
func (s *JSONSchema) HasNumberFields() bool {
	for _, property := range s.Properties {
		if property.GoType == IntGoType || property.GoType == FloatGoType {
			return true
		}
	}
	return false
}

//Reference object represents many to many relationships between resources.
type Reference struct {
	GoName      string        `yaml:"-" json:"-"`
	Table       string        `yaml:"-" json:"-"`
	Index       int           `yaml:"-" json:"-"`
	Description string        `yaml:"description" json:"description,omitempty"`
	Operations  string        `yaml:"operations" json:"operations,omitempty"`
	Presence    string        `yaml:"presence" json:"presence,omitempty"`
	RefType     string        `yaml:"-" json:"-"`
	Columns     ColumnConfigs `yaml:"-" json:"-"`
	Attr        *JSONSchema   `yaml:"-" json:"attr"`
	AttrSlice   yaml.MapSlice `yaml:"attr" json:"-"`
	LinkTo      *Schema       `yaml:"-" json:"-"`
	Ref         string        `yaml:"$ref" json:"$ref,omitempty"`
}

//BackReference for representing backward references.
type BackReference struct {
	Index       int     `yaml:"-" json:"-"`
	Description string  `yaml:"description" json:"description,omitempty"`
	LinkTo      *Schema `yaml:"-" json:"-"`
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

// Copy copies a json schema.
//
// Note that non pointer receiver is used to copy the object.
func (s JSONSchema) Copy() *JSONSchema {
	properties := map[string]*JSONSchema{}
	for name, property := range s.Properties {
		properties[name] = property.Copy()
	}
	s.Properties = properties

	if s.Items != nil {
		s.Items = s.Items.Copy()
	}
	return &s
}

//Update merges two JSONSchema
// nolint: gocyclo
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
	if s.Format == "" {
		s.Format = s2.Format
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

func (s *JSONSchema) resolveSQL(
	parentColumn []string, columnName string,
	goPath string, getPath string, updatePath string, columns *ColumnConfigs) error {
	if s == nil {
		return nil
	}
	if len(s.Properties) == 0 || s.Type == ArrayType {
		if s.SQL == "" {
			s.SQL = sqlTypeMap[s.Type]
		}
		bind := ""
		if s.GoType != "" {
			bind = sqlBindMap[s.Type]
		}

		*columns = append(*columns, &ColumnConfig{
			Path:         goPath,
			GetPath:      getPath,
			UpdatePath:   updatePath,
			Bind:         bind,
			Column:       strings.ToLower(columnName),
			ParentColumn: parentColumn,
			Name:         columnName,
			JSONSchema:   s,
		})
		return nil
	}
	for name, property := range s.Properties {
		nextParentColumn := make([]string, len(parentColumn))
		copy(nextParentColumn, parentColumn)
		nextParentColumn = append(nextParentColumn, columnName)

		newUpdatePath := name
		if updatePath != "" {
			newUpdatePath = updatePath + "." + name
		}
		err := property.resolveSQL(nextParentColumn,
			name, goPath+"."+property.GoName,
			getPath+".Get"+property.GoName+"()", newUpdatePath, columns)
		if err != nil {
			return err
		}
	}
	return nil
}

// nolint: gocyclo
func (s *JSONSchema) resolveGoName(name string) error {
	if s == nil {
		return nil
	}
	s.GoName = common.SnakeToCamel(name)
	if s.GoName == "Size" {
		s.GoName = "Size_"
	}

	protoType := ""
	s.GoPremitive = true
	goType := ""
	switch s.Type {
	case IntegerType:
		goType = IntGoType
		protoType = IntProtoType
	case NumberType:
		goType = FloatGoType
		protoType = FloatProtoType
	case StringType:
		goType = stringType
		protoType = stringType
	case Base64Type:
		goType = "base64"
		protoType = stringType
	case BooleanType:
		goType = "bool"
		protoType = "bool"
	case ObjectType:
		goType = s.getRefType()
		if s.Properties == nil {
			goType = "map[string]interface{}"
		}

		if goType != "" {
			protoType = goType
		}
		if s.Properties == nil {
			protoType = "bytes"
		}
	case ArrayType:
		err := s.Items.resolveGoName(name)
		if err != nil {
			return err
		}
		if s.Items == nil {
			goType = "[]string"
			protoType = "repeated string"
		} else {
			if s.Items.Type == IntegerType || s.Items.Type == NumberType || s.Items.Type == BooleanType ||
				s.Items.Type == StringType {
				goType = "[]" + s.Items.GoType
			} else {
				goType = "[]*" + s.Items.GoType
			}
			protoType = "repeated " + s.Items.ProtoType
		}
	}

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

//SchemaByID return schema by ID.
func (api *API) SchemaByID(id string) *Schema {
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
		for _, d := range api.Definitions {
			log.Info(d.FileName)
		}
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
	if schema.Type == ArrayType {
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
	}
	return nil
}

func (api *API) resolveAllSQL() error {
	for _, s := range api.Schemas {
		err := s.JSONSchema.resolveSQL([]string{}, "", "", "", "", &s.Columns)
		if err != nil {
			return err
		}
		s.Columns.shortenColumn()
	}
	return nil
}

func (api *API) resolveRelation(linkToSchema *Schema, reference *Reference) error {
	linkTo := linkToSchema.ID
	reference.GoName = common.SnakeToCamel(linkTo)
	reference.Attr = mapSlice(reference.AttrSlice).JSONSchema()

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
	return definition.resolveSQL([]string{}, "", "", "", "", &reference.Columns)
}

func makeShort(id string) string {
	id = strings.Replace(id, "virtual", "v", -1)
	id = strings.Replace(id, "network", "net", -1)
	id = strings.Replace(id, "interface", "i", -1)
	id = strings.Replace(id, "machine", "m", -1)
	id = strings.Replace(id, "router", "r", -1)
	id = strings.Replace(id, "structured_syslog", "log", -1)
	return id
}

//ReferenceTableName make reference table name.
func ReferenceTableName(prefix, id, linkTo string) string {
	table := prefix + "_" + id + "_" + linkTo
	if len(table) < maxColumnLen {
		return strings.ToLower(table)
	}
	return strings.ToLower(prefix + "_" + makeShort(id) + "_" + makeShort(linkTo))
}

// ChildColumnName makes child column name.
func ChildColumnName(childSchemaID, schemaID string) string {
	return strings.ToLower("child_" + childSchemaID + "_" + schemaID)
}

// BackRefColumnName makes back reference column name.
func BackRefColumnName(fromID, toID string) string {
	return strings.ToLower("backref_" + fromID + "_" + toID)
}

// nolint: gocyclo
func (api *API) resolveAllRelation() error {
	for _, s := range api.Schemas {
		if s.Type == AbstractType {
			continue
		}
		s.References = map[string]*Reference{}

		s.Parents = map[string]*Reference{}
		for _, m := range mapSlice(s.ReferencesSlice) {
			linkTo := m.Key.(string)
			referenceMap := mapSlice(m.Value.(yaml.MapSlice))
			reference := referenceMap.Reference()
			s.References[linkTo] = reference
			linkToSchema := api.SchemaByID(linkTo)
			if linkToSchema == nil {
				return fmt.Errorf("Can't find linked schema %s", linkTo)
			}
			linkToSchema.BackReferences[s.ID] = &BackReference{
				LinkTo:      s,
				Description: reference.Description,
			}
			if err := api.resolveRelation(linkToSchema, reference); err != nil {
				return err
			}
			reference.Table = ReferenceTableName(RefPrefix, s.ID, linkTo)
		}
		s.IsConfigRootInParents = false
		for _, m := range mapSlice(s.ParentsSlice) {
			linkTo := m.Key.(string)
			if linkTo == configRoot {
				s.IsConfigRootInParents = true
				s.ParentOptional = true
				continue
			}
			referenceMap := mapSlice(m.Value.(yaml.MapSlice))
			reference := referenceMap.Reference()
			s.DefaultParent = reference
			if reference.Presence == optional {
				s.ParentOptional = true
			}
			s.Parents[linkTo] = reference
			reference.Table = ReferenceTableName(ParentPrefix, s.ID, linkTo)
			parentSchema := api.SchemaByID(linkTo)
			if parentSchema == nil {
				return fmt.Errorf("Parent schema %s not found", linkTo)
			}
			if err := api.resolveRelation(parentSchema, reference); err != nil {
				return err
			}
			parentSchema.Children = append(parentSchema.Children, &BackReference{LinkTo: s, Description: reference.Description})
		}
		s.HasParents = len(s.Parents) > 0
	}
	return nil
}

func (api *API) resolveIndex() error {
	schemaIndex := schemaStartIndex
	for _, s := range api.Schemas {
		if s.Type == AbstractType {
			continue
		}
		s.Index = schemaIndex
		schemaIndex += schemaIndexOffset
		index := 1
		for _, property := range s.JSONSchema.OrderedProperties {
			property.Index = index
			index++
		}
		index = index + propertyIndexOffset
		for _, key := range mapSlice(s.ReferencesSlice).keys() {
			reference := s.References[key]
			reference.Index = index
			index++
		}
		index = index + propertyIndexOffset
		for _, backReference := range s.Children {
			backReference.Index = index
			index++
		}
		index = index + propertyIndexOffset
		for _, backReference := range s.BackReferences {
			backReference.Index = index
			index++
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
			baseSchema := api.SchemaByID(baseSchemaID)
			if baseSchema == nil {
				continue
			}
			s.JSONSchema.Update(baseSchema.JSONSchema)
		}
	}
	return nil
}

func (api *API) resolveCollectionTypes() error {
	for _, s := range api.Schemas {
		for propertyName, property := range s.JSONSchema.Properties {
			if property.CollectionType != "" {
				propertyType := api.Types[property.ProtoType]

				if err := checkCollectionTypes(property, propertyType); err != nil {
					return err
				}
				propertyType.CollectionType = property.CollectionType

				if propertyType.CollectionType == "map" {
					if err := resolveMapCollectionType(property, propertyType); err != nil {
						return errors.Wrapf(err, "invalid %q property of %q schema", propertyName, s.ID)
					}
				}
			}
		}
	}
	return nil
}

func checkCollectionTypes(property, propertyType *JSONSchema) error {
	if propertyType.CollectionType != "" && propertyType.CollectionType != property.CollectionType {
		return errors.Errorf(
			"type %v is used as multiple collection types - %v and %v",
			property.ProtoType, property.CollectionType, propertyType.CollectionType,
		)
	}
	return nil
}

func resolveMapCollectionType(property, propertyType *JSONSchema) error {
	if property.MapKey == "" {
		return errors.New("empty mapKey field")
	}

	propertyType.MapKeyProperty = propertyType.OrderedProperties[0].Items.Properties[property.MapKey]
	return nil
}

//MakeAPI load directory and generate API definitions.
// nolint: gocyclo
func MakeAPI(dirs []string) (*API, error) {
	api := &API{
		Schemas:     []*Schema{},
		Definitions: []*Schema{},
		Types:       map[string]*JSONSchema{},
	}
	for _, dir := range dirs {
		err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
			if f.IsDir() {
				return nil
			}
			var schema Schema
			err = common.LoadFile(path, &schema)
			if err != nil {
				log.Warn(fmt.Sprintf("[%s] %s", path, err))
				return nil
			}
			if &schema == nil {
				return nil
			}
			schema.FileName = strings.Replace(filepath.Base(path), ".yml", ".json", 1)
			schema.JSONSchema = mapSlice(schema.JSONSchemaSlice).JSONSchema()
			schema.Definitions = map[string]*JSONSchema{}
			for key, definitionSlice := range schema.DefinitionsSlice {
				schema.Definitions[key] = mapSlice(definitionSlice).JSONSchema()
			}
			schema.TypeName = strings.Replace(schema.ID, "_", "-", -1)
			schema.Path = schema.TypeName
			schema.PluralPath = strings.Replace(schema.Plural, "_", "-", -1)
			schema.BackReferences = map[string]*BackReference{}
			if schema.ID != "" {
				api.Schemas = append(api.Schemas, &schema)
			}
			if len(schema.Definitions) > 0 {
				api.Definitions = append(api.Definitions, &schema)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	err := api.resolveAllRef()
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
	err = api.resolveCollectionTypes()
	if err != nil {
		return nil, err
	}
	return api, err
}
