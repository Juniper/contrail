package models

// Node

import "encoding/json"

type Node struct {
	IPAddress                      string         `json:"ip_address"`
	Password                       string         `json:"password"`
	Username                       string         `json:"username"`
	GCPMachineType                 string         `json:"gcp_machine_type"`
	PrivatePowerManagementIP       string         `json:"private_power_management_ip"`
	FQName                         []string       `json:"fq_name"`
	Perms2                         *PermType2     `json:"perms2"`
	MacAddress                     string         `json:"mac_address"`
	SSHKey                         string         `json:"ssh_key"`
	AwsAmi                         string         `json:"aws_ami"`
	GCPImage                       string         `json:"gcp_image"`
	PrivateMachineState            string         `json:"private_machine_state"`
	DisplayName                    string         `json:"display_name"`
	Annotations                    *KeyValuePairs `json:"annotations"`
	UUID                           string         `json:"uuid"`
	Hostname                       string         `json:"hostname"`
	PrivateMachineProperties       string         `json:"private_machine_properties"`
	Type                           string         `json:"type"`
	AwsInstanceType                string         `json:"aws_instance_type"`
	PrivatePowerManagementPassword string         `json:"private_power_management_password"`
	PrivatePowerManagementUsername string         `json:"private_power_management_username"`
	IDPerms                        *IdPermsType   `json:"id_perms"`
}

func (model *Node) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeNode() *Node {
	return &Node{
		//TODO(nati): Apply default
		Hostname:                 "",
		PrivateMachineProperties: "",
		Type:                           "",
		AwsInstanceType:                "",
		PrivatePowerManagementPassword: "",
		PrivatePowerManagementUsername: "",
		IDPerms:                  MakeIdPermsType(),
		IPAddress:                "",
		Password:                 "",
		Username:                 "",
		GCPMachineType:           "",
		PrivatePowerManagementIP: "",
		FQName:              []string{},
		Perms2:              MakePermType2(),
		MacAddress:          "",
		SSHKey:              "",
		AwsAmi:              "",
		GCPImage:            "",
		PrivateMachineState: "",
		DisplayName:         "",
		Annotations:         MakeKeyValuePairs(),
		UUID:                "",
	}
}

func InterfaceToNode(iData interface{}) *Node {
	data := iData.(map[string]interface{})
	return &Node{
		AwsInstanceType: data["aws_instance_type"].(string),

		//{"Title":"Instance Type","Description":"","SQL":"varchar(255)","Default":"t2.micro","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"aws_instance_type","Item":null,"GoName":"AwsInstanceType","GoType":"string"}
		PrivatePowerManagementPassword: data["private_power_management_password"].(string),

		//{"Title":"Power Management UserPassword","Description":"UserPassword for PowerManagement","SQL":"varchar(255)","Default":"ADMIN","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_power_management_password","Item":null,"GoName":"PrivatePowerManagementPassword","GoType":"string"}
		PrivatePowerManagementUsername: data["private_power_management_username"].(string),

		//{"Title":"Power Management User Name","Description":"User Name for PowerManagement","SQL":"varchar(255)","Default":"ADMIN","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_power_management_username","Item":null,"GoName":"PrivatePowerManagementUsername","GoType":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"created":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"created","Item":null,"GoName":"Created","GoType":"string"},"creator":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"creator","Item":null,"GoName":"Creator","GoType":"string"},"description":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"description","Item":null,"GoName":"Description","GoType":"string"},"enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"true","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"enable","Item":null,"GoName":"Enable","GoType":"bool"},"last_modified":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"last_modified","Item":null,"GoName":"LastModified","GoType":"string"},"permissions":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"group":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"group","Item":null,"GoName":"Group","GoType":"string"},"group_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"group_access","Item":null,"GoName":"GroupAccess","GoType":"AccessType"},"other_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"other_access","Item":null,"GoName":"OtherAccess","GoType":"AccessType"},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"owner","Item":null,"GoName":"Owner","GoType":"string"},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType","CollectionType":"","Column":"","Item":null,"GoName":"Permissions","GoType":"PermType"},"user_visible":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"system-only","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"user_visible","Item":null,"GoName":"UserVisible","GoType":"bool"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IdPermsType","CollectionType":"","Column":"","Item":null,"GoName":"IDPerms","GoType":"IdPermsType"}
		Type: data["type"].(string),

		//{"Title":"Machine Type","Description":"Type of machine resource","SQL":"varchar(255)","Default":"private","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":["private","virtual","aws","container","gcp"],"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"type","Item":null,"GoName":"Type","GoType":"string"}
		Password: data["password"].(string),

		//{"Title":"UserPassword","Description":"UserPassword","SQL":"varchar(255)","Default":"ADMIN","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"password","Item":null,"GoName":"Password","GoType":"string"}
		Username: data["username"].(string),

		//{"Title":"User Name","Description":"User Name","SQL":"varchar(255)","Default":"ADMIN","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"username","Item":null,"GoName":"Username","GoType":"string"}
		GCPMachineType: data["gcp_machine_type"].(string),

		//{"Title":"Machine Type","Description":"","SQL":"varchar(255)","Default":"n1-standard-1","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"gcp_machine_type","Item":null,"GoName":"GCPMachineType","GoType":"string"}
		PrivatePowerManagementIP: data["private_power_management_ip"].(string),

		//{"Title":"Power Management IP","Description":"IP address used for power management (IPMI)","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_power_management_ip","Item":null,"GoName":"PrivatePowerManagementIP","GoType":"string"}
		FQName: data["fq_name"].([]string),

		//{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"fq_name","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"FQName","GoType":"string"},"GoName":"FQName","GoType":"[]string"}
		IPAddress: data["ip_address"].(string),

		//{"Title":"IP Address","Description":"IP Address","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"ip_address","Item":null,"GoName":"IPAddress","GoType":"string"}
		SSHKey: data["ssh_key"].(string),

		//{"Title":"SSH public key","Description":"SSH Public Key","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"ssh_key","Item":null,"GoName":"SSHKey","GoType":"string"}
		AwsAmi: data["aws_ami"].(string),

		//{"Title":"AMI","Description":"","SQL":"varchar(255)","Default":"ami-73f7da13","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"aws_ami","Item":null,"GoName":"AwsAmi","GoType":"string"}
		GCPImage: data["gcp_image"].(string),

		//{"Title":"Image","Description":"","SQL":"varchar(255)","Default":"ubuntu-os-cloud/ubuntu-1604-lts","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"gcp_image","Item":null,"GoName":"GCPImage","GoType":"string"}
		PrivateMachineState: data["private_machine_state"].(string),

		//{"Title":"Machine State","Description":"Machine State","SQL":"varchar(255)","Default":"enroll","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":["enroll","manageable","available","assigned"],"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_machine_state","Item":null,"GoName":"PrivateMachineState","GoType":"string"}
		DisplayName: data["display_name"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"display_name","Item":null,"GoName":"DisplayName","GoType":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"key_value_pair":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"key_value_pair","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"string"},"value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePair","CollectionType":"","Column":"","Item":null,"GoName":"KeyValuePair","GoType":"KeyValuePair"},"GoName":"KeyValuePair","GoType":"[]*KeyValuePair"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePairs","CollectionType":"","Column":"","Item":null,"GoName":"Annotations","GoType":"KeyValuePairs"}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"global_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"global_access","Item":null,"GoName":"GlobalAccess","GoType":"AccessType"},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"perms2_owner","Item":null,"GoName":"Owner","GoType":"string"},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"perms2_owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType"},"share":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"share","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"tenant":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tenant","GoType":"string"},"tenant_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"TenantAccess","GoType":"AccessType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ShareType","CollectionType":"","Column":"","Item":null,"GoName":"Share","GoType":"ShareType"},"GoName":"Share","GoType":"[]*ShareType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType2","CollectionType":"","Column":"","Item":null,"GoName":"Perms2","GoType":"PermType2"}
		MacAddress: data["mac_address"].(string),

		//{"Title":"Interface MAC Address","Description":"Provisioning Interface MAC Address","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"mac_address","Item":null,"GoName":"MacAddress","GoType":"string"}
		UUID: data["uuid"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"uuid","Item":null,"GoName":"UUID","GoType":"string"}
		PrivateMachineProperties: data["private_machine_properties"].(string),

		//{"Title":"Machine Properties","Description":"Machine Properties from ironic","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_machine_properties","Item":null,"GoName":"PrivateMachineProperties","GoType":"string"}
		Hostname: data["hostname"].(string),

		//{"Title":"Hostname","Description":"Fully qualified host name","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"hostname","Item":null,"GoName":"Hostname","GoType":"string"}

	}
}

func InterfaceToNodeSlice(data interface{}) []*Node {
	list := data.([]interface{})
	result := MakeNodeSlice()
	for _, item := range list {
		result = append(result, InterfaceToNode(item))
	}
	return result
}

func MakeNodeSlice() []*Node {
	return []*Node{}
}
