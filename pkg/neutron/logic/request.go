package logic

import (
	"encoding/json"

	"github.com/pkg/errors"
)

const (
	OperationCreate       = "CREATE"
	OperationRead         = "READ"
	OperationReadAll      = "READALL"
	OperationUpdate       = "UPDATE"
	OperationDelete       = "DELETE"
	OperationReadCount    = "READCOUNT"
	OperationAddInterface = "ADDINTERFACE"
	OperationDelInterface = "DELINTERFACE"
)

// Request defines an API request.
type Request struct {
	Data    Data           `json:"data" yaml:"data"`
	Context RequestContext `json:"context" yaml:"context"`
}

// Data defines API request data.
type Data struct {
	Filters  Filters  `json:"filters" yaml:"filters"`
	ID       string   `json:"id" yaml:"id"`
	Fields   Fields   `json:"fields" yaml:"fields"`
	Resource Resource `json:"resource" yaml:"resource"`
}

// GetType returns resource type of the Request.
func (r *Request) GetType() string {
	return r.Context.Type
}

// UnmarshalJSON custom unmarshalling of Request.
func (r *Request) UnmarshalJSON(data []byte) error {
	var rawJSON map[string]json.RawMessage
	err := json.Unmarshal(data, &rawJSON)
	if err != nil {
		return err
	}

	err = parseField(rawJSON, "context", &r.Context)
	if err != nil {
		return err
	}

	resource, err := MakeResource(r.Context.Type)
	if err != nil {
		return err
	}

	r.Data.Resource = resource
	return parseField(rawJSON, "data", &r.Data)
}

func parseField(rawJSON map[string]json.RawMessage, key string, dst interface{}) error {
	if val, ok := rawJSON[key]; ok {
		if err := json.Unmarshal(val, dst); err != nil {
			return errors.Errorf("invalid '%s' format: %v", key, err)
		}
		delete(rawJSON, key)
	}
	return nil
}
