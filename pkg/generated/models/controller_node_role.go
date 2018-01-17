package models

// ControllerNodeRole

import "encoding/json"

// ControllerNodeRole
type ControllerNodeRole struct {
	InternalapiBondInterfaceMembers       string         `json:"internalapi_bond_interface_members,omitempty"`
	Perms2                                *PermType2     `json:"perms2,omitempty"`
	UUID                                  string         `json:"uuid,omitempty"`
	FQName                                []string       `json:"fq_name,omitempty"`
	ProvisioningProgress                  int            `json:"provisioning_progress,omitempty"`
	ParentUUID                            string         `json:"parent_uuid,omitempty"`
	PerformanceDrives                     string         `json:"performance_drives,omitempty"`
	StorageManagementBondInterfaceMembers string         `json:"storage_management_bond_interface_members,omitempty"`
	Annotations                           *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningLog                       string         `json:"provisioning_log,omitempty"`
	ProvisioningProgressStage             string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningState                     string         `json:"provisioning_state,omitempty"`
	CapacityDrives                        string         `json:"capacity_drives,omitempty"`
	DisplayName                           string         `json:"display_name,omitempty"`
	ParentType                            string         `json:"parent_type,omitempty"`
	IDPerms                               *IdPermsType   `json:"id_perms,omitempty"`
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
		InternalapiBondInterfaceMembers: "",
		Perms2:                                MakePermType2(),
		UUID:                                  "",
		FQName:                                []string{},
		ProvisioningProgress:                  0,
		ParentUUID:                            "",
		PerformanceDrives:                     "",
		StorageManagementBondInterfaceMembers: "",
		Annotations:                           MakeKeyValuePairs(),
		ProvisioningLog:                       "",
		ProvisioningProgressStage:             "",
		ProvisioningState:                     "",
		CapacityDrives:                        "",
		DisplayName:                           "",
		ParentType:                            "",
		IDPerms:                               MakeIdPermsType(),
		ProvisioningStartTime:                 "",
	}
}

// MakeControllerNodeRoleSlice() makes a slice of ControllerNodeRole
func MakeControllerNodeRoleSlice() []*ControllerNodeRole {
	return []*ControllerNodeRole{}
}
