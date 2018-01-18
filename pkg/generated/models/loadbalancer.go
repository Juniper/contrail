package models

// Loadbalancer

import "encoding/json"

// Loadbalancer
type Loadbalancer struct {
	UUID                   string            `json:"uuid,omitempty"`
	LoadbalancerProvider   string            `json:"loadbalancer_provider,omitempty"`
	FQName                 []string          `json:"fq_name,omitempty"`
	DisplayName            string            `json:"display_name,omitempty"`
	Perms2                 *PermType2        `json:"perms2,omitempty"`
	Annotations            *KeyValuePairs    `json:"annotations,omitempty"`
	LoadbalancerProperties *LoadbalancerType `json:"loadbalancer_properties,omitempty"`
	ParentUUID             string            `json:"parent_uuid,omitempty"`
	ParentType             string            `json:"parent_type,omitempty"`
	IDPerms                *IdPermsType      `json:"id_perms,omitempty"`

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

// String returns json representation of the object
func (model *Loadbalancer) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLoadbalancer makes Loadbalancer
func MakeLoadbalancer() *Loadbalancer {
	return &Loadbalancer{
		//TODO(nati): Apply default
		LoadbalancerProperties: MakeLoadbalancerType(),
		ParentUUID:             "",
		ParentType:             "",
		IDPerms:                MakeIdPermsType(),
		Annotations:            MakeKeyValuePairs(),
		LoadbalancerProvider:   "",
		FQName:                 []string{},
		DisplayName:            "",
		Perms2:                 MakePermType2(),
		UUID:                   "",
	}
}

// MakeLoadbalancerSlice() makes a slice of Loadbalancer
func MakeLoadbalancerSlice() []*Loadbalancer {
	return []*Loadbalancer{}
}
