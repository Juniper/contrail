
package models
// PluginProperties



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propPluginProperties_plugin_property int = iota
)

// PluginProperties 
type PluginProperties struct {

    PluginProperty []*PluginProperty `json:"plugin_property,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *PluginProperties) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakePluginProperties makes PluginProperties
func MakePluginProperties() *PluginProperties{
    return &PluginProperties{
    //TODO(nati): Apply default
    
            
                PluginProperty:  MakePluginPropertySlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakePluginPropertiesSlice makes a slice of PluginProperties
func MakePluginPropertiesSlice() []*PluginProperties {
    return []*PluginProperties{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *PluginProperties) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *PluginProperties) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *PluginProperties) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *PluginProperties) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *PluginProperties) GetFQName() []string {
    return model.FQName
}

func (model *PluginProperties) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *PluginProperties) GetParentType() string {
    return model.ParentType
}

func (model *PluginProperties) GetUuid() string {
    return model.UUID
}

func (model *PluginProperties) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *PluginProperties) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *PluginProperties) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *PluginProperties) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *PluginProperties) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propPluginProperties_plugin_property) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PluginProperty); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PluginProperty as plugin_property")
        }
        msg["plugin_property"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *PluginProperties) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *PluginProperties) UpdateReferences() error {
    return nil
}


