package models

// ControllerNodeRole

import "encoding/json"

// ControllerNodeRole
type ControllerNodeRole struct {
	ProvisioningLog                       string         `json:"provisioning_log,omitempty"`
	ProvisioningProgress                  int            `json:"provisioning_progress,omitempty"`
	InternalapiBondInterfaceMembers       string         `json:"internalapi_bond_interface_members,omitempty"`
	Perms2                                *PermType2     `json:"perms2,omitempty"`
	ProvisioningStartTime                 string         `json:"provisioning_start_time,omitempty"`
	ProvisioningState                     string         `json:"provisioning_state,omitempty"`
	IDPerms                               *IdPermsType   `json:"id_perms,omitempty"`
	Annotations                           *KeyValuePairs `json:"annotations,omitempty"`
	UUID                                  string         `json:"uuid,omitempty"`
	FQName                                []string       `json:"fq_name,omitempty"`
	PerformanceDrives                     string         `json:"performance_drives,omitempty"`
	ProvisioningProgressStage             string         `json:"provisioning_progress_stage,omitempty"`
	ParentType                            string         `json:"parent_type,omitempty"`
	CapacityDrives                        string         `json:"capacity_drives,omitempty"`
	StorageManagementBondInterfaceMembers string         `json:"storage_management_bond_interface_members,omitempty"`
	DisplayName                           string         `json:"display_name,omitempty"`
	ParentUUID                            string         `json:"parent_uuid,omitempty"`
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
		Annotations:                           MakeKeyValuePairs(),
		UUID:                                  "",
		FQName:                                []string{},
		IDPerms:                               MakeIdPermsType(),
		ProvisioningProgressStage:             "",
		PerformanceDrives:                     "",
		StorageManagementBondInterfaceMembers: "",
		DisplayName:                           "",
		ParentUUID:                            "",
		ParentType:                            "",
		CapacityDrives:                        "",
		Perms2:                                MakePermType2(),
		ProvisioningStartTime:                 "",
		ProvisioningState:                     "",
		ProvisioningLog:                       "",
		ProvisioningProgress:                  0,
		InternalapiBondInterfaceMembers:       "",
	}
}

// MakeControllerNodeRoleSlice() makes a slice of ControllerNodeRole
func MakeControllerNodeRoleSlice() []*ControllerNodeRole {
	return []*ControllerNodeRole{}
}
