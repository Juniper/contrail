package models

// Location

// Location
//proteus:generate
type Location struct {
	ProvisioningLog                  string         `json:"provisioning_log,omitempty"`
	ProvisioningProgress             int            `json:"provisioning_progress,omitempty"`
	ProvisioningProgressStage        string         `json:"provisioning_progress_stage,omitempty"`
	ProvisioningStartTime            string         `json:"provisioning_start_time,omitempty"`
	ProvisioningState                string         `json:"provisioning_state,omitempty"`
	UUID                             string         `json:"uuid,omitempty"`
	ParentUUID                       string         `json:"parent_uuid,omitempty"`
	ParentType                       string         `json:"parent_type,omitempty"`
	FQName                           []string       `json:"fq_name,omitempty"`
	IDPerms                          *IdPermsType   `json:"id_perms,omitempty"`
	DisplayName                      string         `json:"display_name,omitempty"`
	Annotations                      *KeyValuePairs `json:"annotations,omitempty"`
	Perms2                           *PermType2     `json:"perms2,omitempty"`
	Type                             string         `json:"type,omitempty"`
	PrivateDNSServers                string         `json:"private_dns_servers,omitempty"`
	PrivateNTPHosts                  string         `json:"private_ntp_hosts,omitempty"`
	PrivateOspdPackageURL            string         `json:"private_ospd_package_url,omitempty"`
	PrivateOspdUserName              string         `json:"private_ospd_user_name,omitempty"`
	PrivateOspdUserPassword          string         `json:"private_ospd_user_password,omitempty"`
	PrivateOspdVMDiskGB              string         `json:"private_ospd_vm_disk_gb,omitempty"`
	PrivateOspdVMName                string         `json:"private_ospd_vm_name,omitempty"`
	PrivateOspdVMRAMMB               string         `json:"private_ospd_vm_ram_mb,omitempty"`
	PrivateOspdVMVcpus               string         `json:"private_ospd_vm_vcpus,omitempty"`
	PrivateRedhatPoolID              string         `json:"private_redhat_pool_id,omitempty"`
	PrivateRedhatSubscriptionKey     string         `json:"private_redhat_subscription_key,omitempty"`
	PrivateRedhatSubscriptionPasword string         `json:"private_redhat_subscription_pasword,omitempty"`
	PrivateRedhatSubscriptionUser    string         `json:"private_redhat_subscription_user,omitempty"`
	GCPAccountInfo                   string         `json:"gcp_account_info,omitempty"`
	GCPAsn                           int            `json:"gcp_asn,omitempty"`
	GCPRegion                        string         `json:"gcp_region,omitempty"`
	GCPSubnet                        string         `json:"gcp_subnet,omitempty"`
	AwsAccessKey                     string         `json:"aws_access_key,omitempty"`
	AwsRegion                        string         `json:"aws_region,omitempty"`
	AwsSecretKey                     string         `json:"aws_secret_key,omitempty"`
	AwsSubnet                        string         `json:"aws_subnet,omitempty"`

	PhysicalRouters []*PhysicalRouter `json:"physical_routers,omitempty"`
}

// MakeLocation makes Location
func MakeLocation() *Location {
	return &Location{
		//TODO(nati): Apply default
		ProvisioningLog:                  "",
		ProvisioningProgress:             0,
		ProvisioningProgressStage:        "",
		ProvisioningStartTime:            "",
		ProvisioningState:                "",
		UUID:                             "",
		ParentUUID:                       "",
		ParentType:                       "",
		FQName:                           []string{},
		IDPerms:                          MakeIdPermsType(),
		DisplayName:                      "",
		Annotations:                      MakeKeyValuePairs(),
		Perms2:                           MakePermType2(),
		Type:                             "",
		PrivateDNSServers:                "",
		PrivateNTPHosts:                  "",
		PrivateOspdPackageURL:            "",
		PrivateOspdUserName:              "",
		PrivateOspdUserPassword:          "",
		PrivateOspdVMDiskGB:              "",
		PrivateOspdVMName:                "",
		PrivateOspdVMRAMMB:               "",
		PrivateOspdVMVcpus:               "",
		PrivateRedhatPoolID:              "",
		PrivateRedhatSubscriptionKey:     "",
		PrivateRedhatSubscriptionPasword: "",
		PrivateRedhatSubscriptionUser:    "",
		GCPAccountInfo:                   "",
		GCPAsn:                           0,
		GCPRegion:                        "",
		GCPSubnet:                        "",
		AwsAccessKey:                     "",
		AwsRegion:                        "",
		AwsSecretKey:                     "",
		AwsSubnet:                        "",
	}
}

// MakeLocationSlice() makes a slice of Location
func MakeLocationSlice() []*Location {
	return []*Location{}
}
