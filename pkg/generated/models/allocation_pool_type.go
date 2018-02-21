
package models
// AllocationPoolType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propAllocationPoolType_vrouter_specific_pool int = iota
    propAllocationPoolType_start int = iota
    propAllocationPoolType_end int = iota
)

// AllocationPoolType 
type AllocationPoolType struct {

    Start string `json:"start,omitempty"`
    End string `json:"end,omitempty"`
    VrouterSpecificPool bool `json:"vrouter_specific_pool"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *AllocationPoolType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeAllocationPoolType makes AllocationPoolType
func MakeAllocationPoolType() *AllocationPoolType{
    return &AllocationPoolType{
    //TODO(nati): Apply default
    VrouterSpecificPool: false,
        Start: "",
        End: "",
        
        modified: big.NewInt(0),
    }
}



// MakeAllocationPoolTypeSlice makes a slice of AllocationPoolType
func MakeAllocationPoolTypeSlice() []*AllocationPoolType {
    return []*AllocationPoolType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *AllocationPoolType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *AllocationPoolType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *AllocationPoolType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *AllocationPoolType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *AllocationPoolType) GetFQName() []string {
    return model.FQName
}

func (model *AllocationPoolType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *AllocationPoolType) GetParentType() string {
    return model.ParentType
}

func (model *AllocationPoolType) GetUuid() string {
    return model.UUID
}

func (model *AllocationPoolType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *AllocationPoolType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *AllocationPoolType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *AllocationPoolType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *AllocationPoolType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propAllocationPoolType_vrouter_specific_pool) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VrouterSpecificPool); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VrouterSpecificPool as vrouter_specific_pool")
        }
        msg["vrouter_specific_pool"] = &val
    }
    
    if model.modified.Bit(propAllocationPoolType_start) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Start); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Start as start")
        }
        msg["start"] = &val
    }
    
    if model.modified.Bit(propAllocationPoolType_end) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.End); err != nil {
            return nil, errors.Wrap(err, "Marshal of: End as end")
        }
        msg["end"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *AllocationPoolType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *AllocationPoolType) UpdateReferences() error {
    return nil
}


