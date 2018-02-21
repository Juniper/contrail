
package models
// ServiceConnectionType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type ServiceConnectionType string

// MakeServiceConnectionType makes ServiceConnectionType
func MakeServiceConnectionType() ServiceConnectionType {
    var data ServiceConnectionType
    return data
}



// MakeServiceConnectionTypeSlice makes a slice of ServiceConnectionType
func MakeServiceConnectionTypeSlice() []ServiceConnectionType {
    return []ServiceConnectionType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceConnectionType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *ServiceConnectionType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceConnectionType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *ServiceConnectionType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceConnectionType) GetFQName() []string {
    return model.FQName
}

func (model *ServiceConnectionType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceConnectionType) GetParentType() string {
    return model.ParentType
}

func (model *ServiceConnectionType) GetUuid() string {
    return model.UUID
}

func (model *ServiceConnectionType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceConnectionType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceConnectionType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceConnectionType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceConnectionType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *ServiceConnectionType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceConnectionType) UpdateReferences() error {
    return nil
}


