
package models
// ServiceModeType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type ServiceModeType string

// MakeServiceModeType makes ServiceModeType
func MakeServiceModeType() ServiceModeType {
    var data ServiceModeType
    return data
}



// MakeServiceModeTypeSlice makes a slice of ServiceModeType
func MakeServiceModeTypeSlice() []ServiceModeType {
    return []ServiceModeType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceModeType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *ServiceModeType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceModeType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *ServiceModeType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceModeType) GetFQName() []string {
    return model.FQName
}

func (model *ServiceModeType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceModeType) GetParentType() string {
    return model.ParentType
}

func (model *ServiceModeType) GetUuid() string {
    return model.UUID
}

func (model *ServiceModeType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceModeType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceModeType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceModeType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceModeType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *ServiceModeType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceModeType) UpdateReferences() error {
    return nil
}


