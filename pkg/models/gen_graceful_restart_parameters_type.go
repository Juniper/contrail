package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeGracefulRestartParametersType makes GracefulRestartParametersType
// nolint
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
// nolint
func InterfaceToGracefulRestartParametersType(i interface{}) *GracefulRestartParametersType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &GracefulRestartParametersType{
		//TODO(nati): Apply default
		Enable:               common.InterfaceToBool(m["enable"]),
		EndOfRibTimeout:      common.InterfaceToInt64(m["end_of_rib_timeout"]),
		BGPHelperEnable:      common.InterfaceToBool(m["bgp_helper_enable"]),
		XMPPHelperEnable:     common.InterfaceToBool(m["xmpp_helper_enable"]),
		RestartTime:          common.InterfaceToInt64(m["restart_time"]),
		LongLivedRestartTime: common.InterfaceToInt64(m["long_lived_restart_time"]),
	}
}

// MakeGracefulRestartParametersTypeSlice() makes a slice of GracefulRestartParametersType
// nolint
func MakeGracefulRestartParametersTypeSlice() []*GracefulRestartParametersType {
	return []*GracefulRestartParametersType{}
}

// InterfaceToGracefulRestartParametersTypeSlice() makes a slice of GracefulRestartParametersType
// nolint
func InterfaceToGracefulRestartParametersTypeSlice(i interface{}) []*GracefulRestartParametersType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*GracefulRestartParametersType{}
	for _, item := range list {
		result = append(result, InterfaceToGracefulRestartParametersType(item))
	}
	return result
}
