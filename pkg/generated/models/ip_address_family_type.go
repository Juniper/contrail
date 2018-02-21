
package models
// IpAddressFamilyType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type IpAddressFamilyType string

// MakeIpAddressFamilyType makes IpAddressFamilyType
func MakeIpAddressFamilyType() IpAddressFamilyType {
    var data IpAddressFamilyType
    return data
}



// MakeIpAddressFamilyTypeSlice makes a slice of IpAddressFamilyType
func MakeIpAddressFamilyTypeSlice() []IpAddressFamilyType {
    return []IpAddressFamilyType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *IpAddressFamilyType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *IpAddressFamilyType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *IpAddressFamilyType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *IpAddressFamilyType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *IpAddressFamilyType) GetFQName() []string {
    return model.FQName
}

func (model *IpAddressFamilyType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *IpAddressFamilyType) GetParentType() string {
    return model.ParentType
}

func (model *IpAddressFamilyType) GetUuid() string {
    return model.UUID
}

func (model *IpAddressFamilyType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *IpAddressFamilyType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *IpAddressFamilyType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *IpAddressFamilyType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *IpAddressFamilyType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *IpAddressFamilyType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *IpAddressFamilyType) UpdateReferences() error {
    return nil
}


