package models

import (
	"github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version

// MakeFloatingIPPool makes FloatingIPPool
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
		FloatingIPPoolSubnets: MakeFloatingIpPoolSubnetType(),
	}
}

// MakeFloatingIPPool makes FloatingIPPool
func InterfaceToFloatingIPPool(i interface{}) *FloatingIPPool {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &FloatingIPPool{
		//TODO(nati): Apply default
		UUID:                  schema.InterfaceToString(m["uuid"]),
		ParentUUID:            schema.InterfaceToString(m["parent_uuid"]),
		ParentType:            schema.InterfaceToString(m["parent_type"]),
		FQName:                schema.InterfaceToStringList(m["fq_name"]),
		IDPerms:               InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:           schema.InterfaceToString(m["display_name"]),
		Annotations:           InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                InterfaceToPermType2(m["perms2"]),
		FloatingIPPoolSubnets: InterfaceToFloatingIpPoolSubnetType(m["floating_ip_pool_subnets"]),
	}
}

// MakeFloatingIPPoolSlice() makes a slice of FloatingIPPool
func MakeFloatingIPPoolSlice() []*FloatingIPPool {
	return []*FloatingIPPool{}
}

// InterfaceToFloatingIPPoolSlice() makes a slice of FloatingIPPool
func InterfaceToFloatingIPPoolSlice(i interface{}) []*FloatingIPPool {
	list := schema.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*FloatingIPPool{}
	for _, item := range list {
		result = append(result, InterfaceToFloatingIPPool(item))
	}
	return result
}
