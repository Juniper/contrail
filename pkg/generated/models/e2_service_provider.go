package models

// E2ServiceProvider

import "encoding/json"

// E2ServiceProvider
type E2ServiceProvider struct {
	E2ServiceProviderPromiscuous bool           `json:"e2_service_provider_promiscuous"`
	DisplayName                  string         `json:"display_name"`
	Annotations                  *KeyValuePairs `json:"annotations"`
	UUID                         string         `json:"uuid"`
	ParentType                   string         `json:"parent_type"`
	FQName                       []string       `json:"fq_name"`
	IDPerms                      *IdPermsType   `json:"id_perms"`
	Perms2                       *PermType2     `json:"perms2"`
	ParentUUID                   string         `json:"parent_uuid"`

	PhysicalRouterRefs []*E2ServiceProviderPhysicalRouterRef `json:"physical_router_refs"`
	PeeringPolicyRefs  []*E2ServiceProviderPeeringPolicyRef  `json:"peering_policy_refs"`
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
		Annotations: MakeKeyValuePairs(),
		UUID:        "",
		ParentType:  "",
		FQName:      []string{},
		E2ServiceProviderPromiscuous: false,
		DisplayName:                  "",
		ParentUUID:                   "",
		IDPerms:                      MakeIdPermsType(),
		Perms2:                       MakePermType2(),
	}
}

// MakeE2ServiceProviderSlice() makes a slice of E2ServiceProvider
func MakeE2ServiceProviderSlice() []*E2ServiceProvider {
	return []*E2ServiceProvider{}
}
