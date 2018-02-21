
package models
// FatFlowProtocols



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propFatFlowProtocols_fat_flow_protocol int = iota
)

// FatFlowProtocols 
type FatFlowProtocols struct {

    FatFlowProtocol []*ProtocolType `json:"fat_flow_protocol,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *FatFlowProtocols) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeFatFlowProtocols makes FatFlowProtocols
func MakeFatFlowProtocols() *FatFlowProtocols{
    return &FatFlowProtocols{
    //TODO(nati): Apply default
    
            
                FatFlowProtocol:  MakeProtocolTypeSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeFatFlowProtocolsSlice makes a slice of FatFlowProtocols
func MakeFatFlowProtocolsSlice() []*FatFlowProtocols {
    return []*FatFlowProtocols{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *FatFlowProtocols) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *FatFlowProtocols) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *FatFlowProtocols) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *FatFlowProtocols) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *FatFlowProtocols) GetFQName() []string {
    return model.FQName
}

func (model *FatFlowProtocols) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *FatFlowProtocols) GetParentType() string {
    return model.ParentType
}

func (model *FatFlowProtocols) GetUuid() string {
    return model.UUID
}

func (model *FatFlowProtocols) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *FatFlowProtocols) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *FatFlowProtocols) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *FatFlowProtocols) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *FatFlowProtocols) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propFatFlowProtocols_fat_flow_protocol) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FatFlowProtocol); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FatFlowProtocol as fat_flow_protocol")
        }
        msg["fat_flow_protocol"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *FatFlowProtocols) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *FatFlowProtocols) UpdateReferences() error {
    return nil
}


