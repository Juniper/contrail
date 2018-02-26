package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeLoadbalancerListenerType makes LoadbalancerListenerType
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
func InterfaceToLoadbalancerListenerType(i interface{}) *LoadbalancerListenerType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &LoadbalancerListenerType{
		//TODO(nati): Apply default
		DefaultTLSContainer: schema.InterfaceToString(m["default_tls_container"]),
		Protocol:            schema.InterfaceToString(m["protocol"]),
		ConnectionLimit:     schema.InterfaceToInt64(m["connection_limit"]),
		AdminState:          schema.InterfaceToBool(m["admin_state"]),
		SniContainers:       schema.InterfaceToStringList(m["sni_containers"]),
		ProtocolPort:        schema.InterfaceToInt64(m["protocol_port"]),
	}
}

// MakeLoadbalancerListenerTypeSlice() makes a slice of LoadbalancerListenerType
func MakeLoadbalancerListenerTypeSlice() []*LoadbalancerListenerType {
	return []*LoadbalancerListenerType{}
}

// InterfaceToLoadbalancerListenerTypeSlice() makes a slice of LoadbalancerListenerType
func InterfaceToLoadbalancerListenerTypeSlice(i interface{}) []*LoadbalancerListenerType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*LoadbalancerListenerType{}
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerListenerType(item))
	}
	return result
}
