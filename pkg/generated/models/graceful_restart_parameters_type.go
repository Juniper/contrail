package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeGracefulRestartParametersType makes GracefulRestartParametersType
func MakeGracefulRestartParametersType() *GracefulRestartParametersType {
	return &GracefulRestartParametersType{
		//TODO(nati): Apply default
		Enable:               false,
		EndOfRibTimeout:      0,
		BGPHelperEnable:      false,
		XMPPHelperEnable:     false,
		RestartTime:          0,
		LongLivedRestartTime: 0,
	}
}

// MakeGracefulRestartParametersType makes GracefulRestartParametersType
func InterfaceToGracefulRestartParametersType(i interface{}) *GracefulRestartParametersType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &GracefulRestartParametersType{
		//TODO(nati): Apply default
		Enable:               schema.InterfaceToBool(m["enable"]),
		EndOfRibTimeout:      schema.InterfaceToInt64(m["end_of_rib_timeout"]),
		BGPHelperEnable:      schema.InterfaceToBool(m["bgp_helper_enable"]),
		XMPPHelperEnable:     schema.InterfaceToBool(m["xmpp_helper_enable"]),
		RestartTime:          schema.InterfaceToInt64(m["restart_time"]),
		LongLivedRestartTime: schema.InterfaceToInt64(m["long_lived_restart_time"]),
	}
}

// MakeGracefulRestartParametersTypeSlice() makes a slice of GracefulRestartParametersType
func MakeGracefulRestartParametersTypeSlice() []*GracefulRestartParametersType {
	return []*GracefulRestartParametersType{}
}

// InterfaceToGracefulRestartParametersTypeSlice() makes a slice of GracefulRestartParametersType
func InterfaceToGracefulRestartParametersTypeSlice(i interface{}) []*GracefulRestartParametersType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*GracefulRestartParametersType{}
	for _, item := range list {
		result = append(result, InterfaceToGracefulRestartParametersType(item))
	}
	return result
}
