package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeFloatingIPPool makes FloatingIPPool
// nolint
func MakeFloatingIPPool() *FloatingIPPool {
	return &FloatingIPPool{
		//TODO(nati): Apply default
		UUID:                  "",
		ParentUUID:            "",
		ParentType:            "",
		FQName:                []string{},
		IDPerms:               MakeIdPermsType(),
		DisplayName:           "",
		Annotations:           MakeKeyValuePairs(),
		Perms2:                MakePermType2(),
		ConfigurationVersion:  0,
		FloatingIPPoolSubnets: MakeFloatingIpPoolSubnetType(),
	}
}

// MakeFloatingIPPool makes FloatingIPPool
// nolint
func InterfaceToFloatingIPPool(i interface{}) *FloatingIPPool {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &FloatingIPPool{
		//TODO(nati): Apply default
		UUID:                  common.InterfaceToString(m["uuid"]),
		ParentUUID:            common.InterfaceToString(m["parent_uuid"]),
		ParentType:            common.InterfaceToString(m["parent_type"]),
		FQName:                common.InterfaceToStringList(m["fq_name"]),
		IDPerms:               InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:           common.InterfaceToString(m["display_name"]),
		Annotations:           InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                InterfaceToPermType2(m["perms2"]),
		ConfigurationVersion:  common.InterfaceToInt64(m["configuration_version"]),
		FloatingIPPoolSubnets: InterfaceToFloatingIpPoolSubnetType(m["floating_ip_pool_subnets"]),
	}
}

// MakeFloatingIPPoolSlice() makes a slice of FloatingIPPool
// nolint
func MakeFloatingIPPoolSlice() []*FloatingIPPool {
	return []*FloatingIPPool{}
}

// InterfaceToFloatingIPPoolSlice() makes a slice of FloatingIPPool
// nolint
func InterfaceToFloatingIPPoolSlice(i interface{}) []*FloatingIPPool {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FloatingIPPool{}
	for _, item := range list {
		result = append(result, InterfaceToFloatingIPPool(item))
	}
	return result
}
