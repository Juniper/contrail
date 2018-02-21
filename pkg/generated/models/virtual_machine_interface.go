
package models
// VirtualMachineInterface



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propVirtualMachineInterface_parent_type int = iota
    propVirtualMachineInterface_display_name int = iota
    propVirtualMachineInterface_virtual_machine_interface_disable_policy int = iota
    propVirtualMachineInterface_virtual_machine_interface_fat_flow_protocols int = iota
    propVirtualMachineInterface_uuid int = iota
    propVirtualMachineInterface_virtual_machine_interface_device_owner int = iota
    propVirtualMachineInterface_vrf_assign_table int = iota
    propVirtualMachineInterface_port_security_enabled int = iota
    propVirtualMachineInterface_perms2 int = iota
    propVirtualMachineInterface_parent_uuid int = iota
    propVirtualMachineInterface_virtual_machine_interface_dhcp_option_list int = iota
    propVirtualMachineInterface_virtual_machine_interface_bindings int = iota
    propVirtualMachineInterface_virtual_machine_interface_allowed_address_pairs int = iota
    propVirtualMachineInterface_virtual_machine_interface_properties int = iota
    propVirtualMachineInterface_fq_name int = iota
    propVirtualMachineInterface_id_perms int = iota
    propVirtualMachineInterface_annotations int = iota
    propVirtualMachineInterface_ecmp_hashing_include_fields int = iota
    propVirtualMachineInterface_virtual_machine_interface_host_routes int = iota
    propVirtualMachineInterface_virtual_machine_interface_mac_addresses int = iota
    propVirtualMachineInterface_vlan_tag_based_bridge_domain int = iota
)

// VirtualMachineInterface 
type VirtualMachineInterface struct {

    VlanTagBasedBridgeDomain bool `json:"vlan_tag_based_bridge_domain"`
    DisplayName string `json:"display_name,omitempty"`
    VirtualMachineInterfaceDisablePolicy bool `json:"virtual_machine_interface_disable_policy"`
    VirtualMachineInterfaceFatFlowProtocols *FatFlowProtocols `json:"virtual_machine_interface_fat_flow_protocols,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    VRFAssignTable *VrfAssignTableType `json:"vrf_assign_table,omitempty"`
    PortSecurityEnabled bool `json:"port_security_enabled"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    VirtualMachineInterfaceDHCPOptionList *DhcpOptionsListType `json:"virtual_machine_interface_dhcp_option_list,omitempty"`
    VirtualMachineInterfaceBindings *KeyValuePairs `json:"virtual_machine_interface_bindings,omitempty"`
    VirtualMachineInterfaceAllowedAddressPairs *AllowedAddressPairs `json:"virtual_machine_interface_allowed_address_pairs,omitempty"`
    VirtualMachineInterfaceDeviceOwner string `json:"virtual_machine_interface_device_owner,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    EcmpHashingIncludeFields *EcmpHashingIncludeFields `json:"ecmp_hashing_include_fields,omitempty"`
    VirtualMachineInterfaceHostRoutes *RouteTableType `json:"virtual_machine_interface_host_routes,omitempty"`
    VirtualMachineInterfaceMacAddresses *MacAddressesType `json:"virtual_machine_interface_mac_addresses,omitempty"`
    VirtualMachineInterfaceProperties *VirtualMachineInterfacePropertiesType `json:"virtual_machine_interface_properties,omitempty"`

    PortTupleRefs []*VirtualMachineInterfacePortTupleRef `json:"port_tuple_refs,omitempty"`
    SecurityGroupRefs []*VirtualMachineInterfaceSecurityGroupRef `json:"security_group_refs,omitempty"`
    VirtualMachineRefs []*VirtualMachineInterfaceVirtualMachineRef `json:"virtual_machine_refs,omitempty"`
    SecurityLoggingObjectRefs []*VirtualMachineInterfaceSecurityLoggingObjectRef `json:"security_logging_object_refs,omitempty"`
    InterfaceRouteTableRefs []*VirtualMachineInterfaceInterfaceRouteTableRef `json:"interface_route_table_refs,omitempty"`
    RoutingInstanceRefs []*VirtualMachineInterfaceRoutingInstanceRef `json:"routing_instance_refs,omitempty"`
    PhysicalInterfaceRefs []*VirtualMachineInterfacePhysicalInterfaceRef `json:"physical_interface_refs,omitempty"`
    BGPRouterRefs []*VirtualMachineInterfaceBGPRouterRef `json:"bgp_router_refs,omitempty"`
    VirtualNetworkRefs []*VirtualMachineInterfaceVirtualNetworkRef `json:"virtual_network_refs,omitempty"`
    VirtualMachineInterfaceRefs []*VirtualMachineInterfaceVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
    QosConfigRefs []*VirtualMachineInterfaceQosConfigRef `json:"qos_config_refs,omitempty"`
    ServiceHealthCheckRefs []*VirtualMachineInterfaceServiceHealthCheckRef `json:"service_health_check_refs,omitempty"`
    BridgeDomainRefs []*VirtualMachineInterfaceBridgeDomainRef `json:"bridge_domain_refs,omitempty"`
    ServiceEndpointRefs []*VirtualMachineInterfaceServiceEndpointRef `json:"service_endpoint_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// VirtualMachineInterfaceSecurityLoggingObjectRef references each other
type VirtualMachineInterfaceSecurityLoggingObjectRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// VirtualMachineInterfaceInterfaceRouteTableRef references each other
type VirtualMachineInterfaceInterfaceRouteTableRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// VirtualMachineInterfaceRoutingInstanceRef references each other
type VirtualMachineInterfaceRoutingInstanceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
    Attr *PolicyBasedForwardingRuleType
    
}

// VirtualMachineInterfacePhysicalInterfaceRef references each other
type VirtualMachineInterfacePhysicalInterfaceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// VirtualMachineInterfaceVirtualMachineRef references each other
type VirtualMachineInterfaceVirtualMachineRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// VirtualMachineInterfaceVirtualNetworkRef references each other
type VirtualMachineInterfaceVirtualNetworkRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// VirtualMachineInterfaceBGPRouterRef references each other
type VirtualMachineInterfaceBGPRouterRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// VirtualMachineInterfaceQosConfigRef references each other
type VirtualMachineInterfaceQosConfigRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// VirtualMachineInterfaceServiceHealthCheckRef references each other
type VirtualMachineInterfaceServiceHealthCheckRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// VirtualMachineInterfaceBridgeDomainRef references each other
type VirtualMachineInterfaceBridgeDomainRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
    Attr *BridgeDomainMembershipType
    
}

// VirtualMachineInterfaceServiceEndpointRef references each other
type VirtualMachineInterfaceServiceEndpointRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// VirtualMachineInterfaceVirtualMachineInterfaceRef references each other
type VirtualMachineInterfaceVirtualMachineInterfaceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// VirtualMachineInterfaceSecurityGroupRef references each other
type VirtualMachineInterfaceSecurityGroupRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// VirtualMachineInterfacePortTupleRef references each other
type VirtualMachineInterfacePortTupleRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *VirtualMachineInterface) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeVirtualMachineInterface makes VirtualMachineInterface
func MakeVirtualMachineInterface() *VirtualMachineInterface{
    return &VirtualMachineInterface{
    //TODO(nati): Apply default
    EcmpHashingIncludeFields: MakeEcmpHashingIncludeFields(),
        VirtualMachineInterfaceHostRoutes: MakeRouteTableType(),
        VirtualMachineInterfaceMacAddresses: MakeMacAddressesType(),
        VirtualMachineInterfaceProperties: MakeVirtualMachineInterfacePropertiesType(),
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        Annotations: MakeKeyValuePairs(),
        VlanTagBasedBridgeDomain: false,
        VirtualMachineInterfaceDisablePolicy: false,
        VirtualMachineInterfaceFatFlowProtocols: MakeFatFlowProtocols(),
        UUID: "",
        ParentType: "",
        DisplayName: "",
        VirtualMachineInterfaceDHCPOptionList: MakeDhcpOptionsListType(),
        VirtualMachineInterfaceBindings: MakeKeyValuePairs(),
        VirtualMachineInterfaceAllowedAddressPairs: MakeAllowedAddressPairs(),
        VirtualMachineInterfaceDeviceOwner: "",
        VRFAssignTable: MakeVrfAssignTableType(),
        PortSecurityEnabled: false,
        Perms2: MakePermType2(),
        ParentUUID: "",
        
        modified: big.NewInt(0),
    }
}



// MakeVirtualMachineInterfaceSlice makes a slice of VirtualMachineInterface
func MakeVirtualMachineInterfaceSlice() []*VirtualMachineInterface {
    return []*VirtualMachineInterface{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VirtualMachineInterface) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc420147c20 virtual_machine:0xc420147cc0 virtual_router:0xc420147d60])
    fqn := []string{}
    
    fqn = nil
    
    return fqn
}

func (model *VirtualMachineInterface) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VirtualMachineInterface) GetDefaultName() string {
    return strings.Replace("default-virtual_machine_interface", "_", "-", -1)
}

func (model *VirtualMachineInterface) GetType() string {
    return strings.Replace("virtual_machine_interface", "_", "-", -1)
}

func (model *VirtualMachineInterface) GetFQName() []string {
    return model.FQName
}

func (model *VirtualMachineInterface) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VirtualMachineInterface) GetParentType() string {
    return model.ParentType
}

func (model *VirtualMachineInterface) GetUuid() string {
    return model.UUID
}

func (model *VirtualMachineInterface) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VirtualMachineInterface) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VirtualMachineInterface) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VirtualMachineInterface) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VirtualMachineInterface) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propVirtualMachineInterface_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterface_ecmp_hashing_include_fields) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.EcmpHashingIncludeFields); err != nil {
            return nil, errors.Wrap(err, "Marshal of: EcmpHashingIncludeFields as ecmp_hashing_include_fields")
        }
        msg["ecmp_hashing_include_fields"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterface_virtual_machine_interface_host_routes) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualMachineInterfaceHostRoutes); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualMachineInterfaceHostRoutes as virtual_machine_interface_host_routes")
        }
        msg["virtual_machine_interface_host_routes"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterface_virtual_machine_interface_mac_addresses) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualMachineInterfaceMacAddresses); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualMachineInterfaceMacAddresses as virtual_machine_interface_mac_addresses")
        }
        msg["virtual_machine_interface_mac_addresses"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterface_virtual_machine_interface_properties) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualMachineInterfaceProperties); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualMachineInterfaceProperties as virtual_machine_interface_properties")
        }
        msg["virtual_machine_interface_properties"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterface_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterface_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterface_vlan_tag_based_bridge_domain) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VlanTagBasedBridgeDomain); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VlanTagBasedBridgeDomain as vlan_tag_based_bridge_domain")
        }
        msg["vlan_tag_based_bridge_domain"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterface_virtual_machine_interface_disable_policy) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualMachineInterfaceDisablePolicy); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualMachineInterfaceDisablePolicy as virtual_machine_interface_disable_policy")
        }
        msg["virtual_machine_interface_disable_policy"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterface_virtual_machine_interface_fat_flow_protocols) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualMachineInterfaceFatFlowProtocols); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualMachineInterfaceFatFlowProtocols as virtual_machine_interface_fat_flow_protocols")
        }
        msg["virtual_machine_interface_fat_flow_protocols"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterface_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterface_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterface_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterface_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterface_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterface_virtual_machine_interface_dhcp_option_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualMachineInterfaceDHCPOptionList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualMachineInterfaceDHCPOptionList as virtual_machine_interface_dhcp_option_list")
        }
        msg["virtual_machine_interface_dhcp_option_list"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterface_virtual_machine_interface_bindings) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualMachineInterfaceBindings); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualMachineInterfaceBindings as virtual_machine_interface_bindings")
        }
        msg["virtual_machine_interface_bindings"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterface_virtual_machine_interface_allowed_address_pairs) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualMachineInterfaceAllowedAddressPairs); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualMachineInterfaceAllowedAddressPairs as virtual_machine_interface_allowed_address_pairs")
        }
        msg["virtual_machine_interface_allowed_address_pairs"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterface_virtual_machine_interface_device_owner) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualMachineInterfaceDeviceOwner); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualMachineInterfaceDeviceOwner as virtual_machine_interface_device_owner")
        }
        msg["virtual_machine_interface_device_owner"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterface_vrf_assign_table) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VRFAssignTable); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VRFAssignTable as vrf_assign_table")
        }
        msg["vrf_assign_table"] = &val
    }
    
    if model.modified.Bit(propVirtualMachineInterface_port_security_enabled) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PortSecurityEnabled); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PortSecurityEnabled as port_security_enabled")
        }
        msg["port_security_enabled"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *VirtualMachineInterface) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VirtualMachineInterface) UpdateReferences() error {
    return nil
}


