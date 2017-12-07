package models

// Node

import "encoding/json"

// Node
type Node struct {
	PrivateMachineState            string         `json:"private_machine_state"`
	PrivatePowerManagementPassword string         `json:"private_power_management_password"`
	FQName                         []string       `json:"fq_name"`
	Hostname                       string         `json:"hostname"`
	IPAddress                      string         `json:"ip_address"`
	MacAddress                     string         `json:"mac_address"`
	GCPImage                       string         `json:"gcp_image"`
	PrivateMachineProperties       string         `json:"private_machine_properties"`
	DisplayName                    string         `json:"display_name"`
	Username                       string         `json:"username"`
	AwsInstanceType                string         `json:"aws_instance_type"`
	PrivatePowerManagementUsername string         `json:"private_power_management_username"`
	IDPerms                        *IdPermsType   `json:"id_perms"`
	AwsAmi                         string         `json:"aws_ami"`
	GCPMachineType                 string         `json:"gcp_machine_type"`
	UUID                           string         `json:"uuid"`
	Annotations                    *KeyValuePairs `json:"annotations"`
	Type                           string         `json:"type"`
	Password                       string         `json:"password"`
	SSHKey                         string         `json:"ssh_key"`
	PrivatePowerManagementIP       string         `json:"private_power_management_ip"`
	Perms2                         *PermType2     `json:"perms2"`
}

// Node parents relation object

// String returns json representation of the object
func (model *Node) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeNode makes Node
func MakeNode() *Node {
	return &Node{
		//TODO(nati): Apply default
		PrivatePowerManagementPassword: "",
		FQName:                         []string{},
		Hostname:                       "",
		IPAddress:                      "",
		MacAddress:                     "",
		GCPImage:                       "",
		PrivateMachineProperties:       "",
		PrivateMachineState:            "",
		DisplayName:                    "",
		Username:                       "",
		AwsInstanceType:                "",
		PrivatePowerManagementUsername: "",
		IDPerms:        MakeIdPermsType(),
		AwsAmi:         "",
		GCPMachineType: "",
		UUID:           "",
		Annotations:    MakeKeyValuePairs(),
		Type:           "",
		Password:       "",
		SSHKey:         "",
		PrivatePowerManagementIP: "",
		Perms2: MakePermType2(),
	}
}

// InterfaceToNode makes Node from interface
func InterfaceToNode(iData interface{}) *Node {
	data := iData.(map[string]interface{})
	return &Node{
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"key_value_pair":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"key_value_pair","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"string","GoPremitive":true},"value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePair","CollectionType":"","Column":"","Item":null,"GoName":"KeyValuePair","GoType":"KeyValuePair","GoPremitive":false},"GoName":"KeyValuePair","GoType":"[]*KeyValuePair","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePairs","CollectionType":"","Column":"","Item":null,"GoName":"Annotations","GoType":"KeyValuePairs","GoPremitive":false}
		AwsAmi: data["aws_ami"].(string),

		//{"Title":"AMI","Description":"","SQL":"varchar(255)","Default":"ami-73f7da13","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"aws_ami","Item":null,"GoName":"AwsAmi","GoType":"string","GoPremitive":true}
		GCPMachineType: data["gcp_machine_type"].(string),

		//{"Title":"Machine Type","Description":"","SQL":"varchar(255)","Default":"n1-standard-1","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"gcp_machine_type","Item":null,"GoName":"GCPMachineType","GoType":"string","GoPremitive":true}
		UUID: data["uuid"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"uuid","Item":null,"GoName":"UUID","GoType":"string","GoPremitive":true}
		PrivatePowerManagementIP: data["private_power_management_ip"].(string),

		//{"Title":"Power Management IP","Description":"IP address used for power management (IPMI)","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_power_management_ip","Item":null,"GoName":"PrivatePowerManagementIP","GoType":"string","GoPremitive":true}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"global_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"global_access","Item":null,"GoName":"GlobalAccess","GoType":"AccessType","GoPremitive":false},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"owner","Item":null,"GoName":"Owner","GoType":"string","GoPremitive":true},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType","GoPremitive":false},"share":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"share","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"tenant":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tenant","GoType":"string","GoPremitive":true},"tenant_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"TenantAccess","GoType":"AccessType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ShareType","CollectionType":"","Column":"","Item":null,"GoName":"Share","GoType":"ShareType","GoPremitive":false},"GoName":"Share","GoType":"[]*ShareType","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType2","CollectionType":"","Column":"","Item":null,"GoName":"Perms2","GoType":"PermType2","GoPremitive":false}
		Type: data["type"].(string),

		//{"Title":"Machine Type","Description":"Type of machine resource","SQL":"varchar(255)","Default":"private","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":["private","virtual","aws","container","gcp"],"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"type","Item":null,"GoName":"Type","GoType":"string","GoPremitive":true}
		Password: data["password"].(string),

		//{"Title":"UserPassword","Description":"UserPassword","SQL":"varchar(255)","Default":"ADMIN","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"password","Item":null,"GoName":"Password","GoType":"string","GoPremitive":true}
		SSHKey: data["ssh_key"].(string),

		//{"Title":"SSH public key","Description":"SSH Public Key","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"ssh_key","Item":null,"GoName":"SSHKey","GoType":"string","GoPremitive":true}
		GCPImage: data["gcp_image"].(string),

		//{"Title":"Image","Description":"","SQL":"varchar(255)","Default":"ubuntu-os-cloud/ubuntu-1604-lts","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"gcp_image","Item":null,"GoName":"GCPImage","GoType":"string","GoPremitive":true}
		PrivateMachineProperties: data["private_machine_properties"].(string),

		//{"Title":"Machine Properties","Description":"Machine Properties from ironic","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_machine_properties","Item":null,"GoName":"PrivateMachineProperties","GoType":"string","GoPremitive":true}
		PrivateMachineState: data["private_machine_state"].(string),

		//{"Title":"Machine State","Description":"Machine State","SQL":"varchar(255)","Default":"enroll","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":["enroll","manageable","available","assigned"],"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_machine_state","Item":null,"GoName":"PrivateMachineState","GoType":"string","GoPremitive":true}
		PrivatePowerManagementPassword: data["private_power_management_password"].(string),

		//{"Title":"Power Management UserPassword","Description":"UserPassword for PowerManagement","SQL":"varchar(255)","Default":"ADMIN","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_power_management_password","Item":null,"GoName":"PrivatePowerManagementPassword","GoType":"string","GoPremitive":true}
		FQName: data["fq_name"].([]string),

		//{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"fq_name","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"FQName","GoType":"string","GoPremitive":true},"GoName":"FQName","GoType":"[]string","GoPremitive":true}
		Hostname: data["hostname"].(string),

		//{"Title":"Hostname","Description":"Fully qualified host name","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"hostname","Item":null,"GoName":"Hostname","GoType":"string","GoPremitive":true}
		IPAddress: data["ip_address"].(string),

		//{"Title":"IP Address","Description":"IP Address","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"ip_address","Item":null,"GoName":"IPAddress","GoType":"string","GoPremitive":true}
		MacAddress: data["mac_address"].(string),

		//{"Title":"Interface MAC Address","Description":"Provisioning Interface MAC Address","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"mac_address","Item":null,"GoName":"MacAddress","GoType":"string","GoPremitive":true}
		DisplayName: data["display_name"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"display_name","Item":null,"GoName":"DisplayName","GoType":"string","GoPremitive":true}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"created":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"created","Item":null,"GoName":"Created","GoType":"string","GoPremitive":true},"creator":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"creator","Item":null,"GoName":"Creator","GoType":"string","GoPremitive":true},"description":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"description","Item":null,"GoName":"Description","GoType":"string","GoPremitive":true},"enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"true","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"enable","Item":null,"GoName":"Enable","GoType":"bool","GoPremitive":true},"last_modified":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"last_modified","Item":null,"GoName":"LastModified","GoType":"string","GoPremitive":true},"permissions":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"group":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"group","Item":null,"GoName":"Group","GoType":"string","GoPremitive":true},"group_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"group_access","Item":null,"GoName":"GroupAccess","GoType":"AccessType","GoPremitive":false},"other_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"other_access","Item":null,"GoName":"OtherAccess","GoType":"AccessType","GoPremitive":false},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"permissions_owner","Item":null,"GoName":"Owner","GoType":"string","GoPremitive":true},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"permissions_owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType","CollectionType":"","Column":"","Item":null,"GoName":"Permissions","GoType":"PermType","GoPremitive":false},"user_visible":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"system-only","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"user_visible","Item":null,"GoName":"UserVisible","GoType":"bool","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IdPermsType","CollectionType":"","Column":"","Item":null,"GoName":"IDPerms","GoType":"IdPermsType","GoPremitive":false}
		Username: data["username"].(string),

		//{"Title":"User Name","Description":"User Name","SQL":"varchar(255)","Default":"ADMIN","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"username","Item":null,"GoName":"Username","GoType":"string","GoPremitive":true}
		AwsInstanceType: data["aws_instance_type"].(string),

		//{"Title":"Instance Type","Description":"","SQL":"varchar(255)","Default":"t2.micro","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"aws_instance_type","Item":null,"GoName":"AwsInstanceType","GoType":"string","GoPremitive":true}
		PrivatePowerManagementUsername: data["private_power_management_username"].(string),

		//{"Title":"Power Management User Name","Description":"User Name for PowerManagement","SQL":"varchar(255)","Default":"ADMIN","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_power_management_username","Item":null,"GoName":"PrivatePowerManagementUsername","GoType":"string","GoPremitive":true}

	}
}

// InterfaceToNodeSlice makes a slice of Node from interface
func InterfaceToNodeSlice(data interface{}) []*Node {
	list := data.([]interface{})
	result := MakeNodeSlice()
	for _, item := range list {
		result = append(result, InterfaceToNode(item))
	}
	return result
}

// MakeNodeSlice() makes a slice of Node
func MakeNodeSlice() []*Node {
	return []*Node{}
}
