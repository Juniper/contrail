package models

// Location

import "encoding/json"

// Location
type Location struct {
	ProvisioningLog                  string         `json:"provisioning_log,omitempty"`
	UUID                             string         `json:"uuid,omitempty"`
	FQName                           []string       `json:"fq_name,omitempty"`
	ProvisioningProgressStage        string         `json:"provisioning_progress_stage,omitempty"`
	GCPAsn                           int            `json:"gcp_asn,omitempty"`
	Annotations                      *KeyValuePairs `json:"annotations,omitempty"`
	GCPAccountInfo                   string         `json:"gcp_account_info,omitempty"`
	AwsSecretKey                     string         `json:"aws_secret_key,omitempty"`
	Type                             string         `json:"type,omitempty"`
	PrivateOspdUserPassword          string         `json:"private_ospd_user_password,omitempty"`
	PrivateOspdVMName                string         `json:"private_ospd_vm_name,omitempty"`
	PrivateRedhatSubscriptionKey     string         `json:"private_redhat_subscription_key,omitempty"`
	PrivateRedhatSubscriptionPasword string         `json:"private_redhat_subscription_pasword,omitempty"`
	PrivateNTPHosts                  string         `json:"private_ntp_hosts,omitempty"`
	PrivateOspdUserName              string         `json:"private_ospd_user_name,omitempty"`
	PrivateRedhatPoolID              string         `json:"private_redhat_pool_id,omitempty"`
	ProvisioningProgress             int            `json:"provisioning_progress,omitempty"`
	ProvisioningState                string         `json:"provisioning_state,omitempty"`
	ProvisioningStartTime            string         `json:"provisioning_start_time,omitempty"`
	PrivateDNSServers                string         `json:"private_dns_servers,omitempty"`
	PrivateOspdVMDiskGB              string         `json:"private_ospd_vm_disk_gb,omitempty"`
	PrivateRedhatSubscriptionUser    string         `json:"private_redhat_subscription_user,omitempty"`
	GCPRegion                        string         `json:"gcp_region,omitempty"`
	ParentUUID                       string         `json:"parent_uuid,omitempty"`
	GCPSubnet                        string         `json:"gcp_subnet,omitempty"`
	AwsSubnet                        string         `json:"aws_subnet,omitempty"`
	Perms2                           *PermType2     `json:"perms2,omitempty"`
	IDPerms                          *IdPermsType   `json:"id_perms,omitempty"`
	ParentType                       string         `json:"parent_type,omitempty"`
	DisplayName                      string         `json:"display_name,omitempty"`
	PrivateOspdPackageURL            string         `json:"private_ospd_package_url,omitempty"`
	PrivateOspdVMRAMMB               string         `json:"private_ospd_vm_ram_mb,omitempty"`
	PrivateOspdVMVcpus               string         `json:"private_ospd_vm_vcpus,omitempty"`
	AwsAccessKey                     string         `json:"aws_access_key,omitempty"`
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
		ProvisioningState:             "",
		PrivateNTPHosts:               "",
		PrivateOspdUserName:           "",
		PrivateRedhatPoolID:           "",
		ProvisioningProgress:          0,
		ParentUUID:                    "",
		ProvisioningStartTime:         "",
		PrivateDNSServers:             "",
		PrivateOspdVMDiskGB:           "",
		PrivateRedhatSubscriptionUser: "",
		GCPRegion:                     "",
		GCPSubnet:                     "",
		AwsSubnet:                     "",
		Perms2:                        MakePermType2(),
		IDPerms:                       MakeIdPermsType(),
		AwsRegion:                     "",
		ParentType:                    "",
		DisplayName:                   "",
		PrivateOspdPackageURL:         "",
		PrivateOspdVMRAMMB:            "",
		PrivateOspdVMVcpus:            "",
		AwsAccessKey:                  "",
		ProvisioningLog:               "",
		UUID:                          "",
		FQName:                        []string{},
		ProvisioningProgressStage: "",
		GCPAsn:                           0,
		Annotations:                      MakeKeyValuePairs(),
		PrivateRedhatSubscriptionPasword: "",
		GCPAccountInfo:                   "",
		AwsSecretKey:                     "",
		Type:                             "",
		PrivateOspdUserPassword:      "",
		PrivateOspdVMName:            "",
		PrivateRedhatSubscriptionKey: "",
	}
}

// MakeLocationSlice() makes a slice of Location
func MakeLocationSlice() []*Location {
	return []*Location{}
}
