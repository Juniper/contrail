package logic

import (
	"encoding/json"
	"fmt"
)

// NeutronError structure.
type NeutronError struct {
	fields map[string]interface{}
}

// NewNeutronError creates new neutron error.
func NewNeutronError(name string, fields map[string]interface{}) *NeutronError {
	e := &NeutronError{
		fields: fields,
	}
	if fields == nil {
		e.fields = map[string]interface{}{}
	}
	e.fields["exception"] = name
	return e
}

func (e *NeutronError) Error() string {
	bytes, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("failed to marshall error description. error: %v", err)
	}
	return string(bytes)
}

// MarshalJSON custom marshalling code.
func (e *NeutronError) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.fields)
}

// here is the place to define constants for neutron exception names
// https://docs.openstack.org/neutron-lib/queens/reference/modules/neutron_lib.exceptions.html

const (
	// BadRequest exception name
	BadRequest = "BadRequest"
)
