package logic

import (
	"encoding/json"
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

func (r *Request) UnmarshalJSON(data []byte) error {
	var rawJSON map[string]json.RawMessage
	err := json.Unmarshal(data, &rawJSON)
	if err != nil {
		return err
	}

	r.Context = RequestContext{}
	err = ParseField(rawJSON, "context", &r.Context)
	if err != nil {
		return err
	}

	resource, err := MakeResource(r.Context.Type)
	if err != nil {
		return err
	}

	r.Data.Resource = resource
	err = ParseField(rawJSON, "data", &r.Data)
	if err != nil {
		return err
	}
	return err
}
