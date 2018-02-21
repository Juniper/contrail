
package models
// KeyValuePair



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propKeyValuePair_value int = iota
    propKeyValuePair_key int = iota
)

// KeyValuePair 
type KeyValuePair struct {

    Key string `json:"key,omitempty"`
    Value string `json:"value,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *KeyValuePair) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeKeyValuePair makes KeyValuePair
func MakeKeyValuePair() *KeyValuePair{
    return &KeyValuePair{
    //TODO(nati): Apply default
    Value: "",
        Key: "",
        
        modified: big.NewInt(0),
    }
}



// MakeKeyValuePairSlice makes a slice of KeyValuePair
func MakeKeyValuePairSlice() []*KeyValuePair {
    return []*KeyValuePair{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *KeyValuePair) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *KeyValuePair) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *KeyValuePair) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *KeyValuePair) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *KeyValuePair) GetFQName() []string {
    return model.FQName
}

func (model *KeyValuePair) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *KeyValuePair) GetParentType() string {
    return model.ParentType
}

func (model *KeyValuePair) GetUuid() string {
    return model.UUID
}

func (model *KeyValuePair) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *KeyValuePair) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *KeyValuePair) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *KeyValuePair) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *KeyValuePair) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propKeyValuePair_value) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Value); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Value as value")
        }
        msg["value"] = &val
    }
    
    if model.modified.Bit(propKeyValuePair_key) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Key); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Key as key")
        }
        msg["key"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *KeyValuePair) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *KeyValuePair) UpdateReferences() error {
    return nil
}


