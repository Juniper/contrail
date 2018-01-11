package models

// ForwardingClass

import "encoding/json"

// ForwardingClass
type ForwardingClass struct {
	FQName                      []string          `json:"fq_name"`
	Annotations                 *KeyValuePairs    `json:"annotations"`
	UUID                        string            `json:"uuid"`
	ForwardingClassDSCP         DscpValueType     `json:"forwarding_class_dscp"`
	ForwardingClassID           ForwardingClassId `json:"forwarding_class_id"`
	ParentUUID                  string            `json:"parent_uuid"`
	ParentType                  string            `json:"parent_type"`
	Perms2                      *PermType2        `json:"perms2"`
	ForwardingClassVlanPriority VlanPriorityType  `json:"forwarding_class_vlan_priority"`
	ForwardingClassMPLSExp      MplsExpType       `json:"forwarding_class_mpls_exp"`
	IDPerms                     *IdPermsType      `json:"id_perms"`
	DisplayName                 string            `json:"display_name"`

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
		ForwardingClassMPLSExp: MakeMplsExpType(),
		IDPerms:                MakeIdPermsType(),
		DisplayName:            "",
		Perms2:                 MakePermType2(),
		ForwardingClassVlanPriority: MakeVlanPriorityType(),
		ForwardingClassID:           MakeForwardingClassId(),
		ParentUUID:                  "",
		ParentType:                  "",
		FQName:                      []string{},
		Annotations:                 MakeKeyValuePairs(),
		UUID:                        "",
		ForwardingClassDSCP:         MakeDscpValueType(),
	}
}

// MakeForwardingClassSlice() makes a slice of ForwardingClass
func MakeForwardingClassSlice() []*ForwardingClass {
	return []*ForwardingClass{}
}
