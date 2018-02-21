
package models
// AddressType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propAddressType_network_policy int = iota
    propAddressType_subnet_list int = iota
    propAddressType_virtual_network int = iota
    propAddressType_security_group int = iota
    propAddressType_subnet int = iota
)

// AddressType 
type AddressType struct {

    SubnetList []*SubnetType `json:"subnet_list,omitempty"`
    VirtualNetwork string `json:"virtual_network,omitempty"`
    SecurityGroup string `json:"security_group,omitempty"`
    Subnet *SubnetType `json:"subnet,omitempty"`
    NetworkPolicy string `json:"network_policy,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *AddressType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeAddressType makes AddressType
func MakeAddressType() *AddressType{
    return &AddressType{
    //TODO(nati): Apply default
    Subnet: MakeSubnetType(),
        NetworkPolicy: "",
        
            
                SubnetList:  MakeSubnetTypeSlice(),
            
        VirtualNetwork: "",
        SecurityGroup: "",
        
        modified: big.NewInt(0),
    }
}



// MakeAddressTypeSlice makes a slice of AddressType
func MakeAddressTypeSlice() []*AddressType {
    return []*AddressType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *AddressType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *AddressType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *AddressType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *AddressType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *AddressType) GetFQName() []string {
    return model.FQName
}

func (model *AddressType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *AddressType) GetParentType() string {
    return model.ParentType
}

func (model *AddressType) GetUuid() string {
    return model.UUID
}

func (model *AddressType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *AddressType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *AddressType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *AddressType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *AddressType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propAddressType_subnet_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SubnetList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SubnetList as subnet_list")
        }
        msg["subnet_list"] = &val
    }
    
    if model.modified.Bit(propAddressType_virtual_network) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualNetwork); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualNetwork as virtual_network")
        }
        msg["virtual_network"] = &val
    }
    
    if model.modified.Bit(propAddressType_security_group) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SecurityGroup); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SecurityGroup as security_group")
        }
        msg["security_group"] = &val
    }
    
    if model.modified.Bit(propAddressType_subnet) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Subnet); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Subnet as subnet")
        }
        msg["subnet"] = &val
    }
    
    if model.modified.Bit(propAddressType_network_policy) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.NetworkPolicy); err != nil {
            return nil, errors.Wrap(err, "Marshal of: NetworkPolicy as network_policy")
        }
        msg["network_policy"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *AddressType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *AddressType) UpdateReferences() error {
    return nil
}


