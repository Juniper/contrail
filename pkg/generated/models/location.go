package models

// Location

import "encoding/json"

// Location
type Location struct {
	Annotations                      *KeyValuePairs `json:"annotations"`
	PrivateDNSServers                string         `json:"private_dns_servers"`
	PrivateOspdUserPassword          string         `json:"private_ospd_user_password"`
	PrivateRedhatSubscriptionPasword string         `json:"private_redhat_subscription_pasword"`
	GCPAccountInfo                   string         `json:"gcp_account_info"`
	AwsAccessKey                     string         `json:"aws_access_key"`
	ParentType                       string         `json:"parent_type"`
	IDPerms                          *IdPermsType   `json:"id_perms"`
	PrivateOspdUserName              string         `json:"private_ospd_user_name"`
	PrivateOspdVMDiskGB              string         `json:"private_ospd_vm_disk_gb"`
	PrivateRedhatPoolID              string         `json:"private_redhat_pool_id"`
	PrivateRedhatSubscriptionKey     string         `json:"private_redhat_subscription_key"`
	PrivateNTPHosts                  string         `json:"private_ntp_hosts"`
	GCPRegion                        string         `json:"gcp_region"`
	UUID                             string         `json:"uuid"`
	ProvisioningStartTime            string         `json:"provisioning_start_time"`
	PrivateOspdPackageURL            string         `json:"private_ospd_package_url"`
	GCPSubnet                        string         `json:"gcp_subnet"`
	AwsSubnet                        string         `json:"aws_subnet"`
	DisplayName                      string         `json:"display_name"`
	ProvisioningProgressStage        string         `json:"provisioning_progress_stage"`
	Perms2                           *PermType2     `json:"perms2"`
	ProvisioningLog                  string         `json:"provisioning_log"`
	ProvisioningState                string         `json:"provisioning_state"`
	PrivateOspdVMName                string         `json:"private_ospd_vm_name"`
	PrivateOspdVMVcpus               string         `json:"private_ospd_vm_vcpus"`
	PrivateRedhatSubscriptionUser    string         `json:"private_redhat_subscription_user"`
	GCPAsn                           int            `json:"gcp_asn"`
	AwsSecretKey                     string         `json:"aws_secret_key"`
	ProvisioningProgress             int            `json:"provisioning_progress"`
	Type                             string         `json:"type"`
	PrivateOspdVMRAMMB               string         `json:"private_ospd_vm_ram_mb"`
	AwsRegion                        string         `json:"aws_region"`
	FQName                           []string       `json:"fq_name"`
	ParentUUID                       string         `json:"parent_uuid"`

	PhysicalRouters []*PhysicalRouter `json:"physical_routers"`
}

// String returns json representation of the object
func (model *Location) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeLocation makes Location
func MakeLocation() *Location {
	return &Location{
		//TODO(nati): Apply default
		PrivateOspdUserName:          "",
		PrivateOspdVMDiskGB:          "",
		PrivateRedhatPoolID:          "",
		PrivateRedhatSubscriptionKey: "",
		PrivateNTPHosts:              "",
		GCPRegion:                    "",
		UUID:                         "",
		PrivateOspdPackageURL:         "",
		GCPSubnet:                     "",
		AwsSubnet:                     "",
		DisplayName:                   "",
		ProvisioningProgressStage:     "",
		ProvisioningStartTime:         "",
		ProvisioningState:             "",
		PrivateOspdVMName:             "",
		PrivateOspdVMVcpus:            "",
		PrivateRedhatSubscriptionUser: "",
		GCPAsn:                           0,
		AwsSecretKey:                     "",
		Perms2:                           MakePermType2(),
		ProvisioningLog:                  "",
		Type:                             "",
		PrivateOspdVMRAMMB:               "",
		AwsRegion:                        "",
		FQName:                           []string{},
		ParentUUID:                       "",
		ProvisioningProgress:             0,
		PrivateDNSServers:                "",
		PrivateOspdUserPassword:          "",
		PrivateRedhatSubscriptionPasword: "",
		GCPAccountInfo:                   "",
		AwsAccessKey:                     "",
		Annotations:                      MakeKeyValuePairs(),
		ParentType:                       "",
		IDPerms:                          MakeIdPermsType(),
	}
}

// InterfaceToLocation makes Location from interface
func InterfaceToLocation(iData interface{}) *Location {
	data := iData.(map[string]interface{})
	return &Location{
		AwsRegion: data["aws_region"].(string),

		//{"title":"Region","default":"us-west-1","type":"string","permission":["create"]}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ProvisioningProgress: data["provisioning_progress"].(int),

		//{"title":"Provisioning Progress","default":0,"type":"integer","permission":["create","update"]}
		Type: data["type"].(string),

		//{"title":"Location Type","description":"Type of location","default":"private","type":"string","permission":["create","update"],"enum":["private","aws","gcp","openstack"]}
		PrivateOspdVMRAMMB: data["private_ospd_vm_ram_mb"].(string),

		//{"title":"OSPD RAM in megabyts","description":"ram to assign to RH OSPD vm","default":"24576","type":"string","permission":["create","update"]}
		PrivateRedhatSubscriptionPasword: data["private_redhat_subscription_pasword"].(string),

		//{"title":"Redhat Subscription Password","description":"Password for subscription account","default":"","type":"string","permission":["create","update"]}
		GCPAccountInfo: data["gcp_account_info"].(string),

		//{"title":"Account info","description":"copy and paste contents of account.json","default":"","type":"string","permission":["create","update"]}
		AwsAccessKey: data["aws_access_key"].(string),

		//{"title":"Access Key","default":"","type":"string","permission":["create","update"]}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		PrivateDNSServers: data["private_dns_servers"].(string),

		//{"title":"DNS Servers","description":"List of DNS servers","default":"8.8.8.8","type":"string","permission":["create","update"]}
		PrivateOspdUserPassword: data["private_ospd_user_password"].(string),

		//{"title":"OSPD User Passowrd","description":"OSPD Passowrd for account","default":"","type":"string","permission":["create","update"]}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		PrivateRedhatPoolID: data["private_redhat_pool_id"].(string),

		//{"title":"Redhat Pool ID","description":"Repo Pool ID","default":"","type":"string","permission":["create","update"]}
		PrivateRedhatSubscriptionKey: data["private_redhat_subscription_key"].(string),

		//{"title":"Redhat Subscription Acctivation Key","description":"Subscription Activation Key","default":"","type":"string","permission":["create","update"]}
		PrivateOspdUserName: data["private_ospd_user_name"].(string),

		//{"title":"OSPD User Name","description":"OSPD Non-Root User Account","default":"stack","type":"string","permission":["create"]}
		PrivateOspdVMDiskGB: data["private_ospd_vm_disk_gb"].(string),

		//{"title":"OSPD Disk Size in gigabytes","description":"disk spae to assign to RH OSPD vm","default":"100","type":"string","permission":["create","update"]}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		PrivateNTPHosts: data["private_ntp_hosts"].(string),

		//{"title":"NTP Hosts","description":"List of NTP sources","default":"","type":"string","permission":["create","update"]}
		GCPRegion: data["gcp_region"].(string),

		//{"title":"Region","default":"us-west1","type":"string","permission":["create","update"]}
		AwsSubnet: data["aws_subnet"].(string),

		//{"title":"Subnet","default":"10.0.0.0/16","type":"string","permission":["create","update"]}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		ProvisioningProgressStage: data["provisioning_progress_stage"].(string),

		//{"title":"Provisioning Progress Stage","default":"","type":"string","permission":["create","update"]}
		ProvisioningStartTime: data["provisioning_start_time"].(string),

		//{"title":"Time provisioning started","default":"","type":"string","permission":["create","update"]}
		PrivateOspdPackageURL: data["private_ospd_package_url"].(string),

		//{"title":"Location of OSPD Contrail Networking Packages","description":"Location of Contrail Networking Packages","default":"","type":"string","permission":["create","update"]}
		GCPSubnet: data["gcp_subnet"].(string),

		//{"title":"Subnet","default":"10.1.0.0/16","type":"string","permission":["create","update"]}
		PrivateRedhatSubscriptionUser: data["private_redhat_subscription_user"].(string),

		//{"title":"Redhat Subscription User","description":"User name for RedHat subscription account","default":"","type":"string","permission":["create","update"]}
		GCPAsn: data["gcp_asn"].(int),

		//{"title":"ASN","default":65001,"type":"integer","permission":["create","update"]}
		AwsSecretKey: data["aws_secret_key"].(string),

		//{"title":"Secret Key","default":"","type":"string","permission":["create","update"]}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ProvisioningLog: data["provisioning_log"].(string),

		//{"title":"Provisioning Log","default":"","type":"string","permission":["create","update"]}
		ProvisioningState: data["provisioning_state"].(string),

		//{"title":"Provisioning Status","default":"CREATED","type":"string","permission":["create","update"],"enum":["CREATED","IN_CREATE_PROGRESS","UPDATED","IN_UPDATE_PROGRESS","DELETED","IN_DELETE_PROGRESS","ERROR"]}
		PrivateOspdVMName: data["private_ospd_vm_name"].(string),

		//{"title":"OSPD Virtual Machine Name","description":"Name of RH OSPD VM","default":"undercloud","type":"string","permission":["create"]}
		PrivateOspdVMVcpus: data["private_ospd_vm_vcpus"].(string),

		//{"title":"OSPD vCPUs","description":"vcpus to assign to RH OSPD vm","default":"8","type":"string","permission":["create","update"]}

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
