package models

// Loadbalancer

import "encoding/json"

// Loadbalancer
type Loadbalancer struct {
	IDPerms                *IdPermsType      `json:"id_perms"`
	Perms2                 *PermType2        `json:"perms2"`
	UUID                   string            `json:"uuid"`
	LoadbalancerProperties *LoadbalancerType `json:"loadbalancer_properties"`
	ParentUUID             string            `json:"parent_uuid"`
	ParentType             string            `json:"parent_type"`
	FQName                 []string          `json:"fq_name"`
	DisplayName            string            `json:"display_name"`
	Annotations            *KeyValuePairs    `json:"annotations"`
	LoadbalancerProvider   string            `json:"loadbalancer_provider"`

	ServiceInstanceRefs         []*LoadbalancerServiceInstanceRef         `json:"service_instance_refs"`
	ServiceApplianceSetRefs     []*LoadbalancerServiceApplianceSetRef     `json:"service_appliance_set_refs"`
	VirtualMachineInterfaceRefs []*LoadbalancerVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs"`
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
		LoadbalancerProperties: MakeLoadbalancerType(),
		IDPerms:                MakeIdPermsType(),
		Perms2:                 MakePermType2(),
		UUID:                   "",
		DisplayName:            "",
		Annotations:            MakeKeyValuePairs(),
		LoadbalancerProvider:   "",
		ParentUUID:             "",
		ParentType:             "",
		FQName:                 []string{},
	}
}

// MakeLoadbalancerSlice() makes a slice of Loadbalancer
func MakeLoadbalancerSlice() []*Loadbalancer {
	return []*Loadbalancer{}
}
