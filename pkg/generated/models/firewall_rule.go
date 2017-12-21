package models

// FirewallRule

import "encoding/json"

// FirewallRule
type FirewallRule struct {
	DisplayName   string                           `json:"display_name"`
	Annotations   *KeyValuePairs                   `json:"annotations"`
	ParentType    string                           `json:"parent_type"`
	Service       *FirewallServiceType             `json:"service"`
	Endpoint2     *FirewallRuleEndpointType        `json:"endpoint_2"`
	ActionList    *ActionListType                  `json:"action_list"`
	MatchTagTypes *FirewallRuleMatchTagsTypeIdList `json:"match_tag_types"`
	FQName        []string                         `json:"fq_name"`
	Perms2        *PermType2                       `json:"perms2"`
	ParentUUID    string                           `json:"parent_uuid"`
	Endpoint1     *FirewallRuleEndpointType        `json:"endpoint_1"`
	MatchTags     *FirewallRuleMatchTagsType       `json:"match_tags"`
	IDPerms       *IdPermsType                     `json:"id_perms"`
	UUID          string                           `json:"uuid"`
	Direction     FirewallRuleDirectionType        `json:"direction"`

	VirtualNetworkRefs        []*FirewallRuleVirtualNetworkRef        `json:"virtual_network_refs"`
	ServiceGroupRefs          []*FirewallRuleServiceGroupRef          `json:"service_group_refs"`
	AddressGroupRefs          []*FirewallRuleAddressGroupRef          `json:"address_group_refs"`
	SecurityLoggingObjectRefs []*FirewallRuleSecurityLoggingObjectRef `json:"security_logging_object_refs"`
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
		ActionList:    MakeActionListType(),
		MatchTagTypes: MakeFirewallRuleMatchTagsTypeIdList(),
		FQName:        []string{},
		Perms2:        MakePermType2(),
		ParentUUID:    "",
		Endpoint1:     MakeFirewallRuleEndpointType(),
		Endpoint2:     MakeFirewallRuleEndpointType(),
		IDPerms:       MakeIdPermsType(),
		UUID:          "",
		Direction:     MakeFirewallRuleDirectionType(),
		MatchTags:     MakeFirewallRuleMatchTagsType(),
		Annotations:   MakeKeyValuePairs(),
		ParentType:    "",
		Service:       MakeFirewallServiceType(),
		DisplayName:   "",
	}
}

// InterfaceToFirewallRule makes FirewallRule from interface
func InterfaceToFirewallRule(iData interface{}) *FirewallRule {
	data := iData.(map[string]interface{})
	return &FirewallRule{
		Direction: InterfaceToFirewallRuleDirectionType(data["direction"]),

		//{"description":"Direction in the rule","type":"string","enum":["\u003c","\u003e","\u003c\u003e"]}
		MatchTags: InterfaceToFirewallRuleMatchTagsType(data["match_tags"]),

		//{"description":"matching tags for source and destination endpoints","type":"object","properties":{"tag_list":{"type":"array","item":{"type":"string"}}}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		Service: InterfaceToFirewallServiceType(data["service"]),

		//{"description":"Service (port, protocol) for packets match condition","type":"object","properties":{"dst_ports":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}},"protocol":{"type":"string"},"protocol_id":{"type":"integer"},"src_ports":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		Endpoint1: InterfaceToFirewallRuleEndpointType(data["endpoint_1"]),

		//{"description":"match condition for traffic source","type":"object","properties":{"address_group":{"type":"string"},"any":{"type":"boolean"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"tag_ids":{"type":"array","item":{"type":"integer"}},"tags":{"type":"array","item":{"type":"string"}},"virtual_network":{"type":"string"}}}
		Endpoint2: InterfaceToFirewallRuleEndpointType(data["endpoint_2"]),

		//{"description":"match condition for traffic destination","type":"object","properties":{"address_group":{"type":"string"},"any":{"type":"boolean"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"tag_ids":{"type":"array","item":{"type":"integer"}},"tags":{"type":"array","item":{"type":"string"}},"virtual_network":{"type":"string"}}}
		ActionList: InterfaceToActionListType(data["action_list"]),

		//{"description":"Actions to be performed if packets match condition","type":"object","properties":{"alert":{"type":"boolean"},"apply_service":{"type":"array","item":{"type":"string"}},"assign_routing_instance":{"type":"string"},"gateway_name":{"type":"string"},"log":{"type":"boolean"},"mirror_to":{"type":"object","properties":{"analyzer_ip_address":{"type":"string"},"analyzer_mac_address":{"type":"string"},"analyzer_name":{"type":"string"},"encapsulation":{"type":"string"},"juniper_header":{"type":"boolean"},"nh_mode":{"type":"string","enum":["dynamic","static"]},"nic_assisted_mirroring":{"type":"boolean"},"nic_assisted_mirroring_vlan":{"type":"integer","minimum":1,"maximum":4094},"routing_instance":{"type":"string"},"static_nh_header":{"type":"object","properties":{"vni":{"type":"integer","minimum":1,"maximum":16777215},"vtep_dst_ip_address":{"type":"string"},"vtep_dst_mac_address":{"type":"string"}}},"udp_port":{"type":"integer"}}},"qos_action":{"type":"string"},"simple_action":{"type":"string","enum":["deny","pass"]}}}
		MatchTagTypes: InterfaceToFirewallRuleMatchTagsTypeIdList(data["match_tag_types"]),

		//{"description":"matching tags ids for source and destination endpoints","type":"object","properties":{"tag_type":{"type":"array","item":{"type":"integer"}}}}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}

	}
}

// InterfaceToFirewallRuleSlice makes a slice of FirewallRule from interface
func InterfaceToFirewallRuleSlice(data interface{}) []*FirewallRule {
	list := data.([]interface{})
	result := MakeFirewallRuleSlice()
	for _, item := range list {
		result = append(result, InterfaceToFirewallRule(item))
	}
	return result
}

// MakeFirewallRuleSlice() makes a slice of FirewallRule
func MakeFirewallRuleSlice() []*FirewallRule {
	return []*FirewallRule{}
}
