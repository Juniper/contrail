package models

// OpenstackStorageNodeRole

import "encoding/json"

// OpenstackStorageNodeRole
type OpenstackStorageNodeRole struct {
	ProvisioningLog                    string         `json:"provisioning_log,omitempty"`
	ProvisioningState                  string         `json:"provisioning_state,omitempty"`
	ParentType                         string         `json:"parent_type,omitempty"`
	IDPerms                            *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningProgressStage          string         `json:"provisioning_progress_stage,omitempty"`
	OsdDrives                          string         `json:"osd_drives,omitempty"`
	DisplayName                        string         `json:"display_name,omitempty"`
	UUID                               string         `json:"uuid,omitempty"`
	ProvisioningProgress               int            `json:"provisioning_progress,omitempty"`
	StorageAccessBondInterfaceMembers  string         `json:"storage_access_bond_interface_members,omitempty"`
	StorageBackendBondInterfaceMembers string         `json:"storage_backend_bond_interface_members,omitempty"`
	Annotations                        *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                             *PermType2     `json:"perms2,omitempty"`
	ParentUUID                         string         `json:"parent_uuid,omitempty"`
	FQName                             []string       `json:"fq_name,omitempty"`
	ProvisioningStartTime              string         `json:"provisioning_start_time,omitempty"`
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
		Perms2:                             MakePermType2(),
		ParentUUID:                         "",
		FQName:                             []string{},
		ProvisioningStartTime:              "",
		JournalDrives:                      "",
		StorageBackendBondInterfaceMembers: "",
		Annotations:                        MakeKeyValuePairs(),
		ParentType:                         "",
		ProvisioningLog:                    "",
		ProvisioningState:                  "",
		OsdDrives:                          "",
		IDPerms:                            MakeIdPermsType(),
		ProvisioningProgressStage:          "",
		ProvisioningProgress:               0,
		StorageAccessBondInterfaceMembers:  "",
		DisplayName:                        "",
		UUID:                               "",
	}
}

// MakeOpenstackStorageNodeRoleSlice() makes a slice of OpenstackStorageNodeRole
func MakeOpenstackStorageNodeRoleSlice() []*OpenstackStorageNodeRole {
	return []*OpenstackStorageNodeRole{}
}
