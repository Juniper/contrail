
package models
// VirtualNetworkIdType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type VirtualNetworkIdType int

// MakeVirtualNetworkIdType makes VirtualNetworkIdType
func MakeVirtualNetworkIdType() VirtualNetworkIdType {
    var data VirtualNetworkIdType
    return data
}



// MakeVirtualNetworkIdTypeSlice makes a slice of VirtualNetworkIdType
func MakeVirtualNetworkIdTypeSlice() []VirtualNetworkIdType {
    return []VirtualNetworkIdType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VirtualNetworkIdType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *VirtualNetworkIdType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VirtualNetworkIdType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *VirtualNetworkIdType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *VirtualNetworkIdType) GetFQName() []string {
    return model.FQName
}

func (model *VirtualNetworkIdType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VirtualNetworkIdType) GetParentType() string {
    return model.ParentType
}

func (model *VirtualNetworkIdType) GetUuid() string {
    return model.UUID
}

func (model *VirtualNetworkIdType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VirtualNetworkIdType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VirtualNetworkIdType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VirtualNetworkIdType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VirtualNetworkIdType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *VirtualNetworkIdType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VirtualNetworkIdType) UpdateReferences() error {
    return nil
}


