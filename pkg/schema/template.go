package schema

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/flosch/pongo2"
)

//TemplateConfig is configuration option for templates.
type TemplateConfig struct {
	TemplateType string `yaml:"type"`
	TemplatePath string `yaml:"template_path"`
	OutputPath   string `yaml:"output_path"`
}

//TemplateOption for template
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

func (config *TemplateConfig) load(base string) (*pongo2.Template, error) {
	path := filepath.Join(base, config.TemplatePath)
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
		err = ioutil.WriteFile(tc.outputPath("", option), []byte(output), 0644)
		if err != nil {
			return err
		}
	} else if tc.TemplateType == "type" {
		for goName, typeJSONSchema := range api.Types {
			output, err := tpl.Execute(pongo2.Context{
				"type": typeJSONSchema, "name": goName, "option": option})
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(
				tc.outputPath(goName, option),
				[]byte(output), 0644)
			if err != nil {
				return err
			}
		}
		for _, schema := range api.Schemas {
			if schema.Type == AbstractType || schema.ID == "" {
				continue
			}
			output, err := tpl.Execute(pongo2.Context{
				"type":       schema.JSONSchema,
				"typename":   schema.TypeName,
				"name":       schema.JSONSchema.GoName,
				"references": schema.References,
				"parents":    schema.Parents,
				"children":   schema.Children,
			})
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(
				tc.outputPath(schema.JSONSchema.GoName, option),
				[]byte(output), 0644)
			if err != nil {
				return err
			}
		}
	} else if tc.TemplateType == "alltype" {
		types := []interface{}{}
		for goName, typeDef := range api.Types {
			types = append(types,
				pongo2.Context{"type": typeDef, "name": goName})
		}
		for _, schema := range api.Schemas {
			if schema.Type == AbstractType || schema.ID == "" {
				continue
			}
			goName := schema.JSONSchema.GoName
			typeDef := schema.JSONSchema
			typeName := schema.TypeName
			types = append(types,
				pongo2.Context{"type": typeDef, "typename": typeName, "name": goName,
					"references": schema.References, "parents": schema.Parents, "children": schema.Children})
		}
		output, err := tpl.Execute(pongo2.Context{"schemas": api.Schemas, "option": option})
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(tc.outputPath("", option), []byte(output), 0644)
		if err != nil {
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
			err = ioutil.WriteFile(
				tc.outputPath(schema.ID, option),
				[]byte(output), 0644)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//LoadTemplates loads templates from config path.
func LoadTemplates(path string) ([]*TemplateConfig, error) {
	var config []*TemplateConfig
	err := common.LoadFile(path, &config)
	return config, err
}

//ApplyTemplates applies templates and generate codes.
func ApplyTemplates(api *API, templateBase string, config []*TemplateConfig, option *TemplateOption) error {

	// Make custom filters available for everyone

	/* When called like this: {{ dict_value|dict_get_JSONSchema_by_string_key:key_var }}
	then: dict_value is here as `in' variable and key_var is here as `param'
	This is needed to obtain value from map with a key in variable (not as a hardcoded string)
	*/
	pongo2.RegisterFilter("dict_get_JSONSchema_by_string_key",
		func(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
			m, _ := in.Interface().(map[string]*JSONSchema)
			return pongo2.AsValue(m[param.String()]), nil
		})

	for _, templateConfig := range config {
		err := templateConfig.apply(templateBase, api, option)
		if err != nil {
			return err
		}
	}
	return nil
}
