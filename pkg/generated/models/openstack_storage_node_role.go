package models

// OpenstackStorageNodeRole

import "encoding/json"

// OpenstackStorageNodeRole
type OpenstackStorageNodeRole struct {
	StorageBackendBondInterfaceMembers string         `json:"storage_backend_bond_interface_members,omitempty"`
	Annotations                        *KeyValuePairs `json:"annotations,omitempty"`
	ParentType                         string         `json:"parent_type,omitempty"`
	ProvisioningProgressStage          string         `json:"provisioning_progress_stage,omitempty"`
	OsdDrives                          string         `json:"osd_drives,omitempty"`
	Perms2                             *PermType2     `json:"perms2,omitempty"`
	ParentUUID                         string         `json:"parent_uuid,omitempty"`
	UUID                               string         `json:"uuid,omitempty"`
	FQName                             []string       `json:"fq_name,omitempty"`
	ProvisioningProgress               int            `json:"provisioning_progress,omitempty"`
	ProvisioningStartTime              string         `json:"provisioning_start_time,omitempty"`
	ProvisioningState                  string         `json:"provisioning_state,omitempty"`
	JournalDrives                      string         `json:"journal_drives,omitempty"`
	StorageAccessBondInterfaceMembers  string         `json:"storage_access_bond_interface_members,omitempty"`
	IDPerms                            *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName                        string         `json:"display_name,omitempty"`
	ProvisioningLog                    string         `json:"provisioning_log,omitempty"`
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
		UUID:                              "",
		FQName:                            []string{},
		ProvisioningProgress:              0,
		ProvisioningStartTime:             "",
		ProvisioningState:                 "",
		JournalDrives:                     "",
		StorageAccessBondInterfaceMembers: "",
		IDPerms:                            MakeIdPermsType(),
		DisplayName:                        "",
		ProvisioningLog:                    "",
		StorageBackendBondInterfaceMembers: "",
		Annotations:                        MakeKeyValuePairs(),
		ParentType:                         "",
		ProvisioningProgressStage:          "",
		OsdDrives:                          "",
		Perms2:                             MakePermType2(),
		ParentUUID:                         "",
	}
}

// MakeOpenstackStorageNodeRoleSlice() makes a slice of OpenstackStorageNodeRole
func MakeOpenstackStorageNodeRoleSlice() []*OpenstackStorageNodeRole {
	return []*OpenstackStorageNodeRole{}
}
