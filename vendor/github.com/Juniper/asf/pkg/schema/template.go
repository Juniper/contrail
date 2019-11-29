package schema

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/flosch/pongo2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	dictGetJSONSchemaByStringKeyFilter = "dict_get_JSONSchema_by_string_key"
)

// TemplateConfig contains a configuration for the template.
type TemplateConfig struct {
	TemplateType string `yaml:"type"`
	TemplatePath string `yaml:"template_path"`
	OutputPath   string `yaml:"output_path"`
}

// GenerateFiles generates files by applying API data to templates specified in template config.
func GenerateFiles(api *API, config []*TemplateConfig) error {
	if api == nil {
		return errors.New("received API is nil")
	}

	if err := registerCustomFilters(); err != nil {
		return errors.Wrap(err, "register filters")
	}

	for _, tc := range config {
		tc.resolveOutputPath()
		if !tc.isOutdated(api) {
			logrus.WithField(
				"template-config", fmt.Sprintf("%+v", tc),
			).Debug("Target file is up to date - skipping generation")
			continue
		}
		err := tc.generateFile(api)
		if err != nil {
			return errors.Wrap(err, "generate file")
		}
	}
	return nil
}

func (tc *TemplateConfig) resolveOutputPath() {
	if tc.OutputPath == "" {
		tc.OutputPath = generatedFilePath(tc.TemplatePath)
	}
	return
}

func (tc *TemplateConfig) isOutdated(api *API) bool {
	if api.Timestamp.IsZero() {
		return true
	}
	sourceInfo, err := os.Stat(tc.TemplatePath)
	if err != nil {
		return true
	}
	targetInfo, err := os.Stat(tc.OutputPath)
	if err != nil {
		return true
	}
	return sourceInfo.ModTime().After(targetInfo.ModTime()) || api.Timestamp.After(targetInfo.ModTime())
}

// nolint: gocyclo
func (tc *TemplateConfig) generateFile(api *API) error {
	tpl, err := loadTemplate(tc.TemplatePath)
	if err != nil {
		return errors.Wrapf(err, "load template %q", tc.TemplatePath)
	}

	if err = ensureDirectoryExists(tc.OutputPath); err != nil {
		return errors.Wrapf(err, "ensure the directory exists for output path: %q", tc.OutputPath)
	}

	if tc.TemplateType == "all" {
		data, err := tpl.Execute(pongo2.Context{
			"schemas": api.Schemas,
			"types":   api.Types,
		})
		if err != nil {
			return errors.Wrapf(err, "execute template %q", tc.TemplatePath)
		}

		if err = generateFile(tc.OutputPath, data, tc.TemplatePath); err != nil {
			return errors.Wrapf(err, "generate the file from template %q", tc.TemplatePath)
		}
	} else if tc.TemplateType == "alltype" {
		var schemas []*Schema
		for typeName, typeJSONSchema := range api.Types {
			typeJSONSchema.GoName = typeName
			schemas = append(schemas, &Schema{
				JSONSchema:     typeJSONSchema,
				Children:       map[string]*BackReference{},
				BackReferences: map[string]*BackReference{},
			})
		}
		for _, schema := range api.Schemas {
			if schema.Type == AbstractType || schema.ID == "" {
				continue
			}
			schemas = append(schemas, schema)
		}
		data, err := tpl.Execute(pongo2.Context{
			"schemas": schemas,
		})
		if err != nil {
			return errors.Wrapf(err, "execute template %q", tc.TemplatePath)
		}

		if err = generateFile(tc.OutputPath, data, tc.TemplatePath); err != nil {
			return errors.Wrapf(err, "generate the file from template %q", tc.TemplatePath)
		}
	}
	return nil
}

func loadTemplate(templatePath string) (*pongo2.Template, error) {
	o, err := fileutil.GetContent(templatePath)
	if err != nil {
		return nil, errors.Wrapf(err, "get content of template %q", templatePath)
	}
	return pongo2.FromString(string(o))
}

// LoadTemplateConfig loads template configurations from given path.
func LoadTemplateConfig(path string) ([]*TemplateConfig, error) {
	var config []*TemplateConfig
	err := fileutil.LoadFile(path, &config)
	return config, err
}

func generatedFilePath(tmplFilePath string) string {
	dir, file := filepath.Split(tmplFilePath)
	return filepath.Join(dir, generatedFileName(file))
}

func generatedFileName(tmplFile string) string {
	return "gen_" + strings.TrimSuffix(tmplFile, ".tmpl")
}

func ensureDirectoryExists(path string) error {
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
			return errors.Wrapf(err, "register filter %q", dictGetJSONSchemaByStringKeyFilter)
		}
	}
	return nil
}

func generateFile(outputPath, data, templatePath string) error {
	logrus.WithFields(logrus.Fields{
		"output-path":   outputPath,
		"template-path": templatePath,
	}).Debug("Generating file")
	if err := ioutil.WriteFile(outputPath, []byte(generationPrefix(outputPath, templatePath)+data), 0644); err != nil {
		return errors.Wrapf(err, "write the file to path %q", outputPath)
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
