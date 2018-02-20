package agent

import (
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/Juniper/contrail/pkg/schema"
	"github.com/flosch/pongo2"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const schemaTemplate = `
{% for schema in schemas %}
# {{ schema.Title }} {{ schema.Description }}
- kind: {{ schema.ID }}
  data: {% for key, value in schema.JSONSchema.Properties %}
    {{ key }}: {{ value.Default }} # {{ value.Title }} ({{ value.Type }}) {% endfor %}
{% endfor %}`

const showHelpTemplate = `Show command possible usages:

{% for schema in schemas %}contrail show {{ schema.ID }} $UUID
{% endfor %}`

const listHelpTemplate = `List command possible usages:

{% for schema in schemas %}contrail list {{ schema.ID }}
{% endfor %}`

const setHelpTemplate = `Set command possible usages:

{% for schema in schemas %}contrail set {{ schema.ID }} $UUID $YAML
{% endfor %}`

const removeHelpTemplate = `Remove command possible usages:

{% for schema in schemas %}contrail rm {{ schema.ID }} $UUID
{% endfor %}`

// ResourceData is data output by CLI tool.
type ResourceData struct {
	Kind string      `yaml:"kind"`
	Data interface{} `yaml:"data"`
}

// SchemaCLI returns YAML-formatted schema information of resource with given schemaID.
func (a *Agent) SchemaCLI(schemaID string) (string, error) {
	schemas := a.serverAPI.Schemas
	if schemaID != "" {
		s, ok := a.schemas[schemaID]
		if !ok {
			return "", fmt.Errorf("schema %s not found", schemaID)
		}
		schemas = []*schema.Schema{s}

	}
	tpl, err := pongo2.FromString(schemaTemplate)
	if err != nil {
		return "", err
	}
	o, err := tpl.Execute(pongo2.Context{"schemas": schemas})
	if err != nil {
		return "", err
	}
	return o, nil
}

// ShowCLI returns YAML-formatted data resource with given schemaID and UUID.
// TODO(daniel): FIXME duplication
// nolint: dupl
func (a *Agent) ShowCLI(schemaID, uuid string) (string, error) {
	if schemaID == "" {
		tpl, err := pongo2.FromString(showHelpTemplate)
		if err != nil {
			return "", err
		}
		o, err := tpl.Execute(pongo2.Context{"schemas": a.serverAPI.Schemas})
		if err != nil {
			return "", err
		}
		return o, nil
	} else if uuid == "" {
		return "", fmt.Errorf("missing UUID")
	}

	data, err := a.Show(schemaID, uuid)
	if err != nil {
		a.log.Fatalf("error: %v", err)
	}
	output, err := yaml.Marshal(
		&ResourceData{
			Kind: schemaID,
			Data: data,
		},
	)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// ListCLI returns YAML-formatted data of all resources with given schemaID
// that satisfy filter, offset and limit.
func (a *Agent) ListCLI(schemaID string, queryParameters url.Values) (string, error) {
	if schemaID == "" {
		return buildListHelpMessage(a.serverAPI.Schemas)
	}

	var schemas []string
	if schemaID == "all" {
		for _, schema := range a.serverAPI.Schemas {
			schemas = append(schemas, schema.ID)
		}
	} else {
		schemas = append(schemas, schemaID)
	}

	var resourceData []*ResourceData
	for _, sID := range schemas {
		data, err := a.List(sID, queryParameters)
		if err != nil {
			return "", err
		}
		resourceData = append(resourceData, &ResourceData{
			Kind: sID,
			Data: data,
		})
	}

	output, err := yaml.Marshal(resourceData)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func buildListHelpMessage(schemas []*schema.Schema) (string, error) {
	tpl, err := pongo2.FromString(listHelpTemplate)
	if err != nil {
		return "", err
	}
	o, err := tpl.Execute(pongo2.Context{"schemas": schemas})
	if err != nil {
		return "", err
	}
	return o, nil
}

// CreateCLI creates resources with data specified in file with given dataPath.
// TODO(daniel): FIXME duplication
// nolint: dupl
func (a *Agent) CreateCLI(dataPath string) (string, error) {
	data, err := ioutil.ReadFile(dataPath)
	if err != nil {
		return "", err
	}
	resources, err := getInputResources(data)
	if err != nil {
		return "", err
	}
	for _, resource := range resources {
		list, ok := resource.Data.([]interface{})
		if !ok {
			list = []interface{}{resource.Data}
		}
		var createdData []interface{}
		for _, resourceData := range list {
			data, cErr := a.Create(resource.Kind, resourceData)
			if cErr != nil {
				return "", cErr
			}
			createdData = append(createdData, data)
		}
		resource.Data = createdData
	}
	output, err := yaml.Marshal(&resources)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// SetCLI updates properties specified in yamlString of resource with given schemaID and UUID.
// TODO(daniel): FIXME duplication
// nolint: dupl
func (a *Agent) SetCLI(schemaID, uuid, yamlString string) (string, error) {
	if schemaID == "" {
		tpl, err := pongo2.FromString(setHelpTemplate)
		if err != nil {
			return "", err
		}

		o, err := tpl.Execute(pongo2.Context{"schemas": a.serverAPI.Schemas})
		if err != nil {
			return "", err
		}
		return o, nil
	} else if uuid == "" {
		return "", fmt.Errorf("missing UUID")
	}

	var data map[interface{}]interface{}
	err := yaml.Unmarshal([]byte(yamlString), &data)
	if err != nil {
		return "", err
	}

	data["uuid"] = uuid
	updated, err := a.Update(schemaID, data)
	if err != nil {
		return "", err
	}

	output, err := yaml.Marshal(&updated)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// UpdateCLI updates resources with data specified in file with given dataPath.
// TODO(daniel): FIXME duplication
// nolint: dupl
func (a *Agent) UpdateCLI(dataPath string) (string, error) {
	data, err := ioutil.ReadFile(dataPath)
	if err != nil {
		return "", errors.Wrap(err, "failed to load datafile")
	}
	resources, err := getInputResources(data)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse data file")
	}
	for _, resource := range resources {
		list, ok := resource.Data.([]interface{})
		if !ok {
			list = []interface{}{resource.Data}
		}
		var updatedData []interface{}
		for _, resourceData := range list {
			data, uErr := a.Update(resource.Kind, resourceData)
			if uErr != nil {
				return "", uErr
			}
			updatedData = append(updatedData, data)
		}
		resource.Data = updatedData
	}
	output, err := yaml.Marshal(&resources)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// SyncCLI updates resources with data specified in given dataPath.
// It creates new resource for every non-existing resource UUID.
// TODO(daniel): FIXME duplication
// nolint: dupl
func (a *Agent) SyncCLI(dataPath string) (string, error) {
	data, err := ioutil.ReadFile(dataPath)
	if err != nil {
		return "", errors.Wrap(err, "failed to load datafile")
	}
	resources, err := getInputResources(data)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse data file")
	}
	for _, resource := range resources {
		list, ok := resource.Data.([]interface{})
		if !ok {
			list = []interface{}{resource.Data}
		}
		var updatedData []interface{}
		for _, resourceData := range list {
			data, sErr := a.resourceSync(resource.Kind, resourceData)
			if sErr != nil {
				return "", sErr
			}
			updatedData = append(updatedData, data)
		}
		resource.Data = updatedData
	}
	output, err := yaml.Marshal(&resources)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// RemoveCLI removes a resource with given schemaID and UUID.
func (a *Agent) RemoveCLI(schemaID, uuid string) (string, error) {
	if schemaID == "" {
		tpl, err := pongo2.FromString(removeHelpTemplate)
		if err != nil {
			return "", err
		}
		o, err := tpl.Execute(pongo2.Context{"schemas": a.serverAPI.Schemas})
		if err != nil {
			return "", err
		}
		return o, nil
	} else if uuid == "" {
		return "", fmt.Errorf("missing UUID")
	}
	err := a.Delete(schemaID, uuid)
	if err != nil {
		return "", err
	}
	return "", nil
}

// DeleteCLI deletes resources with IDs specified in file with given dataPath.
func (a *Agent) DeleteCLI(datafile string) error {
	data, err := ioutil.ReadFile(datafile)
	if err != nil {
		return err
	}
	resources, err := getInputResources(data)
	if err != nil {
		return err
	}
	for i := len(resources) - 1; i >= 0; i-- {
		resource := resources[i]
		list, ok := resource.Data.([]interface{})
		if !ok {
			list = []interface{}{resource.Data}
		}
		for _, resourceData := range list {
			uuid, err := uuidFromRawProperties(resourceData)
			if err != nil {
				return err
			}

			if err = a.Delete(resource.Kind, uuid); err != nil {
				return err
			}
		}
	}
	return nil
}

// getInputResources decodes single or array of input data from YAML.
func getInputResources(yamlData []byte) ([]*ResourceData, error) {
	var resources []*ResourceData
	var err error
	if err = yaml.Unmarshal(yamlData, &resources); err != nil {
		var resource *ResourceData
		err = yaml.Unmarshal(yamlData, &resource)
		if err == nil {
			resources = append(resources, resource)
		}
	}
	return resources, err
}
