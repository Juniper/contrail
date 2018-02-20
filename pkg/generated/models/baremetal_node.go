package models

// BaremetalNode

// BaremetalNode
//proteus:generate
type BaremetalNode struct {
	UUID                 string               `json:"uuid,omitempty"`
	ParentUUID           string               `json:"parent_uuid,omitempty"`
	ParentType           string               `json:"parent_type,omitempty"`
	FQName               []string             `json:"fq_name,omitempty"`
	IDPerms              *IdPermsType         `json:"id_perms,omitempty"`
	DisplayName          string               `json:"display_name,omitempty"`
	Annotations          *KeyValuePairs       `json:"annotations,omitempty"`
	Perms2               *PermType2           `json:"perms2,omitempty"`
	Name                 string               `json:"name,omitempty"`
	DriverInfo           *DriverInfo          `json:"driver_info,omitempty"`
	BMProperties         *BaremetalProperties `json:"bm_properties,omitempty"`
	InstanceUUID         string               `json:"instance_uuid,omitempty"`
	InstanceInfo         *InstanceInfo        `json:"instance_info,omitempty"`
	Maintenance          bool                 `json:"maintenance"`
	MaintenanceReason    string               `json:"maintenance_reason,omitempty"`
	PowerState           string               `json:"power_state,omitempty"`
	TargetPowerState     string               `json:"target_power_state,omitempty"`
	ProvisionState       string               `json:"provision_state,omitempty"`
	TargetProvisionState string               `json:"target_provision_state,omitempty"`
	ConsoleEnabled       bool                 `json:"console_enabled"`
	CreatedAt            string               `json:"created_at,omitempty"`
	UpdatedAt            string               `json:"updated_at,omitempty"`
	LastError            string               `json:"last_error,omitempty"`
}

// MakeBaremetalNode makes BaremetalNode
func MakeBaremetalNode() *BaremetalNode {
	return &BaremetalNode{
		//TODO(nati): Apply default
		UUID:                 "",
		ParentUUID:           "",
		ParentType:           "",
		FQName:               []string{},
		IDPerms:              MakeIdPermsType(),
		DisplayName:          "",
		Annotations:          MakeKeyValuePairs(),
		Perms2:               MakePermType2(),
		Name:                 "",
		DriverInfo:           MakeDriverInfo(),
		BMProperties:         MakeBaremetalProperties(),
		InstanceUUID:         "",
		InstanceInfo:         MakeInstanceInfo(),
		Maintenance:          false,
		MaintenanceReason:    "",
		PowerState:           "",
		TargetPowerState:     "",
		ProvisionState:       "",
		TargetProvisionState: "",
		ConsoleEnabled:       false,
		CreatedAt:            "",
		UpdatedAt:            "",
		LastError:            "",
	}
}

// MakeBaremetalNodeSlice() makes a slice of BaremetalNode
func MakeBaremetalNodeSlice() []*BaremetalNode {
	return []*BaremetalNode{}
}
