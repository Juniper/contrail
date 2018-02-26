package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeLoadbalancerHealthmonitorType makes LoadbalancerHealthmonitorType
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
func InterfaceToLoadbalancerHealthmonitorType(i interface{}) *LoadbalancerHealthmonitorType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &LoadbalancerHealthmonitorType{
		//TODO(nati): Apply default
		Delay:         schema.InterfaceToInt64(m["delay"]),
		ExpectedCodes: schema.InterfaceToString(m["expected_codes"]),
		MaxRetries:    schema.InterfaceToInt64(m["max_retries"]),
		HTTPMethod:    schema.InterfaceToString(m["http_method"]),
		AdminState:    schema.InterfaceToBool(m["admin_state"]),
		Timeout:       schema.InterfaceToInt64(m["timeout"]),
		URLPath:       schema.InterfaceToString(m["url_path"]),
		MonitorType:   schema.InterfaceToString(m["monitor_type"]),
	}
}

// MakeLoadbalancerHealthmonitorTypeSlice() makes a slice of LoadbalancerHealthmonitorType
func MakeLoadbalancerHealthmonitorTypeSlice() []*LoadbalancerHealthmonitorType {
	return []*LoadbalancerHealthmonitorType{}
}

// InterfaceToLoadbalancerHealthmonitorTypeSlice() makes a slice of LoadbalancerHealthmonitorType
func InterfaceToLoadbalancerHealthmonitorTypeSlice(i interface{}) []*LoadbalancerHealthmonitorType {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*LoadbalancerHealthmonitorType{}
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerHealthmonitorType(item))
	}
	return result
}
