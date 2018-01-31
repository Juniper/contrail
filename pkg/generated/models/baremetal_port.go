package models

// BaremetalPort

import "encoding/json"

// BaremetalPort
type BaremetalPort struct {
	MacAddress  string         `json:"mac_address,omitempty"`
	Node        string         `json:"node,omitempty"`
	PxeEnabled  bool           `json:"pxe_enabled"`
	DisplayName string         `json:"display_name,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	PortID      string         `json:"port_id,omitempty"`
	SwitchInfo  string         `json:"switch_info,omitempty"`
	SwitchID    string         `json:"switch_id,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
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
		Perms2:      MakePermType2(),
		MacAddress:  "",
		Node:        "",
		PxeEnabled:  false,
		DisplayName: "",
		PortID:      "",
		SwitchInfo:  "",
		SwitchID:    "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		Annotations: MakeKeyValuePairs(),
		UUID:        "",
		ParentUUID:  "",
	}
}

// MakeBaremetalPortSlice() makes a slice of BaremetalPort
func MakeBaremetalPortSlice() []*BaremetalPort {
	return []*BaremetalPort{}
}
