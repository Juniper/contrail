package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeServiceAppliance makes ServiceAppliance
// nolint
func MakeServiceAppliance() *ServiceAppliance {
	return &ServiceAppliance{
		//TODO(nati): Apply default
		UUID:                            "",
		ParentUUID:                      "",
		ParentType:                      "",
		FQName:                          []string{},
		IDPerms:                         MakeIdPermsType(),
		DisplayName:                     "",
		Annotations:                     MakeKeyValuePairs(),
		Perms2:                          MakePermType2(),
		ConfigurationVersion:            0,
		ServiceApplianceUserCredentials: MakeUserCredentials(),
		ServiceApplianceIPAddress:       "",
		ServiceApplianceProperties:      MakeKeyValuePairs(),
	}
}

// MakeServiceAppliance makes ServiceAppliance
// nolint
func InterfaceToServiceAppliance(i interface{}) *ServiceAppliance {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ServiceAppliance{
		//TODO(nati): Apply default
		UUID:                            common.InterfaceToString(m["uuid"]),
		ParentUUID:                      common.InterfaceToString(m["parent_uuid"]),
		ParentType:                      common.InterfaceToString(m["parent_type"]),
		FQName:                          common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                         InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:                     common.InterfaceToString(m["display_name"]),
		Annotations:                     InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                          InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:            common.InterfaceToInt64(m["configuration_version"]),
		ServiceApplianceUserCredentials: InterfaceToUserCredentials(m["service_appliance_user_credentials"]),
		ServiceApplianceIPAddress:       common.InterfaceToString(m["service_appliance_ip_address"]),
		ServiceApplianceProperties:      InterfaceToKeyValuePairs(m["service_appliance_properties"]),
	}
}

// MakeServiceApplianceSlice() makes a slice of ServiceAppliance
// nolint
func MakeServiceApplianceSlice() []*ServiceAppliance {
	return []*ServiceAppliance{}
}

// InterfaceToServiceApplianceSlice() makes a slice of ServiceAppliance
// nolint
func InterfaceToServiceApplianceSlice(i interface{}) []*ServiceAppliance {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ServiceAppliance{}
	for _, item := range list {
		result = append(result, InterfaceToServiceAppliance(item))
	}
	return result
}
