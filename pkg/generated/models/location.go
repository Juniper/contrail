package models

// Location

import "encoding/json"

// Location
type Location struct {
	PrivateOspdUserName              string         `json:"private_ospd_user_name,omitempty"`
	PrivateOspdUserPassword          string         `json:"private_ospd_user_password,omitempty"`
	PrivateRedhatSubscriptionPasword string         `json:"private_redhat_subscription_pasword,omitempty"`
	Type                             string         `json:"type,omitempty"`
	PrivateOspdVMRAMMB               string         `json:"private_ospd_vm_ram_mb,omitempty"`
	GCPAccountInfo                   string         `json:"gcp_account_info,omitempty"`
	FQName                           []string       `json:"fq_name,omitempty"`
	ParentUUID                       string         `json:"parent_uuid,omitempty"`
	AwsRegion                        string         `json:"aws_region,omitempty"`
	Annotations                      *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningStartTime            string         `json:"provisioning_start_time,omitempty"`
	PrivateOspdVMVcpus               string         `json:"private_ospd_vm_vcpus,omitempty"`
	ProvisioningState                string         `json:"provisioning_state,omitempty"`
	ProvisioningProgress             int            `json:"provisioning_progress,omitempty"`
	ProvisioningProgressStage        string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningLog                  string         `json:"provisioning_log,omitempty"`
	PrivateDNSServers                string         `json:"private_dns_servers,omitempty"`
	PrivateNTPHosts                  string         `json:"private_ntp_hosts,omitempty"`
	PrivateOspdVMName                string         `json:"private_ospd_vm_name,omitempty"`
	PrivateRedhatPoolID              string         `json:"private_redhat_pool_id,omitempty"`
	PrivateRedhatSubscriptionKey     string         `json:"private_redhat_subscription_key,omitempty"`
	GCPSubnet                        string         `json:"gcp_subnet,omitempty"`
	AwsSecretKey                     string         `json:"aws_secret_key,omitempty"`
	AwsSubnet                        string         `json:"aws_subnet,omitempty"`
	UUID                             string         `json:"uuid,omitempty"`
	IDPerms                          *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName                      string         `json:"display_name,omitempty"`
	PrivateOspdVMDiskGB              string         `json:"private_ospd_vm_disk_gb,omitempty"`
	PrivateRedhatSubscriptionUser    string         `json:"private_redhat_subscription_user,omitempty"`
	GCPAsn                           int            `json:"gcp_asn,omitempty"`
	GCPRegion                        string         `json:"gcp_region,omitempty"`
	AwsAccessKey                     string         `json:"aws_access_key,omitempty"`
	ParentType                       string         `json:"parent_type,omitempty"`
	Perms2                           *PermType2     `json:"perms2,omitempty"`
	PrivateOspdPackageURL            string         `json:"private_ospd_package_url,omitempty"`

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
		PrivateOspdUserName:              "",
		PrivateOspdUserPassword:          "",
		PrivateRedhatSubscriptionPasword: "",
		Type:                          "",
		PrivateOspdVMRAMMB:            "",
		GCPAccountInfo:                "",
		FQName:                        []string{},
		ParentUUID:                    "",
		AwsRegion:                     "",
		Annotations:                   MakeKeyValuePairs(),
		ProvisioningStartTime:         "",
		PrivateOspdVMVcpus:            "",
		ProvisioningState:             "",
		ProvisioningProgress:          0,
		ProvisioningProgressStage:     "",
		PrivateDNSServers:             "",
		PrivateNTPHosts:               "",
		PrivateOspdVMName:             "",
		PrivateRedhatPoolID:           "",
		PrivateRedhatSubscriptionKey:  "",
		GCPSubnet:                     "",
		ProvisioningLog:               "",
		AwsSecretKey:                  "",
		AwsSubnet:                     "",
		UUID:                          "",
		PrivateOspdVMDiskGB:           "",
		PrivateRedhatSubscriptionUser: "",
		GCPAsn:                0,
		GCPRegion:             "",
		AwsAccessKey:          "",
		ParentType:            "",
		IDPerms:               MakeIdPermsType(),
		DisplayName:           "",
		Perms2:                MakePermType2(),
		PrivateOspdPackageURL: "",
	}
}

// MakeLocationSlice() makes a slice of Location
func MakeLocationSlice() []*Location {
	return []*Location{}
}
