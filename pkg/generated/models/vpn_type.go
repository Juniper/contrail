
package models
// VpnType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type VpnType string

// MakeVpnType makes VpnType
func MakeVpnType() VpnType {
    var data VpnType
    return data
}



// MakeVpnTypeSlice makes a slice of VpnType
func MakeVpnTypeSlice() []VpnType {
    return []VpnType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VpnType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *VpnType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VpnType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *VpnType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *VpnType) GetFQName() []string {
    return model.FQName
}

func (model *VpnType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VpnType) GetParentType() string {
    return model.ParentType
}

func (model *VpnType) GetUuid() string {
    return model.UUID
}

func (model *VpnType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VpnType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VpnType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VpnType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VpnType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *VpnType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VpnType) UpdateReferences() error {
    return nil
}


