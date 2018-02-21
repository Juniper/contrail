
package models
// HealthCheckProtocolType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type HealthCheckProtocolType string

// MakeHealthCheckProtocolType makes HealthCheckProtocolType
func MakeHealthCheckProtocolType() HealthCheckProtocolType {
    var data HealthCheckProtocolType
    return data
}



// MakeHealthCheckProtocolTypeSlice makes a slice of HealthCheckProtocolType
func MakeHealthCheckProtocolTypeSlice() []HealthCheckProtocolType {
    return []HealthCheckProtocolType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *HealthCheckProtocolType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *HealthCheckProtocolType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *HealthCheckProtocolType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *HealthCheckProtocolType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *HealthCheckProtocolType) GetFQName() []string {
    return model.FQName
}

func (model *HealthCheckProtocolType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *HealthCheckProtocolType) GetParentType() string {
    return model.ParentType
}

func (model *HealthCheckProtocolType) GetUuid() string {
    return model.UUID
}

func (model *HealthCheckProtocolType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *HealthCheckProtocolType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *HealthCheckProtocolType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *HealthCheckProtocolType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *HealthCheckProtocolType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *HealthCheckProtocolType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *HealthCheckProtocolType) UpdateReferences() error {
    return nil
}


