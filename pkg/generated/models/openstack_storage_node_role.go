package models

// OpenstackStorageNodeRole

import "encoding/json"

// OpenstackStorageNodeRole
type OpenstackStorageNodeRole struct {
	FQName                             []string       `json:"fq_name,omitempty"`
	OsdDrives                          string         `json:"osd_drives,omitempty"`
	StorageBackendBondInterfaceMembers string         `json:"storage_backend_bond_interface_members,omitempty"`
	Perms2                             *PermType2     `json:"perms2,omitempty"`
	UUID                               string         `json:"uuid,omitempty"`
	ParentUUID                         string         `json:"parent_uuid,omitempty"`
	ProvisioningState                  string         `json:"provisioning_state,omitempty"`
	DisplayName                        string         `json:"display_name,omitempty"`
	Annotations                        *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningProgressStage          string         `json:"provisioning_progress_stage,omitempty"`
	JournalDrives                      string         `json:"journal_drives,omitempty"`
	StorageAccessBondInterfaceMembers  string         `json:"storage_access_bond_interface_members,omitempty"`
	IDPerms                            *IdPermsType   `json:"id_perms,omitempty"`
	ParentType                         string         `json:"parent_type,omitempty"`
	ProvisioningLog                    string         `json:"provisioning_log,omitempty"`
	ProvisioningProgress               int            `json:"provisioning_progress,omitempty"`
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
		DisplayName:                       "",
		Annotations:                       MakeKeyValuePairs(),
		ProvisioningProgressStage:         "",
		ProvisioningStartTime:             "",
		JournalDrives:                     "",
		StorageAccessBondInterfaceMembers: "",
		IDPerms:                            MakeIdPermsType(),
		ParentType:                         "",
		ProvisioningLog:                    "",
		ProvisioningProgress:               0,
		FQName:                             []string{},
		OsdDrives:                          "",
		StorageBackendBondInterfaceMembers: "",
		Perms2:            MakePermType2(),
		UUID:              "",
		ParentUUID:        "",
		ProvisioningState: "",
	}
}

// MakeOpenstackStorageNodeRoleSlice() makes a slice of OpenstackStorageNodeRole
func MakeOpenstackStorageNodeRoleSlice() []*OpenstackStorageNodeRole {
	return []*OpenstackStorageNodeRole{}
}
