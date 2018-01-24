package models

// ControllerNodeRole

import "encoding/json"

// ControllerNodeRole
type ControllerNodeRole struct {
	FQName                                []string       `json:"fq_name,omitempty"`
	IDPerms                               *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName                           string         `json:"display_name,omitempty"`
	UUID                                  string         `json:"uuid,omitempty"`
	ProvisioningProgress                  int            `json:"provisioning_progress,omitempty"`
	CapacityDrives                        string         `json:"capacity_drives,omitempty"`
	ParentType                            string         `json:"parent_type,omitempty"`
	Annotations                           *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                                *PermType2     `json:"perms2,omitempty"`
	PerformanceDrives                     string         `json:"performance_drives,omitempty"`
	ParentUUID                            string         `json:"parent_uuid,omitempty"`
	ProvisioningLog                       string         `json:"provisioning_log,omitempty"`
	ProvisioningProgressStage             string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningStartTime                 string         `json:"provisioning_start_time,omitempty"`
	InternalapiBondInterfaceMembers       string         `json:"internalapi_bond_interface_members,omitempty"`
	StorageManagementBondInterfaceMembers string         `json:"storage_management_bond_interface_members,omitempty"`
	ProvisioningState                     string         `json:"provisioning_state,omitempty"`
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
		ProvisioningLog:                       "",
		ProvisioningProgressStage:             "",
		ProvisioningStartTime:                 "",
		PerformanceDrives:                     "",
		ParentUUID:                            "",
		ProvisioningState:                     "",
		InternalapiBondInterfaceMembers:       "",
		StorageManagementBondInterfaceMembers: "",
		DisplayName:                           "",
		UUID:                                  "",
		ProvisioningProgress:                  0,
		FQName:                                []string{},
		IDPerms:                               MakeIdPermsType(),
		Annotations:                           MakeKeyValuePairs(),
		Perms2:                                MakePermType2(),
		CapacityDrives:                        "",
		ParentType:                            "",
	}
}

// MakeControllerNodeRoleSlice() makes a slice of ControllerNodeRole
func MakeControllerNodeRoleSlice() []*ControllerNodeRole {
	return []*ControllerNodeRole{}
}
