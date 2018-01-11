package models

// FirewallRule

import "encoding/json"

// FirewallRule
type FirewallRule struct {
	Annotations   *KeyValuePairs                   `json:"annotations"`
	Endpoint2     *FirewallRuleEndpointType        `json:"endpoint_2"`
	Direction     FirewallRuleDirectionType        `json:"direction"`
	MatchTagTypes *FirewallRuleMatchTagsTypeIdList `json:"match_tag_types"`
	MatchTags     *FirewallRuleMatchTagsType       `json:"match_tags"`
	FQName        []string                         `json:"fq_name"`
	DisplayName   string                           `json:"display_name"`
	Service       *FirewallServiceType             `json:"service"`
	ParentType    string                           `json:"parent_type"`
	IDPerms       *IdPermsType                     `json:"id_perms"`
	ActionList    *ActionListType                  `json:"action_list"`
	ParentUUID    string                           `json:"parent_uuid"`
	Endpoint1     *FirewallRuleEndpointType        `json:"endpoint_1"`
	Perms2        *PermType2                       `json:"perms2"`
	UUID          string                           `json:"uuid"`

	SecurityLoggingObjectRefs []*FirewallRuleSecurityLoggingObjectRef `json:"security_logging_object_refs"`
	VirtualNetworkRefs        []*FirewallRuleVirtualNetworkRef        `json:"virtual_network_refs"`
	ServiceGroupRefs          []*FirewallRuleServiceGroupRef          `json:"service_group_refs"`
	AddressGroupRefs          []*FirewallRuleAddressGroupRef          `json:"address_group_refs"`
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
		Endpoint1:     MakeFirewallRuleEndpointType(),
		Perms2:        MakePermType2(),
		UUID:          "",
		Endpoint2:     MakeFirewallRuleEndpointType(),
		Direction:     MakeFirewallRuleDirectionType(),
		MatchTagTypes: MakeFirewallRuleMatchTagsTypeIdList(),
		MatchTags:     MakeFirewallRuleMatchTagsType(),
		FQName:        []string{},
		DisplayName:   "",
		Annotations:   MakeKeyValuePairs(),
		Service:       MakeFirewallServiceType(),
		ParentType:    "",
		IDPerms:       MakeIdPermsType(),
		ActionList:    MakeActionListType(),
		ParentUUID:    "",
	}
}

// MakeFirewallRuleSlice() makes a slice of FirewallRule
func MakeFirewallRuleSlice() []*FirewallRule {
	return []*FirewallRule{}
}
