package models

// ServiceApplianceSet

import "encoding/json"

// ServiceApplianceSet
type ServiceApplianceSet struct {
	ServiceApplianceDriver        string         `json:"service_appliance_driver,omitempty"`
	ParentType                    string         `json:"parent_type,omitempty"`
	FQName                        []string       `json:"fq_name,omitempty"`
	DisplayName                   string         `json:"display_name,omitempty"`
	Perms2                        *PermType2     `json:"perms2,omitempty"`
	ServiceApplianceSetProperties *KeyValuePairs `json:"service_appliance_set_properties,omitempty"`
	ServiceApplianceHaMode        string         `json:"service_appliance_ha_mode,omitempty"`
	Annotations                   *KeyValuePairs `json:"annotations,omitempty"`
	UUID                          string         `json:"uuid,omitempty"`
	ParentUUID                    string         `json:"parent_uuid,omitempty"`
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
		ServiceApplianceDriver: "",
		ParentType:             "",
		FQName:                 []string{},
		DisplayName:            "",
		Perms2:                 MakePermType2(),
		ServiceApplianceSetProperties: MakeKeyValuePairs(),
		ServiceApplianceHaMode:        "",
		Annotations:                   MakeKeyValuePairs(),
		UUID:                          "",
		ParentUUID:                    "",
		IDPerms:                       MakeIdPermsType(),
	}
}

// MakeServiceApplianceSetSlice() makes a slice of ServiceApplianceSet
func MakeServiceApplianceSetSlice() []*ServiceApplianceSet {
	return []*ServiceApplianceSet{}
}
