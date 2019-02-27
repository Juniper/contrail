package schema

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/format"
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
	UintType     = "uint64"
	ArrayType    = "array"
	BooleanType  = "boolean"
	NumberType   = "number"
	StringType   = stringType
	Base64Type   = "base64"
)

// Available Go type values.
const (
	IntGoType   = "int64"
	UintGoType  = "uint64"
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
	ParentPrefix    = "parent"
	configRoot      = "config_root"
	optional        = "optional"
	serviceProperty = "service"
)

const (
	definitionsInFile = "definitions"
	schemasInFile     = "schemas"
)

var sqlTypeMap = map[string]string{
	ObjectType:  "json",
	IntegerType: "bigint",
	UintType:    "numeric(21)",
	ArrayType:   "json",
	BooleanType: "bool",
	NumberType:  "float",
	StringType:  "varchar(255)",
	Base64Type:  "varchar(255)",
}

var sqlBindMap = map[string]string{
	ObjectType:  "json",
	IntegerType: "int",
	UintType:    "uint64",
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
	Timestamp   time.Time
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
	Table                 string                    `yaml:"table" json:"table,omitempty"`
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
	JSONTag           string                 `yaml:"json_tag" json:"-"`
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
	GoType            string                 `yaml:"go_type" json:"go_type"`
	ProtoType         string                 `yaml:"proto_type" json:"proto_type"`
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
		logrus.WithError(err).Debug("Could not stringify JSONSchema")
	}
	return string(data)
}

// IsInt returns true if schema is of int type.
func (s *JSONSchema) IsInt() bool {
	return s.GoType == IntGoType
}

// IsUint returns true if schema is of int type.
func (s *JSONSchema) IsUint() bool {
	return s.GoType == UintGoType
}

// IsFloat returns true if schema is of float type.
func (s *JSONSchema) IsFloat() bool {
	return s.GoType == FloatGoType
}

// HasNumberFields returns true if JSONSchema has any number fields (int or float).
func (s *JSONSchema) HasNumberFields() bool {
	for _, property := range s.Properties {
		if property.IsInt() || property.IsFloat() || property.IsUint() {
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
	Derived     bool          `yaml:"derived" json:"derived,omitempty"`
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

func parseRef(ref string) (file, section, goType string) {
	if ref == "" {
		return "", "", ""
	}
	refs := strings.Split(ref, "#")
	types := strings.Split(ref, "/")
	return refs[0], types[1], types[len(types)-1]
}

func (s *JSONSchema) getRefType() string {
	_, _, goType := parseRef(s.Ref)
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
	if s.JSONTag == "" {
		s.JSONTag = s2.JSONTag
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
	if s.GoType == "" {
		s.GoType = s2.GoType
	}
	if s.ProtoType == "" {
		s.ProtoType = s2.ProtoType
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
	if s.Items == nil {
		s.Items = s2.Items
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

func (s *JSONSchema) resolveSQLForArray(
	parentColumn []string,
	columnName string,
	goPath string,
	getPath string,
	updatePath string,
	columns *ColumnConfigs,
) bool {
	if len(s.Properties) != 0 && s.Type != ArrayType {
		return false
	}

	if s.SQL == "" {
		s.SQL = sqlTypeMap[s.Type]
	}
	var bind string
	if s.GoType != "" {
		if s.IsUint() {
			bind = sqlBindMap[UintType]
		} else {
			bind = sqlBindMap[s.Type]
		}
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

	return true
}

func (s *JSONSchema) resolveSQL(
	parentColumn []string,
	columnName string,
	goPath string,
	getPath string,
	updatePath string,
	columns *ColumnConfigs,
) error {
	if s == nil || s.Presence == serviceProperty {
		return nil
	}

	if s.resolveSQLForArray(parentColumn, columnName, goPath, getPath, updatePath, columns) {
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
	s.GoName = format.SnakeToCamel(name)
	if s.GoName == "Size" {
		s.GoName = "Size_"
	}
	if s.ProtoType == "" {
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
				logrus.Errorf("Got <nil> Items for array in schema '%v': %+#v", name, s)
				goType = "[]string"
				protoType = "repeated string"
			} else {
				if s.Items.Type == IntegerType || s.Items.Type == NumberType || s.Items.Type == BooleanType ||
					s.Items.Type == StringType || s.Items.Type == ArrayType {
					goType = "[]" + s.Items.GoType
				} else {
					goType = "[]*" + s.Items.GoType
				}
				protoType = "repeated " + s.Items.ProtoType
			}
		}
		s.GoType = goType
		s.ProtoType = protoType
	}
	for pname, property := range s.Properties {
		err := property.resolveGoName(pname)
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

func (api *API) readDefinitionFromDefinitions(schemaFile, typeName string) (*JSONSchema, error) {
	definitions := api.definitionByFileName(schemaFile)
	if definitions == nil {
		logrus.Info("definitions read from following files:")
		for _, d := range api.Definitions {
			logrus.Info(d.FileName)
		}
		return nil, fmt.Errorf("can't find file '%s' (with type %s)", schemaFile, typeName)
	}
	definition, ok := definitions.Definitions[typeName]
	if !ok {
		return nil, fmt.Errorf("%s isn't defined in %s", typeName, schemaFile)
	}
	return definition, nil
}

func (api *API) readDefinitionFromSchemas(typeName string) (*JSONSchema, error) {
	schema := api.SchemaByID(typeName)
	if schema == nil {
		return nil, fmt.Errorf("can find schema with id: %v", typeName)
	}
	return schema.JSONSchema, nil
}

func (api *API) readDefinition(schemaFile, section, typeName string) (*JSONSchema, error) {
	switch section {
	case definitionsInFile:
		return api.readDefinitionFromDefinitions(schemaFile, typeName)
	case schemasInFile:
		return api.readDefinitionFromSchemas(typeName)
	}
	return nil, fmt.Errorf("section '%v' not handled for reading definitions", section)
}

func (api *API) loadType(schemaFile, section, typeName string) (*JSONSchema, error) {
	if definition, ok := api.Types[typeName]; ok {
		return definition, nil
	}
	definition, err := api.readDefinition(schemaFile, section, typeName)
	if err != nil {
		return nil, errors.Wrapf(err, "reading definiton from %v", section)
	}
	err = definition.Walk(api.resolveRef)
	if err != nil {
		return nil, err
	}
	err = definition.resolveGoName("")
	if err != nil {
		return nil, err
	}
	err = definition.Walk(api.resolveJSONTag)
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
		return errors.Wrapf(err, "resolve ref required by %v (ref: %v)", schema.ID, schema.Ref)
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
	reference.GoName = format.SnakeToCamel(linkTo)
	reference.Attr = mapSlice(reference.AttrSlice).JSONSchema()

	reference.LinkTo = linkToSchema
	ref := reference.Ref
	if ref == "" {
		return nil
	}
	file, section, jsonType := parseRef(ref)
	reference.RefType = jsonType
	definition, err := api.loadType(file, section, jsonType)
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
			linkTo := m.Key.(string)                          //nolint: errcheck
			referenceMap := mapSlice(m.Value.(yaml.MapSlice)) //nolint: errcheck
			reference := referenceMap.Reference()
			s.References[linkTo] = reference
			linkToSchema := api.SchemaByID(linkTo)
			if linkToSchema == nil {
				return fmt.Errorf("missing linked schema '%s' for reference '%v' in schema %v [%v]",
					linkTo, linkTo, s.ID, s.FileName)
			}
			linkToSchema.BackReferences[s.ID] = &BackReference{
				LinkTo:      s,
				Description: reference.Description,
			}
			if err := api.resolveRelation(linkToSchema, reference); err != nil {
				return err
			}
			reference.Table = ReferenceTableName(RefPrefix, s.Table, linkTo)
		}
		s.IsConfigRootInParents = false
		for _, m := range mapSlice(s.ParentsSlice) {
			linkTo := m.Key.(string) //nolint: errcheck
			if linkTo == configRoot {
				s.IsConfigRootInParents = true
				s.ParentOptional = true
				continue
			}
			referenceMap := mapSlice(m.Value.(yaml.MapSlice)) //nolint: errcheck
			reference := referenceMap.Reference()
			s.DefaultParent = reference
			if reference.Presence == optional {
				s.ParentOptional = true
			}
			s.Parents[linkTo] = reference
			reference.Table = ReferenceTableName(ParentPrefix, s.Table, linkTo)
			parentSchema := api.SchemaByID(linkTo)
			if parentSchema == nil {
				return fmt.Errorf("parent schema %s not found", linkTo)
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
			s.ReferencesSlice = append(s.ReferencesSlice, baseSchema.ReferencesSlice...)
		}
	}
	return nil
}

func (api *API) resolveAllJSONTag() error {
	for _, schema := range api.Schemas {
		err := schema.JSONSchema.Walk(api.resolveJSONTag)
		if err != nil {
			return err
		}
	}
	return nil
}

func (api *API) resolveJSONTag(schema *JSONSchema) error {
	if schema == nil {
		return nil
	}
	if len(schema.JSONTag) == 0 {
		schema.JSONTag = schema.ID
	}

	for _, prop := range schema.Properties {
		if len(prop.JSONTag) == 0 {
			prop.JSONTag = prop.ID
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

func (api *API) loadSchemaFromPath(path string) (*Schema, error) {
	var schema Schema
	err := fileutil.LoadFile(path, &schema)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load file %s", path)
	}
	info, err := os.Stat(path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to stat file %s", path)
	}
	if info.ModTime().After(api.Timestamp) {
		api.Timestamp = info.ModTime()
	}
	logrus.Printf("Loading schema from %v - %v", path, schema.ID)
	return &schema, nil
}

func (api *API) readOverrides(dir string) (*Schema, error) {
	var schemaOverrides = &Schema{DefinitionsSlice: map[string]yaml.MapSlice{}}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if (f != nil && f.IsDir()) || err != nil {
			if os.IsNotExist(err) {
				return nil // Silence not-exist as if overrides dir not exist it is not an error
			}
			return err
		}
		// This is as a Warning because overrides fixes schema problems that should be fixed in upstream schema definition
		logrus.Warnf("Reading overrides from %v file", path)
		schema, err := api.loadSchemaFromPath(path)
		if err == nil && len(schema.DefinitionsSlice) > 0 {
			for key, def := range schema.DefinitionsSlice {
				schemaOverrides.DefinitionsSlice[key] = def
			}
		}
		return err
	})
	return schemaOverrides, err
}

func processSchema(schema, overrides *Schema, api *API) error {
	schema.JSONSchema = mapSlice(schema.JSONSchemaSlice).JSONSchema()
	schema.Definitions = map[string]*JSONSchema{}
	for key, definitionSlice := range schema.DefinitionsSlice {
		schema.Definitions[key] = mapSlice(definitionSlice).JSONSchema()
		overDef, ok := overrides.DefinitionsSlice[key]
		if ok {
			schema.Definitions[key] = mapSlice(overDef).JSONSchema()
		}
	}
	schema.TypeName = strings.Replace(schema.ID, "_", "-", -1)
	if schema.Table == "" {
		schema.Table = strings.ToLower(schema.ID)
	}
	schema.Path = schema.TypeName
	schema.PluralPath = strings.Replace(schema.Plural, "_", "-", -1)
	schema.BackReferences = map[string]*BackReference{}
	if schema.ID != "" {
		api.Schemas = append(api.Schemas, schema)
	}
	if len(schema.Definitions) > 0 {
		api.Definitions = append(api.Definitions, schema)
	}
	return nil
}

func walkSchemaFile(overridePath string, overrides *Schema, api *API, path string, f os.FileInfo, err error) error {
	// Don't walk over override schema files
	if path == overridePath && f.IsDir() {
		return filepath.SkipDir
	}
	if f == nil || f.IsDir() || err != nil {
		return err
	}
	schema, err := api.loadSchemaFromPath(path)
	if err != nil {
		return err
	}
	if schema == nil {
		return nil
	}
	r := strings.NewReplacer(".yml", ".json", ".yaml", ".json")
	schema.FileName = r.Replace(filepath.Base(path))
	return processSchema(schema, overrides, api)
}

func (api *API) process() error {
	err := api.resolveAllRef()
	if err != nil {
		return err
	}
	err = api.resolveExtend()
	if err != nil {
		return err
	}
	err = api.resolveAllGoName()
	if err != nil {
		return err
	}
	err = api.resolveAllSQL()
	if err != nil {
		return err
	}
	err = api.resolveAllRelation()
	if err != nil {
		return err
	}
	err = api.resolveIndex()
	if err != nil {
		return err
	}
	err = api.resolveCollectionTypes()
	if err != nil {
		return err
	}
	err = api.resolveAllJSONTag()
	return err
}

func (api *API) loadOverrides(dir string) (*Schema, error) {
	overrides, err := api.readOverrides(dir)
	if overrides == nil {
		overrides = &Schema{}
	}
	if err != nil {
		return overrides, err
	}
	return overrides, nil
}

// MakeAPI load directory and generate API definitions.
func MakeAPI(dirs []string, overrideSubdir string) (*API, error) {
	api := &API{
		Schemas:     []*Schema{},
		Definitions: []*Schema{},
		Types:       map[string]*JSONSchema{},
	}
	logrus.Printf("Making API from schema dirs: %v", dirs)
	for _, dir := range dirs {
		overrides := &Schema{}
		overridePath := ""
		if overrideSubdir != "" {
			overridePath = dir + string(os.PathSeparator) + overrideSubdir
			var err error
			overrides, err = api.loadOverrides(overridePath)
			if err != nil {
				return api, err
			}
		}
		err := filepath.Walk(dir, func(p string, f os.FileInfo, e error) error {
			return walkSchemaFile(overridePath, overrides, api, p, f, e)
		})
		if err != nil {
			return nil, err
		}
	}
	err := api.process()
	return api, err
}
