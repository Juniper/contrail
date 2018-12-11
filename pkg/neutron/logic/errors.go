package logic

import (
	"encoding/json"
	"fmt"
)

// Error structure.
type Error struct {
	fields map[string]interface{}
}

// NewNeutronError creates new Neutron error.
func NewNeutronError(name string, fields map[string]interface{}) *Error {
	e := &Error{
		fields: fields,
	}
	if fields == nil {
		e.fields = map[string]interface{}{}
	}
	e.fields["exception"] = name
	return e
}

func (e *Error) Error() string {
	bytes, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("failed to marshall error description: %v", err)
	}
	return string(bytes)
}

// MarshalJSON custom marshalling code.
func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.fields)
}

// constants for Neutron API exception names
// https://docs.openstack.org/neutron-lib/queens/reference/modules/neutron_lib.exceptions.html
const (
	BadRequest            = "BadRequest"
	SecurityGroupNotFound = "SecurityGroupNotFound"
)
