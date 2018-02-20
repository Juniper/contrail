package models

// NetworkDeviceConfig

// NetworkDeviceConfig
//proteus:generate
type NetworkDeviceConfig struct {
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`

	PhysicalRouterRefs []*NetworkDeviceConfigPhysicalRouterRef `json:"physical_router_refs,omitempty"`
}

// NetworkDeviceConfigPhysicalRouterRef references each other
type NetworkDeviceConfigPhysicalRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// MakeNetworkDeviceConfig makes NetworkDeviceConfig
func MakeNetworkDeviceConfig() *NetworkDeviceConfig {
	return &NetworkDeviceConfig{
		//TODO(nati): Apply default
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
	}
}

// MakeNetworkDeviceConfigSlice() makes a slice of NetworkDeviceConfig
func MakeNetworkDeviceConfigSlice() []*NetworkDeviceConfig {
	return []*NetworkDeviceConfig{}
}
