package models

// Location

import "encoding/json"

// Location
type Location struct {
	IDPerms                          *IdPermsType   `json:"id_perms"`
	ProvisioningProgress             int            `json:"provisioning_progress"`
	ProvisioningState                string         `json:"provisioning_state"`
	PrivateOspdVMDiskGB              string         `json:"private_ospd_vm_disk_gb"`
	PrivateOspdVMVcpus               string         `json:"private_ospd_vm_vcpus"`
	GCPRegion                        string         `json:"gcp_region"`
	AwsSubnet                        string         `json:"aws_subnet"`
	FQName                           []string       `json:"fq_name"`
	PrivateRedhatSubscriptionPasword string         `json:"private_redhat_subscription_pasword"`
	GCPAsn                           int            `json:"gcp_asn"`
	AwsAccessKey                     string         `json:"aws_access_key"`
	PrivateRedhatSubscriptionKey     string         `json:"private_redhat_subscription_key"`
	GCPSubnet                        string         `json:"gcp_subnet"`
	Perms2                           *PermType2     `json:"perms2"`
	PrivateOspdVMRAMMB               string         `json:"private_ospd_vm_ram_mb"`
	AwsSecretKey                     string         `json:"aws_secret_key"`
	PrivateRedhatPoolID              string         `json:"private_redhat_pool_id"`
	ProvisioningLog                  string         `json:"provisioning_log"`
	ProvisioningStartTime            string         `json:"provisioning_start_time"`
	Type                             string         `json:"type"`
	PrivateOspdPackageURL            string         `json:"private_ospd_package_url"`
	PrivateOspdUserName              string         `json:"private_ospd_user_name"`
	DisplayName                      string         `json:"display_name"`
	ProvisioningProgressStage        string         `json:"provisioning_progress_stage"`
	PrivateDNSServers                string         `json:"private_dns_servers"`
	PrivateNTPHosts                  string         `json:"private_ntp_hosts"`
	AwsRegion                        string         `json:"aws_region"`
	PrivateOspdVMName                string         `json:"private_ospd_vm_name"`
	PrivateRedhatSubscriptionUser    string         `json:"private_redhat_subscription_user"`
	Annotations                      *KeyValuePairs `json:"annotations"`
	PrivateOspdUserPassword          string         `json:"private_ospd_user_password"`
	GCPAccountInfo                   string         `json:"gcp_account_info"`
	UUID                             string         `json:"uuid"`
}

// Location parents relation object

// String returns json representation of the object
func (model *Location) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLocation makes Location
func MakeLocation() *Location {
	return &Location{
		//TODO(nati): Apply default
		PrivateDNSServers:                "",
		PrivateNTPHosts:                  "",
		AwsRegion:                        "",
		DisplayName:                      "",
		ProvisioningProgressStage:        "",
		PrivateOspdVMName:                "",
		PrivateRedhatSubscriptionUser:    "",
		Annotations:                      MakeKeyValuePairs(),
		PrivateOspdUserPassword:          "",
		GCPAccountInfo:                   "",
		UUID:                             "",
		PrivateOspdVMDiskGB:              "",
		PrivateOspdVMVcpus:               "",
		GCPRegion:                        "",
		IDPerms:                          MakeIdPermsType(),
		ProvisioningProgress:             0,
		ProvisioningState:                "",
		PrivateRedhatSubscriptionPasword: "",
		GCPAsn:       0,
		AwsAccessKey: "",
		AwsSubnet:    "",
		FQName:       []string{},
		PrivateRedhatSubscriptionKey: "",
		GCPSubnet:                    "",
		Perms2:                       MakePermType2(),
		PrivateOspdVMRAMMB:           "",
		AwsSecretKey:                 "",
		Type:                         "",
		PrivateOspdPackageURL: "",
		PrivateOspdUserName:   "",
		PrivateRedhatPoolID:   "",
		ProvisioningLog:       "",
		ProvisioningStartTime: "",
	}
}

// InterfaceToLocation makes Location from interface
func InterfaceToLocation(iData interface{}) *Location {
	data := iData.(map[string]interface{})
	return &Location{
		PrivateOspdVMName: data["private_ospd_vm_name"].(string),

		//{"Title":"OSPD Virtual Machine Name","Description":"Name of RH OSPD VM","SQL":"varchar(255)","Default":"undercloud","Operation":"","Presence":"","Type":"string","Permission":["create"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_ospd_vm_name","Item":null,"GoName":"PrivateOspdVMName","GoType":"string","GoPremitive":true}
		PrivateRedhatSubscriptionUser: data["private_redhat_subscription_user"].(string),

		//{"Title":"Redhat Subscription User","Description":"User name for RedHat subscription account","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_redhat_subscription_user","Item":null,"GoName":"PrivateRedhatSubscriptionUser","GoType":"string","GoPremitive":true}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"key_value_pair":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"key_value_pair","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"key":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Key","GoType":"string","GoPremitive":true},"value":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Value","GoType":"string","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePair","CollectionType":"","Column":"","Item":null,"GoName":"KeyValuePair","GoType":"KeyValuePair","GoPremitive":false},"GoName":"KeyValuePair","GoType":"[]*KeyValuePair","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/KeyValuePairs","CollectionType":"","Column":"","Item":null,"GoName":"Annotations","GoType":"KeyValuePairs","GoPremitive":false}
		PrivateOspdUserPassword: data["private_ospd_user_password"].(string),

		//{"Title":"OSPD User Passowrd","Description":"OSPD Passowrd for account","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_ospd_user_password","Item":null,"GoName":"PrivateOspdUserPassword","GoType":"string","GoPremitive":true}
		GCPAccountInfo: data["gcp_account_info"].(string),

		//{"Title":"Account info","Description":"copy and paste contents of account.json","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"gcp_account_info","Item":null,"GoName":"GCPAccountInfo","GoType":"string","GoPremitive":true}
		UUID: data["uuid"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"uuid","Item":null,"GoName":"UUID","GoType":"string","GoPremitive":true}
		ProvisioningState: data["provisioning_state"].(string),

		//{"Title":"Provisioning Status","Description":"","SQL":"varchar(255)","Default":"CREATED","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":{},"Enum":["CREATED","IN_CREATE_PROGRESS","UPDATED","IN_UPDATE_PROGRESS","DELETED","IN_DELETE_PROGRESS","ERROR"],"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"provisioning_state","Item":null,"GoName":"ProvisioningState","GoType":"string","GoPremitive":true}
		PrivateOspdVMDiskGB: data["private_ospd_vm_disk_gb"].(string),

		//{"Title":"OSPD Disk Size in gigabytes","Description":"disk spae to assign to RH OSPD vm","SQL":"varchar(255)","Default":"100","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_ospd_vm_disk_gb","Item":null,"GoName":"PrivateOspdVMDiskGB","GoType":"string","GoPremitive":true}
		PrivateOspdVMVcpus: data["private_ospd_vm_vcpus"].(string),

		//{"Title":"OSPD vCPUs","Description":"vcpus to assign to RH OSPD vm","SQL":"varchar(255)","Default":"8","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_ospd_vm_vcpus","Item":null,"GoName":"PrivateOspdVMVcpus","GoType":"string","GoPremitive":true}
		GCPRegion: data["gcp_region"].(string),

		//{"Title":"Region","Description":"","SQL":"varchar(255)","Default":"us-west1","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"gcp_region","Item":null,"GoName":"GCPRegion","GoType":"string","GoPremitive":true}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"created":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"created","Item":null,"GoName":"Created","GoType":"string","GoPremitive":true},"creator":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"creator","Item":null,"GoName":"Creator","GoType":"string","GoPremitive":true},"description":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"description","Item":null,"GoName":"Description","GoType":"string","GoPremitive":true},"enable":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"true","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"enable","Item":null,"GoName":"Enable","GoType":"bool","GoPremitive":true},"last_modified":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"system-only","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"last_modified","Item":null,"GoName":"LastModified","GoType":"string","GoPremitive":true},"permissions":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"group":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"group","Item":null,"GoName":"Group","GoType":"string","GoPremitive":true},"group_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"group_access","Item":null,"GoName":"GroupAccess","GoType":"AccessType","GoPremitive":false},"other_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"other_access","Item":null,"GoName":"OtherAccess","GoType":"AccessType","GoPremitive":false},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"owner","Item":null,"GoName":"Owner","GoType":"string","GoPremitive":true},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType","CollectionType":"","Column":"","Item":null,"GoName":"Permissions","GoType":"PermType","GoPremitive":false},"user_visible":{"Title":"","Description":"","SQL":"bool","Default":null,"Operation":"","Presence":"system-only","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"user_visible","Item":null,"GoName":"UserVisible","GoType":"bool","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IdPermsType","CollectionType":"","Column":"","Item":null,"GoName":"IDPerms","GoType":"IdPermsType","GoPremitive":false}
		ProvisioningProgress: data["provisioning_progress"].(int),

		//{"Title":"Provisioning Progress","Description":"","SQL":"int","Default":0,"Operation":"","Presence":"","Type":"integer","Permission":["create","update"],"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"provisioning_progress","Item":null,"GoName":"ProvisioningProgress","GoType":"int","GoPremitive":true}
		PrivateRedhatSubscriptionPasword: data["private_redhat_subscription_pasword"].(string),

		//{"Title":"Redhat Subscription Password","Description":"Password for subscription account","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_redhat_subscription_pasword","Item":null,"GoName":"PrivateRedhatSubscriptionPasword","GoType":"string","GoPremitive":true}
		GCPAsn: data["gcp_asn"].(int),

		//{"Title":"ASN","Description":"","SQL":"int","Default":65001,"Operation":"","Presence":"","Type":"integer","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"gcp_asn","Item":null,"GoName":"GCPAsn","GoType":"int","GoPremitive":true}
		AwsAccessKey: data["aws_access_key"].(string),

		//{"Title":"Access Key","Description":"","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"aws_access_key","Item":null,"GoName":"AwsAccessKey","GoType":"string","GoPremitive":true}
		AwsSubnet: data["aws_subnet"].(string),

		//{"Title":"Subnet","Description":"","SQL":"varchar(255)","Default":"10.0.0.0/16","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"aws_subnet","Item":null,"GoName":"AwsSubnet","GoType":"string","GoPremitive":true}
		FQName: data["fq_name"].([]string),

		//{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"fq_name","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"FQName","GoType":"string","GoPremitive":true},"GoName":"FQName","GoType":"[]string","GoPremitive":true}
		PrivateRedhatSubscriptionKey: data["private_redhat_subscription_key"].(string),

		//{"Title":"Redhat Subscription Acctivation Key","Description":"Subscription Activation Key","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_redhat_subscription_key","Item":null,"GoName":"PrivateRedhatSubscriptionKey","GoType":"string","GoPremitive":true}
		GCPSubnet: data["gcp_subnet"].(string),

		//{"Title":"Subnet","Description":"","SQL":"varchar(255)","Default":"10.1.0.0/16","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"gcp_subnet","Item":null,"GoName":"GCPSubnet","GoType":"string","GoPremitive":true}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"system-only","Type":"object","Permission":null,"Properties":{"global_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"global_access","Item":null,"GoName":"GlobalAccess","GoType":"AccessType","GoPremitive":false},"owner":{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"perms2_owner","Item":null,"GoName":"Owner","GoType":"string","GoPremitive":true},"owner_access":{"Title":"","Description":"","SQL":"int","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"perms2_owner_access","Item":null,"GoName":"OwnerAccess","GoType":"AccessType","GoPremitive":false},"share":{"Title":"","Description":"","SQL":"text","Default":null,"Operation":"","Presence":"optional","Type":"array","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"share","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"tenant":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Tenant","GoType":"string","GoPremitive":true},"tenant_access":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":0,"Maximum":7,"Ref":"types.json#/definitions/AccessType","CollectionType":"","Column":"","Item":null,"GoName":"TenantAccess","GoType":"AccessType","GoPremitive":false}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ShareType","CollectionType":"","Column":"","Item":null,"GoName":"Share","GoType":"ShareType","GoPremitive":false},"GoName":"Share","GoType":"[]*ShareType","GoPremitive":true}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/PermType2","CollectionType":"","Column":"","Item":null,"GoName":"Perms2","GoType":"PermType2","GoPremitive":false}
		PrivateOspdVMRAMMB: data["private_ospd_vm_ram_mb"].(string),

		//{"Title":"OSPD RAM in megabyts","Description":"ram to assign to RH OSPD vm","SQL":"varchar(255)","Default":"24576","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_ospd_vm_ram_mb","Item":null,"GoName":"PrivateOspdVMRAMMB","GoType":"string","GoPremitive":true}
		AwsSecretKey: data["aws_secret_key"].(string),

		//{"Title":"Secret Key","Description":"","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"aws_secret_key","Item":null,"GoName":"AwsSecretKey","GoType":"string","GoPremitive":true}
		ProvisioningStartTime: data["provisioning_start_time"].(string),

		//{"Title":"Time provisioning started","Description":"","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"provisioning_start_time","Item":null,"GoName":"ProvisioningStartTime","GoType":"string","GoPremitive":true}
		Type: data["type"].(string),

		//{"Title":"Location Type","Description":"Type of location","SQL":"varchar(255)","Default":"private","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":["private","aws","gcp","openstack"],"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"type","Item":null,"GoName":"Type","GoType":"string","GoPremitive":true}
		PrivateOspdPackageURL: data["private_ospd_package_url"].(string),

		//{"Title":"Location of OSPD Contrail Networking Packages","Description":"Location of Contrail Networking Packages","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_ospd_package_url","Item":null,"GoName":"PrivateOspdPackageURL","GoType":"string","GoPremitive":true}
		PrivateOspdUserName: data["private_ospd_user_name"].(string),

		//{"Title":"OSPD User Name","Description":"OSPD Non-Root User Account","SQL":"varchar(255)","Default":"stack","Operation":"","Presence":"","Type":"string","Permission":["create"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_ospd_user_name","Item":null,"GoName":"PrivateOspdUserName","GoType":"string","GoPremitive":true}
		PrivateRedhatPoolID: data["private_redhat_pool_id"].(string),

		//{"Title":"Redhat Pool ID","Description":"Repo Pool ID","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_redhat_pool_id","Item":null,"GoName":"PrivateRedhatPoolID","GoType":"string","GoPremitive":true}
		ProvisioningLog: data["provisioning_log"].(string),

		//{"Title":"Provisioning Log","Description":"","SQL":"text","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"provisioning_log","Item":null,"GoName":"ProvisioningLog","GoType":"string","GoPremitive":true}
		PrivateDNSServers: data["private_dns_servers"].(string),

		//{"Title":"DNS Servers","Description":"List of DNS servers","SQL":"varchar(255)","Default":"8.8.8.8","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_dns_servers","Item":null,"GoName":"PrivateDNSServers","GoType":"string","GoPremitive":true}
		PrivateNTPHosts: data["private_ntp_hosts"].(string),

		//{"Title":"NTP Hosts","Description":"List of NTP sources","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"private_ntp_hosts","Item":null,"GoName":"PrivateNTPHosts","GoType":"string","GoPremitive":true}
		AwsRegion: data["aws_region"].(string),

		//{"Title":"Region","Description":"","SQL":"varchar(255)","Default":"us-west-1","Operation":"","Presence":"","Type":"string","Permission":["create"],"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"aws_region","Item":null,"GoName":"AwsRegion","GoType":"string","GoPremitive":true}
		DisplayName: data["display_name"].(string),

		//{"Title":"","Description":"","SQL":"varchar(255)","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"display_name","Item":null,"GoName":"DisplayName","GoType":"string","GoPremitive":true}
		ProvisioningProgressStage: data["provisioning_progress_stage"].(string),

		//{"Title":"Provisioning Progress Stage","Description":"","SQL":"varchar(255)","Default":"","Operation":"","Presence":"","Type":"string","Permission":["create","update"],"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"provisioning_progress_stage","Item":null,"GoName":"ProvisioningProgressStage","GoType":"string","GoPremitive":true}

	}
}

// InterfaceToLocationSlice makes a slice of Location from interface
func InterfaceToLocationSlice(data interface{}) []*Location {
	list := data.([]interface{})
	result := MakeLocationSlice()
	for _, item := range list {
		result = append(result, InterfaceToLocation(item))
	}
	return result
}

// MakeLocationSlice() makes a slice of Location
func MakeLocationSlice() []*Location {
	return []*Location{}
}
