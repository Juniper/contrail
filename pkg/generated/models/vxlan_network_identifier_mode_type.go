
package models
// VxlanNetworkIdentifierModeType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type VxlanNetworkIdentifierModeType string

// MakeVxlanNetworkIdentifierModeType makes VxlanNetworkIdentifierModeType
func MakeVxlanNetworkIdentifierModeType() VxlanNetworkIdentifierModeType {
    var data VxlanNetworkIdentifierModeType
    return data
}



// MakeVxlanNetworkIdentifierModeTypeSlice makes a slice of VxlanNetworkIdentifierModeType
func MakeVxlanNetworkIdentifierModeTypeSlice() []VxlanNetworkIdentifierModeType {
    return []VxlanNetworkIdentifierModeType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VxlanNetworkIdentifierModeType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *VxlanNetworkIdentifierModeType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VxlanNetworkIdentifierModeType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *VxlanNetworkIdentifierModeType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *VxlanNetworkIdentifierModeType) GetFQName() []string {
    return model.FQName
}

func (model *VxlanNetworkIdentifierModeType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VxlanNetworkIdentifierModeType) GetParentType() string {
    return model.ParentType
}

func (model *VxlanNetworkIdentifierModeType) GetUuid() string {
    return model.UUID
}

func (model *VxlanNetworkIdentifierModeType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VxlanNetworkIdentifierModeType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VxlanNetworkIdentifierModeType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VxlanNetworkIdentifierModeType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VxlanNetworkIdentifierModeType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *VxlanNetworkIdentifierModeType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VxlanNetworkIdentifierModeType) UpdateReferences() error {
    return nil
}


