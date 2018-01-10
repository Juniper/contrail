package models

// ControllerNodeRole

import "encoding/json"

// ControllerNodeRole
type ControllerNodeRole struct {
	CapacityDrives                        string         `json:"capacity_drives"`
	InternalapiBondInterfaceMembers       string         `json:"internalapi_bond_interface_members"`
	PerformanceDrives                     string         `json:"performance_drives"`
	ProvisioningLog                       string         `json:"provisioning_log"`
	StorageManagementBondInterfaceMembers string         `json:"storage_management_bond_interface_members"`
	ParentUUID                            string         `json:"parent_uuid"`
	FQName                                []string       `json:"fq_name"`
	ProvisioningStartTime                 string         `json:"provisioning_start_time"`
	DisplayName                           string         `json:"display_name"`
	UUID                                  string         `json:"uuid"`
	ProvisioningProgress                  int            `json:"provisioning_progress"`
	ProvisioningProgressStage             string         `json:"provisioning_progress_stage"`
	ProvisioningState                     string         `json:"provisioning_state"`
	Annotations                           *KeyValuePairs `json:"annotations"`
	Perms2                                *PermType2     `json:"perms2"`
	ParentType                            string         `json:"parent_type"`
	IDPerms                               *IdPermsType   `json:"id_perms"`
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
		CapacityDrives:                        "",
		InternalapiBondInterfaceMembers:       "",
		PerformanceDrives:                     "",
		ProvisioningLog:                       "",
		StorageManagementBondInterfaceMembers: "",
		ParentUUID:                            "",
		FQName:                                []string{},
		DisplayName:                           "",
		UUID:                                  "",
		ProvisioningProgress:                  0,
		ProvisioningProgressStage:             "",
		ProvisioningStartTime:                 "",
		Annotations:                           MakeKeyValuePairs(),
		Perms2:                                MakePermType2(),
		ParentType:                            "",
		IDPerms:                               MakeIdPermsType(),
		ProvisioningState:                     "",
	}
}

// InterfaceToControllerNodeRole makes ControllerNodeRole from interface
func InterfaceToControllerNodeRole(iData interface{}) *ControllerNodeRole {
	data := iData.(map[string]interface{})
	return &ControllerNodeRole{
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		ProvisioningState: data["provisioning_state"].(string),

		//{"title":"Provisioning Status","default":"CREATED","type":"string","permission":["create","update"],"enum":["CREATED","IN_CREATE_PROGRESS","UPDATED","IN_UPDATE_PROGRESS","DELETED","IN_DELETE_PROGRESS","ERROR"]}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		PerformanceDrives: data["performance_drives"].(string),

		//{"title":"Performance Drive","description":"Drives for performance oriented application such as journaling","default":"sdf","type":"string","permission":["create","update"]}
		ProvisioningLog: data["provisioning_log"].(string),

		//{"title":"Provisioning Log","default":"","type":"string","permission":["create","update"]}
		CapacityDrives: data["capacity_drives"].(string),

		//{"title":"Capacity Drives","description":"Drives for capacity oriented applications such as logging","default":"sdc,sdd,sde","type":"string","permission":["create","update"]}
		InternalapiBondInterfaceMembers: data["internalapi_bond_interface_members"].(string),

		//{"title":"Internal API Bond Interface Members","description":"Internal API Bond Interface Members","default":"ens7f0,ens7f1","type":"string","permission":["create","update"]}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		StorageManagementBondInterfaceMembers: data["storage_management_bond_interface_members"].(string),

		//{"title":"Storage Management  Bond Interface Members","description":"Storage Management  Bond Interface Members","default":"ens8f0,ens8f1","type":"string","permission":["create","update"]}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ProvisioningProgress: data["provisioning_progress"].(int),

		//{"title":"Provisioning Progress","default":0,"type":"integer","permission":["create","update"]}
		ProvisioningProgressStage: data["provisioning_progress_stage"].(string),

		//{"title":"Provisioning Progress Stage","default":"","type":"string","permission":["create","update"]}
		ProvisioningStartTime: data["provisioning_start_time"].(string),

		//{"title":"Time provisioning started","default":"","type":"string","permission":["create","update"]}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
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
