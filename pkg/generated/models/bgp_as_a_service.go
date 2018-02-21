
package models
// BGPAsAService



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propBGPAsAService_annotations int = iota
    propBGPAsAService_bgpaas_ipv4_mapped_ipv6_nexthop int = iota
    propBGPAsAService_perms2 int = iota
    propBGPAsAService_uuid int = iota
    propBGPAsAService_id_perms int = iota
    propBGPAsAService_bgpaas_session_attributes int = iota
    propBGPAsAService_bgpaas_ip_address int = iota
    propBGPAsAService_autonomous_system int = iota
    propBGPAsAService_parent_uuid int = iota
    propBGPAsAService_bgpaas_shared int = iota
    propBGPAsAService_bgpaas_suppress_route_advertisement int = iota
    propBGPAsAService_parent_type int = iota
    propBGPAsAService_fq_name int = iota
    propBGPAsAService_display_name int = iota
)

// BGPAsAService 
type BGPAsAService struct {

    BgpaasSessionAttributes string `json:"bgpaas_session_attributes,omitempty"`
    BgpaasIPAddress IpAddressType `json:"bgpaas_ip_address,omitempty"`
    AutonomousSystem AutonomousSystemType `json:"autonomous_system,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    BgpaasShared bool `json:"bgpaas_shared"`
    BgpaasSuppressRouteAdvertisement bool `json:"bgpaas_suppress_route_advertisement"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    BgpaasIpv4MappedIpv6Nexthop bool `json:"bgpaas_ipv4_mapped_ipv6_nexthop"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`

    VirtualMachineInterfaceRefs []*BGPAsAServiceVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`
    ServiceHealthCheckRefs []*BGPAsAServiceServiceHealthCheckRef `json:"service_health_check_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// BGPAsAServiceVirtualMachineInterfaceRef references each other
type BGPAsAServiceVirtualMachineInterfaceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// BGPAsAServiceServiceHealthCheckRef references each other
type BGPAsAServiceServiceHealthCheckRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *BGPAsAService) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeBGPAsAService makes BGPAsAService
func MakeBGPAsAService() *BGPAsAService{
    return &BGPAsAService{
    //TODO(nati): Apply default
    Annotations: MakeKeyValuePairs(),
        UUID: "",
        IDPerms: MakeIdPermsType(),
        BgpaasIpv4MappedIpv6Nexthop: false,
        Perms2: MakePermType2(),
        AutonomousSystem: MakeAutonomousSystemType(),
        ParentUUID: "",
        BgpaasSessionAttributes: "",
        BgpaasIPAddress: MakeIpAddressType(),
        ParentType: "",
        FQName: []string{},
        DisplayName: "",
        BgpaasShared: false,
        BgpaasSuppressRouteAdvertisement: false,
        
        modified: big.NewInt(0),
    }
}



// MakeBGPAsAServiceSlice makes a slice of BGPAsAService
func MakeBGPAsAServiceSlice() []*BGPAsAService {
    return []*BGPAsAService{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *BGPAsAService) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc420183400])
    fqn := []string{}
    
    fqn = Project{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *BGPAsAService) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-project", "_", "-", -1)
}

func (model *BGPAsAService) GetDefaultName() string {
    return strings.Replace("default-bgp_as_a_service", "_", "-", -1)
}

func (model *BGPAsAService) GetType() string {
    return strings.Replace("bgp_as_a_service", "_", "-", -1)
}

func (model *BGPAsAService) GetFQName() []string {
    return model.FQName
}

func (model *BGPAsAService) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *BGPAsAService) GetParentType() string {
    return model.ParentType
}

func (model *BGPAsAService) GetUuid() string {
    return model.UUID
}

func (model *BGPAsAService) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *BGPAsAService) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *BGPAsAService) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *BGPAsAService) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *BGPAsAService) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propBGPAsAService_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propBGPAsAService_bgpaas_ipv4_mapped_ipv6_nexthop) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.BgpaasIpv4MappedIpv6Nexthop); err != nil {
            return nil, errors.Wrap(err, "Marshal of: BgpaasIpv4MappedIpv6Nexthop as bgpaas_ipv4_mapped_ipv6_nexthop")
        }
        msg["bgpaas_ipv4_mapped_ipv6_nexthop"] = &val
    }
    
    if model.modified.Bit(propBGPAsAService_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propBGPAsAService_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propBGPAsAService_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propBGPAsAService_bgpaas_session_attributes) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.BgpaasSessionAttributes); err != nil {
            return nil, errors.Wrap(err, "Marshal of: BgpaasSessionAttributes as bgpaas_session_attributes")
        }
        msg["bgpaas_session_attributes"] = &val
    }
    
    if model.modified.Bit(propBGPAsAService_bgpaas_ip_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.BgpaasIPAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: BgpaasIPAddress as bgpaas_ip_address")
        }
        msg["bgpaas_ip_address"] = &val
    }
    
    if model.modified.Bit(propBGPAsAService_autonomous_system) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AutonomousSystem); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AutonomousSystem as autonomous_system")
        }
        msg["autonomous_system"] = &val
    }
    
    if model.modified.Bit(propBGPAsAService_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propBGPAsAService_bgpaas_shared) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.BgpaasShared); err != nil {
            return nil, errors.Wrap(err, "Marshal of: BgpaasShared as bgpaas_shared")
        }
        msg["bgpaas_shared"] = &val
    }
    
    if model.modified.Bit(propBGPAsAService_bgpaas_suppress_route_advertisement) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.BgpaasSuppressRouteAdvertisement); err != nil {
            return nil, errors.Wrap(err, "Marshal of: BgpaasSuppressRouteAdvertisement as bgpaas_suppress_route_advertisement")
        }
        msg["bgpaas_suppress_route_advertisement"] = &val
    }
    
    if model.modified.Bit(propBGPAsAService_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propBGPAsAService_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propBGPAsAService_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *BGPAsAService) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *BGPAsAService) UpdateReferences() error {
    return nil
}


