package models

// OpenstackStorageNodeRole

import "encoding/json"

// OpenstackStorageNodeRole
type OpenstackStorageNodeRole struct {
	ProvisioningProgressStage          string         `json:"provisioning_progress_stage"`
	ProvisioningStartTime              string         `json:"provisioning_start_time"`
	ProvisioningState                  string         `json:"provisioning_state"`
	ParentType                         string         `json:"parent_type"`
	ProvisioningLog                    string         `json:"provisioning_log"`
	DisplayName                        string         `json:"display_name"`
	OsdDrives                          string         `json:"osd_drives"`
	Perms2                             *PermType2     `json:"perms2"`
	Annotations                        *KeyValuePairs `json:"annotations"`
	ParentUUID                         string         `json:"parent_uuid"`
	StorageBackendBondInterfaceMembers string         `json:"storage_backend_bond_interface_members"`
	UUID                               string         `json:"uuid"`
	FQName                             []string       `json:"fq_name"`
	IDPerms                            *IdPermsType   `json:"id_perms"`
	ProvisioningProgress               int            `json:"provisioning_progress"`
	JournalDrives                      string         `json:"journal_drives"`
	StorageAccessBondInterfaceMembers  string         `json:"storage_access_bond_interface_members"`
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
		StorageAccessBondInterfaceMembers:  "",
		StorageBackendBondInterfaceMembers: "",
		UUID:                      "",
		FQName:                    []string{},
		IDPerms:                   MakeIdPermsType(),
		ProvisioningProgress:      0,
		ParentType:                "",
		ProvisioningLog:           "",
		ProvisioningProgressStage: "",
		ProvisioningStartTime:     "",
		ProvisioningState:         "",
		OsdDrives:                 "",
		Perms2:                    MakePermType2(),
		DisplayName:               "",
		Annotations:               MakeKeyValuePairs(),
		ParentUUID:                "",
	}
}

// InterfaceToOpenstackStorageNodeRole makes OpenstackStorageNodeRole from interface
func InterfaceToOpenstackStorageNodeRole(iData interface{}) *OpenstackStorageNodeRole {
	data := iData.(map[string]interface{})
	return &OpenstackStorageNodeRole{
		UUID: data["uuid"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		ProvisioningProgress: data["provisioning_progress"].(int),

		//{"title":"Provisioning Progress","default":0,"type":"integer","permission":["create","update"]}
		JournalDrives: data["journal_drives"].(string),

		//{"title":"Journal Drives","description":"SSD Drives to use for journals","default":"sdf","type":"string","permission":["create","update"]}
		StorageAccessBondInterfaceMembers: data["storage_access_bond_interface_members"].(string),

		//{"title":"Storage Access  Bond Interface Members","description":"Storage Management  Bond Interface Members","default":"ens8f0,ens8f1","type":"string","permission":["create","update"]}
		StorageBackendBondInterfaceMembers: data["storage_backend_bond_interface_members"].(string),

		//{"title":"Storage Backend Bond Interface Members","description":"Storage Backend Bond Interface Members","default":"ens9f0,ens9f1","type":"string","permission":["create","update"]}
		ProvisioningStartTime: data["provisioning_start_time"].(string),

		//{"title":"Time provisioning started","default":"","type":"string","permission":["create","update"]}
		ProvisioningState: data["provisioning_state"].(string),

		//{"title":"Provisioning Status","default":"CREATED","type":"string","permission":["create","update"],"enum":["CREATED","IN_CREATE_PROGRESS","UPDATED","IN_UPDATE_PROGRESS","DELETED","IN_DELETE_PROGRESS","ERROR"]}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		ProvisioningLog: data["provisioning_log"].(string),

		//{"title":"Provisioning Log","default":"","type":"string","permission":["create","update"]}
		ProvisioningProgressStage: data["provisioning_progress_stage"].(string),

		//{"title":"Provisioning Progress Stage","default":"","type":"string","permission":["create","update"]}
		OsdDrives: data["osd_drives"].(string),

		//{"title":"Stoage Drives","description":"Drives to use for cloud storage","default":"sdc,sdd,sde","type":"string","permission":["create","update"]}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}

	}
}

// InterfaceToOpenstackStorageNodeRoleSlice makes a slice of OpenstackStorageNodeRole from interface
func InterfaceToOpenstackStorageNodeRoleSlice(data interface{}) []*OpenstackStorageNodeRole {
	list := data.([]interface{})
	result := MakeOpenstackStorageNodeRoleSlice()
	for _, item := range list {
		result = append(result, InterfaceToOpenstackStorageNodeRole(item))
	}
	return result
}

// MakeOpenstackStorageNodeRoleSlice() makes a slice of OpenstackStorageNodeRole
func MakeOpenstackStorageNodeRoleSlice() []*OpenstackStorageNodeRole {
	return []*OpenstackStorageNodeRole{}
}
