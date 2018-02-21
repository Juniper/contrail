
package models
// UuidStringType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type UuidStringType string

// MakeUuidStringType makes UuidStringType
func MakeUuidStringType() UuidStringType {
    var data UuidStringType
    return data
}



// MakeUuidStringTypeSlice makes a slice of UuidStringType
func MakeUuidStringTypeSlice() []UuidStringType {
    return []UuidStringType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *UuidStringType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *UuidStringType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *UuidStringType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *UuidStringType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *UuidStringType) GetFQName() []string {
    return model.FQName
}

func (model *UuidStringType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *UuidStringType) GetParentType() string {
    return model.ParentType
}

func (model *UuidStringType) GetUuid() string {
    return model.UUID
}

func (model *UuidStringType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *UuidStringType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *UuidStringType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *UuidStringType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *UuidStringType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *UuidStringType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *UuidStringType) UpdateReferences() error {
    return nil
}


