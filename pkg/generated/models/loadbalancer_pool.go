package models

// LoadbalancerPool

import "encoding/json"

// LoadbalancerPool
type LoadbalancerPool struct {
	LoadbalancerPoolProperties       *LoadbalancerPoolType `json:"loadbalancer_pool_properties"`
	LoadbalancerPoolProvider         string                `json:"loadbalancer_pool_provider"`
	IDPerms                          *IdPermsType          `json:"id_perms"`
	FQName                           []string              `json:"fq_name"`
	DisplayName                      string                `json:"display_name"`
	Annotations                      *KeyValuePairs        `json:"annotations"`
	LoadbalancerPoolCustomAttributes *KeyValuePairs        `json:"loadbalancer_pool_custom_attributes"`
	Perms2                           *PermType2            `json:"perms2"`
	UUID                             string                `json:"uuid"`
	ParentUUID                       string                `json:"parent_uuid"`
	ParentType                       string                `json:"parent_type"`

	LoadbalancerListenerRefs      []*LoadbalancerPoolLoadbalancerListenerRef      `json:"loadbalancer_listener_refs"`
	ServiceInstanceRefs           []*LoadbalancerPoolServiceInstanceRef           `json:"service_instance_refs"`
	LoadbalancerHealthmonitorRefs []*LoadbalancerPoolLoadbalancerHealthmonitorRef `json:"loadbalancer_healthmonitor_refs"`
	ServiceApplianceSetRefs       []*LoadbalancerPoolServiceApplianceSetRef       `json:"service_appliance_set_refs"`
	VirtualMachineInterfaceRefs   []*LoadbalancerPoolVirtualMachineInterfaceRef   `json:"virtual_machine_interface_refs"`

	LoadbalancerMembers []*LoadbalancerMember `json:"loadbalancer_members"`
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

// String returns json representation of the object
func (model *LoadbalancerPool) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLoadbalancerPool makes LoadbalancerPool
func MakeLoadbalancerPool() *LoadbalancerPool {
	return &LoadbalancerPool{
		//TODO(nati): Apply default
		LoadbalancerPoolProperties: MakeLoadbalancerPoolType(),
		LoadbalancerPoolProvider:   "",
		IDPerms:                    MakeIdPermsType(),
		LoadbalancerPoolCustomAttributes: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
	}
}

// MakeLoadbalancerPoolSlice() makes a slice of LoadbalancerPool
func MakeLoadbalancerPoolSlice() []*LoadbalancerPool {
	return []*LoadbalancerPool{}
}
