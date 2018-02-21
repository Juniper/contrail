
package models
// UveKeysType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propUveKeysType_uve_key int = iota
)

// UveKeysType 
type UveKeysType struct {

    UveKey []string `json:"uve_key,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *UveKeysType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeUveKeysType makes UveKeysType
func MakeUveKeysType() *UveKeysType{
    return &UveKeysType{
    //TODO(nati): Apply default
    UveKey: []string{},
        
        modified: big.NewInt(0),
    }
}



// MakeUveKeysTypeSlice makes a slice of UveKeysType
func MakeUveKeysTypeSlice() []*UveKeysType {
    return []*UveKeysType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *UveKeysType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *UveKeysType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *UveKeysType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *UveKeysType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *UveKeysType) GetFQName() []string {
    return model.FQName
}

func (model *UveKeysType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *UveKeysType) GetParentType() string {
    return model.ParentType
}

func (model *UveKeysType) GetUuid() string {
    return model.UUID
}

func (model *UveKeysType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *UveKeysType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *UveKeysType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *UveKeysType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *UveKeysType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propUveKeysType_uve_key) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UveKey); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UveKey as uve_key")
        }
        msg["uve_key"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *UveKeysType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *UveKeysType) UpdateReferences() error {
    return nil
}


