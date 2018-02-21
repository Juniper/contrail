
package models
// EncapsulationPrioritiesType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propEncapsulationPrioritiesType_encapsulation int = iota
)

// EncapsulationPrioritiesType 
type EncapsulationPrioritiesType struct {

    Encapsulation EncapsulationType `json:"encapsulation,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *EncapsulationPrioritiesType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeEncapsulationPrioritiesType makes EncapsulationPrioritiesType
func MakeEncapsulationPrioritiesType() *EncapsulationPrioritiesType{
    return &EncapsulationPrioritiesType{
    //TODO(nati): Apply default
    
            
                Encapsulation:  MakeEncapsulationType(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeEncapsulationPrioritiesTypeSlice makes a slice of EncapsulationPrioritiesType
func MakeEncapsulationPrioritiesTypeSlice() []*EncapsulationPrioritiesType {
    return []*EncapsulationPrioritiesType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *EncapsulationPrioritiesType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *EncapsulationPrioritiesType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *EncapsulationPrioritiesType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *EncapsulationPrioritiesType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *EncapsulationPrioritiesType) GetFQName() []string {
    return model.FQName
}

func (model *EncapsulationPrioritiesType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *EncapsulationPrioritiesType) GetParentType() string {
    return model.ParentType
}

func (model *EncapsulationPrioritiesType) GetUuid() string {
    return model.UUID
}

func (model *EncapsulationPrioritiesType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *EncapsulationPrioritiesType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *EncapsulationPrioritiesType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *EncapsulationPrioritiesType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *EncapsulationPrioritiesType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propEncapsulationPrioritiesType_encapsulation) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Encapsulation); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Encapsulation as encapsulation")
        }
        msg["encapsulation"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *EncapsulationPrioritiesType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *EncapsulationPrioritiesType) UpdateReferences() error {
    return nil
}


