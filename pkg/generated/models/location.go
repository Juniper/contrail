package models

// Location

import "encoding/json"

// Location
type Location struct {
	Perms2                           *PermType2     `json:"perms2,omitempty"`
	ProvisioningLog                  string         `json:"provisioning_log,omitempty"`
	AwsSecretKey                     string         `json:"aws_secret_key,omitempty"`
	PrivateOspdUserName              string         `json:"private_ospd_user_name,omitempty"`
	PrivateRedhatSubscriptionUser    string         `json:"private_redhat_subscription_user,omitempty"`
	AwsAccessKey                     string         `json:"aws_access_key,omitempty"`
	FQName                           []string       `json:"fq_name,omitempty"`
	ProvisioningProgressStage        string         `json:"provisioning_progress_stage,omitempty"`
	PrivateOspdPackageURL            string         `json:"private_ospd_package_url,omitempty"`
	PrivateOspdVMName                string         `json:"private_ospd_vm_name,omitempty"`
	PrivateOspdVMRAMMB               string         `json:"private_ospd_vm_ram_mb,omitempty"`
	PrivateRedhatPoolID              string         `json:"private_redhat_pool_id,omitempty"`
	PrivateRedhatSubscriptionKey     string         `json:"private_redhat_subscription_key,omitempty"`
	GCPAccountInfo                   string         `json:"gcp_account_info,omitempty"`
	GCPSubnet                        string         `json:"gcp_subnet,omitempty"`
	ProvisioningState                string         `json:"provisioning_state,omitempty"`
	PrivateDNSServers                string         `json:"private_dns_servers,omitempty"`
	ProvisioningProgress             int            `json:"provisioning_progress,omitempty"`
	GCPAsn                           int            `json:"gcp_asn,omitempty"`
	AwsRegion                        string         `json:"aws_region,omitempty"`
	DisplayName                      string         `json:"display_name,omitempty"`
	UUID                             string         `json:"uuid,omitempty"`
	PrivateNTPHosts                  string         `json:"private_ntp_hosts,omitempty"`
	GCPRegion                        string         `json:"gcp_region,omitempty"`
	Type                             string         `json:"type,omitempty"`
	PrivateRedhatSubscriptionPasword string         `json:"private_redhat_subscription_pasword,omitempty"`
	AwsSubnet                        string         `json:"aws_subnet,omitempty"`
	Annotations                      *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningStartTime            string         `json:"provisioning_start_time,omitempty"`
	PrivateOspdVMDiskGB              string         `json:"private_ospd_vm_disk_gb,omitempty"`
	PrivateOspdVMVcpus               string         `json:"private_ospd_vm_vcpus,omitempty"`
	ParentUUID                       string         `json:"parent_uuid,omitempty"`
	ParentType                       string         `json:"parent_type,omitempty"`
	IDPerms                          *IdPermsType   `json:"id_perms,omitempty"`
	PrivateOspdUserPassword          string         `json:"private_ospd_user_password,omitempty"`

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
		PrivateNTPHosts:                  "",
		GCPAsn:                           0,
		AwsRegion:                        "",
		DisplayName:                      "",
		UUID:                             "",
		Type:                             "",
		GCPRegion:                        "",
		PrivateOspdVMDiskGB:              "",
		PrivateRedhatSubscriptionPasword: "",
		AwsSubnet:                        "",
		Annotations:                      MakeKeyValuePairs(),
		ProvisioningStartTime:            "",
		PrivateOspdUserPassword:          "",
		PrivateOspdVMVcpus:               "",
		ParentUUID:                       "",
		ParentType:                       "",
		IDPerms:                          MakeIdPermsType(),
		AwsSecretKey:                     "",
		Perms2:                           MakePermType2(),
		ProvisioningLog:                  "",
		PrivateOspdPackageURL:            "",
		PrivateOspdUserName:              "",
		PrivateRedhatSubscriptionUser:    "",
		AwsAccessKey:                     "",
		FQName:                           []string{},
		ProvisioningProgressStage:    "",
		ProvisioningState:            "",
		PrivateDNSServers:            "",
		PrivateOspdVMName:            "",
		PrivateOspdVMRAMMB:           "",
		PrivateRedhatPoolID:          "",
		PrivateRedhatSubscriptionKey: "",
		GCPAccountInfo:               "",
		GCPSubnet:                    "",
		ProvisioningProgress:         0,
	}
}

// MakeLocationSlice() makes a slice of Location
func MakeLocationSlice() []*Location {
	return []*Location{}
}
