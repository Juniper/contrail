
package models
// SequenceType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propSequenceType_major int = iota
    propSequenceType_minor int = iota
)

// SequenceType 
type SequenceType struct {

    Major int `json:"major,omitempty"`
    Minor int `json:"minor,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *SequenceType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeSequenceType makes SequenceType
func MakeSequenceType() *SequenceType{
    return &SequenceType{
    //TODO(nati): Apply default
    Major: 0,
        Minor: 0,
        
        modified: big.NewInt(0),
    }
}



// MakeSequenceTypeSlice makes a slice of SequenceType
func MakeSequenceTypeSlice() []*SequenceType {
    return []*SequenceType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *SequenceType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *SequenceType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *SequenceType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *SequenceType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *SequenceType) GetFQName() []string {
    return model.FQName
}

func (model *SequenceType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *SequenceType) GetParentType() string {
    return model.ParentType
}

func (model *SequenceType) GetUuid() string {
    return model.UUID
}

func (model *SequenceType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *SequenceType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *SequenceType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *SequenceType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *SequenceType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propSequenceType_major) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Major); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Major as major")
        }
        msg["major"] = &val
    }
    
    if model.modified.Bit(propSequenceType_minor) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Minor); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Minor as minor")
        }
        msg["minor"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *SequenceType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *SequenceType) UpdateReferences() error {
    return nil
}


