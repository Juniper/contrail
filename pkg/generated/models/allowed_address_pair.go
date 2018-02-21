
package models
// AllowedAddressPair



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propAllowedAddressPair_ip int = iota
    propAllowedAddressPair_mac int = iota
    propAllowedAddressPair_address_mode int = iota
)

// AllowedAddressPair 
type AllowedAddressPair struct {

    IP *SubnetType `json:"ip,omitempty"`
    Mac string `json:"mac,omitempty"`
    AddressMode AddressMode `json:"address_mode,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *AllowedAddressPair) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeAllowedAddressPair makes AllowedAddressPair
func MakeAllowedAddressPair() *AllowedAddressPair{
    return &AllowedAddressPair{
    //TODO(nati): Apply default
    IP: MakeSubnetType(),
        Mac: "",
        AddressMode: MakeAddressMode(),
        
        modified: big.NewInt(0),
    }
}



// MakeAllowedAddressPairSlice makes a slice of AllowedAddressPair
func MakeAllowedAddressPairSlice() []*AllowedAddressPair {
    return []*AllowedAddressPair{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *AllowedAddressPair) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *AllowedAddressPair) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *AllowedAddressPair) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *AllowedAddressPair) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *AllowedAddressPair) GetFQName() []string {
    return model.FQName
}

func (model *AllowedAddressPair) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *AllowedAddressPair) GetParentType() string {
    return model.ParentType
}

func (model *AllowedAddressPair) GetUuid() string {
    return model.UUID
}

func (model *AllowedAddressPair) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *AllowedAddressPair) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *AllowedAddressPair) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *AllowedAddressPair) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *AllowedAddressPair) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propAllowedAddressPair_ip) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IP as ip")
        }
        msg["ip"] = &val
    }
    
    if model.modified.Bit(propAllowedAddressPair_mac) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Mac); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Mac as mac")
        }
        msg["mac"] = &val
    }
    
    if model.modified.Bit(propAllowedAddressPair_address_mode) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AddressMode); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AddressMode as address_mode")
        }
        msg["address_mode"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *AllowedAddressPair) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *AllowedAddressPair) UpdateReferences() error {
    return nil
}


