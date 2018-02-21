
package models
// PluginProperty



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propPluginProperty_value int = iota
    propPluginProperty_property int = iota
)

// PluginProperty 
type PluginProperty struct {

    Property string `json:"property,omitempty"`
    Value string `json:"value,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *PluginProperty) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakePluginProperty makes PluginProperty
func MakePluginProperty() *PluginProperty{
    return &PluginProperty{
    //TODO(nati): Apply default
    Property: "",
        Value: "",
        
        modified: big.NewInt(0),
    }
}



// MakePluginPropertySlice makes a slice of PluginProperty
func MakePluginPropertySlice() []*PluginProperty {
    return []*PluginProperty{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *PluginProperty) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *PluginProperty) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *PluginProperty) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *PluginProperty) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *PluginProperty) GetFQName() []string {
    return model.FQName
}

func (model *PluginProperty) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *PluginProperty) GetParentType() string {
    return model.ParentType
}

func (model *PluginProperty) GetUuid() string {
    return model.UUID
}

func (model *PluginProperty) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *PluginProperty) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *PluginProperty) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *PluginProperty) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *PluginProperty) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propPluginProperty_property) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Property); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Property as property")
        }
        msg["property"] = &val
    }
    
    if model.modified.Bit(propPluginProperty_value) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Value); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Value as value")
        }
        msg["value"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *PluginProperty) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *PluginProperty) UpdateReferences() error {
    return nil
}


