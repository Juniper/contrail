package logic

import (
	"encoding/json"
<<<<<<< HEAD
=======
	"fmt"
>>>>>>> 477652b2... Unmarshalling of boolean filters

	"github.com/pkg/errors"
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

// UnmarshalJSON Filters.
func (f *Filters) UnmarshalJSON(data []byte) error {
	if *f == nil {
		*f = Filters{}
	}
	var m map[string]interface{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		var ss []string
		switch s := v.(type) {
		case []interface{}:
			for _, i := range s {
				switch c := i.(type) {
				case bool:
					ss = append(ss, fmt.Sprintf("%t", c))
				case string:
					ss = append(ss, fmt.Sprintf("%s", c))
				default:
					return errors.Errorf("%T filter not supported", v)
				}
			}
		default:
			return errors.Errorf("%T filter not supported", v)
		}

		(*f)[k] = ss
	}
	return nil
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
