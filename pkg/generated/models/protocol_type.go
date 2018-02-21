
package models
// ProtocolType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propProtocolType_protocol int = iota
    propProtocolType_port int = iota
)

// ProtocolType 
type ProtocolType struct {

    Protocol string `json:"protocol,omitempty"`
    Port int `json:"port,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *ProtocolType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeProtocolType makes ProtocolType
func MakeProtocolType() *ProtocolType{
    return &ProtocolType{
    //TODO(nati): Apply default
    Protocol: "",
        Port: 0,
        
        modified: big.NewInt(0),
    }
}



// MakeProtocolTypeSlice makes a slice of ProtocolType
func MakeProtocolTypeSlice() []*ProtocolType {
    return []*ProtocolType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ProtocolType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *ProtocolType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ProtocolType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *ProtocolType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *ProtocolType) GetFQName() []string {
    return model.FQName
}

func (model *ProtocolType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ProtocolType) GetParentType() string {
    return model.ParentType
}

func (model *ProtocolType) GetUuid() string {
    return model.UUID
}

func (model *ProtocolType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ProtocolType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ProtocolType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ProtocolType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ProtocolType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propProtocolType_protocol) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Protocol); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Protocol as protocol")
        }
        msg["protocol"] = &val
    }
    
    if model.modified.Bit(propProtocolType_port) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Port); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Port as port")
        }
        msg["port"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ProtocolType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ProtocolType) UpdateReferences() error {
    return nil
}


