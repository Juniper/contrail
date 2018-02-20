package models

// BaremetalPort

// BaremetalPort
//proteus:generate
type BaremetalPort struct {
	UUID                string               `json:"uuid,omitempty"`
	ParentUUID          string               `json:"parent_uuid,omitempty"`
	ParentType          string               `json:"parent_type,omitempty"`
	FQName              []string             `json:"fq_name,omitempty"`
	IDPerms             *IdPermsType         `json:"id_perms,omitempty"`
	DisplayName         string               `json:"display_name,omitempty"`
	Annotations         *KeyValuePairs       `json:"annotations,omitempty"`
	Perms2              *PermType2           `json:"perms2,omitempty"`
	MacAddress          string               `json:"mac_address,omitempty"`
	CreatedAt           string               `json:"created_at,omitempty"`
	UpdatedAt           string               `json:"updated_at,omitempty"`
	Node                string               `json:"node,omitempty"`
	PxeEnabled          bool                 `json:"pxe_enabled"`
	LocalLinkConnection *LocalLinkConnection `json:"local_link_connection,omitempty"`
}

// MakeBaremetalPort makes BaremetalPort
func MakeBaremetalPort() *BaremetalPort {
	return &BaremetalPort{
		//TODO(nati): Apply default
		UUID:                "",
		ParentUUID:          "",
		ParentType:          "",
		FQName:              []string{},
		IDPerms:             MakeIdPermsType(),
		DisplayName:         "",
		Annotations:         MakeKeyValuePairs(),
		Perms2:              MakePermType2(),
		MacAddress:          "",
		CreatedAt:           "",
		UpdatedAt:           "",
		Node:                "",
		PxeEnabled:          false,
		LocalLinkConnection: MakeLocalLinkConnection(),
	}
}

// MakeBaremetalPortSlice() makes a slice of BaremetalPort
func MakeBaremetalPortSlice() []*BaremetalPort {
	return []*BaremetalPort{}
}
