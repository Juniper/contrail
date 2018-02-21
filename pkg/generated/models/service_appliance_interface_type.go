
package models
// ServiceApplianceInterfaceType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propServiceApplianceInterfaceType_interface_type int = iota
)

// ServiceApplianceInterfaceType 
type ServiceApplianceInterfaceType struct {

    InterfaceType ServiceInterfaceType `json:"interface_type,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *ServiceApplianceInterfaceType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeServiceApplianceInterfaceType makes ServiceApplianceInterfaceType
func MakeServiceApplianceInterfaceType() *ServiceApplianceInterfaceType{
    return &ServiceApplianceInterfaceType{
    //TODO(nati): Apply default
    InterfaceType: MakeServiceInterfaceType(),
        
        modified: big.NewInt(0),
    }
}



// MakeServiceApplianceInterfaceTypeSlice makes a slice of ServiceApplianceInterfaceType
func MakeServiceApplianceInterfaceTypeSlice() []*ServiceApplianceInterfaceType {
    return []*ServiceApplianceInterfaceType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceApplianceInterfaceType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *ServiceApplianceInterfaceType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceApplianceInterfaceType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *ServiceApplianceInterfaceType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceApplianceInterfaceType) GetFQName() []string {
    return model.FQName
}

func (model *ServiceApplianceInterfaceType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceApplianceInterfaceType) GetParentType() string {
    return model.ParentType
}

func (model *ServiceApplianceInterfaceType) GetUuid() string {
    return model.UUID
}

func (model *ServiceApplianceInterfaceType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceApplianceInterfaceType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceApplianceInterfaceType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceApplianceInterfaceType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceApplianceInterfaceType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propServiceApplianceInterfaceType_interface_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.InterfaceType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: InterfaceType as interface_type")
        }
        msg["interface_type"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ServiceApplianceInterfaceType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceApplianceInterfaceType) UpdateReferences() error {
    return nil
}


