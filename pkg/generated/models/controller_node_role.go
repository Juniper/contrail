package models

// ControllerNodeRole

import "encoding/json"

// ControllerNodeRole
type ControllerNodeRole struct {
	InternalapiBondInterfaceMembers       string         `json:"internalapi_bond_interface_members,omitempty"`
	ParentUUID                            string         `json:"parent_uuid,omitempty"`
	ProvisioningLog                       string         `json:"provisioning_log,omitempty"`
	ProvisioningState                     string         `json:"provisioning_state,omitempty"`
	StorageManagementBondInterfaceMembers string         `json:"storage_management_bond_interface_members,omitempty"`
	UUID                                  string         `json:"uuid,omitempty"`
	Annotations                           *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                                *PermType2     `json:"perms2,omitempty"`
	ProvisioningProgressStage             string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningStartTime                 string         `json:"provisioning_start_time,omitempty"`
	CapacityDrives                        string         `json:"capacity_drives,omitempty"`
	PerformanceDrives                     string         `json:"performance_drives,omitempty"`
	IDPerms                               *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName                           string         `json:"display_name,omitempty"`
	ProvisioningProgress                  int            `json:"provisioning_progress,omitempty"`
	ParentType                            string         `json:"parent_type,omitempty"`
	FQName                                []string       `json:"fq_name,omitempty"`
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
		Perms2: MakePermType2(),
		ProvisioningProgressStage:             "",
		ProvisioningStartTime:                 "",
		StorageManagementBondInterfaceMembers: "",
		UUID:                            "",
		Annotations:                     MakeKeyValuePairs(),
		DisplayName:                     "",
		ProvisioningProgress:            0,
		CapacityDrives:                  "",
		PerformanceDrives:               "",
		IDPerms:                         MakeIdPermsType(),
		ParentType:                      "",
		FQName:                          []string{},
		ProvisioningState:               "",
		InternalapiBondInterfaceMembers: "",
		ParentUUID:                      "",
		ProvisioningLog:                 "",
	}
}

// MakeControllerNodeRoleSlice() makes a slice of ControllerNodeRole
func MakeControllerNodeRoleSlice() []*ControllerNodeRole {
	return []*ControllerNodeRole{}
}
