
package models
// ServiceType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type ServiceType string

// MakeServiceType makes ServiceType
func MakeServiceType() ServiceType {
    var data ServiceType
    return data
}



// MakeServiceTypeSlice makes a slice of ServiceType
func MakeServiceTypeSlice() []ServiceType {
    return []ServiceType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *ServiceType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *ServiceType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceType) GetFQName() []string {
    return model.FQName
}

func (model *ServiceType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceType) GetParentType() string {
    return model.ParentType
}

func (model *ServiceType) GetUuid() string {
    return model.UUID
}

func (model *ServiceType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *ServiceType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceType) UpdateReferences() error {
    return nil
}


