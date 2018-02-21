
package models
// ServiceScaleOutType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propServiceScaleOutType_auto_scale int = iota
    propServiceScaleOutType_max_instances int = iota
)

// ServiceScaleOutType 
type ServiceScaleOutType struct {

    AutoScale bool `json:"auto_scale"`
    MaxInstances int `json:"max_instances,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *ServiceScaleOutType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeServiceScaleOutType makes ServiceScaleOutType
func MakeServiceScaleOutType() *ServiceScaleOutType{
    return &ServiceScaleOutType{
    //TODO(nati): Apply default
    AutoScale: false,
        MaxInstances: 0,
        
        modified: big.NewInt(0),
    }
}



// MakeServiceScaleOutTypeSlice makes a slice of ServiceScaleOutType
func MakeServiceScaleOutTypeSlice() []*ServiceScaleOutType {
    return []*ServiceScaleOutType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceScaleOutType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *ServiceScaleOutType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceScaleOutType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *ServiceScaleOutType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceScaleOutType) GetFQName() []string {
    return model.FQName
}

func (model *ServiceScaleOutType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceScaleOutType) GetParentType() string {
    return model.ParentType
}

func (model *ServiceScaleOutType) GetUuid() string {
    return model.UUID
}

func (model *ServiceScaleOutType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceScaleOutType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceScaleOutType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceScaleOutType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceScaleOutType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propServiceScaleOutType_auto_scale) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AutoScale); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AutoScale as auto_scale")
        }
        msg["auto_scale"] = &val
    }
    
    if model.modified.Bit(propServiceScaleOutType_max_instances) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MaxInstances); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MaxInstances as max_instances")
        }
        msg["max_instances"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ServiceScaleOutType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceScaleOutType) UpdateReferences() error {
    return nil
}


