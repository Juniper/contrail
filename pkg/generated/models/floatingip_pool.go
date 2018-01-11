package models

// FloatingIPPool

import "encoding/json"

// FloatingIPPool
type FloatingIPPool struct {
	DisplayName           string                    `json:"display_name"`
	Annotations           *KeyValuePairs            `json:"annotations"`
	Perms2                *PermType2                `json:"perms2"`
	FQName                []string                  `json:"fq_name"`
	ParentUUID            string                    `json:"parent_uuid"`
	ParentType            string                    `json:"parent_type"`
	IDPerms               *IdPermsType              `json:"id_perms"`
	FloatingIPPoolSubnets *FloatingIpPoolSubnetType `json:"floating_ip_pool_subnets"`
	UUID                  string                    `json:"uuid"`

	FloatingIPs []*FloatingIP `json:"floating_ips"`
}

// String returns json representation of the object
func (model *FloatingIPPool) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeFloatingIPPool makes FloatingIPPool
func MakeFloatingIPPool() *FloatingIPPool {
	return &FloatingIPPool{
		//TODO(nati): Apply default
		Perms2:                MakePermType2(),
		FQName:                []string{},
		DisplayName:           "",
		Annotations:           MakeKeyValuePairs(),
		IDPerms:               MakeIdPermsType(),
		FloatingIPPoolSubnets: MakeFloatingIpPoolSubnetType(),
		UUID:       "",
		ParentUUID: "",
		ParentType: "",
	}
}

// MakeFloatingIPPoolSlice() makes a slice of FloatingIPPool
func MakeFloatingIPPoolSlice() []*FloatingIPPool {
	return []*FloatingIPPool{}
}
