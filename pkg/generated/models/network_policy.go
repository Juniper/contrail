package models

// NetworkPolicy

import "encoding/json"

// NetworkPolicy
type NetworkPolicy struct {
	NetworkPolicyEntries *PolicyEntriesType `json:"network_policy_entries"`
	DisplayName          string             `json:"display_name"`
	Annotations          *KeyValuePairs     `json:"annotations"`
	ParentUUID           string             `json:"parent_uuid"`
	ParentType           string             `json:"parent_type"`
	FQName               []string           `json:"fq_name"`
	IDPerms              *IdPermsType       `json:"id_perms"`
	Perms2               *PermType2         `json:"perms2"`
	UUID                 string             `json:"uuid"`
}

// String returns json representation of the object
func (model *NetworkPolicy) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeNetworkPolicy makes NetworkPolicy
func MakeNetworkPolicy() *NetworkPolicy {
	return &NetworkPolicy{
		//TODO(nati): Apply default
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		Perms2:               MakePermType2(),
		UUID:                 "",
		NetworkPolicyEntries: MakePolicyEntriesType(),
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		ParentUUID:           "",
	}
}

// InterfaceToNetworkPolicy makes NetworkPolicy from interface
func InterfaceToNetworkPolicy(iData interface{}) *NetworkPolicy {
	data := iData.(map[string]interface{})
	return &NetworkPolicy{
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		NetworkPolicyEntries: InterfaceToPolicyEntriesType(data["network_policy_entries"]),

		//{"description":"Network policy rule entries.","type":"object","properties":{"policy_rule":{"type":"array","item":{"type":"object","properties":{"action_list":{"type":"object","properties":{"alert":{"type":"boolean"},"apply_service":{"type":"array","item":{"type":"string"}},"assign_routing_instance":{"type":"string"},"gateway_name":{"type":"string"},"log":{"type":"boolean"},"mirror_to":{"type":"object","properties":{"analyzer_ip_address":{"type":"string"},"analyzer_mac_address":{"type":"string"},"analyzer_name":{"type":"string"},"encapsulation":{"type":"string"},"juniper_header":{"type":"boolean"},"nh_mode":{"type":"string","enum":["dynamic","static"]},"nic_assisted_mirroring":{"type":"boolean"},"nic_assisted_mirroring_vlan":{"type":"integer","minimum":1,"maximum":4094},"routing_instance":{"type":"string"},"static_nh_header":{"type":"object","properties":{"vni":{"type":"integer","minimum":1,"maximum":16777215},"vtep_dst_ip_address":{"type":"string"},"vtep_dst_mac_address":{"type":"string"}}},"udp_port":{"type":"integer"}}},"qos_action":{"type":"string"},"simple_action":{"type":"string","enum":["deny","pass"]}}},"application":{"type":"array","item":{"type":"string"}},"created":{"type":"string"},"direction":{"type":"string","enum":["\u003e","\u003c\u003e"]},"dst_addresses":{"type":"array","item":{"type":"object","properties":{"network_policy":{"type":"string"},"security_group":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_list":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}},"virtual_network":{"type":"string"}}}},"dst_ports":{"type":"array","item":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}}},"ethertype":{"type":"string","enum":["IPv4","IPv6"]},"last_modified":{"type":"string"},"protocol":{"type":"string"},"rule_sequence":{"type":"object","properties":{"major":{"type":"integer"},"minor":{"type":"integer"}}},"rule_uuid":{"type":"string"},"src_addresses":{"type":"array","item":{"type":"object","properties":{"network_policy":{"type":"string"},"security_group":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_list":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}},"virtual_network":{"type":"string"}}}},"src_ports":{"type":"array","item":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}}}}}}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}

	}
}

// InterfaceToNetworkPolicySlice makes a slice of NetworkPolicy from interface
func InterfaceToNetworkPolicySlice(data interface{}) []*NetworkPolicy {
	list := data.([]interface{})
	result := MakeNetworkPolicySlice()
	for _, item := range list {
		result = append(result, InterfaceToNetworkPolicy(item))
	}
	return result
}

// MakeNetworkPolicySlice() makes a slice of NetworkPolicy
func MakeNetworkPolicySlice() []*NetworkPolicy {
	return []*NetworkPolicy{}
}
