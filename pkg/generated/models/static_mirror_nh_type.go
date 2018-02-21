
package models
// StaticMirrorNhType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propStaticMirrorNhType_vni int = iota
    propStaticMirrorNhType_vtep_dst_ip_address int = iota
    propStaticMirrorNhType_vtep_dst_mac_address int = iota
)

// StaticMirrorNhType 
type StaticMirrorNhType struct {

    VtepDSTMacAddress string `json:"vtep_dst_mac_address,omitempty"`
    Vni VxlanNetworkIdentifierType `json:"vni,omitempty"`
    VtepDSTIPAddress string `json:"vtep_dst_ip_address,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *StaticMirrorNhType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeStaticMirrorNhType makes StaticMirrorNhType
func MakeStaticMirrorNhType() *StaticMirrorNhType{
    return &StaticMirrorNhType{
    //TODO(nati): Apply default
    Vni: MakeVxlanNetworkIdentifierType(),
        VtepDSTIPAddress: "",
        VtepDSTMacAddress: "",
        
        modified: big.NewInt(0),
    }
}



// MakeStaticMirrorNhTypeSlice makes a slice of StaticMirrorNhType
func MakeStaticMirrorNhTypeSlice() []*StaticMirrorNhType {
    return []*StaticMirrorNhType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *StaticMirrorNhType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *StaticMirrorNhType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *StaticMirrorNhType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *StaticMirrorNhType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *StaticMirrorNhType) GetFQName() []string {
    return model.FQName
}

func (model *StaticMirrorNhType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *StaticMirrorNhType) GetParentType() string {
    return model.ParentType
}

func (model *StaticMirrorNhType) GetUuid() string {
    return model.UUID
}

func (model *StaticMirrorNhType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *StaticMirrorNhType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *StaticMirrorNhType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *StaticMirrorNhType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *StaticMirrorNhType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propStaticMirrorNhType_vtep_dst_ip_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VtepDSTIPAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VtepDSTIPAddress as vtep_dst_ip_address")
        }
        msg["vtep_dst_ip_address"] = &val
    }
    
    if model.modified.Bit(propStaticMirrorNhType_vtep_dst_mac_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VtepDSTMacAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VtepDSTMacAddress as vtep_dst_mac_address")
        }
        msg["vtep_dst_mac_address"] = &val
    }
    
    if model.modified.Bit(propStaticMirrorNhType_vni) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Vni); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Vni as vni")
        }
        msg["vni"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *StaticMirrorNhType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *StaticMirrorNhType) UpdateReferences() error {
    return nil
}


