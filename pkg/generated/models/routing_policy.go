package models

// RoutingPolicy

// RoutingPolicy
//proteus:generate
type RoutingPolicy struct {
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`

	ServiceInstanceRefs []*RoutingPolicyServiceInstanceRef `json:"service_instance_refs,omitempty"`
}

// RoutingPolicyServiceInstanceRef references each other
type RoutingPolicyServiceInstanceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *RoutingPolicyServiceInstanceType
}

// MakeRoutingPolicy makes RoutingPolicy
func MakeRoutingPolicy() *RoutingPolicy {
	return &RoutingPolicy{
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

// MakeRoutingPolicySlice() makes a slice of RoutingPolicy
func MakeRoutingPolicySlice() []*RoutingPolicy {
	return []*RoutingPolicy{}
}
