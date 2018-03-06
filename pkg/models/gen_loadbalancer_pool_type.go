package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeLoadbalancerPoolType makes LoadbalancerPoolType
// nolint
func MakeLoadbalancerPoolType() *LoadbalancerPoolType {
	return &LoadbalancerPoolType{
		//TODO(nati): Apply default
		Status:                "",
		Protocol:              "",
		SubnetID:              "",
		SessionPersistence:    "",
		AdminState:            false,
		PersistenceCookieName: "",
		StatusDescription:     "",
		LoadbalancerMethod:    "",
	}
}

// MakeLoadbalancerPoolType makes LoadbalancerPoolType
// nolint
func InterfaceToLoadbalancerPoolType(i interface{}) *LoadbalancerPoolType {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &LoadbalancerPoolType{
		//TODO(nati): Apply default
		Status:                common.InterfaceToString(m["status"]),
		Protocol:              common.InterfaceToString(m["protocol"]),
		SubnetID:              common.InterfaceToString(m["subnet_id"]),
		SessionPersistence:    common.InterfaceToString(m["session_persistence"]),
		AdminState:            common.InterfaceToBool(m["admin_state"]),
		PersistenceCookieName: common.InterfaceToString(m["persistence_cookie_name"]),
		StatusDescription:     common.InterfaceToString(m["status_description"]),
		LoadbalancerMethod:    common.InterfaceToString(m["loadbalancer_method"]),
	}
}

// MakeLoadbalancerPoolTypeSlice() makes a slice of LoadbalancerPoolType
// nolint
func MakeLoadbalancerPoolTypeSlice() []*LoadbalancerPoolType {
	return []*LoadbalancerPoolType{}
}

// InterfaceToLoadbalancerPoolTypeSlice() makes a slice of LoadbalancerPoolType
// nolint
func InterfaceToLoadbalancerPoolTypeSlice(i interface{}) []*LoadbalancerPoolType {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*LoadbalancerPoolType{}
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerPoolType(item))
	}
	return result
}
