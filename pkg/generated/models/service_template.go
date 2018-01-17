package models

// ServiceTemplate

import "encoding/json"

// ServiceTemplate
type ServiceTemplate struct {
	Annotations               *KeyValuePairs       `json:"annotations,omitempty"`
	Perms2                    *PermType2           `json:"perms2,omitempty"`
	ServiceTemplateProperties *ServiceTemplateType `json:"service_template_properties,omitempty"`
	ParentUUID                string               `json:"parent_uuid,omitempty"`
	ParentType                string               `json:"parent_type,omitempty"`
	IDPerms                   *IdPermsType         `json:"id_perms,omitempty"`
	DisplayName               string               `json:"display_name,omitempty"`
	FQName                    []string             `json:"fq_name,omitempty"`
	UUID                      string               `json:"uuid,omitempty"`

	ServiceApplianceSetRefs []*ServiceTemplateServiceApplianceSetRef `json:"service_appliance_set_refs,omitempty"`
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
		Perms2: MakePermType2(),
		ServiceTemplateProperties: MakeServiceTemplateType(),
		ParentUUID:                "",
		ParentType:                "",
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		FQName:                    []string{},
		UUID:                      "",
	}
}

// MakeServiceTemplateSlice() makes a slice of ServiceTemplate
func MakeServiceTemplateSlice() []*ServiceTemplate {
	return []*ServiceTemplate{}
}
