package models

// E2ServiceProvider

import "encoding/json"

// E2ServiceProvider
type E2ServiceProvider struct {
	ParentUUID                   string         `json:"parent_uuid"`
	ParentType                   string         `json:"parent_type"`
	E2ServiceProviderPromiscuous bool           `json:"e2_service_provider_promiscuous"`
	DisplayName                  string         `json:"display_name"`
	Annotations                  *KeyValuePairs `json:"annotations"`
	Perms2                       *PermType2     `json:"perms2"`
	UUID                         string         `json:"uuid"`
	FQName                       []string       `json:"fq_name"`
	IDPerms                      *IdPermsType   `json:"id_perms"`

	PeeringPolicyRefs  []*E2ServiceProviderPeeringPolicyRef  `json:"peering_policy_refs"`
	PhysicalRouterRefs []*E2ServiceProviderPhysicalRouterRef `json:"physical_router_refs"`
}

// E2ServiceProviderPeeringPolicyRef references each other
type E2ServiceProviderPeeringPolicyRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// E2ServiceProviderPhysicalRouterRef references each other
type E2ServiceProviderPhysicalRouterRef struct {
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
		FQName:                       []string{},
		IDPerms:                      MakeIdPermsType(),
		Annotations:                  MakeKeyValuePairs(),
		Perms2:                       MakePermType2(),
		UUID:                         "",
		ParentUUID:                   "",
		ParentType:                   "",
		E2ServiceProviderPromiscuous: false,
		DisplayName:                  "",
	}
}

// InterfaceToE2ServiceProvider makes E2ServiceProvider from interface
func InterfaceToE2ServiceProvider(iData interface{}) *E2ServiceProvider {
	data := iData.(map[string]interface{})
	return &E2ServiceProvider{
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		E2ServiceProviderPromiscuous: data["e2_service_provider_promiscuous"].(bool),

		//{"description":"This service provider is connected to all other service providers.","type":"boolean"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}

	}
}

// InterfaceToE2ServiceProviderSlice makes a slice of E2ServiceProvider from interface
func InterfaceToE2ServiceProviderSlice(data interface{}) []*E2ServiceProvider {
	list := data.([]interface{})
	result := MakeE2ServiceProviderSlice()
	for _, item := range list {
		result = append(result, InterfaceToE2ServiceProvider(item))
	}
	return result
}

// MakeE2ServiceProviderSlice() makes a slice of E2ServiceProvider
func MakeE2ServiceProviderSlice() []*E2ServiceProvider {
	return []*E2ServiceProvider{}
}
