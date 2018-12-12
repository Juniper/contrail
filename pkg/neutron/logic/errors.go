package logic

import (
	"encoding/json"
	"fmt"
)

// ErrorFields neutron error fields.
type ErrorFields map[string]interface{}

// Error structure.
type Error struct {
	fields ErrorFields
}

// NewNeutronError creates new Neutron error.
func NewNeutronError(name string, fields ErrorFields) *Error {
	e := &Error{
		fields: fields,
	}
	if fields == nil {
		e.fields = ErrorFields{}
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
	BadRequest      = "BadRequest"
	NetworkNotFound = "NetworkNotFound"
)
