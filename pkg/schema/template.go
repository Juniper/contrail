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

func (config *TemplateConfig) load(base string) (*pongo2.Template, error) {
	path := filepath.Join(base, config.TemplatePath)
	templateCode, err := common.GetContent(path)
	if err != nil {
		return nil, err
	}
	return pongo2.FromString(string(templateCode))
}

func (config *TemplateConfig) apply(templateBase string, api *API) error {
	tpl, err := config.load(templateBase)
	if err != nil {
		return err
	}
	if err = ensureDir(config.OutputPath); err != nil {
		return err
	}
	if config.TemplateType == "all" {
		output, err :=
			tpl.Execute(pongo2.Context{"schemas": api.Schemas, "types": api.Types})
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(config.OutputPath, []byte(output), 0644)
		if err != nil {
			return err
		}
	} else if config.TemplateType == "type" {
		for goName, typeDef := range api.Types {
			output, err :=
				tpl.Execute(pongo2.Context{"type": typeDef, "name": goName})
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(
				strings.Replace(config.OutputPath, "__resource__", common.CamelToSnake(goName), 1),
				[]byte(output), 0644)
			if err != nil {
				return err
			}
		}
		for _, schema := range api.Schemas {
			if schema.Type == AbstractType || schema.ID == "" {
				continue
			}
			goName := schema.JSONSchema.GoName
			typeDef := schema.JSONSchema
			typeName := schema.TypeName
			output, err :=
				tpl.Execute(pongo2.Context{"type": typeDef, "typename": typeName, "name": goName,
					"references": schema.References, "parents": schema.Parents, "children": schema.Children})
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(
				strings.Replace(config.OutputPath, "__resource__", common.CamelToSnake(goName), 1),
				[]byte(output), 0644)
			if err != nil {
				return err
			}
		}
	} else if config.TemplateType == "alltype" {
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
		output, err :=
			tpl.Execute(pongo2.Context{"types": types})
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(config.OutputPath, []byte(output), 0644)
		if err != nil {
			return err
		}
	} else {
		for _, schema := range api.Schemas {
			if schema.Type == AbstractType || schema.ID == "" {
				continue
			}
			output, err :=
				tpl.Execute(pongo2.Context{"schema": schema, "types": api.Types})
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(
				strings.Replace(config.OutputPath, "__resource__", schema.ID, 1),
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
func ApplyTemplates(api *API, templateBase string, config []*TemplateConfig) error {
	for _, templateConfig := range config {
		err := templateConfig.apply(templateBase, api)
		if err != nil {
			return err
		}
	}
	return nil
}
