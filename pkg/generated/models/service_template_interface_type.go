
package models
// ServiceTemplateInterfaceType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propServiceTemplateInterfaceType_static_route_enable int = iota
    propServiceTemplateInterfaceType_shared_ip int = iota
    propServiceTemplateInterfaceType_service_interface_type int = iota
)

// ServiceTemplateInterfaceType 
type ServiceTemplateInterfaceType struct {

    StaticRouteEnable bool `json:"static_route_enable"`
    SharedIP bool `json:"shared_ip"`
    ServiceInterfaceType ServiceInterfaceType `json:"service_interface_type,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *ServiceTemplateInterfaceType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeServiceTemplateInterfaceType makes ServiceTemplateInterfaceType
func MakeServiceTemplateInterfaceType() *ServiceTemplateInterfaceType{
    return &ServiceTemplateInterfaceType{
    //TODO(nati): Apply default
    StaticRouteEnable: false,
        SharedIP: false,
        ServiceInterfaceType: MakeServiceInterfaceType(),
        
        modified: big.NewInt(0),
    }
}



// MakeServiceTemplateInterfaceTypeSlice makes a slice of ServiceTemplateInterfaceType
func MakeServiceTemplateInterfaceTypeSlice() []*ServiceTemplateInterfaceType {
    return []*ServiceTemplateInterfaceType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceTemplateInterfaceType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *ServiceTemplateInterfaceType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceTemplateInterfaceType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *ServiceTemplateInterfaceType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceTemplateInterfaceType) GetFQName() []string {
    return model.FQName
}

func (model *ServiceTemplateInterfaceType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceTemplateInterfaceType) GetParentType() string {
    return model.ParentType
}

func (model *ServiceTemplateInterfaceType) GetUuid() string {
    return model.UUID
}

func (model *ServiceTemplateInterfaceType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceTemplateInterfaceType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceTemplateInterfaceType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceTemplateInterfaceType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceTemplateInterfaceType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propServiceTemplateInterfaceType_static_route_enable) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.StaticRouteEnable); err != nil {
            return nil, errors.Wrap(err, "Marshal of: StaticRouteEnable as static_route_enable")
        }
        msg["static_route_enable"] = &val
    }
    
    if model.modified.Bit(propServiceTemplateInterfaceType_shared_ip) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SharedIP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SharedIP as shared_ip")
        }
        msg["shared_ip"] = &val
    }
    
    if model.modified.Bit(propServiceTemplateInterfaceType_service_interface_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServiceInterfaceType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServiceInterfaceType as service_interface_type")
        }
        msg["service_interface_type"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ServiceTemplateInterfaceType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceTemplateInterfaceType) UpdateReferences() error {
    return nil
}


