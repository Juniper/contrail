package models

// Loadbalancer

// Loadbalancer
//proteus:generate
type Loadbalancer struct {
	UUID                   string            `json:"uuid,omitempty"`
	ParentUUID             string            `json:"parent_uuid,omitempty"`
	ParentType             string            `json:"parent_type,omitempty"`
	FQName                 []string          `json:"fq_name,omitempty"`
	IDPerms                *IdPermsType      `json:"id_perms,omitempty"`
	DisplayName            string            `json:"display_name,omitempty"`
	Annotations            *KeyValuePairs    `json:"annotations,omitempty"`
	Perms2                 *PermType2        `json:"perms2,omitempty"`
	LoadbalancerProperties *LoadbalancerType `json:"loadbalancer_properties,omitempty"`
	LoadbalancerProvider   string            `json:"loadbalancer_provider,omitempty"`

	ServiceApplianceSetRefs     []*LoadbalancerServiceApplianceSetRef     `json:"service_appliance_set_refs,omitempty"`
	VirtualMachineInterfaceRefs []*LoadbalancerVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
	ServiceInstanceRefs         []*LoadbalancerServiceInstanceRef         `json:"service_instance_refs,omitempty"`
}

// LoadbalancerServiceInstanceRef references each other
type LoadbalancerServiceInstanceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LoadbalancerServiceApplianceSetRef references each other
type LoadbalancerServiceApplianceSetRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LoadbalancerVirtualMachineInterfaceRef references each other
type LoadbalancerVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// MakeLoadbalancer makes Loadbalancer
func MakeLoadbalancer() *Loadbalancer {
	return &Loadbalancer{
		//TODO(nati): Apply default
		UUID:                   "",
		ParentUUID:             "",
		ParentType:             "",
		FQName:                 []string{},
		IDPerms:                MakeIdPermsType(),
		DisplayName:            "",
		Annotations:            MakeKeyValuePairs(),
		Perms2:                 MakePermType2(),
		LoadbalancerProperties: MakeLoadbalancerType(),
		LoadbalancerProvider:   "",
	}
}

// MakeLoadbalancerSlice() makes a slice of Loadbalancer
func MakeLoadbalancerSlice() []*Loadbalancer {
	return []*Loadbalancer{}
}
