
package models
// LogicalRouter



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propLogicalRouter_vxlan_network_identifier int = iota
    propLogicalRouter_display_name int = iota
    propLogicalRouter_parent_uuid int = iota
    propLogicalRouter_fq_name int = iota
    propLogicalRouter_id_perms int = iota
    propLogicalRouter_configured_route_target_list int = iota
    propLogicalRouter_annotations int = iota
    propLogicalRouter_perms2 int = iota
    propLogicalRouter_uuid int = iota
    propLogicalRouter_parent_type int = iota
)

// LogicalRouter 
type LogicalRouter struct {

    VxlanNetworkIdentifier string `json:"vxlan_network_identifier,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    ConfiguredRouteTargetList *RouteTargetList `json:"configured_route_target_list,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`

    PhysicalRouterRefs []*LogicalRouterPhysicalRouterRef `json:"physical_router_refs,omitempty"`
    BGPVPNRefs []*LogicalRouterBGPVPNRef `json:"bgpvpn_refs,omitempty"`
    RouteTargetRefs []*LogicalRouterRouteTargetRef `json:"route_target_refs,omitempty"`
    VirtualMachineInterfaceRefs []*LogicalRouterVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
    ServiceInstanceRefs []*LogicalRouterServiceInstanceRef `json:"service_instance_refs,omitempty"`
    RouteTableRefs []*LogicalRouterRouteTableRef `json:"route_table_refs,omitempty"`
    VirtualNetworkRefs []*LogicalRouterVirtualNetworkRef `json:"virtual_network_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// LogicalRouterServiceInstanceRef references each other
type LogicalRouterServiceInstanceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// LogicalRouterRouteTableRef references each other
type LogicalRouterRouteTableRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// LogicalRouterVirtualNetworkRef references each other
type LogicalRouterVirtualNetworkRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// LogicalRouterPhysicalRouterRef references each other
type LogicalRouterPhysicalRouterRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// LogicalRouterBGPVPNRef references each other
type LogicalRouterBGPVPNRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// LogicalRouterRouteTargetRef references each other
type LogicalRouterRouteTargetRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// LogicalRouterVirtualMachineInterfaceRef references each other
type LogicalRouterVirtualMachineInterfaceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *LogicalRouter) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeLogicalRouter makes LogicalRouter
func MakeLogicalRouter() *LogicalRouter{
    return &LogicalRouter{
    //TODO(nati): Apply default
    VxlanNetworkIdentifier: "",
        DisplayName: "",
        ParentUUID: "",
        FQName: []string{},
        ConfiguredRouteTargetList: MakeRouteTargetList(),
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        UUID: "",
        ParentType: "",
        IDPerms: MakeIdPermsType(),
        
        modified: big.NewInt(0),
    }
}



// MakeLogicalRouterSlice makes a slice of LogicalRouter
func MakeLogicalRouterSlice() []*LogicalRouter {
    return []*LogicalRouter{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *LogicalRouter) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc42024bc20])
    fqn := []string{}
    
    fqn = Project{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *LogicalRouter) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-project", "_", "-", -1)
}

func (model *LogicalRouter) GetDefaultName() string {
    return strings.Replace("default-logical_router", "_", "-", -1)
}

func (model *LogicalRouter) GetType() string {
    return strings.Replace("logical_router", "_", "-", -1)
}

func (model *LogicalRouter) GetFQName() []string {
    return model.FQName
}

func (model *LogicalRouter) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *LogicalRouter) GetParentType() string {
    return model.ParentType
}

func (model *LogicalRouter) GetUuid() string {
    return model.UUID
}

func (model *LogicalRouter) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *LogicalRouter) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *LogicalRouter) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *LogicalRouter) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *LogicalRouter) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propLogicalRouter_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propLogicalRouter_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propLogicalRouter_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propLogicalRouter_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propLogicalRouter_configured_route_target_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ConfiguredRouteTargetList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ConfiguredRouteTargetList as configured_route_target_list")
        }
        msg["configured_route_target_list"] = &val
    }
    
    if model.modified.Bit(propLogicalRouter_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propLogicalRouter_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propLogicalRouter_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propLogicalRouter_vxlan_network_identifier) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VxlanNetworkIdentifier); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VxlanNetworkIdentifier as vxlan_network_identifier")
        }
        msg["vxlan_network_identifier"] = &val
    }
    
    if model.modified.Bit(propLogicalRouter_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *LogicalRouter) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *LogicalRouter) UpdateReferences() error {
    return nil
}


