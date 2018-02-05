package models

// OpenstackStorageNodeRole

import "encoding/json"

// OpenstackStorageNodeRole
type OpenstackStorageNodeRole struct {
	JournalDrives                      string         `json:"journal_drives,omitempty"`
	StorageAccessBondInterfaceMembers  string         `json:"storage_access_bond_interface_members,omitempty"`
	FQName                             []string       `json:"fq_name,omitempty"`
	ProvisioningLog                    string         `json:"provisioning_log,omitempty"`
	ProvisioningProgress               int            `json:"provisioning_progress,omitempty"`
	DisplayName                        string         `json:"display_name,omitempty"`
	ParentType                         string         `json:"parent_type,omitempty"`
	Annotations                        *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                             *PermType2     `json:"perms2,omitempty"`
	ProvisioningState                  string         `json:"provisioning_state,omitempty"`
	OsdDrives                          string         `json:"osd_drives,omitempty"`
	StorageBackendBondInterfaceMembers string         `json:"storage_backend_bond_interface_members,omitempty"`
	ParentUUID                         string         `json:"parent_uuid,omitempty"`
	IDPerms                            *IdPermsType   `json:"id_perms,omitempty"`
	UUID                               string         `json:"uuid,omitempty"`
	ProvisioningProgressStage          string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningStartTime              string         `json:"provisioning_start_time,omitempty"`
}

// String returns json representation of the object
func (model *OpenstackStorageNodeRole) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeOpenstackStorageNodeRole makes OpenstackStorageNodeRole
func MakeOpenstackStorageNodeRole() *OpenstackStorageNodeRole {
	return &OpenstackStorageNodeRole{
		//TODO(nati): Apply default
		Perms2:            MakePermType2(),
		ProvisioningState: "",
		ParentType:        "",
		Annotations:       MakeKeyValuePairs(),
		ParentUUID:        "",
		IDPerms:           MakeIdPermsType(),
		UUID:              "",
		ProvisioningProgressStage:          "",
		ProvisioningStartTime:              "",
		OsdDrives:                          "",
		StorageBackendBondInterfaceMembers: "",
		FQName:                            []string{},
		ProvisioningLog:                   "",
		ProvisioningProgress:              0,
		JournalDrives:                     "",
		StorageAccessBondInterfaceMembers: "",
		DisplayName:                       "",
	}
}

// MakeOpenstackStorageNodeRoleSlice() makes a slice of OpenstackStorageNodeRole
func MakeOpenstackStorageNodeRoleSlice() []*OpenstackStorageNodeRole {
	return []*OpenstackStorageNodeRole{}
}
