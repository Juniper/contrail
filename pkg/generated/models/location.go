package models

// Location

import "encoding/json"

// Location
type Location struct {
	PrivateRedhatPoolID              string         `json:"private_redhat_pool_id,omitempty"`
	PrivateRedhatSubscriptionPasword string         `json:"private_redhat_subscription_pasword,omitempty"`
	GCPAsn                           int            `json:"gcp_asn,omitempty"`
	DisplayName                      string         `json:"display_name,omitempty"`
	Annotations                      *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                           *PermType2     `json:"perms2,omitempty"`
	AwsAccessKey                     string         `json:"aws_access_key,omitempty"`
	ParentType                       string         `json:"parent_type,omitempty"`
	PrivateOspdVMRAMMB               string         `json:"private_ospd_vm_ram_mb,omitempty"`
	GCPSubnet                        string         `json:"gcp_subnet,omitempty"`
	IDPerms                          *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningStartTime            string         `json:"provisioning_start_time,omitempty"`
	PrivateOspdPackageURL            string         `json:"private_ospd_package_url,omitempty"`
	PrivateOspdUserPassword          string         `json:"private_ospd_user_password,omitempty"`
	PrivateOspdVMVcpus               string         `json:"private_ospd_vm_vcpus,omitempty"`
	PrivateRedhatSubscriptionUser    string         `json:"private_redhat_subscription_user,omitempty"`
	GCPAccountInfo                   string         `json:"gcp_account_info,omitempty"`
	GCPRegion                        string         `json:"gcp_region,omitempty"`
	UUID                             string         `json:"uuid,omitempty"`
	PrivateDNSServers                string         `json:"private_dns_servers,omitempty"`
	PrivateNTPHosts                  string         `json:"private_ntp_hosts,omitempty"`
	PrivateRedhatSubscriptionKey     string         `json:"private_redhat_subscription_key,omitempty"`
	AwsRegion                        string         `json:"aws_region,omitempty"`
	AwsSubnet                        string         `json:"aws_subnet,omitempty"`
	ParentUUID                       string         `json:"parent_uuid,omitempty"`
	Type                             string         `json:"type,omitempty"`
	PrivateOspdUserName              string         `json:"private_ospd_user_name,omitempty"`
	ProvisioningState                string         `json:"provisioning_state,omitempty"`
	ProvisioningLog                  string         `json:"provisioning_log,omitempty"`
	PrivateOspdVMName                string         `json:"private_ospd_vm_name,omitempty"`
	AwsSecretKey                     string         `json:"aws_secret_key,omitempty"`
	PrivateOspdVMDiskGB              string         `json:"private_ospd_vm_disk_gb,omitempty"`
	FQName                           []string       `json:"fq_name,omitempty"`
	ProvisioningProgress             int            `json:"provisioning_progress,omitempty"`
	ProvisioningProgressStage        string         `json:"provisioning_progress_stage,omitempty"`

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
		ProvisioningProgress:             0,
		ProvisioningProgressStage:        "",
		PrivateOspdVMDiskGB:              "",
		FQName:                           []string{},
		GCPAsn:                           0,
		DisplayName:                      "",
		Annotations:                      MakeKeyValuePairs(),
		Perms2:                           MakePermType2(),
		PrivateRedhatPoolID:              "",
		PrivateRedhatSubscriptionPasword: "",
		AwsAccessKey:                     "",
		ParentType:                       "",
		IDPerms:                          MakeIdPermsType(),
		PrivateOspdVMRAMMB:               "",
		GCPSubnet:                        "",
		PrivateOspdVMVcpus:               "",
		PrivateRedhatSubscriptionUser:    "",
		GCPAccountInfo:                   "",
		GCPRegion:                        "",
		ProvisioningStartTime:            "",
		PrivateOspdPackageURL:            "",
		PrivateOspdUserPassword:          "",
		PrivateRedhatSubscriptionKey:     "",
		AwsRegion:                        "",
		AwsSubnet:                        "",
		ParentUUID:                       "",
		UUID:                             "",
		PrivateDNSServers:                "",
		PrivateNTPHosts:                  "",
		ProvisioningState:                "",
		ProvisioningLog:                  "",
		Type:                             "",
		PrivateOspdUserName:              "",
		PrivateOspdVMName:                "",
		AwsSecretKey:                     "",
	}
}

// MakeLocationSlice() makes a slice of Location
func MakeLocationSlice() []*Location {
	return []*Location{}
}
