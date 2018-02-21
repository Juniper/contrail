
package models
// ServiceInterfaceTag



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propServiceInterfaceTag_interface_type int = iota
)

// ServiceInterfaceTag 
type ServiceInterfaceTag struct {

    InterfaceType ServiceInterfaceType `json:"interface_type,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *ServiceInterfaceTag) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeServiceInterfaceTag makes ServiceInterfaceTag
func MakeServiceInterfaceTag() *ServiceInterfaceTag{
    return &ServiceInterfaceTag{
    //TODO(nati): Apply default
    InterfaceType: MakeServiceInterfaceType(),
        
        modified: big.NewInt(0),
    }
}



// MakeServiceInterfaceTagSlice makes a slice of ServiceInterfaceTag
func MakeServiceInterfaceTagSlice() []*ServiceInterfaceTag {
    return []*ServiceInterfaceTag{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceInterfaceTag) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *ServiceInterfaceTag) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceInterfaceTag) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *ServiceInterfaceTag) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceInterfaceTag) GetFQName() []string {
    return model.FQName
}

func (model *ServiceInterfaceTag) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceInterfaceTag) GetParentType() string {
    return model.ParentType
}

func (model *ServiceInterfaceTag) GetUuid() string {
    return model.UUID
}

func (model *ServiceInterfaceTag) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceInterfaceTag) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceInterfaceTag) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceInterfaceTag) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceInterfaceTag) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propServiceInterfaceTag_interface_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.InterfaceType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: InterfaceType as interface_type")
        }
        msg["interface_type"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ServiceInterfaceTag) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceInterfaceTag) UpdateReferences() error {
    return nil
}


