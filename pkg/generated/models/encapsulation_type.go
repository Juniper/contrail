
package models
// EncapsulationType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type EncapsulationType []string

// MakeEncapsulationType makes EncapsulationType
func MakeEncapsulationType() EncapsulationType {
    var data EncapsulationType
    return data
}



// MakeEncapsulationTypeSlice makes a slice of EncapsulationType
func MakeEncapsulationTypeSlice() []EncapsulationType {
    return []EncapsulationType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *EncapsulationType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *EncapsulationType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *EncapsulationType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *EncapsulationType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *EncapsulationType) GetFQName() []string {
    return model.FQName
}

func (model *EncapsulationType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *EncapsulationType) GetParentType() string {
    return model.ParentType
}

func (model *EncapsulationType) GetUuid() string {
    return model.UUID
}

func (model *EncapsulationType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *EncapsulationType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *EncapsulationType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *EncapsulationType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *EncapsulationType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *EncapsulationType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *EncapsulationType) UpdateReferences() error {
    return nil
}


