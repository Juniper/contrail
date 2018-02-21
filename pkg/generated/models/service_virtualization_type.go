
package models
// ServiceVirtualizationType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type ServiceVirtualizationType string

// MakeServiceVirtualizationType makes ServiceVirtualizationType
func MakeServiceVirtualizationType() ServiceVirtualizationType {
    var data ServiceVirtualizationType
    return data
}



// MakeServiceVirtualizationTypeSlice makes a slice of ServiceVirtualizationType
func MakeServiceVirtualizationTypeSlice() []ServiceVirtualizationType {
    return []ServiceVirtualizationType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ServiceVirtualizationType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *ServiceVirtualizationType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceVirtualizationType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *ServiceVirtualizationType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *ServiceVirtualizationType) GetFQName() []string {
    return model.FQName
}

func (model *ServiceVirtualizationType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ServiceVirtualizationType) GetParentType() string {
    return model.ParentType
}

func (model *ServiceVirtualizationType) GetUuid() string {
    return model.UUID
}

func (model *ServiceVirtualizationType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ServiceVirtualizationType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ServiceVirtualizationType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ServiceVirtualizationType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ServiceVirtualizationType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *ServiceVirtualizationType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ServiceVirtualizationType) UpdateReferences() error {
    return nil
}


