package models

// ForwardingClass

import "encoding/json"

// ForwardingClass
type ForwardingClass struct {
	ForwardingClassDSCP         DscpValueType     `json:"forwarding_class_dscp,omitempty"`
	ParentType                  string            `json:"parent_type,omitempty"`
	DisplayName                 string            `json:"display_name,omitempty"`
	Annotations                 *KeyValuePairs    `json:"annotations,omitempty"`
	Perms2                      *PermType2        `json:"perms2,omitempty"`
	ForwardingClassVlanPriority VlanPriorityType  `json:"forwarding_class_vlan_priority,omitempty"`
	ForwardingClassMPLSExp      MplsExpType       `json:"forwarding_class_mpls_exp,omitempty"`
	ForwardingClassID           ForwardingClassId `json:"forwarding_class_id,omitempty"`
	UUID                        string            `json:"uuid,omitempty"`
	ParentUUID                  string            `json:"parent_uuid,omitempty"`
	FQName                      []string          `json:"fq_name,omitempty"`
	IDPerms                     *IdPermsType      `json:"id_perms,omitempty"`

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
		Annotations:                 MakeKeyValuePairs(),
		Perms2:                      MakePermType2(),
		ForwardingClassDSCP:         MakeDscpValueType(),
		ParentType:                  "",
		DisplayName:                 "",
		UUID:                        "",
		ParentUUID:                  "",
		FQName:                      []string{},
		IDPerms:                     MakeIdPermsType(),
		ForwardingClassVlanPriority: MakeVlanPriorityType(),
		ForwardingClassMPLSExp:      MakeMplsExpType(),
		ForwardingClassID:           MakeForwardingClassId(),
	}
}

// MakeForwardingClassSlice() makes a slice of ForwardingClass
func MakeForwardingClassSlice() []*ForwardingClass {
	return []*ForwardingClass{}
}
