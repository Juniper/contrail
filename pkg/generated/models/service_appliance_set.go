package models

// ServiceApplianceSet

import "encoding/json"

// ServiceApplianceSet
type ServiceApplianceSet struct {
	ServiceApplianceHaMode        string         `json:"service_appliance_ha_mode"`
	ServiceApplianceDriver        string         `json:"service_appliance_driver"`
	FQName                        []string       `json:"fq_name"`
	IDPerms                       *IdPermsType   `json:"id_perms"`
	ServiceApplianceSetProperties *KeyValuePairs `json:"service_appliance_set_properties"`
	DisplayName                   string         `json:"display_name"`
	Annotations                   *KeyValuePairs `json:"annotations"`
	Perms2                        *PermType2     `json:"perms2"`
	UUID                          string         `json:"uuid"`
	ParentUUID                    string         `json:"parent_uuid"`
	ParentType                    string         `json:"parent_type"`

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
		ServiceApplianceHaMode: "",
		ServiceApplianceDriver: "",
		FQName:                 []string{},
		IDPerms:                MakeIdPermsType(),
		ServiceApplianceSetProperties: MakeKeyValuePairs(),
		DisplayName:                   "",
		Annotations:                   MakeKeyValuePairs(),
		Perms2:                        MakePermType2(),
		UUID:                          "",
		ParentUUID:                    "",
		ParentType:                    "",
	}
}

// InterfaceToServiceApplianceSet makes ServiceApplianceSet from interface
func InterfaceToServiceApplianceSet(iData interface{}) *ServiceApplianceSet {
	data := iData.(map[string]interface{})
	return &ServiceApplianceSet{
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		ServiceApplianceSetProperties: InterfaceToKeyValuePairs(data["service_appliance_set_properties"]),

		//{"description":"List of Key:Value pairs that are used by the provider driver and opaque to system.","type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ServiceApplianceHaMode: data["service_appliance_ha_mode"].(string),

		//{"description":"High availability mode for the service appliance set, active-active or active-backup.","type":"string"}
		ServiceApplianceDriver: data["service_appliance_driver"].(string),

		//{"description":"Name of the provider driver for this service appliance set.","type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}

	}
}

// InterfaceToServiceApplianceSetSlice makes a slice of ServiceApplianceSet from interface
func InterfaceToServiceApplianceSetSlice(data interface{}) []*ServiceApplianceSet {
	list := data.([]interface{})
	result := MakeServiceApplianceSetSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceApplianceSet(item))
	}
	return result
}

// MakeServiceApplianceSetSlice() makes a slice of ServiceApplianceSet
func MakeServiceApplianceSetSlice() []*ServiceApplianceSet {
	return []*ServiceApplianceSet{}
}
