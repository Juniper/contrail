
package models
// BGPaaServiceParametersType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propBGPaaServiceParametersType_port_end int = iota
    propBGPaaServiceParametersType_port_start int = iota
)

// BGPaaServiceParametersType 
type BGPaaServiceParametersType struct {

    PortStart L4PortType `json:"port_start,omitempty"`
    PortEnd L4PortType `json:"port_end,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *BGPaaServiceParametersType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeBGPaaServiceParametersType makes BGPaaServiceParametersType
func MakeBGPaaServiceParametersType() *BGPaaServiceParametersType{
    return &BGPaaServiceParametersType{
    //TODO(nati): Apply default
    PortStart: MakeL4PortType(),
        PortEnd: MakeL4PortType(),
        
        modified: big.NewInt(0),
    }
}



// MakeBGPaaServiceParametersTypeSlice makes a slice of BGPaaServiceParametersType
func MakeBGPaaServiceParametersTypeSlice() []*BGPaaServiceParametersType {
    return []*BGPaaServiceParametersType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *BGPaaServiceParametersType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *BGPaaServiceParametersType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *BGPaaServiceParametersType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *BGPaaServiceParametersType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *BGPaaServiceParametersType) GetFQName() []string {
    return model.FQName
}

func (model *BGPaaServiceParametersType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *BGPaaServiceParametersType) GetParentType() string {
    return model.ParentType
}

func (model *BGPaaServiceParametersType) GetUuid() string {
    return model.UUID
}

func (model *BGPaaServiceParametersType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *BGPaaServiceParametersType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *BGPaaServiceParametersType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *BGPaaServiceParametersType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *BGPaaServiceParametersType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propBGPaaServiceParametersType_port_end) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PortEnd); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PortEnd as port_end")
        }
        msg["port_end"] = &val
    }
    
    if model.modified.Bit(propBGPaaServiceParametersType_port_start) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PortStart); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PortStart as port_start")
        }
        msg["port_start"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *BGPaaServiceParametersType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *BGPaaServiceParametersType) UpdateReferences() error {
    return nil
}


