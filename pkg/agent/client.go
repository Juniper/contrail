package agent

import (
	"fmt"
	"net/url"
	"path"

	"github.com/pkg/errors"
)

func (a *Agent) show(schemaID, id string) (interface{}, error) {
	s, ok := a.schemas[schemaID]
	if !ok {
		return nil, fmt.Errorf("%s is undefined in API", schemaID)
	}
	var data interface{}
	_, err := a.APIServer.Read(path.Join(s.Path, id), &data)
	return &data, errors.Wrap(err, "show failed")
}

func (a *Agent) list(schemaID string, params url.Values) (interface{}, error) {
	s, ok := a.schemas[schemaID]
	if !ok {
		return nil, fmt.Errorf("%s is undefined in API", schemaID)
	}
	var data interface{}
	_, err := a.APIServer.Read(fmt.Sprintf("%s?%s", s.PluralPath, params.Encode()), &data)
	return &data, errors.Wrap(err, "list failed")
}

func (a *Agent) create(schemaID string, data interface{}) (interface{}, error) {
	var output interface{}
	s, ok := a.schemas[schemaID]
	if !ok {
		return nil, fmt.Errorf("%s is undefined in API", schemaID)
	}
	dataMap, ok := data.(map[interface{}]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid data format")
	}
	createData := map[string]interface{}{}
	for rawPropertyID, value := range dataMap {
		propertyID := rawPropertyID.(string)
		_, ok := s.JSONSchema.Properties[propertyID]
		if !ok {
			a.log.WithField("property-id", propertyID).Debug("Omitting invalid property")
			continue
		}
		createData[propertyID] = value
	}
	_, err := a.APIServer.Create(s.PluralPath, createData, &output)
	return &output, errors.Wrap(err, "create failed")
}

func (a *Agent) update(schemaID string, data interface{}) (interface{}, error) {
	var output interface{}
	s, ok := a.schemas[schemaID]
	if !ok {
		return nil, fmt.Errorf("%s is undefined in API", schemaID)
	}
	dataMap, ok := data.(map[interface{}]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid data format")
	}
	id := dataMap["id"].(string)
	updateData := map[string]interface{}{}
	for rawPropertyID, value := range dataMap {
		propertyID := rawPropertyID.(string)
		_, ok := s.JSONSchema.Properties[propertyID]
		if !ok {
			a.log.WithField("property-id", propertyID).Debug("Omitting invalid property")
			delete(dataMap, propertyID)
			continue
		}
		updateData[propertyID] = value
	}
	_, err := a.APIServer.Update(path.Join(s.Path, id), updateData, &output)
	return &output, errors.Wrap(err, "update failed")
}

func (a *Agent) resourceSync(schemaID string, data interface{}) (interface{}, error) {
	dataMap, ok := data.(map[interface{}]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid data format")
	}
	id := dataMap["id"].(string)
	_, err := a.show(schemaID, id)
	if err == nil {
		return a.update(schemaID, data)
	}
	return a.create(schemaID, data)
}

func (a *Agent) delete(schemaID, id string) error {
	var output interface{}
	s, ok := a.schemas[schemaID]
	if !ok {
		return fmt.Errorf("%s is undefined in API", schemaID)
	}
	_, err := a.APIServer.Delete(path.Join(s.Path, id), &output)
	return errors.Wrap(err, "delete failed")
}
