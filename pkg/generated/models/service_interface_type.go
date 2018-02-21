
package models
// ServiceInterfaceType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type ServiceInterfaceType string

// MakeServiceInterfaceType makes ServiceInterfaceType
func MakeServiceInterfaceType() ServiceInterfaceType {
    var data ServiceInterfaceType
    return data
}



// MakeServiceInterfaceTypeSlice makes a slice of ServiceInterfaceType
func MakeServiceInterfaceTypeSlice() []ServiceInterfaceType {
    return []ServiceInterfaceType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceInterfaceType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *ServiceInterfaceType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceInterfaceType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *ServiceInterfaceType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceInterfaceType) GetFQName() []string {
    return model.FQName
}

func (model *ServiceInterfaceType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceInterfaceType) GetParentType() string {
    return model.ParentType
}

func (model *ServiceInterfaceType) GetUuid() string {
    return model.UUID
}

func (model *ServiceInterfaceType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceInterfaceType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceInterfaceType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceInterfaceType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceInterfaceType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *ServiceInterfaceType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceInterfaceType) UpdateReferences() error {
    return nil
}


