
package models
// PortMap



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propPortMap_src_port int = iota
    propPortMap_protocol int = iota
    propPortMap_dst_port int = iota
)

// PortMap 
type PortMap struct {

    SRCPort int `json:"src_port,omitempty"`
    Protocol string `json:"protocol,omitempty"`
    DSTPort int `json:"dst_port,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *PortMap) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakePortMap makes PortMap
func MakePortMap() *PortMap{
    return &PortMap{
    //TODO(nati): Apply default
    DSTPort: 0,
        SRCPort: 0,
        Protocol: "",
        
        modified: big.NewInt(0),
    }
}



// MakePortMapSlice makes a slice of PortMap
func MakePortMapSlice() []*PortMap {
    return []*PortMap{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *PortMap) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *PortMap) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *PortMap) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *PortMap) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *PortMap) GetFQName() []string {
    return model.FQName
}

func (model *PortMap) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *PortMap) GetParentType() string {
    return model.ParentType
}

func (model *PortMap) GetUuid() string {
    return model.UUID
}

func (model *PortMap) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *PortMap) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *PortMap) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *PortMap) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *PortMap) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propPortMap_src_port) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SRCPort); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SRCPort as src_port")
        }
        msg["src_port"] = &val
    }
    
    if model.modified.Bit(propPortMap_protocol) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Protocol); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Protocol as protocol")
        }
        msg["protocol"] = &val
    }
    
    if model.modified.Bit(propPortMap_dst_port) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DSTPort); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DSTPort as dst_port")
        }
        msg["dst_port"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *PortMap) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *PortMap) UpdateReferences() error {
    return nil
}


