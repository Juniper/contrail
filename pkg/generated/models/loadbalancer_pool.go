package models

// LoadbalancerPool

import "encoding/json"

// LoadbalancerPool
type LoadbalancerPool struct {
	LoadbalancerPoolCustomAttributes *KeyValuePairs        `json:"loadbalancer_pool_custom_attributes"`
	LoadbalancerPoolProvider         string                `json:"loadbalancer_pool_provider"`
	Perms2                           *PermType2            `json:"perms2"`
	UUID                             string                `json:"uuid"`
	LoadbalancerPoolProperties       *LoadbalancerPoolType `json:"loadbalancer_pool_properties"`
	Annotations                      *KeyValuePairs        `json:"annotations"`
	ParentUUID                       string                `json:"parent_uuid"`
	ParentType                       string                `json:"parent_type"`
	FQName                           []string              `json:"fq_name"`
	IDPerms                          *IdPermsType          `json:"id_perms"`
	DisplayName                      string                `json:"display_name"`

	ServiceInstanceRefs           []*LoadbalancerPoolServiceInstanceRef           `json:"service_instance_refs"`
	LoadbalancerHealthmonitorRefs []*LoadbalancerPoolLoadbalancerHealthmonitorRef `json:"loadbalancer_healthmonitor_refs"`
	ServiceApplianceSetRefs       []*LoadbalancerPoolServiceApplianceSetRef       `json:"service_appliance_set_refs"`
	VirtualMachineInterfaceRefs   []*LoadbalancerPoolVirtualMachineInterfaceRef   `json:"virtual_machine_interface_refs"`
	LoadbalancerListenerRefs      []*LoadbalancerPoolLoadbalancerListenerRef      `json:"loadbalancer_listener_refs"`

	LoadbalancerMembers []*LoadbalancerMember `json:"loadbalancer_members"`
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

// LoadbalancerPoolLoadbalancerHealthmonitorRef references each other
type LoadbalancerPoolLoadbalancerHealthmonitorRef struct {
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
		LoadbalancerPoolCustomAttributes: MakeKeyValuePairs(),
		LoadbalancerPoolProvider:         "",
		Perms2:                     MakePermType2(),
		UUID:                       "",
		FQName:                     []string{},
		IDPerms:                    MakeIdPermsType(),
		DisplayName:                "",
		LoadbalancerPoolProperties: MakeLoadbalancerPoolType(),
		Annotations:                MakeKeyValuePairs(),
		ParentUUID:                 "",
		ParentType:                 "",
	}
}

// InterfaceToLoadbalancerPool makes LoadbalancerPool from interface
func InterfaceToLoadbalancerPool(iData interface{}) *LoadbalancerPool {
	data := iData.(map[string]interface{})
	return &LoadbalancerPool{
		LoadbalancerPoolProperties: InterfaceToLoadbalancerPoolType(data["loadbalancer_pool_properties"]),

		//{"description":"Configuration for loadbalancer pool like protocol, subnet, etc.","type":"object","properties":{"admin_state":{"type":"boolean"},"loadbalancer_method":{"type":"string","enum":["ROUND_ROBIN","LEAST_CONNECTIONS","SOURCE_IP"]},"persistence_cookie_name":{"type":"string"},"protocol":{"type":"string","enum":["HTTP","HTTPS","TCP","UDP","TERMINATED_HTTPS"]},"session_persistence":{"type":"string","enum":["SOURCE_IP","HTTP_COOKIE","APP_COOKIE"]},"status":{"type":"string"},"status_description":{"type":"string"},"subnet_id":{"type":"string"}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		LoadbalancerPoolCustomAttributes: InterfaceToKeyValuePairs(data["loadbalancer_pool_custom_attributes"]),

		//{"description":"Custom loadbalancer config, opaque to the system. Specified as list of Key:Value pairs. Applicable to LBaaS V1.","type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		LoadbalancerPoolProvider: data["loadbalancer_pool_provider"].(string),

		//{"description":"Provider field selects backend provider of the LBaaS, Cloudadmin could offer different levels of service like gold, silver, bronze. Provided by  HA-proxy or various HW or SW appliances in the backend. Applicable to LBaaS V1","type":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}

	}
}

// InterfaceToLoadbalancerPoolSlice makes a slice of LoadbalancerPool from interface
func InterfaceToLoadbalancerPoolSlice(data interface{}) []*LoadbalancerPool {
	list := data.([]interface{})
	result := MakeLoadbalancerPoolSlice()
	for _, item := range list {
		result = append(result, InterfaceToLoadbalancerPool(item))
	}
	return result
}

// MakeLoadbalancerPoolSlice() makes a slice of LoadbalancerPool
func MakeLoadbalancerPoolSlice() []*LoadbalancerPool {
	return []*LoadbalancerPool{}
}
