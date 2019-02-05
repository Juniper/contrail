package logic

import (
	"encoding/json"
	"fmt"
)

type errorType string

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

func newNeutronError(name errorType, fields errorFields) *Error {
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
	internalServerError        errorType = "InternalServerError"
	badRequest                 errorType = "BadRequest"
	portNotFound               errorType = "PortNotFound"
	l3PortInUse                errorType = "L3PortInUse"
	macAddressInUse            errorType = "MacAddressInUse"
	ipAddressGenerationFailure errorType = "IpAddressGenerationFailure"
	networkNotFound            errorType = "NetworkNotFound"
	subnetNotFound             errorType = "SubnetNotFound"
	securityGroupNotFound      errorType = "SecurityGroupNotFound"
	securityGroupRuleNotFound  errorType = "SecurityGroupRuleNotFound"
	networkInUse               errorType = "NetworkInUse"
	invalidSharedSettings      errorType = "InvalidSharedSetting"
)
