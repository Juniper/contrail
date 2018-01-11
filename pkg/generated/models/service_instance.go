package models

// ServiceInstance

import "encoding/json"

// ServiceInstance
type ServiceInstance struct {
	ServiceInstanceProperties *ServiceInstanceType `json:"service_instance_properties"`
	FQName                    []string             `json:"fq_name"`
	DisplayName               string               `json:"display_name"`
	Annotations               *KeyValuePairs       `json:"annotations"`
	ParentUUID                string               `json:"parent_uuid"`
	ServiceInstanceBindings   *KeyValuePairs       `json:"service_instance_bindings"`
	ParentType                string               `json:"parent_type"`
	IDPerms                   *IdPermsType         `json:"id_perms"`
	Perms2                    *PermType2           `json:"perms2"`
	UUID                      string               `json:"uuid"`

	ServiceTemplateRefs []*ServiceInstanceServiceTemplateRef `json:"service_template_refs"`
	InstanceIPRefs      []*ServiceInstanceInstanceIPRef      `json:"instance_ip_refs"`

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
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		ParentUUID:                "",
		ServiceInstanceProperties: MakeServiceInstanceType(),
		FQName:  []string{},
		IDPerms: MakeIdPermsType(),
		Perms2:  MakePermType2(),
		UUID:    "",
		ServiceInstanceBindings: MakeKeyValuePairs(),
		ParentType:              "",
	}
}

// MakeServiceInstanceSlice() makes a slice of ServiceInstance
func MakeServiceInstanceSlice() []*ServiceInstance {
	return []*ServiceInstance{}
}
