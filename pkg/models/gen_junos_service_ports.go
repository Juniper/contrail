package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeJunosServicePorts makes JunosServicePorts
// nolint
func MakeJunosServicePorts() *JunosServicePorts {
	return &JunosServicePorts{
		//TODO(nati): Apply default
		ServicePort: []string{},
	}
}

// MakeJunosServicePorts makes JunosServicePorts
// nolint
func InterfaceToJunosServicePorts(i interface{}) *JunosServicePorts {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &JunosServicePorts{
		//TODO(nati): Apply default
		ServicePort: common.InterfaceToStringList(m["service_port"]),
	}
}

// MakeJunosServicePortsSlice() makes a slice of JunosServicePorts
// nolint
func MakeJunosServicePortsSlice() []*JunosServicePorts {
	return []*JunosServicePorts{}
}

// InterfaceToJunosServicePortsSlice() makes a slice of JunosServicePorts
// nolint
func InterfaceToJunosServicePortsSlice(i interface{}) []*JunosServicePorts {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*JunosServicePorts{}
	for _, item := range list {
		result = append(result, InterfaceToJunosServicePorts(item))
	}
	return result
}
