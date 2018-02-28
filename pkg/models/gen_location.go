package models

import (
	"github.com/Juniper/contrail/pkg/common"
)

//To skip import error.
var _ = common.OPERATION

// MakeLocation makes Location
// nolint
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

// MakeLocation makes Location
// nolint
func InterfaceToLocation(i interface{}) *Location {
	m, ok := i.(map[string]interface{})
	_ = m
	if !ok {
		return nil
	}
	return &Location{
		//TODO(nati): Apply default
		ProvisioningLog:                  common.InterfaceToString(m["provisioning_log"]),
		ProvisioningProgress:             common.InterfaceToInt64(m["provisioning_progress"]),
		ProvisioningProgressStage:        common.InterfaceToString(m["provisioning_progress_stage"]),
		ProvisioningStartTime:            common.InterfaceToString(m["provisioning_start_time"]),
		ProvisioningState:                common.InterfaceToString(m["provisioning_state"]),
		UUID:                             common.InterfaceToString(m["uuid"]),
		ParentUUID:                       common.InterfaceToString(m["parent_uuid"]),
		ParentType:                       common.InterfaceToString(m["parent_type"]),
		FQName:                           common.InterfaceToStringList(m["fq_name"]),
		IDPerms:                          InterfaceToIdPermsType(m["id_perms"]),
		DisplayName:                      common.InterfaceToString(m["display_name"]),
		Annotations:                      InterfaceToKeyValuePairs(m["annotations"]),
		Perms2:                           InterfaceToPermType2(m["perms2"]),
		Type:                             common.InterfaceToString(m["type"]),
		PrivateDNSServers:                common.InterfaceToString(m["private_dns_servers"]),
		PrivateNTPHosts:                  common.InterfaceToString(m["private_ntp_hosts"]),
		PrivateOspdPackageURL:            common.InterfaceToString(m["private_ospd_package_url"]),
		PrivateOspdUserName:              common.InterfaceToString(m["private_ospd_user_name"]),
		PrivateOspdUserPassword:          common.InterfaceToString(m["private_ospd_user_password"]),
		PrivateOspdVMDiskGB:              common.InterfaceToString(m["private_ospd_vm_disk_gb"]),
		PrivateOspdVMName:                common.InterfaceToString(m["private_ospd_vm_name"]),
		PrivateOspdVMRAMMB:               common.InterfaceToString(m["private_ospd_vm_ram_mb"]),
		PrivateOspdVMVcpus:               common.InterfaceToString(m["private_ospd_vm_vcpus"]),
		PrivateRedhatPoolID:              common.InterfaceToString(m["private_redhat_pool_id"]),
		PrivateRedhatSubscriptionKey:     common.InterfaceToString(m["private_redhat_subscription_key"]),
		PrivateRedhatSubscriptionPasword: common.InterfaceToString(m["private_redhat_subscription_pasword"]),
		PrivateRedhatSubscriptionUser:    common.InterfaceToString(m["private_redhat_subscription_user"]),
		GCPAccountInfo:                   common.InterfaceToString(m["gcp_account_info"]),
		GCPAsn:                           common.InterfaceToInt64(m["gcp_asn"]),
		GCPRegion:                        common.InterfaceToString(m["gcp_region"]),
		GCPSubnet:                        common.InterfaceToString(m["gcp_subnet"]),
		AwsAccessKey:                     common.InterfaceToString(m["aws_access_key"]),
		AwsRegion:                        common.InterfaceToString(m["aws_region"]),
		AwsSecretKey:                     common.InterfaceToString(m["aws_secret_key"]),
		AwsSubnet:                        common.InterfaceToString(m["aws_subnet"]),
	}
}

// MakeLocationSlice() makes a slice of Location
// nolint
func MakeLocationSlice() []*Location {
	return []*Location{}
}

// InterfaceToLocationSlice() makes a slice of Location
// nolint
func InterfaceToLocationSlice(i interface{}) []*Location {
	list := common.InterfaceToInterfaceList(i)
	if list == nil {
		return nil
	}
	result := []*Location{}
	for _, item := range list {
		result = append(result, InterfaceToLocation(item))
	}
	return result
}
