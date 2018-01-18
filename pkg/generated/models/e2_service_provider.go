package models

// E2ServiceProvider

import "encoding/json"

// E2ServiceProvider
type E2ServiceProvider struct {
	E2ServiceProviderPromiscuous bool           `json:"e2_service_provider_promiscuous"`
	UUID                         string         `json:"uuid,omitempty"`
	ParentType                   string         `json:"parent_type,omitempty"`
	FQName                       []string       `json:"fq_name,omitempty"`
	IDPerms                      *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName                  string         `json:"display_name,omitempty"`
	Perms2                       *PermType2     `json:"perms2,omitempty"`
	ParentUUID                   string         `json:"parent_uuid,omitempty"`
	Annotations                  *KeyValuePairs `json:"annotations,omitempty"`

	PhysicalRouterRefs []*E2ServiceProviderPhysicalRouterRef `json:"physical_router_refs,omitempty"`
	PeeringPolicyRefs  []*E2ServiceProviderPeeringPolicyRef  `json:"peering_policy_refs,omitempty"`
}

// E2ServiceProviderPhysicalRouterRef references each other
type E2ServiceProviderPhysicalRouterRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// E2ServiceProviderPeeringPolicyRef references each other
type E2ServiceProviderPeeringPolicyRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *E2ServiceProvider) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeE2ServiceProvider makes E2ServiceProvider
func MakeE2ServiceProvider() *E2ServiceProvider {
	return &E2ServiceProvider{
		//TODO(nati): Apply default
		E2ServiceProviderPromiscuous: false,
		UUID:        "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Perms2:      MakePermType2(),
		ParentUUID:  "",
		Annotations: MakeKeyValuePairs(),
	}
}

// MakeE2ServiceProviderSlice() makes a slice of E2ServiceProvider
func MakeE2ServiceProviderSlice() []*E2ServiceProvider {
	return []*E2ServiceProvider{}
}
