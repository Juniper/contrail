package models

// OpenstackStorageNodeRole

import "encoding/json"

// OpenstackStorageNodeRole
type OpenstackStorageNodeRole struct {
	ProvisioningProgress               int            `json:"provisioning_progress,omitempty"`
	JournalDrives                      string         `json:"journal_drives,omitempty"`
	StorageAccessBondInterfaceMembers  string         `json:"storage_access_bond_interface_members,omitempty"`
	FQName                             []string       `json:"fq_name,omitempty"`
	Annotations                        *KeyValuePairs `json:"annotations,omitempty"`
	ParentType                         string         `json:"parent_type,omitempty"`
	ProvisioningLog                    string         `json:"provisioning_log,omitempty"`
	ProvisioningStartTime              string         `json:"provisioning_start_time,omitempty"`
	ProvisioningState                  string         `json:"provisioning_state,omitempty"`
	Perms2                             *PermType2     `json:"perms2,omitempty"`
	ProvisioningProgressStage          string         `json:"provisioning_progress_stage,omitempty"`
	OsdDrives                          string         `json:"osd_drives,omitempty"`
	StorageBackendBondInterfaceMembers string         `json:"storage_backend_bond_interface_members,omitempty"`
	UUID                               string         `json:"uuid,omitempty"`
	ParentUUID                         string         `json:"parent_uuid,omitempty"`
	IDPerms                            *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName                        string         `json:"display_name,omitempty"`
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
		ProvisioningStartTime: "",
		ProvisioningState:     "",
		ParentType:            "",
		ProvisioningLog:       "",
		UUID:                  "",
		ParentUUID:            "",
		IDPerms:               MakeIdPermsType(),
		DisplayName:           "",
		Perms2:                MakePermType2(),
		ProvisioningProgressStage:          "",
		OsdDrives:                          "",
		StorageBackendBondInterfaceMembers: "",
		ProvisioningProgress:               0,
		FQName:                             []string{},
		Annotations:                        MakeKeyValuePairs(),
		JournalDrives:                      "",
		StorageAccessBondInterfaceMembers:  "",
	}
}

// MakeOpenstackStorageNodeRoleSlice() makes a slice of OpenstackStorageNodeRole
func MakeOpenstackStorageNodeRoleSlice() []*OpenstackStorageNodeRole {
	return []*OpenstackStorageNodeRole{}
}
