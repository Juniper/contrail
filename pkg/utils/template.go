package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/flosch/pongo2"
)

type TemplateConfig struct {
	TemplateType string `yaml:"type"`
	TemplatePath string `yaml:"template_path"`
	OutputPath   string `yaml:"output_path"`
}

func ensureDir(path string) {
	os.MkdirAll(filepath.Dir(path), os.ModePerm)
}

func (config *TemplateConfig) Load(base string) (*pongo2.Template, error) {
	path := filepath.Join(base, config.TemplatePath)
	templateCode, err := GetContent(path)
	if err != nil {
		return nil, err
	}
	return pongo2.FromString(string(templateCode))
}

func (config *TemplateConfig) Apply(templateBase string, api *API) error {
	tpl, err := config.Load(templateBase)
	if err != nil {
		return err
	}
	ensureDir(config.OutputPath)
	if config.TemplateType == "all" {
		output, err :=
			tpl.Execute(pongo2.Context{"schemas": api.Schemas})
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(config.OutputPath, []byte(output), 0644)
		if err != nil {
			return err
		}
	} else if config.TemplateType == "type" {
		for typeName, typeDef := range api.Types {
			output, err :=
				tpl.Execute(pongo2.Context{"type": typeDef, "name": typeName})
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(
				strings.Replace(config.OutputPath, "__resource__", CamelToSnake(typeName), 1),
				[]byte(output), 0644)
			if err != nil {
				return err
			}
		}
		for _, schema := range api.Schemas {
			if schema.Type == "abstract" || schema.ID == "" {
				continue
			}
			typeName := schema.JSONSchema.GoName
			typeDef := schema.JSONSchema
			output, err :=
				tpl.Execute(pongo2.Context{"type": typeDef, "name": typeName, "references": schema.References, "parents": schema.Parents})
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(
				strings.Replace(config.OutputPath, "__resource__", CamelToSnake(typeName), 1),
				[]byte(output), 0644)
			if err != nil {
				return err
			}
		}
	} else {
		for _, schema := range api.Schemas {
			if schema.Type == "abstract" || schema.ID == "" {
				continue
			}
			output, err :=
				tpl.Execute(pongo2.Context{"schema": schema})
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

func LoadTemplates(path string) ([]*TemplateConfig, error) {
	var config []*TemplateConfig
	err := LoadFile(path, &config)
	return config, err
}

func ApplyTemplates(api *API, templateBase string, config []*TemplateConfig) error {
	for _, templateConfig := range config {
		err := templateConfig.Apply(templateBase, api)
		if err != nil {
			return err
		}
	}
	return nil
}
