package logic

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// Filters used in Neutron read API
type Filters map[string][]string

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
	err = f.ApplyMap(m)
	if err != nil {
		return err
	}
	return nil
}

func (f *Filters) ApplyMap(m map[string]interface{}) error {
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

// haveKeys checks if one or more keys are present in filters.
// Will return true if at least one key has been defined and all keys are present and not empty.
func (f Filters) haveKeys(keys ...string) bool {
	if len(keys) == 0 {
		return false
	}

	for _, key := range keys {
		filter, ok := f[key]
		if !ok || len(filter) == 0 {
			return false
		}
	}

	return true
}

// checkValue check equality of values in filters struct under specific key and provided sequence of strings
func (f Filters) checkValue(key string, values ...string) bool {
	if !f.haveKeys(key) {
		return true
	}
	if len(f[key]) != len(values) {
		return false
	}

	for i, v := range values {
		if f[key][i] != v {
			return false
		}
	}

	return true
}
