package models

// QosConfig

import "encoding/json"

// QosConfig
type QosConfig struct {
	QosConfigType            QosConfigType              `json:"qos_config_type"`
	MPLSExpEntries           *QosIdForwardingClassPairs `json:"mpls_exp_entries"`
	VlanPriorityEntries      *QosIdForwardingClassPairs `json:"vlan_priority_entries"`
	DSCPEntries              *QosIdForwardingClassPairs `json:"dscp_entries"`
	UUID                     string                     `json:"uuid"`
	ParentType               string                     `json:"parent_type"`
	FQName                   []string                   `json:"fq_name"`
	DisplayName              string                     `json:"display_name"`
	Annotations              *KeyValuePairs             `json:"annotations"`
	Perms2                   *PermType2                 `json:"perms2"`
	DefaultForwardingClassID ForwardingClassId          `json:"default_forwarding_class_id"`
	ParentUUID               string                     `json:"parent_uuid"`
	IDPerms                  *IdPermsType               `json:"id_perms"`

	GlobalSystemConfigRefs []*QosConfigGlobalSystemConfigRef `json:"global_system_config_refs"`
}

// QosConfigGlobalSystemConfigRef references each other
type QosConfigGlobalSystemConfigRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *QosConfig) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeQosConfig makes QosConfig
func MakeQosConfig() *QosConfig {
	return &QosConfig{
		//TODO(nati): Apply default
		DisplayName:         "",
		QosConfigType:       MakeQosConfigType(),
		MPLSExpEntries:      MakeQosIdForwardingClassPairs(),
		VlanPriorityEntries: MakeQosIdForwardingClassPairs(),
		DSCPEntries:         MakeQosIdForwardingClassPairs(),
		UUID:                "",
		ParentType:          "",
		FQName:              []string{},
		Annotations:         MakeKeyValuePairs(),
		Perms2:              MakePermType2(),
		DefaultForwardingClassID: MakeForwardingClassId(),
		ParentUUID:               "",
		IDPerms:                  MakeIdPermsType(),
	}
}

// InterfaceToQosConfig makes QosConfig from interface
func InterfaceToQosConfig(iData interface{}) *QosConfig {
	data := iData.(map[string]interface{})
	return &QosConfig{
		MPLSExpEntries: InterfaceToQosIdForwardingClassPairs(data["mpls_exp_entries"]),

		//{"description":"Map of MPLS EXP values to applicable forwarding class for packet.","type":"object","properties":{"qos_id_forwarding_class_pair":{"type":"array","item":{"type":"object","properties":{"forwarding_class_id":{"default":"0","type":"integer","minimum":0,"maximum":255},"key":{"type":"integer"}}}}}}
		VlanPriorityEntries: InterfaceToQosIdForwardingClassPairs(data["vlan_priority_entries"]),

		//{"description":"Map of .1p priority code to applicable forwarding class for packet.","type":"object","properties":{"qos_id_forwarding_class_pair":{"type":"array","item":{"type":"object","properties":{"forwarding_class_id":{"default":"0","type":"integer","minimum":0,"maximum":255},"key":{"type":"integer"}}}}}}
		DSCPEntries: InterfaceToQosIdForwardingClassPairs(data["dscp_entries"]),

		//{"description":"Map of DSCP match condition and applicable forwarding class for packet.","type":"object","properties":{"qos_id_forwarding_class_pair":{"type":"array","item":{"type":"object","properties":{"forwarding_class_id":{"default":"0","type":"integer","minimum":0,"maximum":255},"key":{"type":"integer"}}}}}}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		QosConfigType: InterfaceToQosConfigType(data["qos_config_type"]),

		//{"description":"Specifies if qos-config is for vhost, fabric or for project.","type":"string","enum":["vhost","fabric","project"]}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		DefaultForwardingClassID: InterfaceToForwardingClassId(data["default_forwarding_class_id"]),

		//{"description":"Default forwarding class used for all non-specified QOS bits","default":"0","type":"integer","minimum":0,"maximum":255}

	}
}

// InterfaceToQosConfigSlice makes a slice of QosConfig from interface
func InterfaceToQosConfigSlice(data interface{}) []*QosConfig {
	list := data.([]interface{})
	result := MakeQosConfigSlice()
	for _, item := range list {
		result = append(result, InterfaceToQosConfig(item))
	}
	return result
}

// MakeQosConfigSlice() makes a slice of QosConfig
func MakeQosConfigSlice() []*QosConfig {
	return []*QosConfig{}
}
