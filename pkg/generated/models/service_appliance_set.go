package models

// ServiceApplianceSet

import "encoding/json"

// ServiceApplianceSet
type ServiceApplianceSet struct {
	ServiceApplianceSetProperties *KeyValuePairs `json:"service_appliance_set_properties,omitempty"`
	ServiceApplianceHaMode        string         `json:"service_appliance_ha_mode,omitempty"`
	ServiceApplianceDriver        string         `json:"service_appliance_driver,omitempty"`
	Annotations                   *KeyValuePairs `json:"annotations,omitempty"`
	DisplayName                   string         `json:"display_name,omitempty"`
	Perms2                        *PermType2     `json:"perms2,omitempty"`
	UUID                          string         `json:"uuid,omitempty"`
	ParentUUID                    string         `json:"parent_uuid,omitempty"`
	ParentType                    string         `json:"parent_type,omitempty"`
	FQName                        []string       `json:"fq_name,omitempty"`
	IDPerms                       *IdPermsType   `json:"id_perms,omitempty"`

	ServiceAppliances []*ServiceAppliance `json:"service_appliances,omitempty"`
}

// String returns json representation of the object
func (model *ServiceApplianceSet) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeServiceApplianceSet makes ServiceApplianceSet
func MakeServiceApplianceSet() *ServiceApplianceSet {
	return &ServiceApplianceSet{
		//TODO(nati): Apply default
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Perms2:      MakePermType2(),
		UUID:        "",
		ServiceApplianceSetProperties: MakeKeyValuePairs(),
		ServiceApplianceHaMode:        "",
		ServiceApplianceDriver:        "",
		Annotations:                   MakeKeyValuePairs(),
	}
}

// MakeServiceApplianceSetSlice() makes a slice of ServiceApplianceSet
func MakeServiceApplianceSetSlice() []*ServiceApplianceSet {
	return []*ServiceApplianceSet{}
}
