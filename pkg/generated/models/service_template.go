package models

// ServiceTemplate

// ServiceTemplate
//proteus:generate
type ServiceTemplate struct {
	UUID                      string               `json:"uuid,omitempty"`
	ParentUUID                string               `json:"parent_uuid,omitempty"`
	ParentType                string               `json:"parent_type,omitempty"`
	FQName                    []string             `json:"fq_name,omitempty"`
	IDPerms                   *IdPermsType         `json:"id_perms,omitempty"`
	DisplayName               string               `json:"display_name,omitempty"`
	Annotations               *KeyValuePairs       `json:"annotations,omitempty"`
	Perms2                    *PermType2           `json:"perms2,omitempty"`
	ServiceTemplateProperties *ServiceTemplateType `json:"service_template_properties,omitempty"`

	ServiceApplianceSetRefs []*ServiceTemplateServiceApplianceSetRef `json:"service_appliance_set_refs,omitempty"`
}

// ServiceTemplateServiceApplianceSetRef references each other
type ServiceTemplateServiceApplianceSetRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// MakeServiceTemplate makes ServiceTemplate
func MakeServiceTemplate() *ServiceTemplate {
	return &ServiceTemplate{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		ServiceTemplateProperties: MakeServiceTemplateType(),
	}
}

// MakeServiceTemplateSlice() makes a slice of ServiceTemplate
func MakeServiceTemplateSlice() []*ServiceTemplate {
	return []*ServiceTemplate{}
}
