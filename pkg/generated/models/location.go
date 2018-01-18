package models

// Location

import "encoding/json"

// Location
type Location struct {
	GCPAccountInfo                   string         `json:"gcp_account_info,omitempty"`
	ParentUUID                       string         `json:"parent_uuid,omitempty"`
	ParentType                       string         `json:"parent_type,omitempty"`
	ProvisioningStartTime            string         `json:"provisioning_start_time,omitempty"`
	PrivateOspdUserPassword          string         `json:"private_ospd_user_password,omitempty"`
	PrivateOspdVMVcpus               string         `json:"private_ospd_vm_vcpus,omitempty"`
	GCPSubnet                        string         `json:"gcp_subnet,omitempty"`
	UUID                             string         `json:"uuid,omitempty"`
	AwsSecretKey                     string         `json:"aws_secret_key,omitempty"`
	ProvisioningProgressStage        string         `json:"provisioning_progress_stage,omitempty"`
	Type                             string         `json:"type,omitempty"`
	PrivateDNSServers                string         `json:"private_dns_servers,omitempty"`
	PrivateOspdVMName                string         `json:"private_ospd_vm_name,omitempty"`
	PrivateOspdVMRAMMB               string         `json:"private_ospd_vm_ram_mb,omitempty"`
	PrivateRedhatPoolID              string         `json:"private_redhat_pool_id,omitempty"`
	PrivateRedhatSubscriptionUser    string         `json:"private_redhat_subscription_user,omitempty"`
	PrivateRedhatSubscriptionPasword string         `json:"private_redhat_subscription_pasword,omitempty"`
	AwsAccessKey                     string         `json:"aws_access_key,omitempty"`
	AwsSubnet                        string         `json:"aws_subnet,omitempty"`
	FQName                           []string       `json:"fq_name,omitempty"`
	DisplayName                      string         `json:"display_name,omitempty"`
	PrivateNTPHosts                  string         `json:"private_ntp_hosts,omitempty"`
	PrivateOspdUserName              string         `json:"private_ospd_user_name,omitempty"`
	GCPAsn                           int            `json:"gcp_asn,omitempty"`
	Perms2                           *PermType2     `json:"perms2,omitempty"`
	ProvisioningLog                  string         `json:"provisioning_log,omitempty"`
	PrivateOspdPackageURL            string         `json:"private_ospd_package_url,omitempty"`
	GCPRegion                        string         `json:"gcp_region,omitempty"`
	IDPerms                          *IdPermsType   `json:"id_perms,omitempty"`
	ProvisioningProgress             int            `json:"provisioning_progress,omitempty"`
	PrivateOspdVMDiskGB              string         `json:"private_ospd_vm_disk_gb,omitempty"`
	PrivateRedhatSubscriptionKey     string         `json:"private_redhat_subscription_key,omitempty"`
	AwsRegion                        string         `json:"aws_region,omitempty"`
	Annotations                      *KeyValuePairs `json:"annotations,omitempty"`
	ProvisioningState                string         `json:"provisioning_state,omitempty"`

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
		PrivateRedhatSubscriptionKey: "",
		AwsRegion:                    "",
		Annotations:                  MakeKeyValuePairs(),
		ProvisioningState:            "",
		PrivateOspdVMDiskGB:          "",
		ParentUUID:                   "",
		ParentType:                   "",
		ProvisioningStartTime:        "",
		GCPAccountInfo:               "",
		PrivateOspdVMVcpus:           "",
		GCPSubnet:                    "",
		UUID:                         "",
		PrivateOspdUserPassword:       "",
		PrivateDNSServers:             "",
		PrivateOspdVMName:             "",
		PrivateOspdVMRAMMB:            "",
		PrivateRedhatPoolID:           "",
		PrivateRedhatSubscriptionUser: "",
		AwsSecretKey:                  "",
		ProvisioningProgressStage:     "",
		Type:                             "",
		AwsAccessKey:                     "",
		PrivateRedhatSubscriptionPasword: "",
		FQName:                []string{},
		DisplayName:           "",
		AwsSubnet:             "",
		PrivateOspdUserName:   "",
		GCPAsn:                0,
		Perms2:                MakePermType2(),
		ProvisioningLog:       "",
		PrivateNTPHosts:       "",
		GCPRegion:             "",
		IDPerms:               MakeIdPermsType(),
		ProvisioningProgress:  0,
		PrivateOspdPackageURL: "",
	}
}

// MakeLocationSlice() makes a slice of Location
func MakeLocationSlice() []*Location {
	return []*Location{}
}
