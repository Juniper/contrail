package models

// ServiceTemplate

import "encoding/json"

// ServiceTemplate
type ServiceTemplate struct {
	ServiceTemplateProperties *ServiceTemplateType `json:"service_template_properties"`
	ParentType                string               `json:"parent_type"`
	FQName                    []string             `json:"fq_name"`
	DisplayName               string               `json:"display_name"`
	Perms2                    *PermType2           `json:"perms2"`
	ParentUUID                string               `json:"parent_uuid"`
	IDPerms                   *IdPermsType         `json:"id_perms"`
	Annotations               *KeyValuePairs       `json:"annotations"`
	UUID                      string               `json:"uuid"`

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
		UUID:        "",
		ParentUUID:  "",
		IDPerms:     MakeIdPermsType(),
		Annotations: MakeKeyValuePairs(),
		DisplayName: "",
		Perms2:      MakePermType2(),
		ServiceTemplateProperties: MakeServiceTemplateType(),
		ParentType:                "",
		FQName:                    []string{},
	}
}

// MakeServiceTemplateSlice() makes a slice of ServiceTemplate
func MakeServiceTemplateSlice() []*ServiceTemplate {
	return []*ServiceTemplate{}
}
