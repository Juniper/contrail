package models

// ControllerNodeRole

import "encoding/json"

// ControllerNodeRole
type ControllerNodeRole struct {
	InternalapiBondInterfaceMembers       string         `json:"internalapi_bond_interface_members,omitempty"`
	FQName                                []string       `json:"fq_name,omitempty"`
	ProvisioningProgress                  int            `json:"provisioning_progress,omitempty"`
	UUID                                  string         `json:"uuid,omitempty"`
	IDPerms                               *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningProgressStage             string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningStartTime                 string         `json:"provisioning_start_time,omitempty"`
	ProvisioningState                     string         `json:"provisioning_state,omitempty"`
	CapacityDrives                        string         `json:"capacity_drives,omitempty"`
	PerformanceDrives                     string         `json:"performance_drives,omitempty"`
	Perms2                                *PermType2     `json:"perms2,omitempty"`
	ParentUUID                            string         `json:"parent_uuid,omitempty"`
	ParentType                            string         `json:"parent_type,omitempty"`
	StorageManagementBondInterfaceMembers string         `json:"storage_management_bond_interface_members,omitempty"`
	DisplayName                           string         `json:"display_name,omitempty"`
	Annotations                           *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningLog                       string         `json:"provisioning_log,omitempty"`
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
		ProvisioningState:                     "",
		CapacityDrives:                        "",
		PerformanceDrives:                     "",
		Perms2:                                MakePermType2(),
		ParentUUID:                            "",
		ParentType:                            "",
		ProvisioningProgressStage:             "",
		ProvisioningStartTime:                 "",
		StorageManagementBondInterfaceMembers: "",
		DisplayName:                           "",
		Annotations:                           MakeKeyValuePairs(),
		ProvisioningLog:                       "",
		InternalapiBondInterfaceMembers:       "",
		FQName:               []string{},
		ProvisioningProgress: 0,
		UUID:                 "",
		IDPerms:              MakeIdPermsType(),
	}
}

// MakeControllerNodeRoleSlice() makes a slice of ControllerNodeRole
func MakeControllerNodeRoleSlice() []*ControllerNodeRole {
	return []*ControllerNodeRole{}
}
