package agent

import (
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/pkg/errors"
)

const uuidKey = "uuid"

func (a *Agent) show(schemaID, uuid string) (interface{}, error) {
	s, ok := a.schemas[schemaID]
	if !ok {
		return nil, fmt.Errorf("%s is undefined in API", schemaID)
	}
	var data interface{}
	_, err := a.APIServer.Read(path.Join(s.Path, uuid), &data)
	return &data, errors.Wrap(err, "show operation failed")
}

func (a *Agent) list(schemaID string, queryParameters url.Values) (interface{}, error) {
	s, ok := a.schemas[schemaID]
	if !ok {
		return nil, fmt.Errorf("%s is undefined in API", schemaID)
	}
	var data interface{}
	_, err := a.APIServer.Read(fmt.Sprintf("%s?%s", s.PluralPath, queryParameters.Encode()), &data)
	return &data, errors.Wrap(err, "list operation failed")
}

func (a *Agent) create(schemaID string, data interface{}) (interface{}, error) {
	var output interface{}
	s, ok := a.schemas[schemaID]
	if !ok {
		return nil, fmt.Errorf("%s is undefined in API", schemaID)
	}

	rd, err := a.buildRequestData(data, s)
	if err != nil {
		return nil, err
	}

	_, err = a.APIServer.Create(s.PluralPath, rd, &output)
	if err != nil {
		return nil, fmt.Errorf("create operation failed: %s", err)
	}

	return extractServerOutput(output, schemaID)
}

func (a *Agent) update(schemaID string, data interface{}) (interface{}, error) {
	var output interface{}
	s, ok := a.schemas[schemaID]
	if !ok {
		return nil, fmt.Errorf("%s is undefined in API", schemaID)
	}

	rd, err := a.buildRequestData(data, s)
	if err != nil {
		return nil, err
	}

	uuid, ok := rd[dashedCase(schemaID)][uuidKey].(string)
	if !ok {
		return nil, fmt.Errorf("data does not contain required UUID property")
	}

	_, err = a.APIServer.Update(path.Join(s.Path, uuid), rd, &output)
	if err != nil {
		return nil, fmt.Errorf("update operation failed: %s", err)
	}

	return extractServerOutput(output, schemaID)
}

func extractServerOutput(output interface{}, schemaID string) (interface{}, error) {
	o, ok := output.(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid server output: no top-level mapping")
	}

	return o[dashedCase(schemaID)], nil
}

func (a *Agent) buildRequestData(data interface{}, schema *common.Schema) (map[string]map[string]interface{}, error) {
	properties, ok := data.(map[interface{}]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid data format: no properties mapping")
	}

	requestData := map[string]map[string]interface{}{
		dashedCase(schema.ID): make(map[string]interface{}),
	}
	for rawPropertyID, value := range properties {
		p, ok := rawPropertyID.(string)
		if !ok {
			return nil, fmt.Errorf("invalid data format: property keys should have string type")
		}
		if _, ok = schema.JSONSchema.Properties[p]; !ok {
			a.log.WithField("property-id", p).Info("Omitting invalid property")
			continue
		}
		if a.isCompoundProperty(p, value) {
			continue
		}

		requestData[dashedCase(schema.ID)][p] = value
	}

	return requestData, nil
}

func dashedCase(schemaID string) string {
	return strings.Replace(schemaID, "_", "-", -1)
}

// isCompoundProperty returns true if given value is of map[interface{}]interface{} type.
// TODO(daniel): support compound property objects
// Maps with interface{} keys cannot be encoded to JSON.
// Workarounds are available here: https://github.com/go-yaml/yaml/issues/139
func (a *Agent) isCompoundProperty(propertyID string, value interface{}) bool {
	if _, ok := value.(map[interface{}]interface{}); ok {
		a.log.WithField("property-id", propertyID).Info("Omitting compound property object")
		return true
	}
	return false
}

func (a *Agent) resourceSync(schemaID string, data interface{}) (interface{}, error) {
	uuid, err := uuidFromRawProperties(data)
	if err != nil {
		return nil, err
	}

	_, err = a.show(schemaID, uuid)
	if err == nil {
		return a.update(schemaID, data)
	}
	return a.create(schemaID, data)
}

func uuidFromRawProperties(rawProperties interface{}) (string, error) {
	properties, ok := rawProperties.(map[interface{}]interface{})
	if !ok {
		return "", fmt.Errorf("invalid data format: no properties mapping")
	}

	rawUUID, ok := properties[uuidKey]
	if !ok {
		return "", errors.New("data does not contain required UUID property")
	}

	uuid, ok := rawUUID.(string)
	if !ok {
		return "", fmt.Errorf("UUID should be string instead of %T", uuid)
	}
	return uuid, nil
}

func (a *Agent) delete(schemaID, uuid string) error {
	var output interface{}
	s, ok := a.schemas[schemaID]
	if !ok {
		return fmt.Errorf("%s is undefined in API", schemaID)
	}
	_, err := a.APIServer.Delete(path.Join(s.Path, uuid), &output)
	return errors.Wrap(err, "delete operation failed")
}
