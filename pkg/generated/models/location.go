package models

// Location

import "encoding/json"

// Location
type Location struct {
	AwsAccessKey                     string         `json:"aws_access_key,omitempty"`
	AwsSecretKey                     string         `json:"aws_secret_key,omitempty"`
	ProvisioningProgress             int            `json:"provisioning_progress,omitempty"`
	PrivateNTPHosts                  string         `json:"private_ntp_hosts,omitempty"`
	PrivateRedhatSubscriptionUser    string         `json:"private_redhat_subscription_user,omitempty"`
	PrivateOspdVMVcpus               string         `json:"private_ospd_vm_vcpus,omitempty"`
	ParentType                       string         `json:"parent_type,omitempty"`
	IDPerms                          *IdPermsType   `json:"id_perms,omitempty"`
	PrivateOspdUserName              string         `json:"private_ospd_user_name,omitempty"`
	PrivateOspdUserPassword          string         `json:"private_ospd_user_password,omitempty"`
	PrivateOspdVMDiskGB              string         `json:"private_ospd_vm_disk_gb,omitempty"`
	ProvisioningState                string         `json:"provisioning_state,omitempty"`
	PrivateOspdVMRAMMB               string         `json:"private_ospd_vm_ram_mb,omitempty"`
	Perms2                           *PermType2     `json:"perms2,omitempty"`
	ProvisioningLog                  string         `json:"provisioning_log,omitempty"`
	Type                             string         `json:"type,omitempty"`
	PrivateOspdVMName                string         `json:"private_ospd_vm_name,omitempty"`
	PrivateRedhatSubscriptionKey     string         `json:"private_redhat_subscription_key,omitempty"`
	GCPAccountInfo                   string         `json:"gcp_account_info,omitempty"`
	GCPSubnet                        string         `json:"gcp_subnet,omitempty"`
	Annotations                      *KeyValuePairs `json:"annotations,omitempty"`
	UUID                             string         `json:"uuid,omitempty"`
	FQName                           []string       `json:"fq_name,omitempty"`
	PrivateDNSServers                string         `json:"private_dns_servers,omitempty"`
	PrivateRedhatPoolID              string         `json:"private_redhat_pool_id,omitempty"`
	GCPRegion                        string         `json:"gcp_region,omitempty"`
	ProvisioningProgressStage        string         `json:"provisioning_progress_stage,omitempty"`
	PrivateOspdPackageURL            string         `json:"private_ospd_package_url,omitempty"`
	GCPAsn                           int            `json:"gcp_asn,omitempty"`
	AwsSubnet                        string         `json:"aws_subnet,omitempty"`
	DisplayName                      string         `json:"display_name,omitempty"`
	ParentUUID                       string         `json:"parent_uuid,omitempty"`
	ProvisioningStartTime            string         `json:"provisioning_start_time,omitempty"`
	PrivateRedhatSubscriptionPasword string         `json:"private_redhat_subscription_pasword,omitempty"`
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
		IDPerms:                          MakeIdPermsType(),
		PrivateOspdUserName:              "",
		PrivateOspdUserPassword:          "",
		PrivateOspdVMDiskGB:              "",
		ProvisioningState:                "",
		PrivateOspdVMRAMMB:               "",
		Perms2:                           MakePermType2(),
		ProvisioningLog:                  "",
		Type:                             "",
		PrivateOspdVMName:                "",
		PrivateRedhatSubscriptionKey:     "",
		GCPAccountInfo:                   "",
		GCPSubnet:                        "",
		Annotations:                      MakeKeyValuePairs(),
		UUID:                             "",
		FQName:                           []string{},
		PrivateDNSServers:                "",
		PrivateRedhatPoolID:              "",
		GCPRegion:                        "",
		ProvisioningProgressStage:        "",
		PrivateOspdPackageURL:            "",
		GCPAsn:                           0,
		AwsSubnet:                        "",
		DisplayName:                      "",
		ParentUUID:                       "",
		ProvisioningStartTime:            "",
		PrivateRedhatSubscriptionPasword: "",
		AwsRegion:                        "",
		AwsAccessKey:                     "",
		AwsSecretKey:                     "",
		ProvisioningProgress:             0,
		PrivateNTPHosts:                  "",
		PrivateRedhatSubscriptionUser:    "",
		PrivateOspdVMVcpus:               "",
		ParentType:                       "",
	}
}

// MakeLocationSlice() makes a slice of Location
func MakeLocationSlice() []*Location {
	return []*Location{}
}
