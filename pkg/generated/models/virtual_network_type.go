
package models
// VirtualNetworkType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propVirtualNetworkType_network_id int = iota
    propVirtualNetworkType_mirror_destination int = iota
    propVirtualNetworkType_vxlan_network_identifier int = iota
    propVirtualNetworkType_rpf int = iota
    propVirtualNetworkType_forwarding_mode int = iota
    propVirtualNetworkType_allow_transit int = iota
)

// VirtualNetworkType 
type VirtualNetworkType struct {

    RPF RpfModeType `json:"rpf,omitempty"`
    ForwardingMode ForwardingModeType `json:"forwarding_mode,omitempty"`
    AllowTransit bool `json:"allow_transit"`
    NetworkID int `json:"network_id,omitempty"`
    MirrorDestination bool `json:"mirror_destination"`
    VxlanNetworkIdentifier VxlanNetworkIdentifierType `json:"vxlan_network_identifier,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *VirtualNetworkType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeVirtualNetworkType makes VirtualNetworkType
func MakeVirtualNetworkType() *VirtualNetworkType{
    return &VirtualNetworkType{
    //TODO(nati): Apply default
    MirrorDestination: false,
        VxlanNetworkIdentifier: MakeVxlanNetworkIdentifierType(),
        RPF: MakeRpfModeType(),
        ForwardingMode: MakeForwardingModeType(),
        AllowTransit: false,
        NetworkID: 0,
        
        modified: big.NewInt(0),
    }
}



// MakeVirtualNetworkTypeSlice makes a slice of VirtualNetworkType
func MakeVirtualNetworkTypeSlice() []*VirtualNetworkType {
    return []*VirtualNetworkType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VirtualNetworkType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *VirtualNetworkType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VirtualNetworkType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *VirtualNetworkType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *VirtualNetworkType) GetFQName() []string {
    return model.FQName
}

func (model *VirtualNetworkType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VirtualNetworkType) GetParentType() string {
    return model.ParentType
}

func (model *VirtualNetworkType) GetUuid() string {
    return model.UUID
}

func (model *VirtualNetworkType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VirtualNetworkType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VirtualNetworkType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VirtualNetworkType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VirtualNetworkType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propVirtualNetworkType_forwarding_mode) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ForwardingMode); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ForwardingMode as forwarding_mode")
        }
        msg["forwarding_mode"] = &val
    }
    
    if model.modified.Bit(propVirtualNetworkType_allow_transit) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AllowTransit); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AllowTransit as allow_transit")
        }
        msg["allow_transit"] = &val
    }
    
    if model.modified.Bit(propVirtualNetworkType_network_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.NetworkID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: NetworkID as network_id")
        }
        msg["network_id"] = &val
    }
    
    if model.modified.Bit(propVirtualNetworkType_mirror_destination) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MirrorDestination); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MirrorDestination as mirror_destination")
        }
        msg["mirror_destination"] = &val
    }
    
    if model.modified.Bit(propVirtualNetworkType_vxlan_network_identifier) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VxlanNetworkIdentifier); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VxlanNetworkIdentifier as vxlan_network_identifier")
        }
        msg["vxlan_network_identifier"] = &val
    }
    
    if model.modified.Bit(propVirtualNetworkType_rpf) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RPF); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RPF as rpf")
        }
        msg["rpf"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *VirtualNetworkType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VirtualNetworkType) UpdateReferences() error {
    return nil
}


