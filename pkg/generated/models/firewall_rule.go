package models

// FirewallRule

import "encoding/json"

// FirewallRule
type FirewallRule struct {
	MatchTags     *FirewallRuleMatchTagsType       `json:"match_tags,omitempty"`
	FQName        []string                         `json:"fq_name,omitempty"`
	Perms2        *PermType2                       `json:"perms2,omitempty"`
	Endpoint2     *FirewallRuleEndpointType        `json:"endpoint_2,omitempty"`
	ParentType    string                           `json:"parent_type,omitempty"`
	DisplayName   string                           `json:"display_name,omitempty"`
	Service       *FirewallServiceType             `json:"service,omitempty"`
	Direction     FirewallRuleDirectionType        `json:"direction,omitempty"`
	MatchTagTypes *FirewallRuleMatchTagsTypeIdList `json:"match_tag_types,omitempty"`
	IDPerms       *IdPermsType                     `json:"id_perms,omitempty"`
	ActionList    *ActionListType                  `json:"action_list,omitempty"`
	Annotations   *KeyValuePairs                   `json:"annotations,omitempty"`
	UUID          string                           `json:"uuid,omitempty"`
	ParentUUID    string                           `json:"parent_uuid,omitempty"`
	Endpoint1     *FirewallRuleEndpointType        `json:"endpoint_1,omitempty"`

	VirtualNetworkRefs        []*FirewallRuleVirtualNetworkRef        `json:"virtual_network_refs,omitempty"`
	ServiceGroupRefs          []*FirewallRuleServiceGroupRef          `json:"service_group_refs,omitempty"`
	AddressGroupRefs          []*FirewallRuleAddressGroupRef          `json:"address_group_refs,omitempty"`
	SecurityLoggingObjectRefs []*FirewallRuleSecurityLoggingObjectRef `json:"security_logging_object_refs,omitempty"`
}

// FirewallRuleServiceGroupRef references each other
type FirewallRuleServiceGroupRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// FirewallRuleAddressGroupRef references each other
type FirewallRuleAddressGroupRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// FirewallRuleSecurityLoggingObjectRef references each other
type FirewallRuleSecurityLoggingObjectRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// FirewallRuleVirtualNetworkRef references each other
type FirewallRuleVirtualNetworkRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *FirewallRule) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeFirewallRule makes FirewallRule
func MakeFirewallRule() *FirewallRule {
	return &FirewallRule{
		//TODO(nati): Apply default
		ParentType:    "",
		DisplayName:   "",
		Service:       MakeFirewallServiceType(),
		Direction:     MakeFirewallRuleDirectionType(),
		MatchTagTypes: MakeFirewallRuleMatchTagsTypeIdList(),
		IDPerms:       MakeIdPermsType(),
		ActionList:    MakeActionListType(),
		Annotations:   MakeKeyValuePairs(),
		UUID:          "",
		ParentUUID:    "",
		Endpoint1:     MakeFirewallRuleEndpointType(),
		MatchTags:     MakeFirewallRuleMatchTagsType(),
		FQName:        []string{},
		Perms2:        MakePermType2(),
		Endpoint2:     MakeFirewallRuleEndpointType(),
	}
}

// MakeFirewallRuleSlice() makes a slice of FirewallRule
func MakeFirewallRuleSlice() []*FirewallRule {
	return []*FirewallRule{}
}
