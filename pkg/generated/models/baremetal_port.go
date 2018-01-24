package models

// BaremetalPort

import "encoding/json"

// BaremetalPort
type BaremetalPort struct {
	SwitchInfo  string         `json:"switch_info,omitempty"`
	PxeEnabled  bool           `json:"pxe_enabled"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	MacAddress  string         `json:"mac_address,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	Node        string         `json:"node,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	SwitchID    string         `json:"switch_id,omitempty"`
	PortID      string         `json:"port_id,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
}

// String returns json representation of the object
func (model *BaremetalPort) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeBaremetalPort makes BaremetalPort
func MakeBaremetalPort() *BaremetalPort {
	return &BaremetalPort{
		//TODO(nati): Apply default
		MacAddress:  "",
		DisplayName: "",
		UUID:        "",
		FQName:      []string{},
		Node:        "",
		ParentType:  "",
		SwitchID:    "",
		PortID:      "",
		ParentUUID:  "",
		IDPerms:     MakeIdPermsType(),
		SwitchInfo:  "",
		PxeEnabled:  false,
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
	}
}

// MakeBaremetalPortSlice() makes a slice of BaremetalPort
func MakeBaremetalPortSlice() []*BaremetalPort {
	return []*BaremetalPort{}
}
