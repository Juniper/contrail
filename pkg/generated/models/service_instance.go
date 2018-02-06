package models

// ServiceInstance

import "encoding/json"

// ServiceInstance
type ServiceInstance struct {
	DisplayName               string               `json:"display_name,omitempty"`
	Perms2                    *PermType2           `json:"perms2,omitempty"`
	ParentUUID                string               `json:"parent_uuid,omitempty"`
	ServiceInstanceBindings   *KeyValuePairs       `json:"service_instance_bindings,omitempty"`
	ServiceInstanceProperties *ServiceInstanceType `json:"service_instance_properties,omitempty"`
	FQName                    []string             `json:"fq_name,omitempty"`
	ParentType                string               `json:"parent_type,omitempty"`
	IDPerms                   *IdPermsType         `json:"id_perms,omitempty"`
	Annotations               *KeyValuePairs       `json:"annotations,omitempty"`
	UUID                      string               `json:"uuid,omitempty"`

	InstanceIPRefs      []*ServiceInstanceInstanceIPRef      `json:"instance_ip_refs,omitempty"`
	ServiceTemplateRefs []*ServiceInstanceServiceTemplateRef `json:"service_template_refs,omitempty"`

	PortTuples []*PortTuple `json:"port_tuples,omitempty"`
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
		ServiceInstanceProperties: MakeServiceInstanceType(),
		FQName:                  []string{},
		DisplayName:             "",
		Perms2:                  MakePermType2(),
		ParentUUID:              "",
		ServiceInstanceBindings: MakeKeyValuePairs(),
		Annotations:             MakeKeyValuePairs(),
		UUID:                    "",
		ParentType:              "",
		IDPerms:                 MakeIdPermsType(),
	}
}

// MakeServiceInstanceSlice() makes a slice of ServiceInstance
func MakeServiceInstanceSlice() []*ServiceInstance {
	return []*ServiceInstance{}
}
