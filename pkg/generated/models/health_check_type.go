
package models
// HealthCheckType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type HealthCheckType string

// MakeHealthCheckType makes HealthCheckType
func MakeHealthCheckType() HealthCheckType {
    var data HealthCheckType
    return data
}



// MakeHealthCheckTypeSlice makes a slice of HealthCheckType
func MakeHealthCheckTypeSlice() []HealthCheckType {
    return []HealthCheckType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *HealthCheckType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *HealthCheckType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *HealthCheckType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *HealthCheckType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *HealthCheckType) GetFQName() []string {
    return model.FQName
}

func (model *HealthCheckType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *HealthCheckType) GetParentType() string {
    return model.ParentType
}

func (model *HealthCheckType) GetUuid() string {
    return model.UUID
}

func (model *HealthCheckType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *HealthCheckType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *HealthCheckType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *HealthCheckType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *HealthCheckType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *HealthCheckType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *HealthCheckType) UpdateReferences() error {
    return nil
}


