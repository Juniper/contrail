package schema

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/flosch/pongo2"
	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/fileutil"
)

const (
	dictGetJSONSchemaByStringKeyFilter = "dict_get_JSONSchema_by_string_key"
)

// TemplateConfig contains configuration for template.
type TemplateConfig struct {
	TemplateType        string `yaml:"type"`
	TemplatePath        string `yaml:"template_path"`
	OutputPath          string `yaml:"output_path"`
}

// TemplateOption contains options for template.
type TemplateOption struct {
	SchemasDir        string
	TemplateConfPath  string
	SchemaOutputPath  string
	OpenapiOutputPath string
	OutputDir         string
}

// ApplyTemplates writes files with content generated from templates.
func ApplyTemplates(api *API, config []*TemplateConfig, option *TemplateOption) error {
	if err := registerCustomFilters(); err != nil {
		return err
	}

	for _, tc := range config {
		if err := tc.resolveOutputPath(); err != nil {
			return err
		}
		if !tc.isOutdated(api) {
			continue
		}
		err := tc.apply(api, option)
		if err != nil {
			return err
		}
	}
	return nil
}

func (tc *TemplateConfig) resolveOutputPath() error {
	if tc.OutputPath != "" {
		return nil
	}

	tc.OutputPath = tc.getOutputPathFromTemplatePath()
	return nil
}

func (tc *TemplateConfig) isOutdated(api *API) bool {
	sourceInfo, err := os.Stat(tc.TemplatePath)
	if err != nil {
		return true
	}
	targetInfo, err := os.Stat(tc.OutputPath)
	if err != nil {
		return true
	}
	sourceModTime := sourceInfo.ModTime()
	targetModTime := targetInfo.ModTime()
	return sourceModTime.After(targetModTime) || api.Timestamp.After(targetModTime)
}

// nolint: gocyclo
func (tc *TemplateConfig) apply(api *API, option *TemplateOption) error {
	tpl, err := tc.load()
	if err != nil {
		return err
	}
	if err = ensureDir(tc.OutputPath); err != nil {
		return err
	}
	if tc.TemplateType == "all" {
		output, err := tpl.Execute(pongo2.Context{"schemas": api.Schemas, "types": api.Types,
			"option": option,
		})
		if err != nil {
			return err
		}

		if err = writeGeneratedFile(tc.OutputPath, output, tc.TemplatePath); err != nil {
			return err
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

		if err = writeGeneratedFile(tc.OutputPath, output, tc.TemplatePath); err != nil {
			return err
		}
	}
	return nil
}

func (tc *TemplateConfig) getOutputPathFromTemplatePath() string {
	return filepath.Join(tc.getTemplateFileDirectory(), tc.getGeneratedFileName())
}

func (tc *TemplateConfig) getGeneratedFileName() string {
	return fmt.Sprintf("gen_%s", tc.getTemplateFilenameWithoutExtension())
}

func (tc *TemplateConfig) getTemplateFilenameWithoutExtension() string {
	return strings.TrimSuffix(tc.getTemplateFilename(), filepath.Ext(tc.getTemplateFilename()))
}

func (tc *TemplateConfig) getTemplateFileDirectory() string {
	return filepath.Dir(tc.TemplatePath)
}

func (tc *TemplateConfig) getTemplateFilename() string {
	return filepath.Base(tc.TemplatePath)
}

func (tc *TemplateConfig) load() (*pongo2.Template, error) {
	templateCode, err := fileutil.GetContent(tc.TemplatePath)
	if err != nil {
		return nil, err
	}
	return pongo2.FromString(string(templateCode))
}

// LoadTemplates loads template configurations from given path.
func LoadTemplates(path string) ([]*TemplateConfig, error) {
	var config []*TemplateConfig
	err := fileutil.LoadFile(path, &config)
	return config, err
}

func ensureDir(path string) error {
	return os.MkdirAll(filepath.Dir(path), os.ModePerm)
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
				m, _ := in.Interface().(map[string]*JSONSchema) //nolint: errcheck
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
	return prefix + fmt.Sprintf("Code generated by contrailschema tool from template %s; DO NOT EDIT.\n\n", template)
}
