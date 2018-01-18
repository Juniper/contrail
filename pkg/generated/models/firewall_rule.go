package models

// FirewallRule

import "encoding/json"

// FirewallRule
type FirewallRule struct {
	IDPerms       *IdPermsType                     `json:"id_perms,omitempty"`
	DisplayName   string                           `json:"display_name,omitempty"`
	Endpoint1     *FirewallRuleEndpointType        `json:"endpoint_1,omitempty"`
	Endpoint2     *FirewallRuleEndpointType        `json:"endpoint_2,omitempty"`
	Service       *FirewallServiceType             `json:"service,omitempty"`
	MatchTags     *FirewallRuleMatchTagsType       `json:"match_tags,omitempty"`
	ActionList    *ActionListType                  `json:"action_list,omitempty"`
	Perms2        *PermType2                       `json:"perms2,omitempty"`
	ParentUUID    string                           `json:"parent_uuid,omitempty"`
	ParentType    string                           `json:"parent_type,omitempty"`
	FQName        []string                         `json:"fq_name,omitempty"`
	Direction     FirewallRuleDirectionType        `json:"direction,omitempty"`
	MatchTagTypes *FirewallRuleMatchTagsTypeIdList `json:"match_tag_types,omitempty"`
	UUID          string                           `json:"uuid,omitempty"`
	Annotations   *KeyValuePairs                   `json:"annotations,omitempty"`

	ServiceGroupRefs          []*FirewallRuleServiceGroupRef          `json:"service_group_refs,omitempty"`
	AddressGroupRefs          []*FirewallRuleAddressGroupRef          `json:"address_group_refs,omitempty"`
	SecurityLoggingObjectRefs []*FirewallRuleSecurityLoggingObjectRef `json:"security_logging_object_refs,omitempty"`
	VirtualNetworkRefs        []*FirewallRuleVirtualNetworkRef        `json:"virtual_network_refs,omitempty"`
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
		FQName:        []string{},
		ActionList:    MakeActionListType(),
		Perms2:        MakePermType2(),
		ParentUUID:    "",
		Annotations:   MakeKeyValuePairs(),
		Direction:     MakeFirewallRuleDirectionType(),
		MatchTagTypes: MakeFirewallRuleMatchTagsTypeIdList(),
		UUID:          "",
		IDPerms:       MakeIdPermsType(),
		DisplayName:   "",
		MatchTags:     MakeFirewallRuleMatchTagsType(),
		Endpoint1:     MakeFirewallRuleEndpointType(),
		Endpoint2:     MakeFirewallRuleEndpointType(),
		Service:       MakeFirewallServiceType(),
	}
}

// MakeFirewallRuleSlice() makes a slice of FirewallRule
func MakeFirewallRuleSlice() []*FirewallRule {
	return []*FirewallRule{}
}
