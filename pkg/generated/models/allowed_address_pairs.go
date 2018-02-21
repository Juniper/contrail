
package models
// AllowedAddressPairs



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propAllowedAddressPairs_allowed_address_pair int = iota
)

// AllowedAddressPairs 
type AllowedAddressPairs struct {

    AllowedAddressPair []*AllowedAddressPair `json:"allowed_address_pair,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *AllowedAddressPairs) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeAllowedAddressPairs makes AllowedAddressPairs
func MakeAllowedAddressPairs() *AllowedAddressPairs{
    return &AllowedAddressPairs{
    //TODO(nati): Apply default
    
            
                AllowedAddressPair:  MakeAllowedAddressPairSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeAllowedAddressPairsSlice makes a slice of AllowedAddressPairs
func MakeAllowedAddressPairsSlice() []*AllowedAddressPairs {
    return []*AllowedAddressPairs{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *AllowedAddressPairs) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *AllowedAddressPairs) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *AllowedAddressPairs) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *AllowedAddressPairs) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *AllowedAddressPairs) GetFQName() []string {
    return model.FQName
}

func (model *AllowedAddressPairs) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *AllowedAddressPairs) GetParentType() string {
    return model.ParentType
}

func (model *AllowedAddressPairs) GetUuid() string {
    return model.UUID
}

func (model *AllowedAddressPairs) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *AllowedAddressPairs) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *AllowedAddressPairs) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *AllowedAddressPairs) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *AllowedAddressPairs) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propAllowedAddressPairs_allowed_address_pair) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AllowedAddressPair); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AllowedAddressPair as allowed_address_pair")
        }
        msg["allowed_address_pair"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *AllowedAddressPairs) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *AllowedAddressPairs) UpdateReferences() error {
    return nil
}


