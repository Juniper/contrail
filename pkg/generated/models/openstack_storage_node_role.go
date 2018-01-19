package models

// OpenstackStorageNodeRole

import "encoding/json"

// OpenstackStorageNodeRole
type OpenstackStorageNodeRole struct {
	FQName                             []string       `json:"fq_name,omitempty"`
	DisplayName                        string         `json:"display_name,omitempty"`
	ProvisioningProgress               int            `json:"provisioning_progress,omitempty"`
	ParentUUID                         string         `json:"parent_uuid,omitempty"`
	IDPerms                            *IdPermsType   `json:"id_perms,omitempty"`
	Perms2                             *PermType2     `json:"perms2,omitempty"`
	ProvisioningState                  string         `json:"provisioning_state,omitempty"`
	JournalDrives                      string         `json:"journal_drives,omitempty"`
	StorageAccessBondInterfaceMembers  string         `json:"storage_access_bond_interface_members,omitempty"`
	StorageBackendBondInterfaceMembers string         `json:"storage_backend_bond_interface_members,omitempty"`
	Annotations                        *KeyValuePairs `json:"annotations,omitempty"`
	OsdDrives                          string         `json:"osd_drives,omitempty"`
	ParentType                         string         `json:"parent_type,omitempty"`
	ProvisioningLog                    string         `json:"provisioning_log,omitempty"`
	ProvisioningProgressStage          string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningStartTime              string         `json:"provisioning_start_time,omitempty"`
	UUID                               string         `json:"uuid,omitempty"`
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
		ProvisioningLog:           "",
		ProvisioningProgressStage: "",
		ProvisioningStartTime:     "",
		UUID:                 "",
		ParentType:           "",
		DisplayName:          "",
		ProvisioningProgress: 0,
		ParentUUID:           "",
		FQName:               []string{},
		Perms2:               MakePermType2(),
		ProvisioningState:    "",
		JournalDrives:        "",
		IDPerms:              MakeIdPermsType(),
		StorageBackendBondInterfaceMembers: "",
		Annotations:                        MakeKeyValuePairs(),
		OsdDrives:                          "",
		StorageAccessBondInterfaceMembers:  "",
	}
}

// MakeOpenstackStorageNodeRoleSlice() makes a slice of OpenstackStorageNodeRole
func MakeOpenstackStorageNodeRoleSlice() []*OpenstackStorageNodeRole {
	return []*OpenstackStorageNodeRole{}
}
