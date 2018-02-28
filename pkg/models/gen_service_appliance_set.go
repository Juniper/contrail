package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeServiceApplianceSet makes ServiceApplianceSet
// nolint
func MakeServiceApplianceSet() *ServiceApplianceSet {
	return &ServiceApplianceSet{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		ServiceApplianceSetProperties: MakeKeyValuePairs(),
		ServiceApplianceHaMode:        "",
		ServiceApplianceDriver:        "",
	}
}

// MakeServiceApplianceSet makes ServiceApplianceSet
// nolint
func InterfaceToServiceApplianceSet(i interface{}) *ServiceApplianceSet {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceApplianceSet{
		//TODO(nati): Apply default
		UUID:        common.InterfaceToString(m["uuid"]),
		ParentUUID:  common.InterfaceToString(m["parent_uuid"]),
		ParentType:  common.InterfaceToString(m["parent_type"]),
		FQName:      common.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: common.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
		ServiceApplianceSetProperties: InterfaceToKeyValuePairs(m["service_appliance_set_properties"]),
		ServiceApplianceHaMode:        common.InterfaceToString(m["service_appliance_ha_mode"]),
		ServiceApplianceDriver:        common.InterfaceToString(m["service_appliance_driver"]),
	}
}

// MakeServiceApplianceSetSlice() makes a slice of ServiceApplianceSet
// nolint
func MakeServiceApplianceSetSlice() []*ServiceApplianceSet {
	return []*ServiceApplianceSet{}
}

// InterfaceToServiceApplianceSetSlice() makes a slice of ServiceApplianceSet
// nolint
func InterfaceToServiceApplianceSetSlice(i interface{}) []*ServiceApplianceSet {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceApplianceSet{}
	for _, item := range list {
		result = append(result, InterfaceToServiceApplianceSet(item))
	}
	return result
}
