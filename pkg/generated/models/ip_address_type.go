
package models
// IpAddressType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type IpAddressType string

// MakeIpAddressType makes IpAddressType
func MakeIpAddressType() IpAddressType {
    var data IpAddressType
    return data
}



// MakeIpAddressTypeSlice makes a slice of IpAddressType
func MakeIpAddressTypeSlice() []IpAddressType {
    return []IpAddressType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *IpAddressType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *IpAddressType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *IpAddressType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *IpAddressType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *IpAddressType) GetFQName() []string {
    return model.FQName
}

func (model *IpAddressType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *IpAddressType) GetParentType() string {
    return model.ParentType
}

func (model *IpAddressType) GetUuid() string {
    return model.UUID
}

func (model *IpAddressType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *IpAddressType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *IpAddressType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *IpAddressType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *IpAddressType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *IpAddressType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *IpAddressType) UpdateReferences() error {
    return nil
}


