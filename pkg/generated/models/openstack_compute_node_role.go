package models

// OpenstackComputeNodeRole

import "encoding/json"

// OpenstackComputeNodeRole
type OpenstackComputeNodeRole struct {
	VrouterBondInterfaceMembers string         `json:"vrouter_bond_interface_members"`
	VrouterType                 string         `json:"vrouter_type"`
	IDPerms                     *IdPermsType   `json:"id_perms"`
	ProvisioningProgressStage   string         `json:"provisioning_progress_stage"`
	DisplayName                 string         `json:"display_name"`
	Annotations                 *KeyValuePairs `json:"annotations"`
	ProvisioningStartTime       string         `json:"provisioning_start_time"`
	ProvisioningState           string         `json:"provisioning_state"`
	DefaultGateway              string         `json:"default_gateway"`
	FQName                      []string       `json:"fq_name"`
	Perms2                      *PermType2     `json:"perms2"`
	ProvisioningLog             string         `json:"provisioning_log"`
	ProvisioningProgress        int            `json:"provisioning_progress"`
	VrouterBondInterface        string         `json:"vrouter_bond_interface"`
	UUID                        string         `json:"uuid"`
}

// OpenstackComputeNodeRole parents relation object

// String returns json representation of the object
func (model *OpenstackComputeNodeRole) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeOpenstackComputeNodeRole makes OpenstackComputeNodeRole
func MakeOpenstackComputeNodeRole() *OpenstackComputeNodeRole {
	return &OpenstackComputeNodeRole{
		//TODO(nati): Apply default
		VrouterBondInterfaceMembers: "",
		VrouterType:                 "",
		IDPerms:                     MakeIdPermsType(),
		ProvisioningProgressStage:   "",
		DisplayName:                 "",
		Annotations:                 MakeKeyValuePairs(),
		ProvisioningStartTime:       "",
		ProvisioningState:           "",
		DefaultGateway:              "",
		FQName:                      []string{},
		Perms2:                      MakePermType2(),
		ProvisioningLog:             "",
		ProvisioningProgress:        0,
		VrouterBondInterface:        "",
		UUID:                        "",
	}
}

// InterfaceToOpenstackComputeNodeRole makes OpenstackComputeNodeRole from interface
func InterfaceToOpenstackComputeNodeRole(iData interface{}) *OpenstackComputeNodeRole {
	data := iData.(map[string]interface{})
	return &OpenstackComputeNodeRole{
		ProvisioningState: data["provisioning_state"].(string),

		//{"Title":"Provisioning Status","Description":"","SQL":"varchar(255)","Default":"CREATED","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":{},"Enum":["CREATED","IN_CREATE_PROGRESS","UPDATED","IN_UPDATE_PROGRESS","DELETED","IN_DELETE_PROGRESS","ERROR"],"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"provisioning_state","Item":null,"GoName":"ProvisioningState","GoType":"string","GoPremitive":true}
		DisplayName: data["display_name"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"display_name","Item":null,"GoName":"DisplayName","GoType":"string","GoPremitive":true}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"key_value_pair":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"key_value_pair","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"string","GoPremitive":true},"value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePair","CollectionType":"","Column":"","Item":null,"GoName":"KeyValuePair","GoType":"KeyValuePair","GoPremitive":false},"GoName":"KeyValuePair","GoType":"[]*KeyValuePair","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePairs","CollectionType":"","Column":"","Item":null,"GoName":"Annotations","GoType":"KeyValuePairs","GoPremitive":false}
		ProvisioningStartTime: data["provisioning_start_time"].(string),

		//{"Title":"Time provisioning started","Description":"","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"provisioning_start_time","Item":null,"GoName":"ProvisioningStartTime","GoType":"string","GoPremitive":true}
		ProvisioningLog: data["provisioning_log"].(string),

		//{"Title":"Provisioning Log","Description":"","SQL":"text","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"provisioning_log","Item":null,"GoName":"ProvisioningLog","GoType":"string","GoPremitive":true}
		ProvisioningProgress: data["provisioning_progress"].(int),

		//{"Title":"Provisioning Progress","Description":"","SQL":"int","Default":0,"Operation":"","Presence":"","Type":"integer","Permission":["create","update"],"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"provisioning_progress","Item":null,"GoName":"ProvisioningProgress","GoType":"int","GoPremitive":true}
		DefaultGateway: data["default_gateway"].(string),

		//{"Title":"Default Gateway","Description":"Default Gateway","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"default_gateway","Item":null,"GoName":"DefaultGateway","GoType":"string","GoPremitive":true}
		FQName: data["fq_name"].([]string),

		//{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"fq_name","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"FQName","GoType":"string","GoPremitive":true},"GoName":"FQName","GoType":"[]string","GoPremitive":true}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"global_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"global_access","Item":null,"GoName":"GlobalAccess","GoType":"AccessType","GoPremitive":false},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"perms2_owner","Item":null,"GoName":"Owner","GoType":"string","GoPremitive":true},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"perms2_owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType","GoPremitive":false},"share":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"share","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"tenant":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tenant","GoType":"string","GoPremitive":true},"tenant_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"TenantAccess","GoType":"AccessType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ShareType","CollectionType":"","Column":"","Item":null,"GoName":"Share","GoType":"ShareType","GoPremitive":false},"GoName":"Share","GoType":"[]*ShareType","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType2","CollectionType":"","Column":"","Item":null,"GoName":"Perms2","GoType":"PermType2","GoPremitive":false}
		VrouterBondInterface: data["vrouter_bond_interface"].(string),

		//{"Title":"vRouter Bond Interface","Description":"vRouter Bond Interface","SQL":"varchar(255)","Default":"bond0","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"vrouter_bond_interface","Item":null,"GoName":"VrouterBondInterface","GoType":"string","GoPremitive":true}
		UUID: data["uuid"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"uuid","Item":null,"GoName":"UUID","GoType":"string","GoPremitive":true}
		ProvisioningProgressStage: data["provisioning_progress_stage"].(string),

		//{"Title":"Provisioning Progress Stage","Description":"","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"provisioning_progress_stage","Item":null,"GoName":"ProvisioningProgressStage","GoType":"string","GoPremitive":true}
		VrouterBondInterfaceMembers: data["vrouter_bond_interface_members"].(string),

		//{"Title":"vRouter Bond Interface Members","Description":"vRouter Bond Interface Members","SQL":"varchar(255)","Default":"ens7f0,ens7f1","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"vrouter_bond_interface_members","Item":null,"GoName":"VrouterBondInterfaceMembers","GoType":"string","GoPremitive":true}
		VrouterType: data["vrouter_type"].(string),

		//{"Title":"vRouter Type","Description":"","SQL":"varchar(255)","Default":"kernel","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":["kernel","dpdk","smartNiC"],"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"vrouter_type","Item":null,"GoName":"VrouterType","GoType":"string","GoPremitive":true}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"created":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"created","Item":null,"GoName":"Created","GoType":"string","GoPremitive":true},"creator":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"creator","Item":null,"GoName":"Creator","GoType":"string","GoPremitive":true},"description":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"description","Item":null,"GoName":"Description","GoType":"string","GoPremitive":true},"enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"true","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"enable","Item":null,"GoName":"Enable","GoType":"bool","GoPremitive":true},"last_modified":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"last_modified","Item":null,"GoName":"LastModified","GoType":"string","GoPremitive":true},"permissions":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"group":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"group","Item":null,"GoName":"Group","GoType":"string","GoPremitive":true},"group_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"group_access","Item":null,"GoName":"GroupAccess","GoType":"AccessType","GoPremitive":false},"other_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"other_access","Item":null,"GoName":"OtherAccess","GoType":"AccessType","GoPremitive":false},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"owner","Item":null,"GoName":"Owner","GoType":"string","GoPremitive":true},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType","CollectionType":"","Column":"","Item":null,"GoName":"Permissions","GoType":"PermType","GoPremitive":false},"user_visible":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"system-only","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"user_visible","Item":null,"GoName":"UserVisible","GoType":"bool","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IdPermsType","CollectionType":"","Column":"","Item":null,"GoName":"IDPerms","GoType":"IdPermsType","GoPremitive":false}

	}
}

// InterfaceToOpenstackComputeNodeRoleSlice makes a slice of OpenstackComputeNodeRole from interface
func InterfaceToOpenstackComputeNodeRoleSlice(data interface{}) []*OpenstackComputeNodeRole {
	list := data.([]interface{})
	result := MakeOpenstackComputeNodeRoleSlice()
	for _, item := range list {
		result = append(result, InterfaceToOpenstackComputeNodeRole(item))
	}
	return result
}

// MakeOpenstackComputeNodeRoleSlice() makes a slice of OpenstackComputeNodeRole
func MakeOpenstackComputeNodeRoleSlice() []*OpenstackComputeNodeRole {
	return []*OpenstackComputeNodeRole{}
}
