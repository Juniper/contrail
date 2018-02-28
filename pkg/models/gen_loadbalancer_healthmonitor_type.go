package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeLoadbalancerHealthmonitorType makes LoadbalancerHealthmonitorType
// nolint
func MakeLoadbalancerHealthmonitorType() *LoadbalancerHealthmonitorType {
	return &LoadbalancerHealthmonitorType{
		//TODO(nati): Apply default
		Delay:         0,
		ExpectedCodes: "",
		MaxRetries:    0,
		HTTPMethod:    "",
		AdminState:    false,
		Timeout:       0,
		URLPath:       "",
		MonitorType:   "",
	}
}

// MakeLoadbalancerHealthmonitorType makes LoadbalancerHealthmonitorType
// nolint
func InterfaceToLoadbalancerHealthmonitorType(i interface{}) *LoadbalancerHealthmonitorType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &LoadbalancerHealthmonitorType{
		//TODO(nati): Apply default
		Delay:         common.InterfaceToInt64(m["delay"]),
		ExpectedCodes: common.InterfaceToString(m["expected_codes"]),
		MaxRetries:    common.InterfaceToInt64(m["max_retries"]),
		HTTPMethod:    common.InterfaceToString(m["http_method"]),
		AdminState:    common.InterfaceToBool(m["admin_state"]),
		Timeout:       common.InterfaceToInt64(m["timeout"]),
		URLPath:       common.InterfaceToString(m["url_path"]),
		MonitorType:   common.InterfaceToString(m["monitor_type"]),
	}
}

// MakeLoadbalancerHealthmonitorTypeSlice() makes a slice of LoadbalancerHealthmonitorType
// nolint
func MakeLoadbalancerHealthmonitorTypeSlice() []*LoadbalancerHealthmonitorType {
	return []*LoadbalancerHealthmonitorType{}
}

// InterfaceToLoadbalancerHealthmonitorTypeSlice() makes a slice of LoadbalancerHealthmonitorType
// nolint
func InterfaceToLoadbalancerHealthmonitorTypeSlice(i interface{}) []*LoadbalancerHealthmonitorType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*LoadbalancerHealthmonitorType{}
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerHealthmonitorType(item))
	}
	return result
}
