
package models
// PortMappings



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propPortMappings_port_mappings int = iota
)

// PortMappings 
type PortMappings struct {

    PortMappings []*PortMap `json:"port_mappings,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *PortMappings) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakePortMappings makes PortMappings
func MakePortMappings() *PortMappings{
    return &PortMappings{
    //TODO(nati): Apply default
    
            
                PortMappings:  MakePortMapSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakePortMappingsSlice makes a slice of PortMappings
func MakePortMappingsSlice() []*PortMappings {
    return []*PortMappings{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *PortMappings) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *PortMappings) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *PortMappings) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *PortMappings) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *PortMappings) GetFQName() []string {
    return model.FQName
}

func (model *PortMappings) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *PortMappings) GetParentType() string {
    return model.ParentType
}

func (model *PortMappings) GetUuid() string {
    return model.UUID
}

func (model *PortMappings) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *PortMappings) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *PortMappings) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *PortMappings) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *PortMappings) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propPortMappings_port_mappings) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PortMappings); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PortMappings as port_mappings")
        }
        msg["port_mappings"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *PortMappings) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *PortMappings) UpdateReferences() error {
    return nil
}


