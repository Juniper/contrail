package models

// OpenstackStorageNodeRole

import "encoding/json"

// OpenstackStorageNodeRole
type OpenstackStorageNodeRole struct {
	UUID                               string         `json:"uuid,omitempty"`
	ProvisioningLog                    string         `json:"provisioning_log,omitempty"`
	ProvisioningProgressStage          string         `json:"provisioning_progress_stage,omitempty"`
	OsdDrives                          string         `json:"osd_drives,omitempty"`
	FQName                             []string       `json:"fq_name,omitempty"`
	ParentUUID                         string         `json:"parent_uuid,omitempty"`
	ParentType                         string         `json:"parent_type,omitempty"`
	IDPerms                            *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName                        string         `json:"display_name,omitempty"`
	ProvisioningStartTime              string         `json:"provisioning_start_time,omitempty"`
	ProvisioningState                  string         `json:"provisioning_state,omitempty"`
	Annotations                        *KeyValuePairs `json:"annotations,omitempty"`
	StorageAccessBondInterfaceMembers  string         `json:"storage_access_bond_interface_members,omitempty"`
	StorageBackendBondInterfaceMembers string         `json:"storage_backend_bond_interface_members,omitempty"`
	Perms2                             *PermType2     `json:"perms2,omitempty"`
	ProvisioningProgress               int            `json:"provisioning_progress,omitempty"`
	JournalDrives                      string         `json:"journal_drives,omitempty"`
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
		DisplayName:                        "",
		ProvisioningStartTime:              "",
		ProvisioningState:                  "",
		Annotations:                        MakeKeyValuePairs(),
		ParentType:                         "",
		IDPerms:                            MakeIdPermsType(),
		Perms2:                             MakePermType2(),
		ProvisioningProgress:               0,
		JournalDrives:                      "",
		StorageAccessBondInterfaceMembers:  "",
		StorageBackendBondInterfaceMembers: "",
		ProvisioningProgressStage:          "",
		OsdDrives:                          "",
		UUID:                               "",
		ProvisioningLog:                    "",
		ParentUUID:                         "",
		FQName:                             []string{},
	}
}

// MakeOpenstackStorageNodeRoleSlice() makes a slice of OpenstackStorageNodeRole
func MakeOpenstackStorageNodeRoleSlice() []*OpenstackStorageNodeRole {
	return []*OpenstackStorageNodeRole{}
}
