
package models
// IpamMethodType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type IpamMethodType string

// MakeIpamMethodType makes IpamMethodType
func MakeIpamMethodType() IpamMethodType {
    var data IpamMethodType
    return data
}



// MakeIpamMethodTypeSlice makes a slice of IpamMethodType
func MakeIpamMethodTypeSlice() []IpamMethodType {
    return []IpamMethodType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *IpamMethodType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *IpamMethodType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *IpamMethodType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *IpamMethodType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *IpamMethodType) GetFQName() []string {
    return model.FQName
}

func (model *IpamMethodType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *IpamMethodType) GetParentType() string {
    return model.ParentType
}

func (model *IpamMethodType) GetUuid() string {
    return model.UUID
}

func (model *IpamMethodType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *IpamMethodType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *IpamMethodType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *IpamMethodType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *IpamMethodType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *IpamMethodType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *IpamMethodType) UpdateReferences() error {
    return nil
}


