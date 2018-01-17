package models

// LoadbalancerPool

import "encoding/json"

// LoadbalancerPool
type LoadbalancerPool struct {
	ParentUUID                       string                `json:"parent_uuid,omitempty"`
	FQName                           []string              `json:"fq_name,omitempty"`
	IDPerms                          *IdPermsType          `json:"id_perms,omitempty"`
	LoadbalancerPoolProperties       *LoadbalancerPoolType `json:"loadbalancer_pool_properties,omitempty"`
	LoadbalancerPoolCustomAttributes *KeyValuePairs        `json:"loadbalancer_pool_custom_attributes,omitempty"`
	ParentType                       string                `json:"parent_type,omitempty"`
	DisplayName                      string                `json:"display_name,omitempty"`
	Annotations                      *KeyValuePairs        `json:"annotations,omitempty"`
	Perms2                           *PermType2            `json:"perms2,omitempty"`
	LoadbalancerPoolProvider         string                `json:"loadbalancer_pool_provider,omitempty"`
	UUID                             string                `json:"uuid,omitempty"`

	VirtualMachineInterfaceRefs   []*LoadbalancerPoolVirtualMachineInterfaceRef   `json:"virtual_machine_interface_refs,omitempty"`
	LoadbalancerListenerRefs      []*LoadbalancerPoolLoadbalancerListenerRef      `json:"loadbalancer_listener_refs,omitempty"`
	ServiceInstanceRefs           []*LoadbalancerPoolServiceInstanceRef           `json:"service_instance_refs,omitempty"`
	LoadbalancerHealthmonitorRefs []*LoadbalancerPoolLoadbalancerHealthmonitorRef `json:"loadbalancer_healthmonitor_refs,omitempty"`
	ServiceApplianceSetRefs       []*LoadbalancerPoolServiceApplianceSetRef       `json:"service_appliance_set_refs,omitempty"`

	LoadbalancerMembers []*LoadbalancerMember `json:"loadbalancer_members,omitempty"`
}

// LoadbalancerPoolVirtualMachineInterfaceRef references each other
type LoadbalancerPoolVirtualMachineInterfaceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LoadbalancerPoolLoadbalancerListenerRef references each other
type LoadbalancerPoolLoadbalancerListenerRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LoadbalancerPoolServiceInstanceRef references each other
type LoadbalancerPoolServiceInstanceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LoadbalancerPoolLoadbalancerHealthmonitorRef references each other
type LoadbalancerPoolLoadbalancerHealthmonitorRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// LoadbalancerPoolServiceApplianceSetRef references each other
type LoadbalancerPoolServiceApplianceSetRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *LoadbalancerPool) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLoadbalancerPool makes LoadbalancerPool
func MakeLoadbalancerPool() *LoadbalancerPool {
	return &LoadbalancerPool{
		//TODO(nati): Apply default
		IDPerms:                          MakeIdPermsType(),
		LoadbalancerPoolProperties:       MakeLoadbalancerPoolType(),
		LoadbalancerPoolCustomAttributes: MakeKeyValuePairs(),
		ParentUUID:                       "",
		FQName:                           []string{},
		Annotations:                      MakeKeyValuePairs(),
		Perms2:                           MakePermType2(),
		LoadbalancerPoolProvider: "",
		UUID:        "",
		ParentType:  "",
		DisplayName: "",
	}
}

// MakeLoadbalancerPoolSlice() makes a slice of LoadbalancerPool
func MakeLoadbalancerPoolSlice() []*LoadbalancerPool {
	return []*LoadbalancerPool{}
}
