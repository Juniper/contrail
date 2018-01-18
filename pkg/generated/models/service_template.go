package models

// ServiceTemplate

import "encoding/json"

// ServiceTemplate
type ServiceTemplate struct {
	ParentType                string               `json:"parent_type,omitempty"`
	FQName                    []string             `json:"fq_name,omitempty"`
	IDPerms                   *IdPermsType         `json:"id_perms,omitempty"`
	Perms2                    *PermType2           `json:"perms2,omitempty"`
	ParentUUID                string               `json:"parent_uuid,omitempty"`
	UUID                      string               `json:"uuid,omitempty"`
	ServiceTemplateProperties *ServiceTemplateType `json:"service_template_properties,omitempty"`
	DisplayName               string               `json:"display_name,omitempty"`
	Annotations               *KeyValuePairs       `json:"annotations,omitempty"`

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
		Perms2:      MakePermType2(),
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		UUID:        "",
		ServiceTemplateProperties: MakeServiceTemplateType(),
	}
}

// MakeServiceTemplateSlice() makes a slice of ServiceTemplate
func MakeServiceTemplateSlice() []*ServiceTemplate {
	return []*ServiceTemplate{}
}
