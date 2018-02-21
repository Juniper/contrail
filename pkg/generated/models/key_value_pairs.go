
package models
// KeyValuePairs



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propKeyValuePairs_key_value_pair int = iota
)

// KeyValuePairs 
type KeyValuePairs struct {

    KeyValuePair []*KeyValuePair `json:"key_value_pair,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *KeyValuePairs) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeKeyValuePairs makes KeyValuePairs
func MakeKeyValuePairs() *KeyValuePairs{
    return &KeyValuePairs{
    //TODO(nati): Apply default
    
            
                KeyValuePair:  MakeKeyValuePairSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeKeyValuePairsSlice makes a slice of KeyValuePairs
func MakeKeyValuePairsSlice() []*KeyValuePairs {
    return []*KeyValuePairs{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *KeyValuePairs) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *KeyValuePairs) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *KeyValuePairs) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *KeyValuePairs) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *KeyValuePairs) GetFQName() []string {
    return model.FQName
}

func (model *KeyValuePairs) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *KeyValuePairs) GetParentType() string {
    return model.ParentType
}

func (model *KeyValuePairs) GetUuid() string {
    return model.UUID
}

func (model *KeyValuePairs) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *KeyValuePairs) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *KeyValuePairs) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *KeyValuePairs) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *KeyValuePairs) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propKeyValuePairs_key_value_pair) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.KeyValuePair); err != nil {
            return nil, errors.Wrap(err, "Marshal of: KeyValuePair as key_value_pair")
        }
        msg["key_value_pair"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *KeyValuePairs) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *KeyValuePairs) UpdateReferences() error {
    return nil
}


