
package models
// IpamDnsMethodType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type IpamDnsMethodType string

// MakeIpamDnsMethodType makes IpamDnsMethodType
func MakeIpamDnsMethodType() IpamDnsMethodType {
    var data IpamDnsMethodType
    return data
}



// MakeIpamDnsMethodTypeSlice makes a slice of IpamDnsMethodType
func MakeIpamDnsMethodTypeSlice() []IpamDnsMethodType {
    return []IpamDnsMethodType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *IpamDnsMethodType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *IpamDnsMethodType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *IpamDnsMethodType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *IpamDnsMethodType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *IpamDnsMethodType) GetFQName() []string {
    return model.FQName
}

func (model *IpamDnsMethodType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *IpamDnsMethodType) GetParentType() string {
    return model.ParentType
}

func (model *IpamDnsMethodType) GetUuid() string {
    return model.UUID
}

func (model *IpamDnsMethodType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *IpamDnsMethodType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *IpamDnsMethodType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *IpamDnsMethodType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *IpamDnsMethodType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *IpamDnsMethodType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *IpamDnsMethodType) UpdateReferences() error {
    return nil
}


