package models

// ServiceApplianceSet

import "encoding/json"

// ServiceApplianceSet
type ServiceApplianceSet struct {
	ParentType                    string         `json:"parent_type"`
	FQName                        []string       `json:"fq_name"`
	ServiceApplianceSetProperties *KeyValuePairs `json:"service_appliance_set_properties"`
	ServiceApplianceHaMode        string         `json:"service_appliance_ha_mode"`
	ServiceApplianceDriver        string         `json:"service_appliance_driver"`
	Annotations                   *KeyValuePairs `json:"annotations"`
	Perms2                        *PermType2     `json:"perms2"`
	ParentUUID                    string         `json:"parent_uuid"`
	DisplayName                   string         `json:"display_name"`
	UUID                          string         `json:"uuid"`
	IDPerms                       *IdPermsType   `json:"id_perms"`

	ServiceAppliances []*ServiceAppliance `json:"service_appliances"`
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
		DisplayName: "",
		UUID:        "",
		IDPerms:     MakeIdPermsType(),
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		ServiceApplianceSetProperties: MakeKeyValuePairs(),
		ServiceApplianceHaMode:        "",
		ServiceApplianceDriver:        "",
	}
}

// MakeServiceApplianceSetSlice() makes a slice of ServiceApplianceSet
func MakeServiceApplianceSetSlice() []*ServiceApplianceSet {
	return []*ServiceApplianceSet{}
}
