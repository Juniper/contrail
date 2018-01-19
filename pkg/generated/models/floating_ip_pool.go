package models

// FloatingIPPool

import "encoding/json"

// FloatingIPPool
type FloatingIPPool struct {
	FloatingIPPoolSubnets *FloatingIpPoolSubnetType `json:"floating_ip_pool_subnets,omitempty"`
	ParentUUID            string                    `json:"parent_uuid,omitempty"`
	ParentType            string                    `json:"parent_type,omitempty"`
	FQName                []string                  `json:"fq_name,omitempty"`
	UUID                  string                    `json:"uuid,omitempty"`
	IDPerms               *IdPermsType              `json:"id_perms,omitempty"`
	DisplayName           string                    `json:"display_name,omitempty"`
	Annotations           *KeyValuePairs            `json:"annotations,omitempty"`
	Perms2                *PermType2                `json:"perms2,omitempty"`

	FloatingIPs []*FloatingIP `json:"floating_ips,omitempty"`
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
		UUID:                  "",
		IDPerms:               MakeIdPermsType(),
		DisplayName:           "",
		Annotations:           MakeKeyValuePairs(),
		Perms2:                MakePermType2(),
		FloatingIPPoolSubnets: MakeFloatingIpPoolSubnetType(),
		ParentUUID:            "",
		ParentType:            "",
		FQName:                []string{},
	}
}

// MakeFloatingIPPoolSlice() makes a slice of FloatingIPPool
func MakeFloatingIPPoolSlice() []*FloatingIPPool {
	return []*FloatingIPPool{}
}
