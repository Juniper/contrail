
package models
// ServiceInstanceInterfaceType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propServiceInstanceInterfaceType_ip_address int = iota
    propServiceInstanceInterfaceType_allowed_address_pairs int = iota
    propServiceInstanceInterfaceType_static_routes int = iota
    propServiceInstanceInterfaceType_virtual_network int = iota
)

// ServiceInstanceInterfaceType 
type ServiceInstanceInterfaceType struct {

    VirtualNetwork string `json:"virtual_network,omitempty"`
    IPAddress IpAddressType `json:"ip_address,omitempty"`
    AllowedAddressPairs *AllowedAddressPairs `json:"allowed_address_pairs,omitempty"`
    StaticRoutes *RouteTableType `json:"static_routes,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *ServiceInstanceInterfaceType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeServiceInstanceInterfaceType makes ServiceInstanceInterfaceType
func MakeServiceInstanceInterfaceType() *ServiceInstanceInterfaceType{
    return &ServiceInstanceInterfaceType{
    //TODO(nati): Apply default
    StaticRoutes: MakeRouteTableType(),
        VirtualNetwork: "",
        IPAddress: MakeIpAddressType(),
        AllowedAddressPairs: MakeAllowedAddressPairs(),
        
        modified: big.NewInt(0),
    }
}



// MakeServiceInstanceInterfaceTypeSlice makes a slice of ServiceInstanceInterfaceType
func MakeServiceInstanceInterfaceTypeSlice() []*ServiceInstanceInterfaceType {
    return []*ServiceInstanceInterfaceType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceInstanceInterfaceType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *ServiceInstanceInterfaceType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceInstanceInterfaceType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *ServiceInstanceInterfaceType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceInstanceInterfaceType) GetFQName() []string {
    return model.FQName
}

func (model *ServiceInstanceInterfaceType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceInstanceInterfaceType) GetParentType() string {
    return model.ParentType
}

func (model *ServiceInstanceInterfaceType) GetUuid() string {
    return model.UUID
}

func (model *ServiceInstanceInterfaceType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceInstanceInterfaceType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceInstanceInterfaceType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceInstanceInterfaceType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceInstanceInterfaceType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propServiceInstanceInterfaceType_virtual_network) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualNetwork); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualNetwork as virtual_network")
        }
        msg["virtual_network"] = &val
    }
    
    if model.modified.Bit(propServiceInstanceInterfaceType_ip_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IPAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IPAddress as ip_address")
        }
        msg["ip_address"] = &val
    }
    
    if model.modified.Bit(propServiceInstanceInterfaceType_allowed_address_pairs) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AllowedAddressPairs); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AllowedAddressPairs as allowed_address_pairs")
        }
        msg["allowed_address_pairs"] = &val
    }
    
    if model.modified.Bit(propServiceInstanceInterfaceType_static_routes) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.StaticRoutes); err != nil {
            return nil, errors.Wrap(err, "Marshal of: StaticRoutes as static_routes")
        }
        msg["static_routes"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ServiceInstanceInterfaceType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceInstanceInterfaceType) UpdateReferences() error {
    return nil
}


