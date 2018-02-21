
package models
// ServiceTemplateType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propServiceTemplateType_instance_data int = iota
    propServiceTemplateType_interface_type int = iota
    propServiceTemplateType_image_name int = iota
    propServiceTemplateType_version int = iota
    propServiceTemplateType_service_scaling int = iota
    propServiceTemplateType_availability_zone_enable int = iota
    propServiceTemplateType_service_virtualization_type int = iota
    propServiceTemplateType_service_mode int = iota
    propServiceTemplateType_service_type int = iota
    propServiceTemplateType_flavor int = iota
    propServiceTemplateType_vrouter_instance_type int = iota
    propServiceTemplateType_ordered_interfaces int = iota
)

// ServiceTemplateType 
type ServiceTemplateType struct {

    VrouterInstanceType VRouterInstanceType `json:"vrouter_instance_type,omitempty"`
    OrderedInterfaces bool `json:"ordered_interfaces"`
    ServiceVirtualizationType ServiceVirtualizationType `json:"service_virtualization_type,omitempty"`
    ServiceMode ServiceModeType `json:"service_mode,omitempty"`
    ServiceType ServiceType `json:"service_type,omitempty"`
    Flavor string `json:"flavor,omitempty"`
    ServiceScaling bool `json:"service_scaling"`
    AvailabilityZoneEnable bool `json:"availability_zone_enable"`
    InstanceData string `json:"instance_data,omitempty"`
    InterfaceType []*ServiceTemplateInterfaceType `json:"interface_type,omitempty"`
    ImageName string `json:"image_name,omitempty"`
    Version int `json:"version,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *ServiceTemplateType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeServiceTemplateType makes ServiceTemplateType
func MakeServiceTemplateType() *ServiceTemplateType{
    return &ServiceTemplateType{
    //TODO(nati): Apply default
    VrouterInstanceType: MakeVRouterInstanceType(),
        OrderedInterfaces: false,
        ServiceVirtualizationType: MakeServiceVirtualizationType(),
        ServiceMode: MakeServiceModeType(),
        ServiceType: MakeServiceType(),
        Flavor: "",
        ServiceScaling: false,
        AvailabilityZoneEnable: false,
        InstanceData: "",
        
            
                InterfaceType:  MakeServiceTemplateInterfaceTypeSlice(),
            
        ImageName: "",
        Version: 0,
        
        modified: big.NewInt(0),
    }
}



// MakeServiceTemplateTypeSlice makes a slice of ServiceTemplateType
func MakeServiceTemplateTypeSlice() []*ServiceTemplateType {
    return []*ServiceTemplateType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceTemplateType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *ServiceTemplateType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceTemplateType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *ServiceTemplateType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceTemplateType) GetFQName() []string {
    return model.FQName
}

func (model *ServiceTemplateType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceTemplateType) GetParentType() string {
    return model.ParentType
}

func (model *ServiceTemplateType) GetUuid() string {
    return model.UUID
}

func (model *ServiceTemplateType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceTemplateType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceTemplateType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceTemplateType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceTemplateType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propServiceTemplateType_ordered_interfaces) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.OrderedInterfaces); err != nil {
            return nil, errors.Wrap(err, "Marshal of: OrderedInterfaces as ordered_interfaces")
        }
        msg["ordered_interfaces"] = &val
    }
    
    if model.modified.Bit(propServiceTemplateType_service_virtualization_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceVirtualizationType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceVirtualizationType as service_virtualization_type")
        }
        msg["service_virtualization_type"] = &val
    }
    
    if model.modified.Bit(propServiceTemplateType_service_mode) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceMode); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceMode as service_mode")
        }
        msg["service_mode"] = &val
    }
    
    if model.modified.Bit(propServiceTemplateType_service_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceType as service_type")
        }
        msg["service_type"] = &val
    }
    
    if model.modified.Bit(propServiceTemplateType_flavor) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Flavor); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Flavor as flavor")
        }
        msg["flavor"] = &val
    }
    
    if model.modified.Bit(propServiceTemplateType_vrouter_instance_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VrouterInstanceType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VrouterInstanceType as vrouter_instance_type")
        }
        msg["vrouter_instance_type"] = &val
    }
    
    if model.modified.Bit(propServiceTemplateType_availability_zone_enable) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AvailabilityZoneEnable); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AvailabilityZoneEnable as availability_zone_enable")
        }
        msg["availability_zone_enable"] = &val
    }
    
    if model.modified.Bit(propServiceTemplateType_instance_data) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.InstanceData); err != nil {
            return nil, errors.Wrap(err, "Marshal of: InstanceData as instance_data")
        }
        msg["instance_data"] = &val
    }
    
    if model.modified.Bit(propServiceTemplateType_interface_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.InterfaceType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: InterfaceType as interface_type")
        }
        msg["interface_type"] = &val
    }
    
    if model.modified.Bit(propServiceTemplateType_image_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ImageName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ImageName as image_name")
        }
        msg["image_name"] = &val
    }
    
    if model.modified.Bit(propServiceTemplateType_version) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Version); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Version as version")
        }
        msg["version"] = &val
    }
    
    if model.modified.Bit(propServiceTemplateType_service_scaling) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceScaling); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceScaling as service_scaling")
        }
        msg["service_scaling"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ServiceTemplateType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceTemplateType) UpdateReferences() error {
    return nil
}


