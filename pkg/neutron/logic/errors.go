package logic

import (
	"encoding/json"
	"fmt"
)

// ErrorFields neutron error fields.
type errorFields map[string]interface{}

// Error structure.
type Error struct {
	fields errorFields
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

func newNeutronError(name string, fields errorFields) *Error {
	e := &Error{
		fields: fields,
	}
	if fields == nil {
		e.fields = errorFields{}
	}
	e.fields["exception"] = name
	return e
}

// constants for Neutron API exception names
// https://docs.openstack.org/neutron-lib/queens/reference/modules/neutron_lib.exceptions.html
const (
	internalServerError        = "InternalServerError"
	badRequest                 = "BadRequest"
	portNotFound               = "PortNotFound"
	l3PortInUse                = "L3PortInUse"
	macAddressInUse            = "MacAddressInUse"
	ipAddressGenerationFailure = "IpAddressGenerationFailure"
	networkNotFound            = "NetworkNotFound"
	securityGroupNotFound      = "SecurityGroupNotFound"
	securityGroupRuleNotFound  = "SecurityGroupRuleNotFound"
	floatingIPNotFound         = "FloatingIPNotFound"
)
