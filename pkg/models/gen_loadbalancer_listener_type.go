package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeLoadbalancerListenerType makes LoadbalancerListenerType
// nolint
func MakeLoadbalancerListenerType() *LoadbalancerListenerType {
	return &LoadbalancerListenerType{
		//TODO(nati): Apply default
		DefaultTLSContainer: "",
		Protocol:            "",
		ConnectionLimit:     0,
		AdminState:          false,
		SniContainers:       []string{},
		ProtocolPort:        0,
	}
}

// MakeLoadbalancerListenerType makes LoadbalancerListenerType
// nolint
func InterfaceToLoadbalancerListenerType(i interface{}) *LoadbalancerListenerType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &LoadbalancerListenerType{
		//TODO(nati): Apply default
		DefaultTLSContainer: common.InterfaceToString(m["default_tls_container"]),
		Protocol:            common.InterfaceToString(m["protocol"]),
		ConnectionLimit:     common.InterfaceToInt64(m["connection_limit"]),
		AdminState:          common.InterfaceToBool(m["admin_state"]),
		SniContainers:       common.InterfaceToStringList(m["sni_containers"]),
		ProtocolPort:        common.InterfaceToInt64(m["protocol_port"]),
	}
}

// MakeLoadbalancerListenerTypeSlice() makes a slice of LoadbalancerListenerType
// nolint
func MakeLoadbalancerListenerTypeSlice() []*LoadbalancerListenerType {
	return []*LoadbalancerListenerType{}
}

// InterfaceToLoadbalancerListenerTypeSlice() makes a slice of LoadbalancerListenerType
// nolint
func InterfaceToLoadbalancerListenerTypeSlice(i interface{}) []*LoadbalancerListenerType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*LoadbalancerListenerType{}
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerListenerType(item))
	}
	return result
}
