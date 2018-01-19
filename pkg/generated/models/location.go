package models

// Location

import "encoding/json"

// Location
type Location struct {
	PrivateOspdVMVcpus               string         `json:"private_ospd_vm_vcpus,omitempty"`
	PrivateRedhatSubscriptionUser    string         `json:"private_redhat_subscription_user,omitempty"`
	AwsAccessKey                     string         `json:"aws_access_key,omitempty"`
	IDPerms                          *IdPermsType   `json:"id_perms,omitempty"`
	PrivateNTPHosts                  string         `json:"private_ntp_hosts,omitempty"`
	PrivateRedhatPoolID              string         `json:"private_redhat_pool_id,omitempty"`
	ProvisioningStartTime            string         `json:"provisioning_start_time,omitempty"`
	ProvisioningLog                  string         `json:"provisioning_log,omitempty"`
	ProvisioningProgress             int            `json:"provisioning_progress,omitempty"`
	PrivateOspdPackageURL            string         `json:"private_ospd_package_url,omitempty"`
	PrivateRedhatSubscriptionPasword string         `json:"private_redhat_subscription_pasword,omitempty"`
	GCPSubnet                        string         `json:"gcp_subnet,omitempty"`
	AwsSubnet                        string         `json:"aws_subnet,omitempty"`
	Annotations                      *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                           *PermType2     `json:"perms2,omitempty"`
	Type                             string         `json:"type,omitempty"`
	PrivateOspdVMDiskGB              string         `json:"private_ospd_vm_disk_gb,omitempty"`
	ParentUUID                       string         `json:"parent_uuid,omitempty"`
	ProvisioningProgressStage        string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningState                string         `json:"provisioning_state,omitempty"`
	AwsRegion                        string         `json:"aws_region,omitempty"`
	ParentType                       string         `json:"parent_type,omitempty"`
	GCPAccountInfo                   string         `json:"gcp_account_info,omitempty"`
	FQName                           []string       `json:"fq_name,omitempty"`
	PrivateOspdUserName              string         `json:"private_ospd_user_name,omitempty"`
	PrivateRedhatSubscriptionKey     string         `json:"private_redhat_subscription_key,omitempty"`
	GCPAsn                           int            `json:"gcp_asn,omitempty"`
	GCPRegion                        string         `json:"gcp_region,omitempty"`
	AwsSecretKey                     string         `json:"aws_secret_key,omitempty"`
	PrivateOspdUserPassword          string         `json:"private_ospd_user_password,omitempty"`
	PrivateOspdVMName                string         `json:"private_ospd_vm_name,omitempty"`
	DisplayName                      string         `json:"display_name,omitempty"`
	UUID                             string         `json:"uuid,omitempty"`
	PrivateDNSServers                string         `json:"private_dns_servers,omitempty"`
	PrivateOspdVMRAMMB               string         `json:"private_ospd_vm_ram_mb,omitempty"`

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
		GCPAsn:                           0,
		GCPRegion:                        "",
		AwsSecretKey:                     "",
		PrivateOspdUserPassword:          "",
		PrivateOspdVMName:                "",
		DisplayName:                      "",
		UUID:                             "",
		PrivateDNSServers:                "",
		PrivateOspdVMRAMMB:               "",
		PrivateOspdVMVcpus:               "",
		PrivateRedhatSubscriptionUser:    "",
		AwsAccessKey:                     "",
		IDPerms:                          MakeIdPermsType(),
		PrivateNTPHosts:                  "",
		PrivateRedhatPoolID:              "",
		ProvisioningStartTime:            "",
		ProvisioningLog:                  "",
		ProvisioningProgress:             0,
		PrivateOspdPackageURL:            "",
		PrivateRedhatSubscriptionPasword: "",
		GCPSubnet:                        "",
		AwsSubnet:                        "",
		Annotations:                      MakeKeyValuePairs(),
		Perms2:                           MakePermType2(),
		Type:                             "",
		PrivateOspdVMDiskGB:              "",
		ParentUUID:                       "",
		ProvisioningProgressStage:        "",
		ProvisioningState:                "",
		AwsRegion:                        "",
		ParentType:                       "",
		GCPAccountInfo:                   "",
		FQName:                           []string{},
		PrivateOspdUserName:              "",
		PrivateRedhatSubscriptionKey:     "",
	}
}

// MakeLocationSlice() makes a slice of Location
func MakeLocationSlice() []*Location {
	return []*Location{}
}
