package models

// NetworkDeviceConfig

import "encoding/json"

// NetworkDeviceConfig
type NetworkDeviceConfig struct {
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`

	PhysicalRouterRefs []*NetworkDeviceConfigPhysicalRouterRef `json:"physical_router_refs,omitempty"`
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
