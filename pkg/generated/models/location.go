
package models
// Location



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propLocation_private_ntp_hosts int = iota
    propLocation_private_ospd_vm_vcpus int = iota
    propLocation_display_name int = iota
    propLocation_parent_uuid int = iota
    propLocation_private_ospd_vm_disk_gb int = iota
    propLocation_private_redhat_pool_id int = iota
    propLocation_gcp_region int = iota
    propLocation_gcp_subnet int = iota
    propLocation_private_ospd_user_name int = iota
    propLocation_gcp_asn int = iota
    propLocation_uuid int = iota
    propLocation_provisioning_log int = iota
    propLocation_provisioning_start_time int = iota
    propLocation_provisioning_state int = iota
    propLocation_private_dns_servers int = iota
    propLocation_private_redhat_subscription_user int = iota
    propLocation_type int = iota
    propLocation_private_ospd_package_url int = iota
    propLocation_private_ospd_vm_name int = iota
    propLocation_private_redhat_subscription_key int = iota
    propLocation_annotations int = iota
    propLocation_provisioning_progress_stage int = iota
    propLocation_aws_secret_key int = iota
    propLocation_id_perms int = iota
    propLocation_parent_type int = iota
    propLocation_provisioning_progress int = iota
    propLocation_private_ospd_vm_ram_mb int = iota
    propLocation_gcp_account_info int = iota
    propLocation_aws_access_key int = iota
    propLocation_fq_name int = iota
    propLocation_perms2 int = iota
    propLocation_private_ospd_user_password int = iota
    propLocation_private_redhat_subscription_pasword int = iota
    propLocation_aws_region int = iota
    propLocation_aws_subnet int = iota
)

// Location 
type Location struct {

    PrivateDNSServers string `json:"private_dns_servers,omitempty"`
    PrivateRedhatSubscriptionUser string `json:"private_redhat_subscription_user,omitempty"`
    PrivateRedhatSubscriptionKey string `json:"private_redhat_subscription_key,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    ProvisioningProgressStage string `json:"provisioning_progress_stage,omitempty"`
    Type string `json:"type,omitempty"`
    PrivateOspdPackageURL string `json:"private_ospd_package_url,omitempty"`
    PrivateOspdVMName string `json:"private_ospd_vm_name,omitempty"`
    ProvisioningProgress int `json:"provisioning_progress,omitempty"`
    AwsSecretKey string `json:"aws_secret_key,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    PrivateOspdVMRAMMB string `json:"private_ospd_vm_ram_mb,omitempty"`
    GCPAccountInfo string `json:"gcp_account_info,omitempty"`
    AwsAccessKey string `json:"aws_access_key,omitempty"`
    AwsSubnet string `json:"aws_subnet,omitempty"`
    PrivateOspdUserPassword string `json:"private_ospd_user_password,omitempty"`
    PrivateRedhatSubscriptionPasword string `json:"private_redhat_subscription_pasword,omitempty"`
    AwsRegion string `json:"aws_region,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    PrivateNTPHosts string `json:"private_ntp_hosts,omitempty"`
    PrivateOspdVMVcpus string `json:"private_ospd_vm_vcpus,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    GCPSubnet string `json:"gcp_subnet,omitempty"`
    PrivateOspdVMDiskGB string `json:"private_ospd_vm_disk_gb,omitempty"`
    PrivateRedhatPoolID string `json:"private_redhat_pool_id,omitempty"`
    GCPRegion string `json:"gcp_region,omitempty"`
    ProvisioningLog string `json:"provisioning_log,omitempty"`
    ProvisioningStartTime string `json:"provisioning_start_time,omitempty"`
    ProvisioningState string `json:"provisioning_state,omitempty"`
    PrivateOspdUserName string `json:"private_ospd_user_name,omitempty"`
    GCPAsn int `json:"gcp_asn,omitempty"`
    UUID string `json:"uuid,omitempty"`


    PhysicalRouters []*PhysicalRouter `json:"physical_routers,omitempty"`

    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *Location) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeLocation makes Location
func MakeLocation() *Location{
    return &Location{
    //TODO(nati): Apply default
    PrivateRedhatSubscriptionUser: "",
        PrivateDNSServers: "",
        PrivateOspdPackageURL: "",
        PrivateOspdVMName: "",
        PrivateRedhatSubscriptionKey: "",
        Annotations: MakeKeyValuePairs(),
        ProvisioningProgressStage: "",
        Type: "",
        IDPerms: MakeIdPermsType(),
        ParentType: "",
        ProvisioningProgress: 0,
        AwsSecretKey: "",
        GCPAccountInfo: "",
        AwsAccessKey: "",
        FQName: []string{},
        Perms2: MakePermType2(),
        PrivateOspdVMRAMMB: "",
        PrivateRedhatSubscriptionPasword: "",
        AwsRegion: "",
        AwsSubnet: "",
        PrivateOspdUserPassword: "",
        PrivateOspdVMVcpus: "",
        DisplayName: "",
        ParentUUID: "",
        PrivateNTPHosts: "",
        PrivateRedhatPoolID: "",
        GCPRegion: "",
        GCPSubnet: "",
        PrivateOspdVMDiskGB: "",
        GCPAsn: 0,
        UUID: "",
        ProvisioningLog: "",
        ProvisioningStartTime: "",
        ProvisioningState: "",
        PrivateOspdUserName: "",
        
        modified: big.NewInt(0),
    }
}



// MakeLocationSlice makes a slice of Location
func MakeLocationSlice() []*Location {
    return []*Location{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *Location) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *Location) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *Location) GetDefaultName() string {
    return strings.Replace("default-location", "_", "-", -1)
}

func (model *Location) GetType() string {
    return strings.Replace("location", "_", "-", -1)
}

func (model *Location) GetFQName() []string {
    return model.FQName
}

func (model *Location) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *Location) GetParentType() string {
    return model.ParentType
}

func (model *Location) GetUuid() string {
    return model.UUID
}

func (model *Location) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *Location) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *Location) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *Location) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *Location) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propLocation_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Type); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Type as type")
        }
        msg["type"] = &val
    }
    
    if model.modified.Bit(propLocation_private_ospd_package_url) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PrivateOspdPackageURL); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PrivateOspdPackageURL as private_ospd_package_url")
        }
        msg["private_ospd_package_url"] = &val
    }
    
    if model.modified.Bit(propLocation_private_ospd_vm_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PrivateOspdVMName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PrivateOspdVMName as private_ospd_vm_name")
        }
        msg["private_ospd_vm_name"] = &val
    }
    
    if model.modified.Bit(propLocation_private_redhat_subscription_key) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PrivateRedhatSubscriptionKey); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PrivateRedhatSubscriptionKey as private_redhat_subscription_key")
        }
        msg["private_redhat_subscription_key"] = &val
    }
    
    if model.modified.Bit(propLocation_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propLocation_provisioning_progress_stage) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgressStage); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgressStage as provisioning_progress_stage")
        }
        msg["provisioning_progress_stage"] = &val
    }
    
    if model.modified.Bit(propLocation_aws_secret_key) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AwsSecretKey); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AwsSecretKey as aws_secret_key")
        }
        msg["aws_secret_key"] = &val
    }
    
    if model.modified.Bit(propLocation_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propLocation_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propLocation_provisioning_progress) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningProgress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningProgress as provisioning_progress")
        }
        msg["provisioning_progress"] = &val
    }
    
    if model.modified.Bit(propLocation_private_ospd_vm_ram_mb) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PrivateOspdVMRAMMB); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PrivateOspdVMRAMMB as private_ospd_vm_ram_mb")
        }
        msg["private_ospd_vm_ram_mb"] = &val
    }
    
    if model.modified.Bit(propLocation_gcp_account_info) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.GCPAccountInfo); err != nil {
            return nil, errors.Wrap(err, "Marshal of: GCPAccountInfo as gcp_account_info")
        }
        msg["gcp_account_info"] = &val
    }
    
    if model.modified.Bit(propLocation_aws_access_key) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AwsAccessKey); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AwsAccessKey as aws_access_key")
        }
        msg["aws_access_key"] = &val
    }
    
    if model.modified.Bit(propLocation_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propLocation_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propLocation_private_ospd_user_password) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PrivateOspdUserPassword); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PrivateOspdUserPassword as private_ospd_user_password")
        }
        msg["private_ospd_user_password"] = &val
    }
    
    if model.modified.Bit(propLocation_private_redhat_subscription_pasword) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PrivateRedhatSubscriptionPasword); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PrivateRedhatSubscriptionPasword as private_redhat_subscription_pasword")
        }
        msg["private_redhat_subscription_pasword"] = &val
    }
    
    if model.modified.Bit(propLocation_aws_region) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AwsRegion); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AwsRegion as aws_region")
        }
        msg["aws_region"] = &val
    }
    
    if model.modified.Bit(propLocation_aws_subnet) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AwsSubnet); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AwsSubnet as aws_subnet")
        }
        msg["aws_subnet"] = &val
    }
    
    if model.modified.Bit(propLocation_private_ntp_hosts) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PrivateNTPHosts); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PrivateNTPHosts as private_ntp_hosts")
        }
        msg["private_ntp_hosts"] = &val
    }
    
    if model.modified.Bit(propLocation_private_ospd_vm_vcpus) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PrivateOspdVMVcpus); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PrivateOspdVMVcpus as private_ospd_vm_vcpus")
        }
        msg["private_ospd_vm_vcpus"] = &val
    }
    
    if model.modified.Bit(propLocation_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propLocation_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propLocation_private_ospd_vm_disk_gb) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PrivateOspdVMDiskGB); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PrivateOspdVMDiskGB as private_ospd_vm_disk_gb")
        }
        msg["private_ospd_vm_disk_gb"] = &val
    }
    
    if model.modified.Bit(propLocation_private_redhat_pool_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PrivateRedhatPoolID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PrivateRedhatPoolID as private_redhat_pool_id")
        }
        msg["private_redhat_pool_id"] = &val
    }
    
    if model.modified.Bit(propLocation_gcp_region) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.GCPRegion); err != nil {
            return nil, errors.Wrap(err, "Marshal of: GCPRegion as gcp_region")
        }
        msg["gcp_region"] = &val
    }
    
    if model.modified.Bit(propLocation_gcp_subnet) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.GCPSubnet); err != nil {
            return nil, errors.Wrap(err, "Marshal of: GCPSubnet as gcp_subnet")
        }
        msg["gcp_subnet"] = &val
    }
    
    if model.modified.Bit(propLocation_private_ospd_user_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PrivateOspdUserName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PrivateOspdUserName as private_ospd_user_name")
        }
        msg["private_ospd_user_name"] = &val
    }
    
    if model.modified.Bit(propLocation_gcp_asn) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.GCPAsn); err != nil {
            return nil, errors.Wrap(err, "Marshal of: GCPAsn as gcp_asn")
        }
        msg["gcp_asn"] = &val
    }
    
    if model.modified.Bit(propLocation_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propLocation_provisioning_log) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningLog); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningLog as provisioning_log")
        }
        msg["provisioning_log"] = &val
    }
    
    if model.modified.Bit(propLocation_provisioning_start_time) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningStartTime); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningStartTime as provisioning_start_time")
        }
        msg["provisioning_start_time"] = &val
    }
    
    if model.modified.Bit(propLocation_provisioning_state) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProvisioningState); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProvisioningState as provisioning_state")
        }
        msg["provisioning_state"] = &val
    }
    
    if model.modified.Bit(propLocation_private_dns_servers) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PrivateDNSServers); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PrivateDNSServers as private_dns_servers")
        }
        msg["private_dns_servers"] = &val
    }
    
    if model.modified.Bit(propLocation_private_redhat_subscription_user) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PrivateRedhatSubscriptionUser); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PrivateRedhatSubscriptionUser as private_redhat_subscription_user")
        }
        msg["private_redhat_subscription_user"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *Location) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *Location) UpdateReferences() error {
    return nil
}


