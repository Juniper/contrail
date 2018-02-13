package models

// BaremetalPort

import "encoding/json"

// BaremetalPort
//proteus:generate
type BaremetalPort struct {
	UUID        string         `json:"uuid,omitempty"`
	ParentUUID  string         `json:"parent_uuid,omitempty"`
	ParentType  string         `json:"parent_type,omitempty"`
	FQName      []string       `json:"fq_name,omitempty"`
	IDPerms     *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName string         `json:"display_name,omitempty"`
	Annotations *KeyValuePairs `json:"annotations,omitempty"`
	Perms2      *PermType2     `json:"perms2,omitempty"`
	MacAddress  string         `json:"mac_address,omitempty"`
	Node        string         `json:"node,omitempty"`
	SwitchID    string         `json:"switch_id,omitempty"`
	PortID      string         `json:"port_id,omitempty"`
	SwitchInfo  string         `json:"switch_info,omitempty"`
	PxeEnabled  bool           `json:"pxe_enabled"`
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
		UUID:        "",
		ParentUUID:  "",
		ParentType:  "",
		FQName:      []string{},
		IDPerms:     MakeIdPermsType(),
		DisplayName: "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		MacAddress:  "",
		Node:        "",
		SwitchID:    "",
		PortID:      "",
		SwitchInfo:  "",
		PxeEnabled:  false,
	}
}

// MakeBaremetalPortSlice() makes a slice of BaremetalPort
func MakeBaremetalPortSlice() []*BaremetalPort {
	return []*BaremetalPort{}
}
