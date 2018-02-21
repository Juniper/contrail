
package models
// HealthmonitorType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type HealthmonitorType string

// MakeHealthmonitorType makes HealthmonitorType
func MakeHealthmonitorType() HealthmonitorType {
    var data HealthmonitorType
    return data
}



// MakeHealthmonitorTypeSlice makes a slice of HealthmonitorType
func MakeHealthmonitorTypeSlice() []HealthmonitorType {
    return []HealthmonitorType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *HealthmonitorType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *HealthmonitorType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *HealthmonitorType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *HealthmonitorType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *HealthmonitorType) GetFQName() []string {
    return model.FQName
}

func (model *HealthmonitorType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *HealthmonitorType) GetParentType() string {
    return model.ParentType
}

func (model *HealthmonitorType) GetUuid() string {
    return model.UUID
}

func (model *HealthmonitorType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *HealthmonitorType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *HealthmonitorType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *HealthmonitorType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *HealthmonitorType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *HealthmonitorType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *HealthmonitorType) UpdateReferences() error {
    return nil
}


