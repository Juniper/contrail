package models

// OpenstackStorageNodeRole

import "encoding/json"

// OpenstackStorageNodeRole
type OpenstackStorageNodeRole struct {
	ProvisioningProgressStage          string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningLog                    string         `json:"provisioning_log,omitempty"`
	JournalDrives                      string         `json:"journal_drives,omitempty"`
	UUID                               string         `json:"uuid,omitempty"`
	ProvisioningProgress               int            `json:"provisioning_progress,omitempty"`
	StorageBackendBondInterfaceMembers string         `json:"storage_backend_bond_interface_members,omitempty"`
	DisplayName                        string         `json:"display_name,omitempty"`
	ProvisioningState                  string         `json:"provisioning_state,omitempty"`
	FQName                             []string       `json:"fq_name,omitempty"`
	Perms2                             *PermType2     `json:"perms2,omitempty"`
	ParentUUID                         string         `json:"parent_uuid,omitempty"`
	ParentType                         string         `json:"parent_type,omitempty"`
	IDPerms                            *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningStartTime              string         `json:"provisioning_start_time,omitempty"`
	OsdDrives                          string         `json:"osd_drives,omitempty"`
	StorageAccessBondInterfaceMembers  string         `json:"storage_access_bond_interface_members,omitempty"`
	Annotations                        *KeyValuePairs `json:"annotations,omitempty"`
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
		FQName: []string{},
		StorageAccessBondInterfaceMembers:  "",
		Annotations:                        MakeKeyValuePairs(),
		Perms2:                             MakePermType2(),
		ParentUUID:                         "",
		ParentType:                         "",
		IDPerms:                            MakeIdPermsType(),
		ProvisioningStartTime:              "",
		OsdDrives:                          "",
		UUID:                               "",
		ProvisioningProgress:               0,
		ProvisioningProgressStage:          "",
		ProvisioningLog:                    "",
		JournalDrives:                      "",
		DisplayName:                        "",
		ProvisioningState:                  "",
		StorageBackendBondInterfaceMembers: "",
	}
}

// MakeOpenstackStorageNodeRoleSlice() makes a slice of OpenstackStorageNodeRole
func MakeOpenstackStorageNodeRoleSlice() []*OpenstackStorageNodeRole {
	return []*OpenstackStorageNodeRole{}
}
