package neutronerrors

import (
	"encoding/json"
)

type NeutronError struct {
	fields map[string]interface{}
}

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
	b, err := json.Marshal(e.fields)
	if err != nil {
		return ""
	}
	return string(b)
}

// here is the place to define constants for neutron exception names
// https://docs.openstack.org/neutron-lib/queens/reference/modules/neutron_lib.exceptions.html

const (
	BadRequest = "BadRequest"
)
