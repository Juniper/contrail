
package models
// MirrorActionType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propMirrorActionType_nic_assisted_mirroring int = iota
    propMirrorActionType_nh_mode int = iota
    propMirrorActionType_juniper_header int = iota
    propMirrorActionType_analyzer_ip_address int = iota
    propMirrorActionType_analyzer_mac_address int = iota
    propMirrorActionType_static_nh_header int = iota
    propMirrorActionType_encapsulation int = iota
    propMirrorActionType_nic_assisted_mirroring_vlan int = iota
    propMirrorActionType_analyzer_name int = iota
    propMirrorActionType_udp_port int = iota
    propMirrorActionType_routing_instance int = iota
)

// MirrorActionType 
type MirrorActionType struct {

    JuniperHeader bool `json:"juniper_header"`
    AnalyzerIPAddress string `json:"analyzer_ip_address,omitempty"`
    AnalyzerMacAddress string `json:"analyzer_mac_address,omitempty"`
    NicAssistedMirroring bool `json:"nic_assisted_mirroring"`
    NHMode NHModeType `json:"nh_mode,omitempty"`
    AnalyzerName string `json:"analyzer_name,omitempty"`
    UDPPort int `json:"udp_port,omitempty"`
    RoutingInstance string `json:"routing_instance,omitempty"`
    StaticNHHeader *StaticMirrorNhType `json:"static_nh_header,omitempty"`
    Encapsulation string `json:"encapsulation,omitempty"`
    NicAssistedMirroringVlan VlanIdType `json:"nic_assisted_mirroring_vlan,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *MirrorActionType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeMirrorActionType makes MirrorActionType
func MakeMirrorActionType() *MirrorActionType{
    return &MirrorActionType{
    //TODO(nati): Apply default
    Encapsulation: "",
        NicAssistedMirroringVlan: MakeVlanIdType(),
        AnalyzerName: "",
        UDPPort: 0,
        RoutingInstance: "",
        StaticNHHeader: MakeStaticMirrorNhType(),
        NHMode: MakeNHModeType(),
        JuniperHeader: false,
        AnalyzerIPAddress: "",
        AnalyzerMacAddress: "",
        NicAssistedMirroring: false,
        
        modified: big.NewInt(0),
    }
}



// MakeMirrorActionTypeSlice makes a slice of MirrorActionType
func MakeMirrorActionTypeSlice() []*MirrorActionType {
    return []*MirrorActionType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *MirrorActionType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *MirrorActionType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *MirrorActionType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *MirrorActionType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *MirrorActionType) GetFQName() []string {
    return model.FQName
}

func (model *MirrorActionType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *MirrorActionType) GetParentType() string {
    return model.ParentType
}

func (model *MirrorActionType) GetUuid() string {
    return model.UUID
}

func (model *MirrorActionType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *MirrorActionType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *MirrorActionType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *MirrorActionType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *MirrorActionType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propMirrorActionType_encapsulation) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Encapsulation); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Encapsulation as encapsulation")
        }
        msg["encapsulation"] = &val
    }
    
    if model.modified.Bit(propMirrorActionType_nic_assisted_mirroring_vlan) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.NicAssistedMirroringVlan); err != nil {
            return nil, errors.Wrap(err, "Marshal of: NicAssistedMirroringVlan as nic_assisted_mirroring_vlan")
        }
        msg["nic_assisted_mirroring_vlan"] = &val
    }
    
    if model.modified.Bit(propMirrorActionType_analyzer_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AnalyzerName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AnalyzerName as analyzer_name")
        }
        msg["analyzer_name"] = &val
    }
    
    if model.modified.Bit(propMirrorActionType_udp_port) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UDPPort); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UDPPort as udp_port")
        }
        msg["udp_port"] = &val
    }
    
    if model.modified.Bit(propMirrorActionType_routing_instance) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RoutingInstance); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RoutingInstance as routing_instance")
        }
        msg["routing_instance"] = &val
    }
    
    if model.modified.Bit(propMirrorActionType_static_nh_header) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.StaticNHHeader); err != nil {
            return nil, errors.Wrap(err, "Marshal of: StaticNHHeader as static_nh_header")
        }
        msg["static_nh_header"] = &val
    }
    
    if model.modified.Bit(propMirrorActionType_nh_mode) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.NHMode); err != nil {
            return nil, errors.Wrap(err, "Marshal of: NHMode as nh_mode")
        }
        msg["nh_mode"] = &val
    }
    
    if model.modified.Bit(propMirrorActionType_juniper_header) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.JuniperHeader); err != nil {
            return nil, errors.Wrap(err, "Marshal of: JuniperHeader as juniper_header")
        }
        msg["juniper_header"] = &val
    }
    
    if model.modified.Bit(propMirrorActionType_analyzer_ip_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AnalyzerIPAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AnalyzerIPAddress as analyzer_ip_address")
        }
        msg["analyzer_ip_address"] = &val
    }
    
    if model.modified.Bit(propMirrorActionType_analyzer_mac_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AnalyzerMacAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AnalyzerMacAddress as analyzer_mac_address")
        }
        msg["analyzer_mac_address"] = &val
    }
    
    if model.modified.Bit(propMirrorActionType_nic_assisted_mirroring) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.NicAssistedMirroring); err != nil {
            return nil, errors.Wrap(err, "Marshal of: NicAssistedMirroring as nic_assisted_mirroring")
        }
        msg["nic_assisted_mirroring"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *MirrorActionType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *MirrorActionType) UpdateReferences() error {
    return nil
}


