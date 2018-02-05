package models

// ForwardingClass

import "encoding/json"

// ForwardingClass
type ForwardingClass struct {
	UUID                        string            `json:"uuid,omitempty"`
	ForwardingClassDSCP         DscpValueType     `json:"forwarding_class_dscp,omitempty"`
	ForwardingClassID           ForwardingClassId `json:"forwarding_class_id,omitempty"`
	FQName                      []string          `json:"fq_name,omitempty"`
	IDPerms                     *IdPermsType      `json:"id_perms,omitempty"`
	DisplayName                 string            `json:"display_name,omitempty"`
	Annotations                 *KeyValuePairs    `json:"annotations,omitempty"`
	Perms2                      *PermType2        `json:"perms2,omitempty"`
	ParentUUID                  string            `json:"parent_uuid,omitempty"`
	ForwardingClassVlanPriority VlanPriorityType  `json:"forwarding_class_vlan_priority,omitempty"`
	ForwardingClassMPLSExp      MplsExpType       `json:"forwarding_class_mpls_exp,omitempty"`
	ParentType                  string            `json:"parent_type,omitempty"`

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
		UUID:                        "",
		ForwardingClassDSCP:         MakeDscpValueType(),
		ForwardingClassID:           MakeForwardingClassId(),
		FQName:                      []string{},
		IDPerms:                     MakeIdPermsType(),
		DisplayName:                 "",
		Annotations:                 MakeKeyValuePairs(),
		Perms2:                      MakePermType2(),
		ParentUUID:                  "",
		ForwardingClassVlanPriority: MakeVlanPriorityType(),
		ForwardingClassMPLSExp:      MakeMplsExpType(),
		ParentType:                  "",
	}
}

// MakeForwardingClassSlice() makes a slice of ForwardingClass
func MakeForwardingClassSlice() []*ForwardingClass {
	return []*ForwardingClass{}
}
