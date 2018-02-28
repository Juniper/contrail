package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeProtocolType makes ProtocolType
// nolint
func MakeProtocolType() *ProtocolType {
	return &ProtocolType{
		//TODO(nati): Apply default
		Protocol: "",
		Port:     0,
	}
}

// MakeProtocolType makes ProtocolType
// nolint
func InterfaceToProtocolType(i interface{}) *ProtocolType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ProtocolType{
		//TODO(nati): Apply default
		Protocol: common.InterfaceToString(m["protocol"]),
		Port:     common.InterfaceToInt64(m["port"]),
	}
}

// MakeProtocolTypeSlice() makes a slice of ProtocolType
// nolint
func MakeProtocolTypeSlice() []*ProtocolType {
	return []*ProtocolType{}
}

// InterfaceToProtocolTypeSlice() makes a slice of ProtocolType
// nolint
func InterfaceToProtocolTypeSlice(i interface{}) []*ProtocolType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ProtocolType{}
	for _, item := range list {
		result = append(result, InterfaceToProtocolType(item))
	}
	return result
}
