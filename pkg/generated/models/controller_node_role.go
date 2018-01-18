package models

// ControllerNodeRole

import "encoding/json"

// ControllerNodeRole
type ControllerNodeRole struct {
	PerformanceDrives                     string         `json:"performance_drives,omitempty"`
	Annotations                           *KeyValuePairs `json:"annotations,omitempty"`
	UUID                                  string         `json:"uuid,omitempty"`
	ParentUUID                            string         `json:"parent_uuid,omitempty"`
	ProvisioningProgress                  int            `json:"provisioning_progress,omitempty"`
	CapacityDrives                        string         `json:"capacity_drives,omitempty"`
	InternalapiBondInterfaceMembers       string         `json:"internalapi_bond_interface_members,omitempty"`
	FQName                                []string       `json:"fq_name,omitempty"`
	ProvisioningProgressStage             string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningStartTime                 string         `json:"provisioning_start_time,omitempty"`
	ProvisioningState                     string         `json:"provisioning_state,omitempty"`
	ParentType                            string         `json:"parent_type,omitempty"`
	DisplayName                           string         `json:"display_name,omitempty"`
	ProvisioningLog                       string         `json:"provisioning_log,omitempty"`
	StorageManagementBondInterfaceMembers string         `json:"storage_management_bond_interface_members,omitempty"`
	IDPerms                               *IdPermsType   `json:"id_perms,omitempty"`
	Perms2                                *PermType2     `json:"perms2,omitempty"`
}

// String returns json representation of the object
func (model *ControllerNodeRole) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeControllerNodeRole makes ControllerNodeRole
func MakeControllerNodeRole() *ControllerNodeRole {
	return &ControllerNodeRole{
		//TODO(nati): Apply default
		DisplayName:               "",
		ProvisioningLog:           "",
		ProvisioningProgressStage: "",
		ProvisioningStartTime:     "",
		ProvisioningState:         "",
		ParentType:                "",
		IDPerms:                   MakeIdPermsType(),
		Perms2:                    MakePermType2(),
		StorageManagementBondInterfaceMembers: "",
		Annotations:                           MakeKeyValuePairs(),
		PerformanceDrives:                     "",
		InternalapiBondInterfaceMembers:       "",
		FQName:               []string{},
		UUID:                 "",
		ParentUUID:           "",
		ProvisioningProgress: 0,
		CapacityDrives:       "",
	}
}

// MakeControllerNodeRoleSlice() makes a slice of ControllerNodeRole
func MakeControllerNodeRoleSlice() []*ControllerNodeRole {
	return []*ControllerNodeRole{}
}
