
package models
// VxlanNetworkIdentifierType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type VxlanNetworkIdentifierType int

// MakeVxlanNetworkIdentifierType makes VxlanNetworkIdentifierType
func MakeVxlanNetworkIdentifierType() VxlanNetworkIdentifierType {
    var data VxlanNetworkIdentifierType
    return data
}



// MakeVxlanNetworkIdentifierTypeSlice makes a slice of VxlanNetworkIdentifierType
func MakeVxlanNetworkIdentifierTypeSlice() []VxlanNetworkIdentifierType {
    return []VxlanNetworkIdentifierType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VxlanNetworkIdentifierType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *VxlanNetworkIdentifierType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VxlanNetworkIdentifierType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *VxlanNetworkIdentifierType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *VxlanNetworkIdentifierType) GetFQName() []string {
    return model.FQName
}

func (model *VxlanNetworkIdentifierType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VxlanNetworkIdentifierType) GetParentType() string {
    return model.ParentType
}

func (model *VxlanNetworkIdentifierType) GetUuid() string {
    return model.UUID
}

func (model *VxlanNetworkIdentifierType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VxlanNetworkIdentifierType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VxlanNetworkIdentifierType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VxlanNetworkIdentifierType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VxlanNetworkIdentifierType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *VxlanNetworkIdentifierType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VxlanNetworkIdentifierType) UpdateReferences() error {
    return nil
}


