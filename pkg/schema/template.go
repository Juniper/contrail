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

// nolint: gocyclo
func (tc *TemplateConfig) apply(templateBase string, api *API) error {
	tpl, err := tc.load(templateBase)
	if err != nil {
		return err
	}
	if err = ensureDir(tc.OutputPath); err != nil {
		return err
	}
	if tc.TemplateType == "all" {
		output, err := tpl.Execute(pongo2.Context{"schemas": api.Schemas, "types": api.Types})
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(tc.OutputPath, []byte(output), 0644)
		if err != nil {
			return err
		}
	} else if tc.TemplateType == "type" {
		for goName, typeJSONSchema := range api.Types {
			output, err := tpl.Execute(pongo2.Context{"type": typeJSONSchema, "name": goName})
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(
				strings.Replace(tc.OutputPath, "__resource__", common.CamelToSnake(goName), 1),
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
				"type":            schema.JSONSchema,
				"typename":        schema.TypeName,
				"name":            schema.JSONSchema.GoName,
				"references":      schema.References,
				"back_references": schema.BackReferences,
				"parents":         schema.Parents,
				"children":        schema.Children,
			})
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(
				strings.Replace(tc.OutputPath, "__resource__", common.CamelToSnake(schema.JSONSchema.GoName), 1),
				[]byte(output), 0644)
			if err != nil {
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
		output, err := tpl.Execute(pongo2.Context{"schemas": schemas})
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(tc.OutputPath, []byte(output), 0644)
		if err != nil {
			return err
		}
	} else {
		for _, schema := range api.Schemas {
			if schema.Type == AbstractType || schema.ID == "" {
				continue
			}
			output, err := tpl.Execute(pongo2.Context{"schema": schema, "types": api.Types})
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(
				strings.Replace(tc.OutputPath, "__resource__", schema.ID, 1),
				[]byte(output),
				0644,
			)
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
func ApplyTemplates(api *API, templateBase string, config []*TemplateConfig) error {

	// Make custom filters available for everyone

	/* When called like this: {{ dict_value|dict_get_JSONSchema_by_string_key:key_var }}
	then: dict_value is here as `in' variable and key_var is here as `param'
	This is needed to obtain value from map with a key in variable (not as a hardcoded string)
	*/
	err := pongo2.RegisterFilter("dict_get_JSONSchema_by_string_key",
		func(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
			m := in.Interface().(map[string]*JSONSchema)
			return pongo2.AsValue(m[param.String()]), nil
		})
	if err != nil {
		return err
	}

	for _, templateConfig := range config {
		err := templateConfig.apply(templateBase, api)
		if err != nil {
			return err
		}
	}
	return nil
}
