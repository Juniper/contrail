package models

// ForwardingClass

import "encoding/json"

// ForwardingClass
type ForwardingClass struct {
	Annotations                 *KeyValuePairs    `json:"annotations"`
	Perms2                      *PermType2        `json:"perms2"`
	ForwardingClassVlanPriority VlanPriorityType  `json:"forwarding_class_vlan_priority"`
	ForwardingClassID           ForwardingClassId `json:"forwarding_class_id"`
	ParentUUID                  string            `json:"parent_uuid"`
	FQName                      []string          `json:"fq_name"`
	DisplayName                 string            `json:"display_name"`
	UUID                        string            `json:"uuid"`
	ForwardingClassDSCP         DscpValueType     `json:"forwarding_class_dscp"`
	ForwardingClassMPLSExp      MplsExpType       `json:"forwarding_class_mpls_exp"`
	ParentType                  string            `json:"parent_type"`
	IDPerms                     *IdPermsType      `json:"id_perms"`

	QosQueueRefs []*ForwardingClassQosQueueRef `json:"qos_queue_refs"`
}

// ForwardingClassQosQueueRef references each other
type ForwardingClassQosQueueRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *ForwardingClass) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeForwardingClass makes ForwardingClass
func MakeForwardingClass() *ForwardingClass {
	return &ForwardingClass{
		//TODO(nati): Apply default
		UUID:                   "",
		ForwardingClassDSCP:    MakeDscpValueType(),
		ForwardingClassMPLSExp: MakeMplsExpType(),
		ParentType:             "",
		IDPerms:                MakeIdPermsType(),
		DisplayName:            "",
		Perms2:                 MakePermType2(),
		ForwardingClassVlanPriority: MakeVlanPriorityType(),
		ForwardingClassID:           MakeForwardingClassId(),
		ParentUUID:                  "",
		FQName:                      []string{},
		Annotations:                 MakeKeyValuePairs(),
	}
}

// InterfaceToForwardingClass makes ForwardingClass from interface
func InterfaceToForwardingClass(iData interface{}) *ForwardingClass {
	data := iData.(map[string]interface{})
	return &ForwardingClass{
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ForwardingClassVlanPriority: InterfaceToVlanPriorityType(data["forwarding_class_vlan_priority"]),

		//{"description":"802.1p value to be written on outgoing packet for this forwarding-class.","type":"integer","minimum":0,"maximum":7}
		ForwardingClassID: InterfaceToForwardingClassId(data["forwarding_class_id"]),

		//{"description":"Unique ID for this forwarding class.","default":"0","type":"integer","minimum":0,"maximum":255}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		ForwardingClassDSCP: InterfaceToDscpValueType(data["forwarding_class_dscp"]),

		//{"description":"DSCP value to be written on outgoing packet for this forwarding-class.","type":"integer","minimum":0,"maximum":63}
		ForwardingClassMPLSExp: InterfaceToMplsExpType(data["forwarding_class_mpls_exp"]),

		//{"description":"MPLS exp value to be written on outgoing packet for this forwarding-class.","type":"integer","minimum":0,"maximum":7}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}

	}
}

// InterfaceToForwardingClassSlice makes a slice of ForwardingClass from interface
func InterfaceToForwardingClassSlice(data interface{}) []*ForwardingClass {
	list := data.([]interface{})
	result := MakeForwardingClassSlice()
	for _, item := range list {
		result = append(result, InterfaceToForwardingClass(item))
	}
	return result
}

// MakeForwardingClassSlice() makes a slice of ForwardingClass
func MakeForwardingClassSlice() []*ForwardingClass {
	return []*ForwardingClass{}
}
