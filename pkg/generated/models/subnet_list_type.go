
package models
// SubnetListType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propSubnetListType_subnet int = iota
)

// SubnetListType 
type SubnetListType struct {

    Subnet []*SubnetType `json:"subnet,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *SubnetListType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeSubnetListType makes SubnetListType
func MakeSubnetListType() *SubnetListType{
    return &SubnetListType{
    //TODO(nati): Apply default
    
            
                Subnet:  MakeSubnetTypeSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeSubnetListTypeSlice makes a slice of SubnetListType
func MakeSubnetListTypeSlice() []*SubnetListType {
    return []*SubnetListType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *SubnetListType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *SubnetListType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *SubnetListType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *SubnetListType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *SubnetListType) GetFQName() []string {
    return model.FQName
}

func (model *SubnetListType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *SubnetListType) GetParentType() string {
    return model.ParentType
}

func (model *SubnetListType) GetUuid() string {
    return model.UUID
}

func (model *SubnetListType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *SubnetListType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *SubnetListType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *SubnetListType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *SubnetListType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propSubnetListType_subnet) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Subnet); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Subnet as subnet")
        }
        msg["subnet"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *SubnetListType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *SubnetListType) UpdateReferences() error {
    return nil
}


