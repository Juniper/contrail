package agent

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/flosch/pongo2"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const showTemplate = ` Show resource data

{% for schema in schemas%} contrail show {{ schema.ID }} $ID
{% endfor %}`

const setTemplate = ` Set data on resource property

{% for schema in schemas%} contrail set {{ schema.ID }} $ID $YAML
{% endfor %}`

const listTemplate = ` List resource data

{% for schema in schemas%} contrail list {{ schema.ID }}
{% endfor %}`

const deleteTemplate = ` Delete resource data

{% for schema in schemas%} contrail rm {{ schema.ID }} $ID
{% endfor %}`

const schemaTemplate = `{% for schema in schemas%}
# {{schema.Title}} {{schema.Description}}
- kind: {{ schema.ID }}
  data: {% for key, value in schema.JSONSchema.Properties %}
    {{ key }}: {{ value.Default }} # {{ value.Title}} ({{ value.Type }}) {% endfor %}
{% endfor %}`

// ResourceData is data output by CLI tool.
type ResourceData struct {
	Kind string      `yaml:"kind"`
	Data interface{} `yaml:"data"`
}

// getInputResources decodes from YAML single or array of input data.
func getInputResources(data []byte) ([]*ResourceData, error) {
	var resources []*ResourceData
	var err error
	if err = yaml.Unmarshal(data, &resources); err != nil {
		var resource *ResourceData
		err = yaml.Unmarshal(data, &resource)
		if err == nil {
			resources = append(resources, resource)
		}
	}
	return resources, err
}

// ShowCLI returns YAML-formatted data resource with given schemaID and id.
// TODO(daniel): FIXME duplication
// nolint: dupl
func (a *Agent) ShowCLI(schemaID, id string) (string, error) {
	if schemaID == "" {
		tpl, err := pongo2.FromString(showTemplate)
		if err != nil {
			a.log.Fatal(err)
		}
		out, err := tpl.Execute(pongo2.Context{"schemas": a.serverAPI.Schemas})
		if err != nil {
			a.log.Fatal(err)
		}
		return out, nil
	} else if id == "" {
		return "", fmt.Errorf("missing ID")
	}

	data, err := a.show(schemaID, id)
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

// SchemaCLI returns YAML-formatted schema information of resource with given schemaID.
func (a *Agent) SchemaCLI(schemaID string) (string, error) {
	schemas := a.serverAPI.Schemas
	if schemaID != "" {
		schema, ok := a.schemas[schemaID]
		if !ok {
			return "", fmt.Errorf("schema %s not found", schemaID)
		}
		schemas = []*common.Schema{schema}

	}
	tpl, err := pongo2.FromString(schemaTemplate)
	if err != nil {
		a.log.Fatal(err)
	}
	out, err := tpl.Execute(pongo2.Context{"schemas": schemas})
	if err != nil {
		a.log.Fatal(err)
	}
	return out, nil
}

// ListCLI returns YAML-formatted data of all resources with given schemaID
// that satisfy filter, offset and limit.
func (a *Agent) ListCLI(schemaID string, filter string, offset, limit string) (string, error) {
	if schemaID == "" {
		tpl, err := pongo2.FromString(listTemplate)
		if err != nil {
			a.log.Fatal(err)
		}
		out, err := tpl.Execute(pongo2.Context{"schemas": a.serverAPI.Schemas})
		if err != nil {
			a.log.Fatal(err)
		}
		return out, nil
	}

	var schemas []string
	if schemaID == "all" {
		for _, schema := range a.serverAPI.Schemas {
			schemas = append(schemas, schema.ID)
		}
	} else {
		schemas = append(schemas, schemaID)
	}

	p := getListQueryParameters(filter, limit, offset)
	var resourceData []*ResourceData
	for _, sID := range schemas {
		data, err := a.list(sID, p)
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

func getListQueryParameters(filter string, limit string, offset string) url.Values {
	params := url.Values{}
	for _, kv := range strings.Split(filter, ",") {
		if kv == "" {
			continue
		}
		parts := strings.Split(kv, "=")
		params.Add(parts[0], parts[1])
	}
	params.Add("limit", limit)
	params.Add("offset", offset)
	return params
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
			data, cErr := a.create(resource.Kind, resourceData)
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
			data, uErr := a.update(resource.Kind, resourceData)
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
// It creates new resource for every non-existing resource ID.
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

// SetCLI updates properties specified in yamlString of resource with given schemaID and id.
// TODO(daniel): FIXME duplication
// nolint: dupl
func (a *Agent) SetCLI(schemaID, id, yamlString string) (string, error) {
	if schemaID == "" {
		tpl, err := pongo2.FromString(setTemplate)
		if err != nil {
			a.log.Fatal(err)
		}

		out, err := tpl.Execute(pongo2.Context{"schemas": a.serverAPI.Schemas})
		if err != nil {
			a.log.Fatal(err)
		}
		return out, nil
	} else if id == "" {
		return "", fmt.Errorf("missing ID")
	}

	var data map[interface{}]interface{}
	err := yaml.Unmarshal([]byte(yamlString), &data)
	if err != nil {
		return "", err
	}

	data["id"] = id
	updated, err := a.update(schemaID, data)
	if err != nil {
		return "", err
	}

	output, err := yaml.Marshal(&updated)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// RemoveCLI removes a resource with given schemaID and id.
func (a *Agent) RemoveCLI(schemaID, id string) (string, error) {
	if schemaID == "" {
		tpl, err := pongo2.FromString(deleteTemplate)
		if err != nil {
			return "", err
		}
		out, err := tpl.Execute(pongo2.Context{"schemas": a.serverAPI.Schemas})
		if err != nil {
			return "", err
		}
		return out, nil
	} else if id == "" {
		return "", fmt.Errorf("missing ID")
	}
	err := a.delete(schemaID, id)
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
			dataMap := resourceData.(map[interface{}]interface{})
			err := a.delete(resource.Kind, dataMap["id"].(string))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
