
package models
// VirtualRouterType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type VirtualRouterType string

// MakeVirtualRouterType makes VirtualRouterType
func MakeVirtualRouterType() VirtualRouterType {
    var data VirtualRouterType
    return data
}



// MakeVirtualRouterTypeSlice makes a slice of VirtualRouterType
func MakeVirtualRouterTypeSlice() []VirtualRouterType {
    return []VirtualRouterType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VirtualRouterType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *VirtualRouterType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VirtualRouterType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *VirtualRouterType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *VirtualRouterType) GetFQName() []string {
    return model.FQName
}

func (model *VirtualRouterType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VirtualRouterType) GetParentType() string {
    return model.ParentType
}

func (model *VirtualRouterType) GetUuid() string {
    return model.UUID
}

func (model *VirtualRouterType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VirtualRouterType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VirtualRouterType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VirtualRouterType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VirtualRouterType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *VirtualRouterType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VirtualRouterType) UpdateReferences() error {
    return nil
}


