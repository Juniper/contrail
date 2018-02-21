
package models
// VirtualNetwork



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propVirtualNetwork_route_target_list int = iota
    propVirtualNetwork_fq_name int = iota
    propVirtualNetwork_display_name int = iota
    propVirtualNetwork_pbb_evpn_enable int = iota
    propVirtualNetwork_provider_properties int = iota
    propVirtualNetwork_flood_unknown_unicast int = iota
    propVirtualNetwork_pbb_etree_enable int = iota
    propVirtualNetwork_layer2_control_word int = iota
    propVirtualNetwork_address_allocation_mode int = iota
    propVirtualNetwork_mac_aging_time int = iota
    propVirtualNetwork_mac_learning_enabled int = iota
    propVirtualNetwork_external_ipam int = iota
    propVirtualNetwork_mac_limit_control int = iota
    propVirtualNetwork_virtual_network_network_id int = iota
    propVirtualNetwork_import_route_target_list int = iota
    propVirtualNetwork_uuid int = iota
    propVirtualNetwork_mac_move_control int = iota
    propVirtualNetwork_id_perms int = iota
    propVirtualNetwork_ecmp_hashing_include_fields int = iota
    propVirtualNetwork_port_security_enabled int = iota
    propVirtualNetwork_parent_type int = iota
    propVirtualNetwork_multi_policy_service_chains_enabled int = iota
    propVirtualNetwork_is_shared int = iota
    propVirtualNetwork_parent_uuid int = iota
    propVirtualNetwork_annotations int = iota
    propVirtualNetwork_virtual_network_properties int = iota
    propVirtualNetwork_router_external int = iota
    propVirtualNetwork_export_route_target_list int = iota
    propVirtualNetwork_perms2 int = iota
)

// VirtualNetwork 
type VirtualNetwork struct {

    ExternalIpam bool `json:"external_ipam"`
    MacLimitControl *MACLimitControlType `json:"mac_limit_control,omitempty"`
    AddressAllocationMode AddressAllocationModeType `json:"address_allocation_mode,omitempty"`
    MacAgingTime MACAgingTime `json:"mac_aging_time,omitempty"`
    MacLearningEnabled bool `json:"mac_learning_enabled"`
    VirtualNetworkNetworkID VirtualNetworkIdType `json:"virtual_network_network_id,omitempty"`
    ImportRouteTargetList *RouteTargetList `json:"import_route_target_list,omitempty"`
    UUID string `json:"uuid,omitempty"`
    MacMoveControl *MACMoveLimitControlType `json:"mac_move_control,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    EcmpHashingIncludeFields *EcmpHashingIncludeFields `json:"ecmp_hashing_include_fields,omitempty"`
    PortSecurityEnabled bool `json:"port_security_enabled"`
    ParentType string `json:"parent_type,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    MultiPolicyServiceChainsEnabled bool `json:"multi_policy_service_chains_enabled"`
    IsShared bool `json:"is_shared"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    VirtualNetworkProperties *VirtualNetworkType `json:"virtual_network_properties,omitempty"`
    RouterExternal bool `json:"router_external"`
    ExportRouteTargetList *RouteTargetList `json:"export_route_target_list,omitempty"`
    RouteTargetList *RouteTargetList `json:"route_target_list,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    PBBEtreeEnable bool `json:"pbb_etree_enable"`
    Layer2ControlWord bool `json:"layer2_control_word"`
    PBBEvpnEnable bool `json:"pbb_evpn_enable"`
    ProviderProperties *ProviderDetails `json:"provider_properties,omitempty"`
    FloodUnknownUnicast bool `json:"flood_unknown_unicast"`

    NetworkPolicyRefs []*VirtualNetworkNetworkPolicyRef `json:"network_policy_refs,omitempty"`
    QosConfigRefs []*VirtualNetworkQosConfigRef `json:"qos_config_refs,omitempty"`
    RouteTableRefs []*VirtualNetworkRouteTableRef `json:"route_table_refs,omitempty"`
    VirtualNetworkRefs []*VirtualNetworkVirtualNetworkRef `json:"virtual_network_refs,omitempty"`
    BGPVPNRefs []*VirtualNetworkBGPVPNRef `json:"bgpvpn_refs,omitempty"`
    NetworkIpamRefs []*VirtualNetworkNetworkIpamRef `json:"network_ipam_refs,omitempty"`
    SecurityLoggingObjectRefs []*VirtualNetworkSecurityLoggingObjectRef `json:"security_logging_object_refs,omitempty"`

    AccessControlLists []*AccessControlList `json:"access_control_lists,omitempty"`
    AliasIPPools []*AliasIPPool `json:"alias_ip_pools,omitempty"`
    BridgeDomains []*BridgeDomain `json:"bridge_domains,omitempty"`
    FloatingIPPools []*FloatingIPPool `json:"floating_ip_pools,omitempty"`
    RoutingInstances []*RoutingInstance `json:"routing_instances,omitempty"`

    client controller.ObjectInterface
    modified* big.Int
}


// VirtualNetworkSecurityLoggingObjectRef references each other
type VirtualNetworkSecurityLoggingObjectRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// VirtualNetworkNetworkPolicyRef references each other
type VirtualNetworkNetworkPolicyRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
    Attr *VirtualNetworkPolicyType
    
}

// VirtualNetworkQosConfigRef references each other
type VirtualNetworkQosConfigRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// VirtualNetworkRouteTableRef references each other
type VirtualNetworkRouteTableRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// VirtualNetworkVirtualNetworkRef references each other
type VirtualNetworkVirtualNetworkRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// VirtualNetworkBGPVPNRef references each other
type VirtualNetworkBGPVPNRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// VirtualNetworkNetworkIpamRef references each other
type VirtualNetworkNetworkIpamRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
    Attr *VnSubnetsType
    
}


// String returns json representation of the object
func (model *VirtualNetwork) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeVirtualNetwork makes VirtualNetwork
func MakeVirtualNetwork() *VirtualNetwork{
    return &VirtualNetwork{
    //TODO(nati): Apply default
    RouterExternal: false,
        ExportRouteTargetList: MakeRouteTargetList(),
        Perms2: MakePermType2(),
        VirtualNetworkProperties: MakeVirtualNetworkType(),
        FQName: []string{},
        DisplayName: "",
        RouteTargetList: MakeRouteTargetList(),
        ProviderProperties: MakeProviderDetails(),
        FloodUnknownUnicast: false,
        PBBEtreeEnable: false,
        Layer2ControlWord: false,
        PBBEvpnEnable: false,
        MacAgingTime: MakeMACAgingTime(),
        MacLearningEnabled: false,
        ExternalIpam: false,
        MacLimitControl: MakeMACLimitControlType(),
        AddressAllocationMode: MakeAddressAllocationModeType(),
        ImportRouteTargetList: MakeRouteTargetList(),
        UUID: "",
        VirtualNetworkNetworkID: MakeVirtualNetworkIdType(),
        IDPerms: MakeIdPermsType(),
        MacMoveControl: MakeMACMoveLimitControlType(),
        PortSecurityEnabled: false,
        ParentType: "",
        EcmpHashingIncludeFields: MakeEcmpHashingIncludeFields(),
        IsShared: false,
        ParentUUID: "",
        Annotations: MakeKeyValuePairs(),
        MultiPolicyServiceChainsEnabled: false,
        
        modified: big.NewInt(0),
    }
}



// MakeVirtualNetworkSlice makes a slice of VirtualNetwork
func MakeVirtualNetworkSlice() []*VirtualNetwork {
    return []*VirtualNetwork{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VirtualNetwork) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc420150140])
    fqn := []string{}
    
    fqn = Project{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *VirtualNetwork) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-project", "_", "-", -1)
}

func (model *VirtualNetwork) GetDefaultName() string {
    return strings.Replace("default-virtual_network", "_", "-", -1)
}

func (model *VirtualNetwork) GetType() string {
    return strings.Replace("virtual_network", "_", "-", -1)
}

func (model *VirtualNetwork) GetFQName() []string {
    return model.FQName
}

func (model *VirtualNetwork) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VirtualNetwork) GetParentType() string {
    return model.ParentType
}

func (model *VirtualNetwork) GetUuid() string {
    return model.UUID
}

func (model *VirtualNetwork) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VirtualNetwork) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VirtualNetwork) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VirtualNetwork) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VirtualNetwork) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propVirtualNetwork_external_ipam) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ExternalIpam); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ExternalIpam as external_ipam")
        }
        msg["external_ipam"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_mac_limit_control) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MacLimitControl); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MacLimitControl as mac_limit_control")
        }
        msg["mac_limit_control"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_address_allocation_mode) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AddressAllocationMode); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AddressAllocationMode as address_allocation_mode")
        }
        msg["address_allocation_mode"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_mac_aging_time) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MacAgingTime); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MacAgingTime as mac_aging_time")
        }
        msg["mac_aging_time"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_mac_learning_enabled) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MacLearningEnabled); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MacLearningEnabled as mac_learning_enabled")
        }
        msg["mac_learning_enabled"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_virtual_network_network_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualNetworkNetworkID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualNetworkNetworkID as virtual_network_network_id")
        }
        msg["virtual_network_network_id"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_import_route_target_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ImportRouteTargetList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ImportRouteTargetList as import_route_target_list")
        }
        msg["import_route_target_list"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_mac_move_control) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MacMoveControl); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MacMoveControl as mac_move_control")
        }
        msg["mac_move_control"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_ecmp_hashing_include_fields) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.EcmpHashingIncludeFields); err != nil {
            return nil, errors.Wrap(err, "Marshal of: EcmpHashingIncludeFields as ecmp_hashing_include_fields")
        }
        msg["ecmp_hashing_include_fields"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_port_security_enabled) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PortSecurityEnabled); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PortSecurityEnabled as port_security_enabled")
        }
        msg["port_security_enabled"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_multi_policy_service_chains_enabled) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MultiPolicyServiceChainsEnabled); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MultiPolicyServiceChainsEnabled as multi_policy_service_chains_enabled")
        }
        msg["multi_policy_service_chains_enabled"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_is_shared) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IsShared); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IsShared as is_shared")
        }
        msg["is_shared"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_virtual_network_properties) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualNetworkProperties); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualNetworkProperties as virtual_network_properties")
        }
        msg["virtual_network_properties"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_router_external) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RouterExternal); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RouterExternal as router_external")
        }
        msg["router_external"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_export_route_target_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ExportRouteTargetList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ExportRouteTargetList as export_route_target_list")
        }
        msg["export_route_target_list"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_route_target_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RouteTargetList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RouteTargetList as route_target_list")
        }
        msg["route_target_list"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_pbb_etree_enable) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PBBEtreeEnable); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PBBEtreeEnable as pbb_etree_enable")
        }
        msg["pbb_etree_enable"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_layer2_control_word) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Layer2ControlWord); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Layer2ControlWord as layer2_control_word")
        }
        msg["layer2_control_word"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_pbb_evpn_enable) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PBBEvpnEnable); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PBBEvpnEnable as pbb_evpn_enable")
        }
        msg["pbb_evpn_enable"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_provider_properties) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProviderProperties); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProviderProperties as provider_properties")
        }
        msg["provider_properties"] = &val
    }
    
    if model.modified.Bit(propVirtualNetwork_flood_unknown_unicast) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FloodUnknownUnicast); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FloodUnknownUnicast as flood_unknown_unicast")
        }
        msg["flood_unknown_unicast"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *VirtualNetwork) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VirtualNetwork) UpdateReferences() error {
    return nil
}


