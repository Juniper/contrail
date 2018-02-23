package models

import (
    "github.com/Juniper/contrail/pkg/schema"
)

//To skip import error.
var _ = schema.Version


// MakeLocation makes Location
func MakeLocation() *Location{
    return &Location{
    //TODO(nati): Apply default
    ProvisioningLog: "",
        ProvisioningProgress: 0,
        ProvisioningProgressStage: "",
        ProvisioningStartTime: "",
        ProvisioningState: "",
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        Type: "",
        PrivateDNSServers: "",
        PrivateNTPHosts: "",
        PrivateOspdPackageURL: "",
        PrivateOspdUserName: "",
        PrivateOspdUserPassword: "",
        PrivateOspdVMDiskGB: "",
        PrivateOspdVMName: "",
        PrivateOspdVMRAMMB: "",
        PrivateOspdVMVcpus: "",
        PrivateRedhatPoolID: "",
        PrivateRedhatSubscriptionKey: "",
        PrivateRedhatSubscriptionPasword: "",
        PrivateRedhatSubscriptionUser: "",
        GCPAccountInfo: "",
        GCPAsn: 0,
        GCPRegion: "",
        GCPSubnet: "",
        AwsAccessKey: "",
        AwsRegion: "",
        AwsSecretKey: "",
        AwsSubnet: "",
        
    }
}

// MakeLocation makes Location
func InterfaceToLocation(i interface{}) *Location{
    m, ok := i.(map[string]interface{})
    _ = m
    if !ok {
        return nil 
    } 
    return &Location{
    //TODO(nati): Apply default
    ProvisioningLog: schema.InterfaceToString(m["provisioning_log"]),
        ProvisioningProgress: schema.InterfaceToInt64(m["provisioning_progress"]),
        ProvisioningProgressStage: schema.InterfaceToString(m["provisioning_progress_stage"]),
        ProvisioningStartTime: schema.InterfaceToString(m["provisioning_start_time"]),
        ProvisioningState: schema.InterfaceToString(m["provisioning_state"]),
        UUID: schema.InterfaceToString(m["uuid"]),
        ParentUUID: schema.InterfaceToString(m["parent_uuid"]),
        ParentType: schema.InterfaceToString(m["parent_type"]),
        FQName: schema.InterfaceToStringList(m["fq_name"]),
        IDPerms: InterfaceToIdPermsType(m["id_perms"]),
        DisplayName: schema.InterfaceToString(m["display_name"]),
        Annotations: InterfaceToKeyValuePairs(m["annotations"]),
        Perms2: InterfaceToPermType2(m["perms2"]),
        Type: schema.InterfaceToString(m["type"]),
        PrivateDNSServers: schema.InterfaceToString(m["private_dns_servers"]),
        PrivateNTPHosts: schema.InterfaceToString(m["private_ntp_hosts"]),
        PrivateOspdPackageURL: schema.InterfaceToString(m["private_ospd_package_url"]),
        PrivateOspdUserName: schema.InterfaceToString(m["private_ospd_user_name"]),
        PrivateOspdUserPassword: schema.InterfaceToString(m["private_ospd_user_password"]),
        PrivateOspdVMDiskGB: schema.InterfaceToString(m["private_ospd_vm_disk_gb"]),
        PrivateOspdVMName: schema.InterfaceToString(m["private_ospd_vm_name"]),
        PrivateOspdVMRAMMB: schema.InterfaceToString(m["private_ospd_vm_ram_mb"]),
        PrivateOspdVMVcpus: schema.InterfaceToString(m["private_ospd_vm_vcpus"]),
        PrivateRedhatPoolID: schema.InterfaceToString(m["private_redhat_pool_id"]),
        PrivateRedhatSubscriptionKey: schema.InterfaceToString(m["private_redhat_subscription_key"]),
        PrivateRedhatSubscriptionPasword: schema.InterfaceToString(m["private_redhat_subscription_pasword"]),
        PrivateRedhatSubscriptionUser: schema.InterfaceToString(m["private_redhat_subscription_user"]),
        GCPAccountInfo: schema.InterfaceToString(m["gcp_account_info"]),
        GCPAsn: schema.InterfaceToInt64(m["gcp_asn"]),
        GCPRegion: schema.InterfaceToString(m["gcp_region"]),
        GCPSubnet: schema.InterfaceToString(m["gcp_subnet"]),
        AwsAccessKey: schema.InterfaceToString(m["aws_access_key"]),
        AwsRegion: schema.InterfaceToString(m["aws_region"]),
        AwsSecretKey: schema.InterfaceToString(m["aws_secret_key"]),
        AwsSubnet: schema.InterfaceToString(m["aws_subnet"]),
        
    }
}

// MakeLocationSlice() makes a slice of Location
func MakeLocationSlice() []*Location {
    return []*Location{}
}

// InterfaceToLocationSlice() makes a slice of Location
func InterfaceToLocationSlice(i interface{}) []*Location {
    list := schema.InterfaceToInterfaceList(i)
    if list == nil {
        return nil
    }
    result := []*Location{}
    for _, item := range list {
        result = append(result, InterfaceToLocation(item) )
    }
    return result
}



