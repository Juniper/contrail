
package models
// VlanIdType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type VlanIdType int

// MakeVlanIdType makes VlanIdType
func MakeVlanIdType() VlanIdType {
    var data VlanIdType
    return data
}



// MakeVlanIdTypeSlice makes a slice of VlanIdType
func MakeVlanIdTypeSlice() []VlanIdType {
    return []VlanIdType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VlanIdType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *VlanIdType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VlanIdType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *VlanIdType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *VlanIdType) GetFQName() []string {
    return model.FQName
}

func (model *VlanIdType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VlanIdType) GetParentType() string {
    return model.ParentType
}

func (model *VlanIdType) GetUuid() string {
    return model.UUID
}

func (model *VlanIdType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VlanIdType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VlanIdType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VlanIdType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VlanIdType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *VlanIdType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VlanIdType) UpdateReferences() error {
    return nil
}


