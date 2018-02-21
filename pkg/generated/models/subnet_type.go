
package models
// SubnetType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propSubnetType_ip_prefix int = iota
    propSubnetType_ip_prefix_len int = iota
)

// SubnetType 
type SubnetType struct {

    IPPrefix string `json:"ip_prefix,omitempty"`
    IPPrefixLen int `json:"ip_prefix_len,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *SubnetType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeSubnetType makes SubnetType
func MakeSubnetType() *SubnetType{
    return &SubnetType{
    //TODO(nati): Apply default
    IPPrefix: "",
        IPPrefixLen: 0,
        
        modified: big.NewInt(0),
    }
}



// MakeSubnetTypeSlice makes a slice of SubnetType
func MakeSubnetTypeSlice() []*SubnetType {
    return []*SubnetType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *SubnetType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *SubnetType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *SubnetType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *SubnetType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *SubnetType) GetFQName() []string {
    return model.FQName
}

func (model *SubnetType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *SubnetType) GetParentType() string {
    return model.ParentType
}

func (model *SubnetType) GetUuid() string {
    return model.UUID
}

func (model *SubnetType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *SubnetType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *SubnetType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *SubnetType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *SubnetType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propSubnetType_ip_prefix_len) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IPPrefixLen); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IPPrefixLen as ip_prefix_len")
        }
        msg["ip_prefix_len"] = &val
    }
    
    if model.modified.Bit(propSubnetType_ip_prefix) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IPPrefix); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IPPrefix as ip_prefix")
        }
        msg["ip_prefix"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *SubnetType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *SubnetType) UpdateReferences() error {
    return nil
}


