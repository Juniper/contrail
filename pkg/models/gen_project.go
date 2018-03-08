package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeProject makes Project
// nolint
func MakeProject() *Project {
	return &Project{
		//TODO(nati): Apply default
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		ConfigurationVersion: 0,
		VxlanRouting:         false,
		AlarmEnable:          false,
		Quota:                MakeQuotaType(),
	}
}

// MakeProject makes Project
// nolint
func InterfaceToProject(i interface{}) *Project {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Project{
		//TODO(nati): Apply default
		UUID:                 common.InterfaceToString(m["uuid"]),
		ParentUUID:           common.InterfaceToString(m["parent_uuid"]),
		ParentType:           common.InterfaceToString(m["parent_type"]),
		FQName:               common.InterfaceToStringList(m["fq_name"]),
		IDPerms:              InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:          common.InterfaceToString(m["display_name"]),
		Annotations:          InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:               InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion: common.InterfaceToInt64(m["configuration_version"]),
		VxlanRouting:         common.InterfaceToBool(m["vxlan_routing"]),
		AlarmEnable:          common.InterfaceToBool(m["alarm_enable"]),
		Quota:                InterfaceToQuotaType(m["quota"]),

		ApplicationPolicySetRefs: InterfaceToProjectApplicationPolicySetRefs(m["application_policy_set_refs"]),

		FloatingIPPoolRefs: InterfaceToProjectFloatingIPPoolRefs(m["floating_ip_pool_refs"]),

		AliasIPPoolRefs: InterfaceToProjectAliasIPPoolRefs(m["alias_ip_pool_refs"]),

		NamespaceRefs: InterfaceToProjectNamespaceRefs(m["namespace_refs"]),
	}
}

func InterfaceToProjectAliasIPPoolRefs(i interface{}) []*ProjectAliasIPPoolRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*ProjectAliasIPPoolRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &ProjectAliasIPPoolRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToProjectNamespaceRefs(i interface{}) []*ProjectNamespaceRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*ProjectNamespaceRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &ProjectNamespaceRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),

			Attr: InterfaceToSubnetType(m["attr"]),
		})
	}

	return result
}

func InterfaceToProjectApplicationPolicySetRefs(i interface{}) []*ProjectApplicationPolicySetRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*ProjectApplicationPolicySetRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &ProjectApplicationPolicySetRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

func InterfaceToProjectFloatingIPPoolRefs(i interface{}) []*ProjectFloatingIPPoolRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*ProjectFloatingIPPoolRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &ProjectFloatingIPPoolRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

// MakeProjectSlice() makes a slice of Project
// nolint
func MakeProjectSlice() []*Project {
	return []*Project{}
}

// InterfaceToProjectSlice() makes a slice of Project
// nolint
func InterfaceToProjectSlice(i interface{}) []*Project {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Project{}
	for _, item := range list {
		result = append(result, InterfaceToProject(item))
	}
	return result
}
