package models

// ServiceApplianceSet

import "encoding/json"

// ServiceApplianceSet
type ServiceApplianceSet struct {
	UUID                          string         `json:"uuid,omitempty"`
	FQName                        []string       `json:"fq_name,omitempty"`
	IDPerms                       *IdPermsType   `json:"id_perms,omitempty"`
	ServiceApplianceSetProperties *KeyValuePairs `json:"service_appliance_set_properties,omitempty"`
	ServiceApplianceDriver        string         `json:"service_appliance_driver,omitempty"`
	DisplayName                   string         `json:"display_name,omitempty"`
	Annotations                   *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                        *PermType2     `json:"perms2,omitempty"`
	ParentUUID                    string         `json:"parent_uuid,omitempty"`
	ParentType                    string         `json:"parent_type,omitempty"`
	ServiceApplianceHaMode        string         `json:"service_appliance_ha_mode,omitempty"`

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
		ServiceApplianceHaMode:        "",
		ServiceApplianceDriver:        "",
		DisplayName:                   "",
		Annotations:                   MakeKeyValuePairs(),
		Perms2:                        MakePermType2(),
		ParentUUID:                    "",
		ParentType:                    "",
		ServiceApplianceSetProperties: MakeKeyValuePairs(),
		UUID:    "",
		FQName:  []string{},
		IDPerms: MakeIdPermsType(),
	}
}

// MakeServiceApplianceSetSlice() makes a slice of ServiceApplianceSet
func MakeServiceApplianceSetSlice() []*ServiceApplianceSet {
	return []*ServiceApplianceSet{}
}
