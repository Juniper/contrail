package models

// ControllerNodeRole

import "encoding/json"

// ControllerNodeRole
type ControllerNodeRole struct {
	ProvisioningStartTime                 string         `json:"provisioning_start_time,omitempty"`
	Annotations                           *KeyValuePairs `json:"annotations,omitempty"`
	UUID                                  string         `json:"uuid,omitempty"`
	ProvisioningProgress                  int            `json:"provisioning_progress,omitempty"`
	ProvisioningProgressStage             string         `json:"provisioning_progress_stage,omitempty"`
	StorageManagementBondInterfaceMembers string         `json:"storage_management_bond_interface_members,omitempty"`
	IDPerms                               *IdPermsType   `json:"id_perms,omitempty"`
	Perms2                                *PermType2     `json:"perms2,omitempty"`
	InternalapiBondInterfaceMembers       string         `json:"internalapi_bond_interface_members,omitempty"`
	ParentType                            string         `json:"parent_type,omitempty"`
	FQName                                []string       `json:"fq_name,omitempty"`
	ProvisioningLog                       string         `json:"provisioning_log,omitempty"`
	ProvisioningState                     string         `json:"provisioning_state,omitempty"`
	CapacityDrives                        string         `json:"capacity_drives,omitempty"`
	PerformanceDrives                     string         `json:"performance_drives,omitempty"`
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
		ProvisioningProgress:                  0,
		ProvisioningProgressStage:             "",
		ProvisioningStartTime:                 "",
		StorageManagementBondInterfaceMembers: "",
		IDPerms: MakeIdPermsType(),
		Perms2:  MakePermType2(),
		InternalapiBondInterfaceMembers: "",
		ParentType:                      "",
		ProvisioningState:               "",
		CapacityDrives:                  "",
		PerformanceDrives:               "",
		DisplayName:                     "",
		ParentUUID:                      "",
		FQName:                          []string{},
		ProvisioningLog:                 "",
	}
}

// MakeControllerNodeRoleSlice() makes a slice of ControllerNodeRole
func MakeControllerNodeRoleSlice() []*ControllerNodeRole {
	return []*ControllerNodeRole{}
}
