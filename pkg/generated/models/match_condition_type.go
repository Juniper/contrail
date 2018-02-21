
package models
// MatchConditionType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propMatchConditionType_src_port int = iota
    propMatchConditionType_src_address int = iota
    propMatchConditionType_ethertype int = iota
    propMatchConditionType_dst_address int = iota
    propMatchConditionType_dst_port int = iota
    propMatchConditionType_protocol int = iota
)

// MatchConditionType 
type MatchConditionType struct {

    Ethertype EtherType `json:"ethertype,omitempty"`
    DSTAddress *AddressType `json:"dst_address,omitempty"`
    DSTPort *PortType `json:"dst_port,omitempty"`
    Protocol string `json:"protocol,omitempty"`
    SRCPort *PortType `json:"src_port,omitempty"`
    SRCAddress *AddressType `json:"src_address,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *MatchConditionType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeMatchConditionType makes MatchConditionType
func MakeMatchConditionType() *MatchConditionType{
    return &MatchConditionType{
    //TODO(nati): Apply default
    Protocol: "",
        SRCPort: MakePortType(),
        SRCAddress: MakeAddressType(),
        Ethertype: MakeEtherType(),
        DSTAddress: MakeAddressType(),
        DSTPort: MakePortType(),
        
        modified: big.NewInt(0),
    }
}



// MakeMatchConditionTypeSlice makes a slice of MatchConditionType
func MakeMatchConditionTypeSlice() []*MatchConditionType {
    return []*MatchConditionType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *MatchConditionType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *MatchConditionType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *MatchConditionType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *MatchConditionType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *MatchConditionType) GetFQName() []string {
    return model.FQName
}

func (model *MatchConditionType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *MatchConditionType) GetParentType() string {
    return model.ParentType
}

func (model *MatchConditionType) GetUuid() string {
    return model.UUID
}

func (model *MatchConditionType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *MatchConditionType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *MatchConditionType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *MatchConditionType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *MatchConditionType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propMatchConditionType_dst_port) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DSTPort); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DSTPort as dst_port")
        }
        msg["dst_port"] = &val
    }
    
    if model.modified.Bit(propMatchConditionType_protocol) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Protocol); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Protocol as protocol")
        }
        msg["protocol"] = &val
    }
    
    if model.modified.Bit(propMatchConditionType_src_port) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SRCPort); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SRCPort as src_port")
        }
        msg["src_port"] = &val
    }
    
    if model.modified.Bit(propMatchConditionType_src_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SRCAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SRCAddress as src_address")
        }
        msg["src_address"] = &val
    }
    
    if model.modified.Bit(propMatchConditionType_ethertype) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Ethertype); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Ethertype as ethertype")
        }
        msg["ethertype"] = &val
    }
    
    if model.modified.Bit(propMatchConditionType_dst_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DSTAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DSTAddress as dst_address")
        }
        msg["dst_address"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *MatchConditionType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *MatchConditionType) UpdateReferences() error {
    return nil
}


