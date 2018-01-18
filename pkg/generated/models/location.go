package models

// Location

import "encoding/json"

// Location
type Location struct {
	AwsSubnet                        string         `json:"aws_subnet,omitempty"`
	IDPerms                          *IdPermsType   `json:"id_perms,omitempty"`
	Annotations                      *KeyValuePairs `json:"annotations,omitempty"`
	ParentUUID                       string         `json:"parent_uuid,omitempty"`
	PrivateOspdVMVcpus               string         `json:"private_ospd_vm_vcpus,omitempty"`
	AwsSecretKey                     string         `json:"aws_secret_key,omitempty"`
	Perms2                           *PermType2     `json:"perms2,omitempty"`
	PrivateDNSServers                string         `json:"private_dns_servers,omitempty"`
	DisplayName                      string         `json:"display_name,omitempty"`
	UUID                             string         `json:"uuid,omitempty"`
	ProvisioningProgress             int            `json:"provisioning_progress,omitempty"`
	ProvisioningProgressStage        string         `json:"provisioning_progress_stage,omitempty"`
	PrivateRedhatSubscriptionKey     string         `json:"private_redhat_subscription_key,omitempty"`
	GCPSubnet                        string         `json:"gcp_subnet,omitempty"`
	FQName                           []string       `json:"fq_name,omitempty"`
	ProvisioningState                string         `json:"provisioning_state,omitempty"`
	PrivateOspdUserPassword          string         `json:"private_ospd_user_password,omitempty"`
	PrivateOspdVMName                string         `json:"private_ospd_vm_name,omitempty"`
	PrivateRedhatSubscriptionUser    string         `json:"private_redhat_subscription_user,omitempty"`
	GCPAsn                           int            `json:"gcp_asn,omitempty"`
	GCPRegion                        string         `json:"gcp_region,omitempty"`
	PrivateOspdUserName              string         `json:"private_ospd_user_name,omitempty"`
	GCPAccountInfo                   string         `json:"gcp_account_info,omitempty"`
	AwsRegion                        string         `json:"aws_region,omitempty"`
	ProvisioningLog                  string         `json:"provisioning_log,omitempty"`
	ProvisioningStartTime            string         `json:"provisioning_start_time,omitempty"`
	PrivateRedhatPoolID              string         `json:"private_redhat_pool_id,omitempty"`
	PrivateRedhatSubscriptionPasword string         `json:"private_redhat_subscription_pasword,omitempty"`
	AwsAccessKey                     string         `json:"aws_access_key,omitempty"`
	Type                             string         `json:"type,omitempty"`
	PrivateNTPHosts                  string         `json:"private_ntp_hosts,omitempty"`
	PrivateOspdPackageURL            string         `json:"private_ospd_package_url,omitempty"`
	PrivateOspdVMDiskGB              string         `json:"private_ospd_vm_disk_gb,omitempty"`
	PrivateOspdVMRAMMB               string         `json:"private_ospd_vm_ram_mb,omitempty"`
	ParentType                       string         `json:"parent_type,omitempty"`

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
		PrivateOspdVMVcpus: "",
		AwsSecretKey:       "",
		Perms2:             MakePermType2(),
		ProvisioningProgressStage:     "",
		PrivateDNSServers:             "",
		DisplayName:                   "",
		UUID:                          "",
		ProvisioningProgress:          0,
		PrivateRedhatSubscriptionKey:  "",
		GCPSubnet:                     "",
		FQName:                        []string{},
		GCPRegion:                     "",
		ProvisioningState:             "",
		PrivateOspdUserPassword:       "",
		PrivateOspdVMName:             "",
		PrivateRedhatSubscriptionUser: "",
		GCPAsn:                           0,
		PrivateOspdUserName:              "",
		GCPAccountInfo:                   "",
		AwsRegion:                        "",
		ProvisioningLog:                  "",
		ProvisioningStartTime:            "",
		PrivateOspdVMRAMMB:               "",
		PrivateRedhatPoolID:              "",
		PrivateRedhatSubscriptionPasword: "",
		AwsAccessKey:                     "",
		Type:                             "",
		PrivateNTPHosts:                  "",
		PrivateOspdPackageURL:            "",
		PrivateOspdVMDiskGB:              "",
		ParentType:                       "",
		AwsSubnet:                        "",
		IDPerms:                          MakeIdPermsType(),
		Annotations:                      MakeKeyValuePairs(),
		ParentUUID:                       "",
	}
}

// MakeLocationSlice() makes a slice of Location
func MakeLocationSlice() []*Location {
	return []*Location{}
}
