package models

// ControllerNodeRole

import "encoding/json"

// ControllerNodeRole
type ControllerNodeRole struct {
	DisplayName                           string         `json:"display_name"`
	ProvisioningProgressStage             string         `json:"provisioning_progress_stage"`
	ProvisioningStartTime                 string         `json:"provisioning_start_time"`
	ProvisioningState                     string         `json:"provisioning_state"`
	InternalapiBondInterfaceMembers       string         `json:"internalapi_bond_interface_members"`
	PerformanceDrives                     string         `json:"performance_drives"`
	StorageManagementBondInterfaceMembers string         `json:"storage_management_bond_interface_members"`
	Annotations                           *KeyValuePairs `json:"annotations"`
	IDPerms                               *IdPermsType   `json:"id_perms"`
	ProvisioningLog                       string         `json:"provisioning_log"`
	ProvisioningProgress                  int            `json:"provisioning_progress"`
	CapacityDrives                        string         `json:"capacity_drives"`
	UUID                                  string         `json:"uuid"`
	FQName                                []string       `json:"fq_name"`
	Perms2                                *PermType2     `json:"perms2"`
	ParentUUID                            string         `json:"parent_uuid"`
	ParentType                            string         `json:"parent_type"`
}

// String returns json representation of the object
func (model *ControllerNodeRole) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeControllerNodeRole makes ControllerNodeRole
func MakeControllerNodeRole() *ControllerNodeRole {
	return &ControllerNodeRole{
		//TODO(nati): Apply default
		StorageManagementBondInterfaceMembers: "",
		Annotations:                           MakeKeyValuePairs(),
		ProvisioningLog:                       "",
		ProvisioningProgress:                  0,
		CapacityDrives:                        "",
		UUID:                                  "",
		IDPerms:                               MakeIdPermsType(),
		Perms2:                                MakePermType2(),
		ParentUUID:                            "",
		ParentType:                            "",
		FQName:                                []string{},
		ProvisioningProgressStage:       "",
		ProvisioningStartTime:           "",
		ProvisioningState:               "",
		InternalapiBondInterfaceMembers: "",
		PerformanceDrives:               "",
		DisplayName:                     "",
	}
}

// InterfaceToControllerNodeRole makes ControllerNodeRole from interface
func InterfaceToControllerNodeRole(iData interface{}) *ControllerNodeRole {
	data := iData.(map[string]interface{})
	return &ControllerNodeRole{
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		ProvisioningProgressStage: data["provisioning_progress_stage"].(string),

		//{"title":"Provisioning Progress Stage","default":"","type":"string","permission":["create","update"]}
		ProvisioningStartTime: data["provisioning_start_time"].(string),

		//{"title":"Time provisioning started","default":"","type":"string","permission":["create","update"]}
		ProvisioningState: data["provisioning_state"].(string),

		//{"title":"Provisioning Status","default":"CREATED","type":"string","permission":["create","update"],"enum":["CREATED","IN_CREATE_PROGRESS","UPDATED","IN_UPDATE_PROGRESS","DELETED","IN_DELETE_PROGRESS","ERROR"]}
		InternalapiBondInterfaceMembers: data["internalapi_bond_interface_members"].(string),

		//{"title":"Internal API Bond Interface Members","description":"Internal API Bond Interface Members","default":"ens7f0,ens7f1","type":"string","permission":["create","update"]}
		PerformanceDrives: data["performance_drives"].(string),

		//{"title":"Performance Drive","description":"Drives for performance oriented application such as journaling","default":"sdf","type":"string","permission":["create","update"]}
		StorageManagementBondInterfaceMembers: data["storage_management_bond_interface_members"].(string),

		//{"title":"Storage Management  Bond Interface Members","description":"Storage Management  Bond Interface Members","default":"ens8f0,ens8f1","type":"string","permission":["create","update"]}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		ProvisioningLog: data["provisioning_log"].(string),

		//{"title":"Provisioning Log","default":"","type":"string","permission":["create","update"]}
		ProvisioningProgress: data["provisioning_progress"].(int),

		//{"title":"Provisioning Progress","default":0,"type":"integer","permission":["create","update"]}
		CapacityDrives: data["capacity_drives"].(string),

		//{"title":"Capacity Drives","description":"Drives for capacity oriented applications such as logging","default":"sdc,sdd,sde","type":"string","permission":["create","update"]}
		UUID: data["uuid"].(string),

		//{"type":"string"}

	}
}

// InterfaceToControllerNodeRoleSlice makes a slice of ControllerNodeRole from interface
func InterfaceToControllerNodeRoleSlice(data interface{}) []*ControllerNodeRole {
	list := data.([]interface{})
	result := MakeControllerNodeRoleSlice()
	for _, item := range list {
		result = append(result, InterfaceToControllerNodeRole(item))
	}
	return result
}

// MakeControllerNodeRoleSlice() makes a slice of ControllerNodeRole
func MakeControllerNodeRoleSlice() []*ControllerNodeRole {
	return []*ControllerNodeRole{}
}
