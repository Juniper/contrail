package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeNetworkIpam makes NetworkIpam
// nolint
func MakeNetworkIpam() *NetworkIpam {
	return &NetworkIpam{
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
		NetworkIpamMGMT:      MakeIpamType(),
		IpamSubnets:          MakeIpamSubnets(),
		IpamSubnetMethod:     "",
	}
}

// MakeNetworkIpam makes NetworkIpam
// nolint
func InterfaceToNetworkIpam(i interface{}) *NetworkIpam {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &NetworkIpam{
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
		NetworkIpamMGMT:      InterfaceToIpamType(m["network_ipam_mgmt"]),
		IpamSubnets:          InterfaceToIpamSubnets(m["ipam_subnets"]),
		IpamSubnetMethod:     common.InterfaceToString(m["ipam_subnet_method"]),

		VirtualDNSRefs: InterfaceToNetworkIpamVirtualDNSRefs(m["virtual_DNS_refs"]),
	}
}

func InterfaceToNetworkIpamVirtualDNSRefs(i interface{}) []*NetworkIpamVirtualDNSRef {
	list, ok := i.([]interface{})
	if !ok {
		return nil
	}
	result := []*NetworkIpamVirtualDNSRef{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		_ = m
		if !ok {
			return nil
		}
		result = append(result, &NetworkIpamVirtualDNSRef{
			UUID: common.InterfaceToString(m["uuid"]),
			To:   common.InterfaceToStringList(m["to"]),
		})
	}

	return result
}

// MakeNetworkIpamSlice() makes a slice of NetworkIpam
// nolint
func MakeNetworkIpamSlice() []*NetworkIpam {
	return []*NetworkIpam{}
}

// InterfaceToNetworkIpamSlice() makes a slice of NetworkIpam
// nolint
func InterfaceToNetworkIpamSlice(i interface{}) []*NetworkIpam {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*NetworkIpam{}
	for _, item := range list {
		result = append(result, InterfaceToNetworkIpam(item))
	}
	return result
}
