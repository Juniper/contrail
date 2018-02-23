package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeServiceApplianceSet makes ServiceApplianceSet
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
func InterfaceToServiceApplianceSet(i interface{}) *ServiceApplianceSet {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceApplianceSet{
		//TODO(nati): Apply default
		UUID:        schema.InterfaceToString(m["uuid"]),
		ParentUUID:  schema.InterfaceToString(m["parent_uuid"]),
		ParentType:  schema.InterfaceToString(m["parent_type"]),
		FQName:      schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: schema.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
		ServiceApplianceSetProperties: InterfaceToKeyValuePairs(m["service_appliance_set_properties"]),
		ServiceApplianceHaMode:        schema.InterfaceToString(m["service_appliance_ha_mode"]),
		ServiceApplianceDriver:        schema.InterfaceToString(m["service_appliance_driver"]),
	}
}

// MakeServiceApplianceSetSlice() makes a slice of ServiceApplianceSet
func MakeServiceApplianceSetSlice() []*ServiceApplianceSet {
	return []*ServiceApplianceSet{}
}

// InterfaceToServiceApplianceSetSlice() makes a slice of ServiceApplianceSet
func InterfaceToServiceApplianceSetSlice(i interface{}) []*ServiceApplianceSet {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceApplianceSet{}
	for _, item := range list {
		result = append(result, InterfaceToServiceApplianceSet(item))
	}
	return result
}
