package models

// PortTuple

import "encoding/json"

// PortTuple
type PortTuple struct {
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
}

// String returns json representation of the object
func (model *PortTuple) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakePortTuple makes PortTuple
func MakePortTuple() *PortTuple {
	return &PortTuple{
		//TODO(nati): Apply default
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
	}
}

// MakePortTupleSlice() makes a slice of PortTuple
func MakePortTupleSlice() []*PortTuple {
	return []*PortTuple{}
}
