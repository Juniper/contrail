
package models
// PortType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propPortType_end_port int = iota
    propPortType_start_port int = iota
)

// PortType 
type PortType struct {

    EndPort L4PortType `json:"end_port,omitempty"`
    StartPort L4PortType `json:"start_port,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *PortType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakePortType makes PortType
func MakePortType() *PortType{
    return &PortType{
    //TODO(nati): Apply default
    EndPort: MakeL4PortType(),
        StartPort: MakeL4PortType(),
        
        modified: big.NewInt(0),
    }
}



// MakePortTypeSlice makes a slice of PortType
func MakePortTypeSlice() []*PortType {
    return []*PortType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *PortType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *PortType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *PortType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *PortType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *PortType) GetFQName() []string {
    return model.FQName
}

func (model *PortType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *PortType) GetParentType() string {
    return model.ParentType
}

func (model *PortType) GetUuid() string {
    return model.UUID
}

func (model *PortType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *PortType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *PortType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *PortType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *PortType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propPortType_end_port) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.EndPort); err != nil {
            return nil, errors.Wrap(err, "Marshal of: EndPort as end_port")
        }
        msg["end_port"] = &val
    }
    
    if model.modified.Bit(propPortType_start_port) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.StartPort); err != nil {
            return nil, errors.Wrap(err, "Marshal of: StartPort as start_port")
        }
        msg["start_port"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *PortType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *PortType) UpdateReferences() error {
    return nil
}


