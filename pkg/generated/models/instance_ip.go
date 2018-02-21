
package models
// InstanceIP



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propInstanceIP_instance_ip_address int = iota
    propInstanceIP_fq_name int = iota
    propInstanceIP_id_perms int = iota
    propInstanceIP_annotations int = iota
    propInstanceIP_service_health_check_ip int = iota
    propInstanceIP_secondary_ip_tracking_ip int = iota
    propInstanceIP_subnet_uuid int = iota
    propInstanceIP_instance_ip_secondary int = iota
    propInstanceIP_parent_type int = iota
    propInstanceIP_perms2 int = iota
    propInstanceIP_instance_ip_mode int = iota
    propInstanceIP_service_instance_ip int = iota
    propInstanceIP_display_name int = iota
    propInstanceIP_instance_ip_family int = iota
    propInstanceIP_instance_ip_local_ip int = iota
    propInstanceIP_uuid int = iota
    propInstanceIP_parent_uuid int = iota
)

// InstanceIP 
type InstanceIP struct {

    InstanceIPMode AddressMode `json:"instance_ip_mode,omitempty"`
    ServiceInstanceIP bool `json:"service_instance_ip"`
    DisplayName string `json:"display_name,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    InstanceIPFamily IpAddressFamilyType `json:"instance_ip_family,omitempty"`
    InstanceIPLocalIP bool `json:"instance_ip_local_ip"`
    UUID string `json:"uuid,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    InstanceIPAddress IpAddressType `json:"instance_ip_address,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    InstanceIPSecondary bool `json:"instance_ip_secondary"`
    ParentType string `json:"parent_type,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ServiceHealthCheckIP bool `json:"service_health_check_ip"`
    SecondaryIPTrackingIP *SubnetType `json:"secondary_ip_tracking_ip,omitempty"`
    SubnetUUID string `json:"subnet_uuid,omitempty"`

    NetworkIpamRefs []*InstanceIPNetworkIpamRef `json:"network_ipam_refs,omitempty"`
    VirtualNetworkRefs []*InstanceIPVirtualNetworkRef `json:"virtual_network_refs,omitempty"`
    VirtualMachineInterfaceRefs []*InstanceIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
    PhysicalRouterRefs []*InstanceIPPhysicalRouterRef `json:"physical_router_refs,omitempty"`
    VirtualRouterRefs []*InstanceIPVirtualRouterRef `json:"virtual_router_refs,omitempty"`

    FloatingIPs []*FloatingIP `json:"floating_ips,omitempty"`

    client controller.ObjectInterface
    modified* big.Int
}


// InstanceIPNetworkIpamRef references each other
type InstanceIPNetworkIpamRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// InstanceIPVirtualNetworkRef references each other
type InstanceIPVirtualNetworkRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// InstanceIPVirtualMachineInterfaceRef references each other
type InstanceIPVirtualMachineInterfaceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// InstanceIPPhysicalRouterRef references each other
type InstanceIPPhysicalRouterRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// InstanceIPVirtualRouterRef references each other
type InstanceIPVirtualRouterRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *InstanceIP) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeInstanceIP makes InstanceIP
func MakeInstanceIP() *InstanceIP{
    return &InstanceIP{
    //TODO(nati): Apply default
    ServiceHealthCheckIP: false,
        SecondaryIPTrackingIP: MakeSubnetType(),
        SubnetUUID: "",
        InstanceIPSecondary: false,
        ParentType: "",
        Perms2: MakePermType2(),
        InstanceIPMode: MakeAddressMode(),
        ServiceInstanceIP: false,
        DisplayName: "",
        InstanceIPFamily: MakeIpAddressFamilyType(),
        InstanceIPLocalIP: false,
        UUID: "",
        ParentUUID: "",
        InstanceIPAddress: MakeIpAddressType(),
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        Annotations: MakeKeyValuePairs(),
        
        modified: big.NewInt(0),
    }
}



// MakeInstanceIPSlice makes a slice of InstanceIP
func MakeInstanceIPSlice() []*InstanceIP {
    return []*InstanceIP{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *InstanceIP) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *InstanceIP) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *InstanceIP) GetDefaultName() string {
    return strings.Replace("default-instance_ip", "_", "-", -1)
}

func (model *InstanceIP) GetType() string {
    return strings.Replace("instance_ip", "_", "-", -1)
}

func (model *InstanceIP) GetFQName() []string {
    return model.FQName
}

func (model *InstanceIP) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *InstanceIP) GetParentType() string {
    return model.ParentType
}

func (model *InstanceIP) GetUuid() string {
    return model.UUID
}

func (model *InstanceIP) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *InstanceIP) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *InstanceIP) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *InstanceIP) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *InstanceIP) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propInstanceIP_instance_ip_mode) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.InstanceIPMode); err != nil {
            return nil, errors.Wrap(err, "Marshal of: InstanceIPMode as instance_ip_mode")
        }
        msg["instance_ip_mode"] = &val
    }
    
    if model.modified.Bit(propInstanceIP_service_instance_ip) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceInstanceIP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceInstanceIP as service_instance_ip")
        }
        msg["service_instance_ip"] = &val
    }
    
    if model.modified.Bit(propInstanceIP_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propInstanceIP_instance_ip_family) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.InstanceIPFamily); err != nil {
            return nil, errors.Wrap(err, "Marshal of: InstanceIPFamily as instance_ip_family")
        }
        msg["instance_ip_family"] = &val
    }
    
    if model.modified.Bit(propInstanceIP_instance_ip_local_ip) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.InstanceIPLocalIP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: InstanceIPLocalIP as instance_ip_local_ip")
        }
        msg["instance_ip_local_ip"] = &val
    }
    
    if model.modified.Bit(propInstanceIP_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propInstanceIP_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propInstanceIP_instance_ip_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.InstanceIPAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: InstanceIPAddress as instance_ip_address")
        }
        msg["instance_ip_address"] = &val
    }
    
    if model.modified.Bit(propInstanceIP_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propInstanceIP_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propInstanceIP_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propInstanceIP_service_health_check_ip) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceHealthCheckIP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceHealthCheckIP as service_health_check_ip")
        }
        msg["service_health_check_ip"] = &val
    }
    
    if model.modified.Bit(propInstanceIP_secondary_ip_tracking_ip) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SecondaryIPTrackingIP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SecondaryIPTrackingIP as secondary_ip_tracking_ip")
        }
        msg["secondary_ip_tracking_ip"] = &val
    }
    
    if model.modified.Bit(propInstanceIP_subnet_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SubnetUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SubnetUUID as subnet_uuid")
        }
        msg["subnet_uuid"] = &val
    }
    
    if model.modified.Bit(propInstanceIP_instance_ip_secondary) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.InstanceIPSecondary); err != nil {
            return nil, errors.Wrap(err, "Marshal of: InstanceIPSecondary as instance_ip_secondary")
        }
        msg["instance_ip_secondary"] = &val
    }
    
    if model.modified.Bit(propInstanceIP_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propInstanceIP_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *InstanceIP) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *InstanceIP) UpdateReferences() error {
    return nil
}


