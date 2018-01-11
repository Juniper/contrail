package models

// OpenstackStorageNodeRole

import "encoding/json"

// OpenstackStorageNodeRole
type OpenstackStorageNodeRole struct {
	Perms2                             *PermType2     `json:"perms2"`
	ProvisioningState                  string         `json:"provisioning_state"`
	ParentType                         string         `json:"parent_type"`
	IDPerms                            *IdPermsType   `json:"id_perms"`
	DisplayName                        string         `json:"display_name"`
	ProvisioningProgressStage          string         `json:"provisioning_progress_stage"`
	JournalDrives                      string         `json:"journal_drives"`
	ParentUUID                         string         `json:"parent_uuid"`
	ProvisioningLog                    string         `json:"provisioning_log"`
	ProvisioningProgress               int            `json:"provisioning_progress"`
	FQName                             []string       `json:"fq_name"`
	UUID                               string         `json:"uuid"`
	ProvisioningStartTime              string         `json:"provisioning_start_time"`
	Annotations                        *KeyValuePairs `json:"annotations"`
	OsdDrives                          string         `json:"osd_drives"`
	StorageAccessBondInterfaceMembers  string         `json:"storage_access_bond_interface_members"`
	StorageBackendBondInterfaceMembers string         `json:"storage_backend_bond_interface_members"`
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
		ProvisioningProgress: 0,
		FQName:               []string{},
		UUID:                 "",
		ProvisioningStartTime:              "",
		Annotations:                        MakeKeyValuePairs(),
		OsdDrives:                          "",
		StorageAccessBondInterfaceMembers:  "",
		StorageBackendBondInterfaceMembers: "",
		Perms2:                    MakePermType2(),
		ProvisioningState:         "",
		ParentType:                "",
		IDPerms:                   MakeIdPermsType(),
		DisplayName:               "",
		ProvisioningProgressStage: "",
		JournalDrives:             "",
		ParentUUID:                "",
		ProvisioningLog:           "",
	}
}

// MakeOpenstackStorageNodeRoleSlice() makes a slice of OpenstackStorageNodeRole
func MakeOpenstackStorageNodeRoleSlice() []*OpenstackStorageNodeRole {
	return []*OpenstackStorageNodeRole{}
}
