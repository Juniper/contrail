package logic

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// Filters used in Neutron read API
type Filters map[string][]string

// Keys available in filters.
const (
	idKey             = "id"
	nameKey           = "name"
	fqNameKey         = "fq_name"
	sharedKey         = "shared"
	routerExternalKey = "router:external"
	tenantIDKey       = "tenant_id"
	deviceIDKey	      = "device_id"
)

// UnmarshalJSON Filters.
func (f *Filters) UnmarshalJSON(data []byte) error {
	if f == nil {
		return nil
	}
	var m map[string]interface{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}

	return f.ApplyMap(m)
}

// ApplyMap applies map onto filters.
func (f *Filters) ApplyMap(m map[string]interface{}) error {
	if *f == nil {
		*f = Filters{}
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

// HaveKeys checks if one or more keys are present in filters.
// Will return true if at least one key has been defined and all keys are present and not empty.
func (f Filters) HaveKeys(keys ...string) bool {
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

// Match checks if filters should accept values for given key.
// If key does not exist then it accepts every value and Match returns true.
func (f Filters) Match(key string, values ...string) bool {
	if !f.HaveKeys(key) {
		return true // This is intentional - if filters don't contain key, then we are not filtering out.
	}

	return f.HaveValues(key, values...)
}

// HaveValues check equality of values in filters struct under specific key and provided sequence of strings.
func (f Filters) HaveValues(key string, values ...string) bool {
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
