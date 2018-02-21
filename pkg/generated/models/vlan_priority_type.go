
package models
// VlanPriorityType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type VlanPriorityType int

// MakeVlanPriorityType makes VlanPriorityType
func MakeVlanPriorityType() VlanPriorityType {
    var data VlanPriorityType
    return data
}



// MakeVlanPriorityTypeSlice makes a slice of VlanPriorityType
func MakeVlanPriorityTypeSlice() []VlanPriorityType {
    return []VlanPriorityType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VlanPriorityType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *VlanPriorityType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VlanPriorityType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *VlanPriorityType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *VlanPriorityType) GetFQName() []string {
    return model.FQName
}

func (model *VlanPriorityType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VlanPriorityType) GetParentType() string {
    return model.ParentType
}

func (model *VlanPriorityType) GetUuid() string {
    return model.UUID
}

func (model *VlanPriorityType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VlanPriorityType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VlanPriorityType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VlanPriorityType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VlanPriorityType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *VlanPriorityType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VlanPriorityType) UpdateReferences() error {
    return nil
}


