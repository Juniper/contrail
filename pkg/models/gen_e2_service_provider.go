package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeE2ServiceProvider makes E2ServiceProvider
// nolint
func MakeE2ServiceProvider() *E2ServiceProvider {
	return &E2ServiceProvider{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		E2ServiceProviderPromiscuous: false,
	}
}

// MakeE2ServiceProvider makes E2ServiceProvider
// nolint
func InterfaceToE2ServiceProvider(i interface{}) *E2ServiceProvider {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &E2ServiceProvider{
		//TODO(nati): Apply default
		UUID:        common.InterfaceToString(m["uuid"]),
		ParentUUID:  common.InterfaceToString(m["parent_uuid"]),
		ParentType:  common.InterfaceToString(m["parent_type"]),
		FQName:      common.InterfaceToStringList(m["fq_name"]),
		IDPerms:     InterfaceToIdPermsType(m["id_perms"]),
		DisplayName: common.InterfaceToString(m["display_name"]),
		Annotations: InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:      InterfaceToPermType2(m["perms2"]),
		E2ServiceProviderPromiscuous: common.InterfaceToBool(m["e2_service_provider_promiscuous"]),
	}
}

// MakeE2ServiceProviderSlice() makes a slice of E2ServiceProvider
// nolint
func MakeE2ServiceProviderSlice() []*E2ServiceProvider {
	return []*E2ServiceProvider{}
}

// InterfaceToE2ServiceProviderSlice() makes a slice of E2ServiceProvider
// nolint
func InterfaceToE2ServiceProviderSlice(i interface{}) []*E2ServiceProvider {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*E2ServiceProvider{}
	for _, item := range list {
		result = append(result, InterfaceToE2ServiceProvider(item))
	}
	return result
}
