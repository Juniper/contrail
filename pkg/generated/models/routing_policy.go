package models

// RoutingPolicy

import "encoding/json"

// RoutingPolicy
type RoutingPolicy struct {
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`

	ServiceInstanceRefs []*RoutingPolicyServiceInstanceRef `json:"service_instance_refs"`
}

// RoutingPolicyServiceInstanceRef references each other
type RoutingPolicyServiceInstanceRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

	Attr *RoutingPolicyServiceInstanceType
}

// String returns json representation of the object
func (model *RoutingPolicy) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeRoutingPolicy makes RoutingPolicy
func MakeRoutingPolicy() *RoutingPolicy {
	return &RoutingPolicy{
		//TODO(nati): Apply default
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
	}
}

// MakeRoutingPolicySlice() makes a slice of RoutingPolicy
func MakeRoutingPolicySlice() []*RoutingPolicy {
	return []*RoutingPolicy{}
}
