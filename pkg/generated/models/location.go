package models

// Location

import "encoding/json"

// Location
type Location struct {
	PrivateOspdUserName              string         `json:"private_ospd_user_name,omitempty"`
	PrivateOspdVMVcpus               string         `json:"private_ospd_vm_vcpus,omitempty"`
	GCPRegion                        string         `json:"gcp_region,omitempty"`
	AwsSubnet                        string         `json:"aws_subnet,omitempty"`
	FQName                           []string       `json:"fq_name,omitempty"`
	DisplayName                      string         `json:"display_name,omitempty"`
	UUID                             string         `json:"uuid,omitempty"`
	ParentType                       string         `json:"parent_type,omitempty"`
	ProvisioningState                string         `json:"provisioning_state,omitempty"`
	PrivateOspdVMDiskGB              string         `json:"private_ospd_vm_disk_gb,omitempty"`
	Perms2                           *PermType2     `json:"perms2,omitempty"`
	GCPAsn                           int            `json:"gcp_asn,omitempty"`
	ProvisioningProgress             int            `json:"provisioning_progress,omitempty"`
	PrivateNTPHosts                  string         `json:"private_ntp_hosts,omitempty"`
	PrivateOspdPackageURL            string         `json:"private_ospd_package_url,omitempty"`
	PrivateOspdVMName                string         `json:"private_ospd_vm_name,omitempty"`
	PrivateRedhatPoolID              string         `json:"private_redhat_pool_id,omitempty"`
	PrivateRedhatSubscriptionPasword string         `json:"private_redhat_subscription_pasword,omitempty"`
	ProvisioningStartTime            string         `json:"provisioning_start_time,omitempty"`
	Type                             string         `json:"type,omitempty"`
	PrivateDNSServers                string         `json:"private_dns_servers,omitempty"`
	PrivateOspdUserPassword          string         `json:"private_ospd_user_password,omitempty"`
	PrivateOspdVMRAMMB               string         `json:"private_ospd_vm_ram_mb,omitempty"`
	PrivateRedhatSubscriptionKey     string         `json:"private_redhat_subscription_key,omitempty"`
	PrivateRedhatSubscriptionUser    string         `json:"private_redhat_subscription_user,omitempty"`
	GCPAccountInfo                   string         `json:"gcp_account_info,omitempty"`
	AwsAccessKey                     string         `json:"aws_access_key,omitempty"`
	AwsRegion                        string         `json:"aws_region,omitempty"`
	Annotations                      *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningLog                  string         `json:"provisioning_log,omitempty"`
	GCPSubnet                        string         `json:"gcp_subnet,omitempty"`
	AwsSecretKey                     string         `json:"aws_secret_key,omitempty"`
	ProvisioningProgressStage        string         `json:"provisioning_progress_stage,omitempty"`
	IDPerms                          *IdPermsType   `json:"id_perms,omitempty"`
	ParentUUID                       string         `json:"parent_uuid,omitempty"`

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
		GCPSubnet:                        "",
		AwsSecretKey:                     "",
		ProvisioningProgressStage:        "",
		IDPerms:                          MakeIdPermsType(),
		ParentUUID:                       "",
		UUID:                             "",
		ParentType:                       "",
		PrivateOspdUserName:              "",
		PrivateOspdVMVcpus:               "",
		GCPRegion:                        "",
		AwsSubnet:                        "",
		FQName:                           []string{},
		DisplayName:                      "",
		ProvisioningState:                "",
		PrivateOspdVMDiskGB:              "",
		Perms2:                           MakePermType2(),
		GCPAsn:                           0,
		ProvisioningProgress:             0,
		PrivateNTPHosts:                  "",
		PrivateOspdPackageURL:            "",
		PrivateOspdVMName:                "",
		PrivateRedhatPoolID:              "",
		PrivateRedhatSubscriptionPasword: "",
		ProvisioningStartTime:            "",
		GCPAccountInfo:                   "",
		Type:                             "",
		PrivateDNSServers:                "",
		PrivateOspdUserPassword:          "",
		PrivateOspdVMRAMMB:               "",
		PrivateRedhatSubscriptionKey:     "",
		PrivateRedhatSubscriptionUser:    "",
		AwsAccessKey:                     "",
		AwsRegion:                        "",
		Annotations:                      MakeKeyValuePairs(),
		ProvisioningLog:                  "",
	}
}

// MakeLocationSlice() makes a slice of Location
func MakeLocationSlice() []*Location {
	return []*Location{}
}
