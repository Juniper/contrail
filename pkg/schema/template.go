package schema

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/flosch/pongo2"
	"github.com/pkg/errors"
)

const (
	dictGetJSONSchemaByStringKeyFilter = "dict_get_JSONSchema_by_string_key"
)

// TemplateConfig contains configuration for template.
type TemplateConfig struct {
	TemplateType string `yaml:"type"`
	TemplatePath string `yaml:"template_path"`
	OutputPath   string `yaml:"output_path"`
}

// TemplateOption contains options for template.
type TemplateOption struct {
	SchemasDir        string
	TemplateConfPath  string
	SchemaOutputPath  string
	OpenapiOutputPath string
	PackagePath       string
	ProtoPackage      string
	OutputDir         string
}

func ensureDir(path string) error {
	return os.MkdirAll(filepath.Dir(path), os.ModePerm)
}

func (tc *TemplateConfig) load(base string) (*pongo2.Template, error) {
	path := filepath.Join(base, tc.TemplatePath)
	templateCode, err := common.GetContent(path)
	if err != nil {
		return nil, err
	}
	return pongo2.FromString(string(templateCode))
}

func (tc *TemplateConfig) outputPath(goName string, option *TemplateOption) string {
	path := strings.Replace(tc.OutputPath, "__resource__", common.CamelToSnake(goName), 1)
	path = strings.Replace(path, "__package__", option.PackagePath, 1)
	return path
}

// nolint: gocyclo
func (tc *TemplateConfig) apply(templateBase string, api *API, option *TemplateOption) error {
	tpl, err := tc.load(templateBase)
	if err != nil {
		return err
	}
	if err = ensureDir(tc.outputPath("", option)); err != nil {
		return err
	}
	if tc.TemplateType == "all" {
		output, err := tpl.Execute(pongo2.Context{"schemas": api.Schemas, "types": api.Types,
			"option": option,
		})
		if err != nil {
			return err
		}

		if err = writeGeneratedFile(tc.outputPath("", option), output, tc.TemplatePath); err != nil {
			return err
		}
	} else if tc.TemplateType == "type" {
		for goName, typeJSONSchema := range api.Types {
			output, err := tpl.Execute(pongo2.Context{
				"type": typeJSONSchema, "name": goName, "option": option})
			if err != nil {
				return err
			}

			if err = writeGeneratedFile(tc.outputPath(goName, option), output, tc.TemplatePath); err != nil {
				return err
			}
		}
		for _, schema := range api.Schemas {
			if schema.Type == AbstractType || schema.ID == "" {
				continue
			}

			output, err := tpl.Execute(pongo2.Context{
				"type":            schema.JSONSchema,
				"typename":        schema.TypeName,
				"name":            schema.JSONSchema.GoName,
				"references":      schema.References,
				"back_references": schema.BackReferences,
				"parents":         schema.Parents,
				"children":        schema.Children,
				"option":          option,
			})
			if err != nil {
				return err
			}

			if err = writeGeneratedFile(tc.outputPath(schema.JSONSchema.GoName, option), output, tc.TemplatePath); err != nil {
				return err
			}
		}
	} else if tc.TemplateType == "alltype" {
		var schemas []*Schema
		for typeName, typeJSONSchema := range api.Types {
			typeJSONSchema.GoName = typeName
			schemas = append(schemas, &Schema{
				JSONSchema:     typeJSONSchema,
				Children:       []*BackReference{},
				BackReferences: map[string]*BackReference{},
			})
		}
		for _, schema := range api.Schemas {
			if schema.Type == AbstractType || schema.ID == "" {
				continue
			}
			schemas = append(schemas, schema)
		}
		output, err := tpl.Execute(pongo2.Context{"schemas": schemas, "option": option})
		if err != nil {
			return err
		}

		if err = writeGeneratedFile(tc.outputPath("", option), output, tc.TemplatePath); err != nil {
			return err
		}
	} else {
		for _, schema := range api.Schemas {
			if schema.Type == AbstractType || schema.ID == "" {
				continue
			}
			output, err := tpl.Execute(pongo2.Context{"schema": schema, "types": api.Types, "option": option})
			if err != nil {
				return err
			}

			if err = writeGeneratedFile(tc.outputPath(schema.ID, option), output, tc.TemplatePath); err != nil {
				return err
			}
		}
	}
	return nil
}

// LoadTemplates loads template configurations from given path.
func LoadTemplates(path string) ([]*TemplateConfig, error) {
	var config []*TemplateConfig
	err := common.LoadFile(path, &config)
	return config, err
}

// ApplyTemplates writes files with content generated from templates.
func ApplyTemplates(api *API, templateBase string, config []*TemplateConfig, option *TemplateOption) error {
	if err := registerCustomFilters(); err != nil {
		return err
	}

	for _, templateConfig := range config {
		err := templateConfig.apply(templateBase, api, option)
		if err != nil {
			return err
		}
	}
	return nil
}

func registerCustomFilters() error {
	/* When called like this: {{ dict_value|dict_get_JSONSchema_by_string_key:key_var }}
	then: dict_value is here as `in' variable and key_var is here as `param'
	This is needed to obtain value from map with a key in variable (not as a hardcoded string)
	*/
	if !pongo2.FilterExists(dictGetJSONSchemaByStringKeyFilter) {
		if err := pongo2.RegisterFilter(
			dictGetJSONSchemaByStringKeyFilter,
			func(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
				m, _ := in.Interface().(map[string]*JSONSchema)
				return pongo2.AsValue(m[param.String()]), nil
			},
		); err != nil {
			return err
		}
	}

	return nil
}

func writeGeneratedFile(path, data, template string) error {
	if err := ioutil.WriteFile(path, []byte(generationPrefix(path, template)+data), 0644); err != nil {
		return errors.Wrapf(err, "failed to write generate file to path %q", path)
	}
	return nil
}

func generationPrefix(path, template string) string {
	prefix := "# "
	if strings.HasSuffix(path, ".go") || strings.HasSuffix(path, ".proto") {
		prefix = "// "
	} else if strings.HasSuffix(path, ".sql") {
		prefix = "-- "
	}
	return prefix + fmt.Sprintf("generated by contrailschema tool from template %s; DO NOT EDIT\n\n", template)
}
