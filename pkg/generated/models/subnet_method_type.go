
package models
// SubnetMethodType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type SubnetMethodType string

// MakeSubnetMethodType makes SubnetMethodType
func MakeSubnetMethodType() SubnetMethodType {
    var data SubnetMethodType
    return data
}



// MakeSubnetMethodTypeSlice makes a slice of SubnetMethodType
func MakeSubnetMethodTypeSlice() []SubnetMethodType {
    return []SubnetMethodType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *SubnetMethodType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *SubnetMethodType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *SubnetMethodType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *SubnetMethodType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *SubnetMethodType) GetFQName() []string {
    return model.FQName
}

func (model *SubnetMethodType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *SubnetMethodType) GetParentType() string {
    return model.ParentType
}

func (model *SubnetMethodType) GetUuid() string {
    return model.UUID
}

func (model *SubnetMethodType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *SubnetMethodType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *SubnetMethodType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *SubnetMethodType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *SubnetMethodType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *SubnetMethodType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *SubnetMethodType) UpdateReferences() error {
    return nil
}


