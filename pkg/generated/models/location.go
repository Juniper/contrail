package models

// Location

import "encoding/json"

// Location
type Location struct {
	PrivateOspdUserName              string         `json:"private_ospd_user_name,omitempty"`
	ProvisioningProgress             int            `json:"provisioning_progress,omitempty"`
	PrivateDNSServers                string         `json:"private_dns_servers,omitempty"`
	PrivateRedhatSubscriptionPasword string         `json:"private_redhat_subscription_pasword,omitempty"`
	GCPRegion                        string         `json:"gcp_region,omitempty"`
	DisplayName                      string         `json:"display_name,omitempty"`
	PrivateOspdVMName                string         `json:"private_ospd_vm_name,omitempty"`
	PrivateRedhatSubscriptionKey     string         `json:"private_redhat_subscription_key,omitempty"`
	GCPAsn                           int            `json:"gcp_asn,omitempty"`
	IDPerms                          *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningLog                  string         `json:"provisioning_log,omitempty"`
	PrivateNTPHosts                  string         `json:"private_ntp_hosts,omitempty"`
	GCPAccountInfo                   string         `json:"gcp_account_info,omitempty"`
	AwsSubnet                        string         `json:"aws_subnet,omitempty"`
	ParentType                       string         `json:"parent_type,omitempty"`
	ProvisioningState                string         `json:"provisioning_state,omitempty"`
	ProvisioningProgressStage        string         `json:"provisioning_progress_stage,omitempty"`
	PrivateOspdVMVcpus               string         `json:"private_ospd_vm_vcpus,omitempty"`
	GCPSubnet                        string         `json:"gcp_subnet,omitempty"`
	UUID                             string         `json:"uuid,omitempty"`
	Perms2                           *PermType2     `json:"perms2,omitempty"`
	FQName                           []string       `json:"fq_name,omitempty"`
	Annotations                      *KeyValuePairs `json:"annotations,omitempty"`
	PrivateOspdPackageURL            string         `json:"private_ospd_package_url,omitempty"`
	PrivateOspdVMDiskGB              string         `json:"private_ospd_vm_disk_gb,omitempty"`
	AwsAccessKey                     string         `json:"aws_access_key,omitempty"`
	AwsSecretKey                     string         `json:"aws_secret_key,omitempty"`
	PrivateRedhatPoolID              string         `json:"private_redhat_pool_id,omitempty"`
	PrivateRedhatSubscriptionUser    string         `json:"private_redhat_subscription_user,omitempty"`
	ProvisioningStartTime            string         `json:"provisioning_start_time,omitempty"`
	ParentUUID                       string         `json:"parent_uuid,omitempty"`
	Type                             string         `json:"type,omitempty"`
	PrivateOspdUserPassword          string         `json:"private_ospd_user_password,omitempty"`
	PrivateOspdVMRAMMB               string         `json:"private_ospd_vm_ram_mb,omitempty"`
	AwsRegion                        string         `json:"aws_region,omitempty"`

	PhysicalRouters []*PhysicalRouter `json:"physical_routers,omitempty"`
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
		PrivateOspdVMRAMMB: "",
		AwsRegion:          "",
		ParentUUID:         "",
		Type:               "",
		PrivateOspdUserPassword:          "",
		PrivateOspdUserName:              "",
		ProvisioningProgress:             0,
		GCPRegion:                        "",
		DisplayName:                      "",
		PrivateDNSServers:                "",
		PrivateRedhatSubscriptionPasword: "",
		GCPAsn:                        0,
		IDPerms:                       MakeIdPermsType(),
		PrivateOspdVMName:             "",
		PrivateRedhatSubscriptionKey:  "",
		AwsSubnet:                     "",
		ParentType:                    "",
		ProvisioningLog:               "",
		PrivateNTPHosts:               "",
		GCPAccountInfo:                "",
		UUID:                          "",
		Perms2:                        MakePermType2(),
		ProvisioningState:             "",
		ProvisioningProgressStage:     "",
		PrivateOspdVMVcpus:            "",
		GCPSubnet:                     "",
		AwsAccessKey:                  "",
		AwsSecretKey:                  "",
		FQName:                        []string{},
		Annotations:                   MakeKeyValuePairs(),
		PrivateOspdPackageURL:         "",
		PrivateOspdVMDiskGB:           "",
		ProvisioningStartTime:         "",
		PrivateRedhatPoolID:           "",
		PrivateRedhatSubscriptionUser: "",
	}
}

// MakeLocationSlice() makes a slice of Location
func MakeLocationSlice() []*Location {
	return []*Location{}
}
