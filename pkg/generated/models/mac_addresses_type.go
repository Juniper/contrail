
package models
// MacAddressesType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propMacAddressesType_mac_address int = iota
)

// MacAddressesType 
type MacAddressesType struct {

    MacAddress []string `json:"mac_address,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *MacAddressesType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeMacAddressesType makes MacAddressesType
func MakeMacAddressesType() *MacAddressesType{
    return &MacAddressesType{
    //TODO(nati): Apply default
    MacAddress: []string{},
        
        modified: big.NewInt(0),
    }
}



// MakeMacAddressesTypeSlice makes a slice of MacAddressesType
func MakeMacAddressesTypeSlice() []*MacAddressesType {
    return []*MacAddressesType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *MacAddressesType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *MacAddressesType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *MacAddressesType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *MacAddressesType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *MacAddressesType) GetFQName() []string {
    return model.FQName
}

func (model *MacAddressesType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *MacAddressesType) GetParentType() string {
    return model.ParentType
}

func (model *MacAddressesType) GetUuid() string {
    return model.UUID
}

func (model *MacAddressesType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *MacAddressesType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *MacAddressesType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *MacAddressesType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *MacAddressesType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propMacAddressesType_mac_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MacAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MacAddress as mac_address")
        }
        msg["mac_address"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *MacAddressesType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *MacAddressesType) UpdateReferences() error {
    return nil
}


