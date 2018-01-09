package models

// SecurityGroup

import "encoding/json"

// SecurityGroup
type SecurityGroup struct {
	SecurityGroupID           SecurityGroupIdType           `json:"security_group_id"`
	IDPerms                   *IdPermsType                  `json:"id_perms"`
	DisplayName               string                        `json:"display_name"`
	UUID                      string                        `json:"uuid"`
	ParentUUID                string                        `json:"parent_uuid"`
	ParentType                string                        `json:"parent_type"`
	FQName                    []string                      `json:"fq_name"`
	Annotations               *KeyValuePairs                `json:"annotations"`
	SecurityGroupEntries      *PolicyEntriesType            `json:"security_group_entries"`
	ConfiguredSecurityGroupID ConfiguredSecurityGroupIdType `json:"configured_security_group_id"`
	Perms2                    *PermType2                    `json:"perms2"`

	AccessControlLists []*AccessControlList `json:"access_control_lists"`
}

// String returns json representation of the object
func (model *SecurityGroup) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeSecurityGroup makes SecurityGroup
func MakeSecurityGroup() *SecurityGroup {
	return &SecurityGroup{
		//TODO(nati): Apply default
		SecurityGroupID:           MakeSecurityGroupIdType(),
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		SecurityGroupEntries:      MakePolicyEntriesType(),
		ConfiguredSecurityGroupID: MakeConfiguredSecurityGroupIdType(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		Annotations: MakeKeyValuePairs(),
	}
}

// InterfaceToSecurityGroup makes SecurityGroup from interface
func InterfaceToSecurityGroup(iData interface{}) *SecurityGroup {
	data := iData.(map[string]interface{})
	return &SecurityGroup{
		SecurityGroupID: InterfaceToSecurityGroupIdType(data["security_group_id"]),

		//{"description":"Unique 32 bit ID automatically assigned to this security group [8M+1, 32G].","type":"integer","minimum":8000000,"maximum":4294967296}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		SecurityGroupEntries: InterfaceToPolicyEntriesType(data["security_group_entries"]),

		//{"description":"Security group rule entries.","type":"object","properties":{"policy_rule":{"type":"array","item":{"type":"object","properties":{"action_list":{"type":"object","properties":{"alert":{"type":"boolean"},"apply_service":{"type":"array","item":{"type":"string"}},"assign_routing_instance":{"type":"string"},"gateway_name":{"type":"string"},"log":{"type":"boolean"},"mirror_to":{"type":"object","properties":{"analyzer_ip_address":{"type":"string"},"analyzer_mac_address":{"type":"string"},"analyzer_name":{"type":"string"},"encapsulation":{"type":"string"},"juniper_header":{"type":"boolean"},"nh_mode":{"type":"string","enum":["dynamic","static"]},"nic_assisted_mirroring":{"type":"boolean"},"nic_assisted_mirroring_vlan":{"type":"integer","minimum":1,"maximum":4094},"routing_instance":{"type":"string"},"static_nh_header":{"type":"object","properties":{"vni":{"type":"integer","minimum":1,"maximum":16777215},"vtep_dst_ip_address":{"type":"string"},"vtep_dst_mac_address":{"type":"string"}}},"udp_port":{"type":"integer"}}},"qos_action":{"type":"string"},"simple_action":{"type":"string","enum":["deny","pass"]}}},"application":{"type":"array","item":{"type":"string"}},"created":{"type":"string"},"direction":{"type":"string","enum":["\u003e","\u003c\u003e"]},"dst_addresses":{"type":"array","item":{"type":"object","properties":{"network_policy":{"type":"string"},"security_group":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_list":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}},"virtual_network":{"type":"string"}}}},"dst_ports":{"type":"array","item":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}}},"ethertype":{"type":"string","enum":["IPv4","IPv6"]},"last_modified":{"type":"string"},"protocol":{"type":"string"},"rule_sequence":{"type":"object","properties":{"major":{"type":"integer"},"minor":{"type":"integer"}}},"rule_uuid":{"type":"string"},"src_addresses":{"type":"array","item":{"type":"object","properties":{"network_policy":{"type":"string"},"security_group":{"type":"string"},"subnet":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}},"subnet_list":{"type":"array","item":{"type":"object","properties":{"ip_prefix":{"type":"string"},"ip_prefix_len":{"type":"integer"}}}},"virtual_network":{"type":"string"}}}},"src_ports":{"type":"array","item":{"type":"object","properties":{"end_port":{"type":"integer","minimum":-1,"maximum":65535},"start_port":{"type":"integer","minimum":-1,"maximum":65535}}}}}}}}}
		ConfiguredSecurityGroupID: InterfaceToConfiguredSecurityGroupIdType(data["configured_security_group_id"]),

		//{"description":"Unique 32 bit user defined ID assigned to this security group [1, 8M - 1].","default":"0","type":"integer","minimum":0,"maximum":7999999}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}

	}
}

// InterfaceToSecurityGroupSlice makes a slice of SecurityGroup from interface
func InterfaceToSecurityGroupSlice(data interface{}) []*SecurityGroup {
	list := data.([]interface{})
	result := MakeSecurityGroupSlice()
	for _, item := range list {
		result = append(result, InterfaceToSecurityGroup(item))
	}
	return result
}

// MakeSecurityGroupSlice() makes a slice of SecurityGroup
func MakeSecurityGroupSlice() []*SecurityGroup {
	return []*SecurityGroup{}
}
