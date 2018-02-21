
package models
// VirtualRouter



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propVirtualRouter_virtual_router_type int = iota
    propVirtualRouter_fq_name int = iota
    propVirtualRouter_annotations int = iota
    propVirtualRouter_parent_type int = iota
    propVirtualRouter_virtual_router_dpdk_enabled int = iota
    propVirtualRouter_virtual_router_ip_address int = iota
    propVirtualRouter_id_perms int = iota
    propVirtualRouter_display_name int = iota
    propVirtualRouter_perms2 int = iota
    propVirtualRouter_uuid int = iota
    propVirtualRouter_parent_uuid int = iota
)

// VirtualRouter 
type VirtualRouter struct {

    VirtualRouterIPAddress IpAddressType `json:"virtual_router_ip_address,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    VirtualRouterDPDKEnabled bool `json:"virtual_router_dpdk_enabled"`
    FQName []string `json:"fq_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    VirtualRouterType VirtualRouterType `json:"virtual_router_type,omitempty"`

    VirtualMachineRefs []*VirtualRouterVirtualMachineRef `json:"virtual_machine_refs,omitempty"`
    NetworkIpamRefs []*VirtualRouterNetworkIpamRef `json:"network_ipam_refs,omitempty"`

    VirtualMachineInterfaces []*VirtualMachineInterface `json:"virtual_machine_interfaces,omitempty"`

    client controller.ObjectInterface
    modified* big.Int
}


// VirtualRouterNetworkIpamRef references each other
type VirtualRouterNetworkIpamRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
    Attr *VirtualRouterNetworkIpamType
    
}

// VirtualRouterVirtualMachineRef references each other
type VirtualRouterVirtualMachineRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *VirtualRouter) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeVirtualRouter makes VirtualRouter
func MakeVirtualRouter() *VirtualRouter{
    return &VirtualRouter{
    //TODO(nati): Apply default
    ParentUUID: "",
        ParentType: "",
        VirtualRouterDPDKEnabled: false,
        VirtualRouterIPAddress: MakeIpAddressType(),
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Perms2: MakePermType2(),
        UUID: "",
        VirtualRouterType: MakeVirtualRouterType(),
        FQName: []string{},
        Annotations: MakeKeyValuePairs(),
        
        modified: big.NewInt(0),
    }
}



// MakeVirtualRouterSlice makes a slice of VirtualRouter
func MakeVirtualRouterSlice() []*VirtualRouter {
    return []*VirtualRouter{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VirtualRouter) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[global_system_config:0xc420150320])
    fqn := []string{}
    
    fqn = GlobalSystemConfig{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *VirtualRouter) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-global_system_config", "_", "-", -1)
}

func (model *VirtualRouter) GetDefaultName() string {
    return strings.Replace("default-virtual_router", "_", "-", -1)
}

func (model *VirtualRouter) GetType() string {
    return strings.Replace("virtual_router", "_", "-", -1)
}

func (model *VirtualRouter) GetFQName() []string {
    return model.FQName
}

func (model *VirtualRouter) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VirtualRouter) GetParentType() string {
    return model.ParentType
}

func (model *VirtualRouter) GetUuid() string {
    return model.UUID
}

func (model *VirtualRouter) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VirtualRouter) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VirtualRouter) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VirtualRouter) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VirtualRouter) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propVirtualRouter_virtual_router_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualRouterType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualRouterType as virtual_router_type")
        }
        msg["virtual_router_type"] = &val
    }
    
    if model.modified.Bit(propVirtualRouter_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propVirtualRouter_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propVirtualRouter_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propVirtualRouter_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propVirtualRouter_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propVirtualRouter_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propVirtualRouter_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propVirtualRouter_virtual_router_dpdk_enabled) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualRouterDPDKEnabled); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualRouterDPDKEnabled as virtual_router_dpdk_enabled")
        }
        msg["virtual_router_dpdk_enabled"] = &val
    }
    
    if model.modified.Bit(propVirtualRouter_virtual_router_ip_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualRouterIPAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualRouterIPAddress as virtual_router_ip_address")
        }
        msg["virtual_router_ip_address"] = &val
    }
    
    if model.modified.Bit(propVirtualRouter_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *VirtualRouter) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VirtualRouter) UpdateReferences() error {
    return nil
}


