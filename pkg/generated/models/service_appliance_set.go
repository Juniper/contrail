package models

// ServiceApplianceSet

import "encoding/json"

// ServiceApplianceSet
type ServiceApplianceSet struct {
	Annotations                   *KeyValuePairs `json:"annotations,omitempty"`
	UUID                          string         `json:"uuid,omitempty"`
	FQName                        []string       `json:"fq_name,omitempty"`
	ServiceApplianceHaMode        string         `json:"service_appliance_ha_mode,omitempty"`
	ServiceApplianceDriver        string         `json:"service_appliance_driver,omitempty"`
	ParentUUID                    string         `json:"parent_uuid,omitempty"`
	ParentType                    string         `json:"parent_type,omitempty"`
	IDPerms                       *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName                   string         `json:"display_name,omitempty"`
	ServiceApplianceSetProperties *KeyValuePairs `json:"service_appliance_set_properties,omitempty"`
	Perms2                        *PermType2     `json:"perms2,omitempty"`

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
		Perms2:                        MakePermType2(),
		ParentUUID:                    "",
		ParentType:                    "",
		IDPerms:                       MakeIdPermsType(),
		DisplayName:                   "",
		ServiceApplianceSetProperties: MakeKeyValuePairs(),
		ServiceApplianceDriver:        "",
		Annotations:                   MakeKeyValuePairs(),
		UUID:                          "",
		FQName:                        []string{},
		ServiceApplianceHaMode:        "",
	}
}

// MakeServiceApplianceSetSlice() makes a slice of ServiceApplianceSet
func MakeServiceApplianceSetSlice() []*ServiceApplianceSet {
	return []*ServiceApplianceSet{}
}
