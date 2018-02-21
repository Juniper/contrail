
package models
// IpAddressesType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propIpAddressesType_ip_address int = iota
)

// IpAddressesType 
type IpAddressesType struct {

    IPAddress IpAddressType `json:"ip_address,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *IpAddressesType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeIpAddressesType makes IpAddressesType
func MakeIpAddressesType() *IpAddressesType{
    return &IpAddressesType{
    //TODO(nati): Apply default
    IPAddress: MakeIpAddressType(),
        
        modified: big.NewInt(0),
    }
}



// MakeIpAddressesTypeSlice makes a slice of IpAddressesType
func MakeIpAddressesTypeSlice() []*IpAddressesType {
    return []*IpAddressesType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *IpAddressesType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *IpAddressesType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *IpAddressesType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *IpAddressesType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *IpAddressesType) GetFQName() []string {
    return model.FQName
}

func (model *IpAddressesType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *IpAddressesType) GetParentType() string {
    return model.ParentType
}

func (model *IpAddressesType) GetUuid() string {
    return model.UUID
}

func (model *IpAddressesType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *IpAddressesType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *IpAddressesType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *IpAddressesType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *IpAddressesType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propIpAddressesType_ip_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IPAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IPAddress as ip_address")
        }
        msg["ip_address"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *IpAddressesType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *IpAddressesType) UpdateReferences() error {
    return nil
}


