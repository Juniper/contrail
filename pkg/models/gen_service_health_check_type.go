package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeServiceHealthCheckType makes ServiceHealthCheckType
// nolint
func MakeServiceHealthCheckType() *ServiceHealthCheckType {
	return &ServiceHealthCheckType{
		//TODO(nati): Apply default
		DelayUsecs:      0,
		TimeoutUsecs:    0,
		Enabled:         false,
		Delay:           0,
		ExpectedCodes:   "",
		MaxRetries:      0,
		HealthCheckType: "",
		HTTPMethod:      "",
		Timeout:         0,
		URLPath:         "",
		MonitorType:     "",
	}
}

// MakeServiceHealthCheckType makes ServiceHealthCheckType
// nolint
func InterfaceToServiceHealthCheckType(i interface{}) *ServiceHealthCheckType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceHealthCheckType{
		//TODO(nati): Apply default
		DelayUsecs:      common.InterfaceToInt64(m["delayUsecs"]),
		TimeoutUsecs:    common.InterfaceToInt64(m["timeoutUsecs"]),
		Enabled:         common.InterfaceToBool(m["enabled"]),
		Delay:           common.InterfaceToInt64(m["delay"]),
		ExpectedCodes:   common.InterfaceToString(m["expected_codes"]),
		MaxRetries:      common.InterfaceToInt64(m["max_retries"]),
		HealthCheckType: common.InterfaceToString(m["health_check_type"]),
		HTTPMethod:      common.InterfaceToString(m["http_method"]),
		Timeout:         common.InterfaceToInt64(m["timeout"]),
		URLPath:         common.InterfaceToString(m["url_path"]),
		MonitorType:     common.InterfaceToString(m["monitor_type"]),
	}
}

// MakeServiceHealthCheckTypeSlice() makes a slice of ServiceHealthCheckType
// nolint
func MakeServiceHealthCheckTypeSlice() []*ServiceHealthCheckType {
	return []*ServiceHealthCheckType{}
}

// InterfaceToServiceHealthCheckTypeSlice() makes a slice of ServiceHealthCheckType
// nolint
func InterfaceToServiceHealthCheckTypeSlice(i interface{}) []*ServiceHealthCheckType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceHealthCheckType{}
	for _, item := range list {
		result = append(result, InterfaceToServiceHealthCheckType(item))
	}
	return result
}
