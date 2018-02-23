package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeProtocolType makes ProtocolType
func MakeProtocolType() *ProtocolType {
	return &ProtocolType{
		//TODO(nati): Apply default
		Protocol: "",
		Port:     0,
	}
}

// MakeProtocolType makes ProtocolType
func InterfaceToProtocolType(i interface{}) *ProtocolType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ProtocolType{
		//TODO(nati): Apply default
		Protocol: schema.InterfaceToString(m["protocol"]),
		Port:     schema.InterfaceToInt64(m["port"]),
	}
}

// MakeProtocolTypeSlice() makes a slice of ProtocolType
func MakeProtocolTypeSlice() []*ProtocolType {
	return []*ProtocolType{}
}

// InterfaceToProtocolTypeSlice() makes a slice of ProtocolType
func InterfaceToProtocolTypeSlice(i interface{}) []*ProtocolType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ProtocolType{}
	for _, item := range list {
		result = append(result, InterfaceToProtocolType(item))
	}
	return result
}
