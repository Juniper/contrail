package models

// Loadbalancer

import "encoding/json"

// Loadbalancer
type Loadbalancer struct {
	Annotations            *KeyValuePairs    `json:"annotations,omitempty"`
	ParentUUID             string            `json:"parent_uuid,omitempty"`
	FQName                 []string          `json:"fq_name,omitempty"`
	LoadbalancerProperties *LoadbalancerType `json:"loadbalancer_properties,omitempty"`
	LoadbalancerProvider   string            `json:"loadbalancer_provider,omitempty"`
	IDPerms                *IdPermsType      `json:"id_perms,omitempty"`
	DisplayName            string            `json:"display_name,omitempty"`
	Perms2                 *PermType2        `json:"perms2,omitempty"`
	UUID                   string            `json:"uuid,omitempty"`
	ParentType             string            `json:"parent_type,omitempty"`

	VirtualMachineInterfaceRefs []*LoadbalancerVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
	ServiceInstanceRefs         []*LoadbalancerServiceInstanceRef         `json:"service_instance_refs,omitempty"`
	ServiceApplianceSetRefs     []*LoadbalancerServiceApplianceSetRef     `json:"service_appliance_set_refs,omitempty"`
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

// LoadbalancerServiceInstanceRef references each other
type LoadbalancerServiceInstanceRef struct {
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
		IDPerms:                MakeIdPermsType(),
		DisplayName:            "",
		Perms2:                 MakePermType2(),
		UUID:                   "",
		ParentType:             "",
		LoadbalancerProperties: MakeLoadbalancerType(),
		LoadbalancerProvider:   "",
		FQName:                 []string{},
		Annotations:            MakeKeyValuePairs(),
		ParentUUID:             "",
	}
}

// MakeLoadbalancerSlice() makes a slice of Loadbalancer
func MakeLoadbalancerSlice() []*Loadbalancer {
	return []*Loadbalancer{}
}
