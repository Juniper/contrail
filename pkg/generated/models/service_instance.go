package models

// ServiceInstance

import "encoding/json"

// ServiceInstance
type ServiceInstance struct {
	IDPerms                   *IdPermsType         `json:"id_perms"`
	DisplayName               string               `json:"display_name"`
	Annotations               *KeyValuePairs       `json:"annotations"`
	UUID                      string               `json:"uuid"`
	ServiceInstanceProperties *ServiceInstanceType `json:"service_instance_properties"`
	ParentType                string               `json:"parent_type"`
	FQName                    []string             `json:"fq_name"`
	Perms2                    *PermType2           `json:"perms2"`
	ParentUUID                string               `json:"parent_uuid"`
	ServiceInstanceBindings   *KeyValuePairs       `json:"service_instance_bindings"`

	InstanceIPRefs      []*ServiceInstanceInstanceIPRef      `json:"instance_ip_refs"`
	ServiceTemplateRefs []*ServiceInstanceServiceTemplateRef `json:"service_template_refs"`

	PortTuples []*PortTuple `json:"port_tuples"`
}

// ServiceInstanceServiceTemplateRef references each other
type ServiceInstanceServiceTemplateRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// ServiceInstanceInstanceIPRef references each other
type ServiceInstanceInstanceIPRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *ServiceInterfaceTag
}

// String returns json representation of the object
func (model *ServiceInstance) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceInstance makes ServiceInstance
func MakeServiceInstance() *ServiceInstance {
	return &ServiceInstance{
		//TODO(nati): Apply default
		UUID: "",
		ServiceInstanceProperties: MakeServiceInstanceType(),
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		ParentUUID:                "",
		ServiceInstanceBindings:   MakeKeyValuePairs(),
		ParentType:                "",
		FQName:                    []string{},
		Perms2:                    MakePermType2(),
	}
}

// InterfaceToServiceInstance makes ServiceInstance from interface
func InterfaceToServiceInstance(iData interface{}) *ServiceInstance {
	data := iData.(map[string]interface{})
	return &ServiceInstance{
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ServiceInstanceBindings: InterfaceToKeyValuePairs(data["service_instance_bindings"]),

		//{"description":"Opaque key value pair for generating config for the service instance.","type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ServiceInstanceProperties: InterfaceToServiceInstanceType(data["service_instance_properties"]),

		//{"description":"Service instance configuration parameters.","type":"object","properties":{"auto_policy":{"type":"boolean"},"availability_zone":{"type":"string"},"ha_mode":{"type":"string","enum":["active-active","active-standby"]},"interface_list":{"type":"array","item":{"type":"object","properties":{"allowed_address_pairs":{"type":"object","properties":{"allowed_address_pair":{"type":"array","item":{"type":"object","properties":{"address_mode":{"type":"string","enum":["active-active","active-standby"]},"ip":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"mac":{"type":"string"}}}}}},"ip_address":{"type":"string"},"static_routes":{"type":"object","properties":{"route":{"type":"array","item":{"type":"object","properties":{"community_attributes":{"type":"object","properties":{"community_attribute":{"type":"array"}}},"next_hop":{"type":"string"},"next_hop_type":{"type":"string","enum":["service-instance","ip-address"]},"prefix":{"type":"string"}}}}}},"virtual_network":{"type":"string"}}}},"left_ip_address":{"type":"string"},"left_virtual_network":{"type":"string"},"management_virtual_network":{"type":"string"},"right_ip_address":{"type":"string"},"right_virtual_network":{"type":"string"},"scale_out":{"type":"object","properties":{"auto_scale":{"type":"boolean"},"max_instances":{"type":"integer"}}},"virtual_router_id":{"type":"string"}}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}

	}
}

// InterfaceToServiceInstanceSlice makes a slice of ServiceInstance from interface
func InterfaceToServiceInstanceSlice(data interface{}) []*ServiceInstance {
	list := data.([]interface{})
	result := MakeServiceInstanceSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceInstance(item))
	}
	return result
}

// MakeServiceInstanceSlice() makes a slice of ServiceInstance
func MakeServiceInstanceSlice() []*ServiceInstance {
	return []*ServiceInstance{}
}
