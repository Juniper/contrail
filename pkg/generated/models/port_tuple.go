package models

// PortTuple

import "encoding/json"

// PortTuple
type PortTuple struct {
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
	UUID        string         `json:"uuid"`
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`
	FQName      []string       `json:"fq_name"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	DisplayName string         `json:"display_name"`
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
