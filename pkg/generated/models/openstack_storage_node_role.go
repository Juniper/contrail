package models

// OpenstackStorageNodeRole

import "encoding/json"

// OpenstackStorageNodeRole
type OpenstackStorageNodeRole struct {
	ProvisioningStartTime              string         `json:"provisioning_start_time,omitempty"`
	DisplayName                        string         `json:"display_name,omitempty"`
	ProvisioningProgress               int            `json:"provisioning_progress,omitempty"`
	JournalDrives                      string         `json:"journal_drives,omitempty"`
	OsdDrives                          string         `json:"osd_drives,omitempty"`
	UUID                               string         `json:"uuid,omitempty"`
	ParentUUID                         string         `json:"parent_uuid,omitempty"`
	ParentType                         string         `json:"parent_type,omitempty"`
	FQName                             []string       `json:"fq_name,omitempty"`
	ProvisioningState                  string         `json:"provisioning_state,omitempty"`
	StorageBackendBondInterfaceMembers string         `json:"storage_backend_bond_interface_members,omitempty"`
	Annotations                        *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                             *PermType2     `json:"perms2,omitempty"`
	ProvisioningProgressStage          string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningLog                    string         `json:"provisioning_log,omitempty"`
	StorageAccessBondInterfaceMembers  string         `json:"storage_access_bond_interface_members,omitempty"`
	IDPerms                            *IdPermsType   `json:"id_perms,omitempty"`
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
		JournalDrives:                      "",
		OsdDrives:                          "",
		DisplayName:                        "",
		ProvisioningProgress:               0,
		StorageBackendBondInterfaceMembers: "",
		Annotations:                        MakeKeyValuePairs(),
		UUID:                               "",
		ParentUUID:                         "",
		ParentType:                         "",
		FQName:                             []string{},
		ProvisioningState:                  "",
		StorageAccessBondInterfaceMembers:  "",
		IDPerms: MakeIdPermsType(),
		Perms2:  MakePermType2(),
		ProvisioningProgressStage: "",
		ProvisioningLog:           "",
		ProvisioningStartTime:     "",
	}
}

// MakeOpenstackStorageNodeRoleSlice() makes a slice of OpenstackStorageNodeRole
func MakeOpenstackStorageNodeRoleSlice() []*OpenstackStorageNodeRole {
	return []*OpenstackStorageNodeRole{}
}
