package models

// ControllerNodeRole

import "encoding/json"

// ControllerNodeRole
type ControllerNodeRole struct {
	InternalapiBondInterfaceMembers       string         `json:"internalapi_bond_interface_members,omitempty"`
	ProvisioningProgressStage             string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningState                     string         `json:"provisioning_state,omitempty"`
	CapacityDrives                        string         `json:"capacity_drives,omitempty"`
	StorageManagementBondInterfaceMembers string         `json:"storage_management_bond_interface_members,omitempty"`
	DisplayName                           string         `json:"display_name,omitempty"`
	Perms2                                *PermType2     `json:"perms2,omitempty"`
	ParentType                            string         `json:"parent_type,omitempty"`
	FQName                                []string       `json:"fq_name,omitempty"`
	IDPerms                               *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningLog                       string         `json:"provisioning_log,omitempty"`
	PerformanceDrives                     string         `json:"performance_drives,omitempty"`
	Annotations                           *KeyValuePairs `json:"annotations,omitempty"`
	UUID                                  string         `json:"uuid,omitempty"`
	ParentUUID                            string         `json:"parent_uuid,omitempty"`
	ProvisioningProgress                  int            `json:"provisioning_progress,omitempty"`
	ProvisioningStartTime                 string         `json:"provisioning_start_time,omitempty"`
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
		InternalapiBondInterfaceMembers:       "",
		StorageManagementBondInterfaceMembers: "",
		DisplayName:                           "",
		Perms2:                                MakePermType2(),
		ProvisioningProgressStage: "",
		ProvisioningState:         "",
		CapacityDrives:            "",
		Annotations:               MakeKeyValuePairs(),
		UUID:                      "",
		ParentUUID:                "",
		ParentType:                "",
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		ProvisioningLog:           "",
		PerformanceDrives:         "",
		ProvisioningStartTime:     "",
		ProvisioningProgress:      0,
	}
}

// MakeControllerNodeRoleSlice() makes a slice of ControllerNodeRole
func MakeControllerNodeRoleSlice() []*ControllerNodeRole {
	return []*ControllerNodeRole{}
}
