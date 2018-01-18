package models

// ForwardingClass

import "encoding/json"

// ForwardingClass
type ForwardingClass struct {
	DisplayName                 string            `json:"display_name,omitempty"`
	ForwardingClassVlanPriority VlanPriorityType  `json:"forwarding_class_vlan_priority,omitempty"`
	ForwardingClassMPLSExp      MplsExpType       `json:"forwarding_class_mpls_exp,omitempty"`
	ForwardingClassID           ForwardingClassId `json:"forwarding_class_id,omitempty"`
	Perms2                      *PermType2        `json:"perms2,omitempty"`
	UUID                        string            `json:"uuid,omitempty"`
	ParentUUID                  string            `json:"parent_uuid,omitempty"`
	ParentType                  string            `json:"parent_type,omitempty"`
	ForwardingClassDSCP         DscpValueType     `json:"forwarding_class_dscp,omitempty"`
	FQName                      []string          `json:"fq_name,omitempty"`
	IDPerms                     *IdPermsType      `json:"id_perms,omitempty"`
	Annotations                 *KeyValuePairs    `json:"annotations,omitempty"`

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
		ForwardingClassDSCP:         MakeDscpValueType(),
		FQName:                      []string{},
		IDPerms:                     MakeIdPermsType(),
		Annotations:                 MakeKeyValuePairs(),
		ParentType:                  "",
		DisplayName:                 "",
		ForwardingClassVlanPriority: MakeVlanPriorityType(),
		ForwardingClassMPLSExp:      MakeMplsExpType(),
		ForwardingClassID:           MakeForwardingClassId(),
		Perms2:                      MakePermType2(),
		UUID:                        "",
		ParentUUID:                  "",
	}
}

// MakeForwardingClassSlice() makes a slice of ForwardingClass
func MakeForwardingClassSlice() []*ForwardingClass {
	return []*ForwardingClass{}
}
