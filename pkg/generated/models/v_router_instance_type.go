
package models
// VRouterInstanceType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type VRouterInstanceType string

// MakeVRouterInstanceType makes VRouterInstanceType
func MakeVRouterInstanceType() VRouterInstanceType {
    var data VRouterInstanceType
    return data
}



// MakeVRouterInstanceTypeSlice makes a slice of VRouterInstanceType
func MakeVRouterInstanceTypeSlice() []VRouterInstanceType {
    return []VRouterInstanceType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VRouterInstanceType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *VRouterInstanceType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VRouterInstanceType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *VRouterInstanceType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *VRouterInstanceType) GetFQName() []string {
    return model.FQName
}

func (model *VRouterInstanceType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VRouterInstanceType) GetParentType() string {
    return model.ParentType
}

func (model *VRouterInstanceType) GetUuid() string {
    return model.UUID
}

func (model *VRouterInstanceType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VRouterInstanceType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VRouterInstanceType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VRouterInstanceType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VRouterInstanceType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *VRouterInstanceType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VRouterInstanceType) UpdateReferences() error {
    return nil
}


