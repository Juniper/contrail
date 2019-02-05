package logic

import (
	"encoding/json"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/format"
)

// Request defines an API request.
type Request struct {
	Data    Data           `json:"data" yaml:"data"`
	Context RequestContext `json:"context" yaml:"context"`
}

// Data defines API request data.
type Data struct {
	Filters   Filters  `json:"filters" yaml:"filters"`
	ID        string   `json:"id" yaml:"id"`
	Fields    Fields   `json:"fields" yaml:"fields"`
	Resource  Resource `json:"resource" yaml:"resource"`
	FieldMask types.FieldMask
}

// GetType returns resource type of the Request.
func (r *Request) GetType() string {
	return r.Context.Type
}

// UnmarshalJSON performs custom unmarshalling of Request.
func (r *Request) UnmarshalJSON(data []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	return r.ApplyMap(m)
}

// ApplyMap applies map into request.
func (r *Request) ApplyMap(m map[string]interface{}) error {
	cm, ok := m["context"].(map[string]interface{})
	if !ok {
		return errors.Errorf("got invalid context: %v", m["context"])
	}
	if err := format.ApplyMap(cm, &r.Context); err != nil {
		return err
	}
	resource, err := MakeResource(r.Context.Type)
	if err != nil {
		return err
	}
	r.Data.Resource = resource
	dm, ok := m["data"].(map[string]interface{})
	if !ok {
		return errors.Errorf("got invalid data: %v", m["data"])
	}

	return format.ApplyMap(dm, &r.Data)
}
