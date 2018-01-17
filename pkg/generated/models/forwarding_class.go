package models

// ForwardingClass

import "encoding/json"

// ForwardingClass
type ForwardingClass struct {
	ForwardingClassVlanPriority VlanPriorityType  `json:"forwarding_class_vlan_priority,omitempty"`
	ForwardingClassID           ForwardingClassId `json:"forwarding_class_id,omitempty"`
	Perms2                      *PermType2        `json:"perms2,omitempty"`
	ParentUUID                  string            `json:"parent_uuid,omitempty"`
	ParentType                  string            `json:"parent_type,omitempty"`
	ForwardingClassDSCP         DscpValueType     `json:"forwarding_class_dscp,omitempty"`
	ForwardingClassMPLSExp      MplsExpType       `json:"forwarding_class_mpls_exp,omitempty"`
	IDPerms                     *IdPermsType      `json:"id_perms,omitempty"`
	DisplayName                 string            `json:"display_name,omitempty"`
	Annotations                 *KeyValuePairs    `json:"annotations,omitempty"`
	UUID                        string            `json:"uuid,omitempty"`
	FQName                      []string          `json:"fq_name,omitempty"`

	QosQueueRefs []*ForwardingClassQosQueueRef `json:"qos_queue_refs,omitempty"`
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
		FQName:                 []string{},
		ForwardingClassDSCP:    MakeDscpValueType(),
		ForwardingClassMPLSExp: MakeMplsExpType(),
		IDPerms:                MakeIdPermsType(),
		DisplayName:            "",
		Annotations:            MakeKeyValuePairs(),
		UUID:                   "",
		ForwardingClassVlanPriority: MakeVlanPriorityType(),
		ForwardingClassID:           MakeForwardingClassId(),
		Perms2:                      MakePermType2(),
		ParentUUID:                  "",
		ParentType:                  "",
	}
}

// MakeForwardingClassSlice() makes a slice of ForwardingClass
func MakeForwardingClassSlice() []*ForwardingClass {
	return []*ForwardingClass{}
}
