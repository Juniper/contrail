
package models
// ServiceInstanceType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propServiceInstanceType_right_virtual_network int = iota
    propServiceInstanceType_availability_zone int = iota
    propServiceInstanceType_management_virtual_network int = iota
    propServiceInstanceType_virtual_router_id int = iota
    propServiceInstanceType_left_ip_address int = iota
    propServiceInstanceType_right_ip_address int = iota
    propServiceInstanceType_scale_out int = iota
    propServiceInstanceType_ha_mode int = iota
    propServiceInstanceType_interface_list int = iota
    propServiceInstanceType_left_virtual_network int = iota
    propServiceInstanceType_auto_policy int = iota
)

// ServiceInstanceType 
type ServiceInstanceType struct {

    AutoPolicy bool `json:"auto_policy"`
    RightIPAddress IpAddressType `json:"right_ip_address,omitempty"`
    ScaleOut *ServiceScaleOutType `json:"scale_out,omitempty"`
    HaMode AddressMode `json:"ha_mode,omitempty"`
    InterfaceList []*ServiceInstanceInterfaceType `json:"interface_list,omitempty"`
    LeftVirtualNetwork string `json:"left_virtual_network,omitempty"`
    RightVirtualNetwork string `json:"right_virtual_network,omitempty"`
    AvailabilityZone string `json:"availability_zone,omitempty"`
    ManagementVirtualNetwork string `json:"management_virtual_network,omitempty"`
    VirtualRouterID string `json:"virtual_router_id,omitempty"`
    LeftIPAddress IpAddressType `json:"left_ip_address,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *ServiceInstanceType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeServiceInstanceType makes ServiceInstanceType
func MakeServiceInstanceType() *ServiceInstanceType{
    return &ServiceInstanceType{
    //TODO(nati): Apply default
    ScaleOut: MakeServiceScaleOutType(),
        HaMode: MakeAddressMode(),
        
            
                InterfaceList:  MakeServiceInstanceInterfaceTypeSlice(),
            
        LeftVirtualNetwork: "",
        AutoPolicy: false,
        RightIPAddress: MakeIpAddressType(),
        AvailabilityZone: "",
        ManagementVirtualNetwork: "",
        VirtualRouterID: "",
        LeftIPAddress: MakeIpAddressType(),
        RightVirtualNetwork: "",
        
        modified: big.NewInt(0),
    }
}



// MakeServiceInstanceTypeSlice makes a slice of ServiceInstanceType
func MakeServiceInstanceTypeSlice() []*ServiceInstanceType {
    return []*ServiceInstanceType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceInstanceType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *ServiceInstanceType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceInstanceType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *ServiceInstanceType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceInstanceType) GetFQName() []string {
    return model.FQName
}

func (model *ServiceInstanceType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceInstanceType) GetParentType() string {
    return model.ParentType
}

func (model *ServiceInstanceType) GetUuid() string {
    return model.UUID
}

func (model *ServiceInstanceType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceInstanceType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceInstanceType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceInstanceType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceInstanceType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propServiceInstanceType_availability_zone) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AvailabilityZone); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AvailabilityZone as availability_zone")
        }
        msg["availability_zone"] = &val
    }
    
    if model.modified.Bit(propServiceInstanceType_management_virtual_network) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ManagementVirtualNetwork); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ManagementVirtualNetwork as management_virtual_network")
        }
        msg["management_virtual_network"] = &val
    }
    
    if model.modified.Bit(propServiceInstanceType_virtual_router_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualRouterID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualRouterID as virtual_router_id")
        }
        msg["virtual_router_id"] = &val
    }
    
    if model.modified.Bit(propServiceInstanceType_left_ip_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LeftIPAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LeftIPAddress as left_ip_address")
        }
        msg["left_ip_address"] = &val
    }
    
    if model.modified.Bit(propServiceInstanceType_right_virtual_network) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RightVirtualNetwork); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RightVirtualNetwork as right_virtual_network")
        }
        msg["right_virtual_network"] = &val
    }
    
    if model.modified.Bit(propServiceInstanceType_scale_out) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ScaleOut); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ScaleOut as scale_out")
        }
        msg["scale_out"] = &val
    }
    
    if model.modified.Bit(propServiceInstanceType_ha_mode) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.HaMode); err != nil {
            return nil, errors.Wrap(err, "Marshal of: HaMode as ha_mode")
        }
        msg["ha_mode"] = &val
    }
    
    if model.modified.Bit(propServiceInstanceType_interface_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.InterfaceList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: InterfaceList as interface_list")
        }
        msg["interface_list"] = &val
    }
    
    if model.modified.Bit(propServiceInstanceType_left_virtual_network) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LeftVirtualNetwork); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LeftVirtualNetwork as left_virtual_network")
        }
        msg["left_virtual_network"] = &val
    }
    
    if model.modified.Bit(propServiceInstanceType_auto_policy) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AutoPolicy); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AutoPolicy as auto_policy")
        }
        msg["auto_policy"] = &val
    }
    
    if model.modified.Bit(propServiceInstanceType_right_ip_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RightIPAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RightIPAddress as right_ip_address")
        }
        msg["right_ip_address"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ServiceInstanceType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceInstanceType) UpdateReferences() error {
    return nil
}


