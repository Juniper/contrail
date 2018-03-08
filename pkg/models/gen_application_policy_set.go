package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeApplicationPolicySet makes ApplicationPolicySet
// nolint
func MakeApplicationPolicySet() *ApplicationPolicySet {
	return &ApplicationPolicySet{
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
		AllApplications:      false,
	}
}

// MakeApplicationPolicySet makes ApplicationPolicySet
// nolint
func InterfaceToApplicationPolicySet(i interface{}) *ApplicationPolicySet {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &ApplicationPolicySet{
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
		AllApplications:      common.InterfaceToBool(m["all_applications"]),

		GlobalVrouterConfigRefs: InterfaceToApplicationPolicySetGlobalVrouterConfigRefs(m["global_vrouter_config_refs"]),

		FirewallPolicyRefs: InterfaceToApplicationPolicySetFirewallPolicyRefs(m["firewall_policy_refs"]),
	}
}

func InterfaceToApplicationPolicySetFirewallPolicyRefs(i interface{}) []*ApplicationPolicySetFirewallPolicyRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*ApplicationPolicySetFirewallPolicyRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &ApplicationPolicySetFirewallPolicyRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),

			Attr: InterfaceToFirewallSequence(m["attr"]),
		})
	}

	return result
}

func InterfaceToApplicationPolicySetGlobalVrouterConfigRefs(i interface{}) []*ApplicationPolicySetGlobalVrouterConfigRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*ApplicationPolicySetGlobalVrouterConfigRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &ApplicationPolicySetGlobalVrouterConfigRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

// MakeApplicationPolicySetSlice() makes a slice of ApplicationPolicySet
// nolint
func MakeApplicationPolicySetSlice() []*ApplicationPolicySet {
	return []*ApplicationPolicySet{}
}

// InterfaceToApplicationPolicySetSlice() makes a slice of ApplicationPolicySet
// nolint
func InterfaceToApplicationPolicySetSlice(i interface{}) []*ApplicationPolicySet {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*ApplicationPolicySet{}
	for _, item := range list {
		result = append(result, InterfaceToApplicationPolicySet(item))
	}
	return result
}
