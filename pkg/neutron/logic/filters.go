package logic

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// FilterKey is a type of keys stored in filters.
type FilterKey string

// Filters used in Neutron read API
type Filters map[FilterKey][]string

// Keys available in filters.
const (
	idKey             FilterKey = "id"
	nameKey           FilterKey = "name"
	fqNameKey         FilterKey = "fq_name"
	sharedKey         FilterKey = "shared"
	routerExternalKey FilterKey = "router:external"
	tenantIDKey       FilterKey = "tenant_id"
)

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

		(*f)[FilterKey(k)] = ss
	}
	return nil
}

// haveKeys checks if one or more keys are present in filters.
// Will return true if at least one key has been defined and all keys are present and not empty.
func (f Filters) haveKeys(keys ...FilterKey) bool {
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
func (f Filters) checkValue(key FilterKey, values ...string) bool {
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
