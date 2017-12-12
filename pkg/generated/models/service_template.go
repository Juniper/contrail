package models

// ServiceTemplate

import "encoding/json"

// ServiceTemplate
type ServiceTemplate struct {
	IDPerms                   *IdPermsType         `json:"id_perms"`
	DisplayName               string               `json:"display_name"`
	Perms2                    *PermType2           `json:"perms2"`
	UUID                      string               `json:"uuid"`
	ServiceTemplateProperties *ServiceTemplateType `json:"service_template_properties"`
	ParentUUID                string               `json:"parent_uuid"`
	Annotations               *KeyValuePairs       `json:"annotations"`
	ParentType                string               `json:"parent_type"`
	FQName                    []string             `json:"fq_name"`

	ServiceApplianceSetRefs []*ServiceTemplateServiceApplianceSetRef `json:"service_appliance_set_refs"`
}

// ServiceTemplateServiceApplianceSetRef references each other
type ServiceTemplateServiceApplianceSetRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *ServiceTemplate) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceTemplate makes ServiceTemplate
func MakeServiceTemplate() *ServiceTemplate {
	return &ServiceTemplate{
		//TODO(nati): Apply default
		ServiceTemplateProperties: MakeServiceTemplateType(),
		ParentUUID:                "",
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		Perms2:                    MakePermType2(),
		UUID:                      "",
		ParentType:                "",
		FQName:                    []string{},
		Annotations:               MakeKeyValuePairs(),
	}
}

// InterfaceToServiceTemplate makes ServiceTemplate from interface
func InterfaceToServiceTemplate(iData interface{}) *ServiceTemplate {
	data := iData.(map[string]interface{})
	return &ServiceTemplate{
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ServiceTemplateProperties: InterfaceToServiceTemplateType(data["service_template_properties"]),

		//{"description":"Service template configuration parameters.","type":"object","properties":{"availability_zone_enable":{"type":"boolean"},"flavor":{"type":"string"},"image_name":{"type":"string"},"instance_data":{"type":"string"},"interface_type":{"type":"array","item":{"type":"object","properties":{"service_interface_type":{"type":"string"},"shared_ip":{"type":"boolean"},"static_route_enable":{"type":"boolean"}}}},"ordered_interfaces":{"type":"boolean"},"service_mode":{"type":"string","enum":["transparent","in-network","in-network-nat"]},"service_scaling":{"type":"boolean"},"service_type":{"type":"string","enum":["firewall","analyzer","source-nat","loadbalancer"]},"service_virtualization_type":{"type":"string","enum":["virtual-machine","network-namespace","vrouter-instance","physical-device"]},"version":{"type":"integer"},"vrouter_instance_type":{"type":"string","enum":["libvirt-qemu","docker"]}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}

	}
}

// InterfaceToServiceTemplateSlice makes a slice of ServiceTemplate from interface
func InterfaceToServiceTemplateSlice(data interface{}) []*ServiceTemplate {
	list := data.([]interface{})
	result := MakeServiceTemplateSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceTemplate(item))
	}
	return result
}

// MakeServiceTemplateSlice() makes a slice of ServiceTemplate
func MakeServiceTemplateSlice() []*ServiceTemplate {
	return []*ServiceTemplate{}
}
