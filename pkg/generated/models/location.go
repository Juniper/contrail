package models

// Location

import "encoding/json"

// Location
type Location struct {
	Perms2                           *PermType2     `json:"perms2"`
	Annotations                      *KeyValuePairs `json:"annotations"`
	ProvisioningStartTime            string         `json:"provisioning_start_time"`
	PrivateOspdVMRAMMB               string         `json:"private_ospd_vm_ram_mb"`
	PrivateRedhatSubscriptionPasword string         `json:"private_redhat_subscription_pasword"`
	GCPRegion                        string         `json:"gcp_region"`
	GCPSubnet                        string         `json:"gcp_subnet"`
	PrivateOspdVMVcpus               string         `json:"private_ospd_vm_vcpus"`
	PrivateRedhatPoolID              string         `json:"private_redhat_pool_id"`
	PrivateRedhatSubscriptionKey     string         `json:"private_redhat_subscription_key"`
	UUID                             string         `json:"uuid"`
	PrivateOspdPackageURL            string         `json:"private_ospd_package_url"`
	AwsSecretKey                     string         `json:"aws_secret_key"`
	AwsAccessKey                     string         `json:"aws_access_key"`
	AwsSubnet                        string         `json:"aws_subnet"`
	FQName                           []string       `json:"fq_name"`
	PrivateOspdVMName                string         `json:"private_ospd_vm_name"`
	GCPAsn                           int            `json:"gcp_asn"`
	AwsRegion                        string         `json:"aws_region"`
	ParentUUID                       string         `json:"parent_uuid"`
	PrivateDNSServers                string         `json:"private_dns_servers"`
	PrivateOspdUserName              string         `json:"private_ospd_user_name"`
	PrivateOspdUserPassword          string         `json:"private_ospd_user_password"`
	PrivateOspdVMDiskGB              string         `json:"private_ospd_vm_disk_gb"`
	ProvisioningProgress             int            `json:"provisioning_progress"`
	ProvisioningProgressStage        string         `json:"provisioning_progress_stage"`
	Type                             string         `json:"type"`
	ParentType                       string         `json:"parent_type"`
	IDPerms                          *IdPermsType   `json:"id_perms"`
	ProvisioningState                string         `json:"provisioning_state"`
	ProvisioningLog                  string         `json:"provisioning_log"`
	PrivateNTPHosts                  string         `json:"private_ntp_hosts"`
	PrivateRedhatSubscriptionUser    string         `json:"private_redhat_subscription_user"`
	GCPAccountInfo                   string         `json:"gcp_account_info"`
	DisplayName                      string         `json:"display_name"`

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
		AwsAccessKey:              "",
		AwsSubnet:                 "",
		FQName:                    []string{},
		PrivateDNSServers:         "",
		PrivateOspdUserName:       "",
		PrivateOspdUserPassword:   "",
		PrivateOspdVMDiskGB:       "",
		PrivateOspdVMName:         "",
		GCPAsn:                    0,
		AwsRegion:                 "",
		ParentUUID:                "",
		ProvisioningProgress:      0,
		ProvisioningProgressStage: "",
		Type:                             "",
		ParentType:                       "",
		IDPerms:                          MakeIdPermsType(),
		PrivateNTPHosts:                  "",
		PrivateRedhatSubscriptionUser:    "",
		GCPAccountInfo:                   "",
		DisplayName:                      "",
		ProvisioningState:                "",
		ProvisioningLog:                  "",
		PrivateOspdVMRAMMB:               "",
		PrivateRedhatSubscriptionPasword: "",
		GCPRegion:                        "",
		GCPSubnet:                        "",
		Perms2:                           MakePermType2(),
		Annotations:                      MakeKeyValuePairs(),
		ProvisioningStartTime:            "",
		PrivateOspdVMVcpus:               "",
		PrivateRedhatPoolID:              "",
		PrivateRedhatSubscriptionKey:     "",
		UUID: "",
		PrivateOspdPackageURL: "",
		AwsSecretKey:          "",
	}
}

// MakeLocationSlice() makes a slice of Location
func MakeLocationSlice() []*Location {
	return []*Location{}
}
