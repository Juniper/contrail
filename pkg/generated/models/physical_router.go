
package models
// PhysicalRouter



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propPhysicalRouter_physical_router_vnc_managed int = iota
    propPhysicalRouter_physical_router_snmp int = iota
    propPhysicalRouter_physical_router_junos_service_ports int = iota
    propPhysicalRouter_physical_router_user_credentials int = iota
    propPhysicalRouter_physical_router_vendor_name int = iota
    propPhysicalRouter_physical_router_image_uri int = iota
    propPhysicalRouter_telemetry_info int = iota
    propPhysicalRouter_display_name int = iota
    propPhysicalRouter_annotations int = iota
    propPhysicalRouter_physical_router_product_name int = iota
    propPhysicalRouter_physical_router_lldp int = iota
    propPhysicalRouter_perms2 int = iota
    propPhysicalRouter_uuid int = iota
    propPhysicalRouter_parent_type int = iota
    propPhysicalRouter_physical_router_snmp_credentials int = iota
    propPhysicalRouter_physical_router_dataplane_ip int = iota
    propPhysicalRouter_physical_router_loopback_ip int = iota
    propPhysicalRouter_parent_uuid int = iota
    propPhysicalRouter_fq_name int = iota
    propPhysicalRouter_id_perms int = iota
    propPhysicalRouter_physical_router_management_ip int = iota
    propPhysicalRouter_physical_router_role int = iota
)

// PhysicalRouter 
type PhysicalRouter struct {

    PhysicalRouterManagementIP string `json:"physical_router_management_ip,omitempty"`
    PhysicalRouterRole PhysicalRouterRole `json:"physical_router_role,omitempty"`
    PhysicalRouterLoopbackIP string `json:"physical_router_loopback_ip,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    PhysicalRouterUserCredentials *UserCredentials `json:"physical_router_user_credentials,omitempty"`
    PhysicalRouterVendorName string `json:"physical_router_vendor_name,omitempty"`
    PhysicalRouterVNCManaged bool `json:"physical_router_vnc_managed"`
    PhysicalRouterSNMP bool `json:"physical_router_snmp"`
    PhysicalRouterJunosServicePorts *JunosServicePorts `json:"physical_router_junos_service_ports,omitempty"`
    PhysicalRouterProductName string `json:"physical_router_product_name,omitempty"`
    PhysicalRouterLLDP bool `json:"physical_router_lldp"`
    PhysicalRouterImageURI string `json:"physical_router_image_uri,omitempty"`
    TelemetryInfo *TelemetryStateInfo `json:"telemetry_info,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    PhysicalRouterSNMPCredentials *SNMPCredentials `json:"physical_router_snmp_credentials,omitempty"`
    PhysicalRouterDataplaneIP string `json:"physical_router_dataplane_ip,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`

    VirtualNetworkRefs []*PhysicalRouterVirtualNetworkRef `json:"virtual_network_refs,omitempty"`
    BGPRouterRefs []*PhysicalRouterBGPRouterRef `json:"bgp_router_refs,omitempty"`
    VirtualRouterRefs []*PhysicalRouterVirtualRouterRef `json:"virtual_router_refs,omitempty"`

    LogicalInterfaces []*LogicalInterface `json:"logical_interfaces,omitempty"`
    PhysicalInterfaces []*PhysicalInterface `json:"physical_interfaces,omitempty"`

    client controller.ObjectInterface
    modified* big.Int
}


// PhysicalRouterVirtualNetworkRef references each other
type PhysicalRouterVirtualNetworkRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// PhysicalRouterBGPRouterRef references each other
type PhysicalRouterBGPRouterRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// PhysicalRouterVirtualRouterRef references each other
type PhysicalRouterVirtualRouterRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *PhysicalRouter) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakePhysicalRouter makes PhysicalRouter
func MakePhysicalRouter() *PhysicalRouter{
    return &PhysicalRouter{
    //TODO(nati): Apply default
    IDPerms: MakeIdPermsType(),
        PhysicalRouterManagementIP: "",
        PhysicalRouterRole: MakePhysicalRouterRole(),
        PhysicalRouterLoopbackIP: "",
        ParentUUID: "",
        FQName: []string{},
        PhysicalRouterUserCredentials: MakeUserCredentials(),
        PhysicalRouterVendorName: "",
        PhysicalRouterVNCManaged: false,
        PhysicalRouterSNMP: false,
        PhysicalRouterJunosServicePorts: MakeJunosServicePorts(),
        Annotations: MakeKeyValuePairs(),
        PhysicalRouterProductName: "",
        PhysicalRouterLLDP: false,
        PhysicalRouterImageURI: "",
        TelemetryInfo: MakeTelemetryStateInfo(),
        DisplayName: "",
        PhysicalRouterSNMPCredentials: MakeSNMPCredentials(),
        PhysicalRouterDataplaneIP: "",
        Perms2: MakePermType2(),
        UUID: "",
        ParentType: "",
        
        modified: big.NewInt(0),
    }
}



// MakePhysicalRouterSlice makes a slice of PhysicalRouter
func MakePhysicalRouterSlice() []*PhysicalRouter {
    return []*PhysicalRouter{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *PhysicalRouter) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[global_system_config:0xc4202e4320 location:0xc4202e43c0])
    fqn := []string{}
    
    fqn = nil
    
    return fqn
}

func (model *PhysicalRouter) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *PhysicalRouter) GetDefaultName() string {
    return strings.Replace("default-physical_router", "_", "-", -1)
}

func (model *PhysicalRouter) GetType() string {
    return strings.Replace("physical_router", "_", "-", -1)
}

func (model *PhysicalRouter) GetFQName() []string {
    return model.FQName
}

func (model *PhysicalRouter) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *PhysicalRouter) GetParentType() string {
    return model.ParentType
}

func (model *PhysicalRouter) GetUuid() string {
    return model.UUID
}

func (model *PhysicalRouter) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *PhysicalRouter) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *PhysicalRouter) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *PhysicalRouter) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *PhysicalRouter) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propPhysicalRouter_physical_router_snmp) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PhysicalRouterSNMP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PhysicalRouterSNMP as physical_router_snmp")
        }
        msg["physical_router_snmp"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_physical_router_junos_service_ports) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PhysicalRouterJunosServicePorts); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PhysicalRouterJunosServicePorts as physical_router_junos_service_ports")
        }
        msg["physical_router_junos_service_ports"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_physical_router_user_credentials) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PhysicalRouterUserCredentials); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PhysicalRouterUserCredentials as physical_router_user_credentials")
        }
        msg["physical_router_user_credentials"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_physical_router_vendor_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PhysicalRouterVendorName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PhysicalRouterVendorName as physical_router_vendor_name")
        }
        msg["physical_router_vendor_name"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_physical_router_vnc_managed) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PhysicalRouterVNCManaged); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PhysicalRouterVNCManaged as physical_router_vnc_managed")
        }
        msg["physical_router_vnc_managed"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_telemetry_info) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.TelemetryInfo); err != nil {
            return nil, errors.Wrap(err, "Marshal of: TelemetryInfo as telemetry_info")
        }
        msg["telemetry_info"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_physical_router_product_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PhysicalRouterProductName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PhysicalRouterProductName as physical_router_product_name")
        }
        msg["physical_router_product_name"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_physical_router_lldp) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PhysicalRouterLLDP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PhysicalRouterLLDP as physical_router_lldp")
        }
        msg["physical_router_lldp"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_physical_router_image_uri) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PhysicalRouterImageURI); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PhysicalRouterImageURI as physical_router_image_uri")
        }
        msg["physical_router_image_uri"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_physical_router_snmp_credentials) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PhysicalRouterSNMPCredentials); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PhysicalRouterSNMPCredentials as physical_router_snmp_credentials")
        }
        msg["physical_router_snmp_credentials"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_physical_router_dataplane_ip) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PhysicalRouterDataplaneIP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PhysicalRouterDataplaneIP as physical_router_dataplane_ip")
        }
        msg["physical_router_dataplane_ip"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_physical_router_management_ip) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PhysicalRouterManagementIP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PhysicalRouterManagementIP as physical_router_management_ip")
        }
        msg["physical_router_management_ip"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_physical_router_role) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PhysicalRouterRole); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PhysicalRouterRole as physical_router_role")
        }
        msg["physical_router_role"] = &val
    }
    
    if model.modified.Bit(propPhysicalRouter_physical_router_loopback_ip) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PhysicalRouterLoopbackIP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PhysicalRouterLoopbackIP as physical_router_loopback_ip")
        }
        msg["physical_router_loopback_ip"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *PhysicalRouter) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *PhysicalRouter) UpdateReferences() error {
    return nil
}


