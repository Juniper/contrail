package models

// NetworkDeviceConfig

import "encoding/json"

// NetworkDeviceConfig
type NetworkDeviceConfig struct {
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`

	PhysicalRouterRefs []*NetworkDeviceConfigPhysicalRouterRef `json:"physical_router_refs"`
}

// NetworkDeviceConfigPhysicalRouterRef references each other
type NetworkDeviceConfigPhysicalRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *NetworkDeviceConfig) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeNetworkDeviceConfig makes NetworkDeviceConfig
func MakeNetworkDeviceConfig() *NetworkDeviceConfig {
	return &NetworkDeviceConfig{
		//TODO(nati): Apply default
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
	}
}

// MakeNetworkDeviceConfigSlice() makes a slice of NetworkDeviceConfig
func MakeNetworkDeviceConfigSlice() []*NetworkDeviceConfig {
	return []*NetworkDeviceConfig{}
}
