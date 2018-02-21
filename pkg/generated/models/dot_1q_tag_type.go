
package models
// Dot1QTagType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type Dot1QTagType int

// MakeDot1QTagType makes Dot1QTagType
func MakeDot1QTagType() Dot1QTagType {
    var data Dot1QTagType
    return data
}



// MakeDot1QTagTypeSlice makes a slice of Dot1QTagType
func MakeDot1QTagTypeSlice() []Dot1QTagType {
    return []Dot1QTagType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *Dot1QTagType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *Dot1QTagType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *Dot1QTagType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *Dot1QTagType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *Dot1QTagType) GetFQName() []string {
    return model.FQName
}

func (model *Dot1QTagType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *Dot1QTagType) GetParentType() string {
    return model.ParentType
}

func (model *Dot1QTagType) GetUuid() string {
    return model.UUID
}

func (model *Dot1QTagType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *Dot1QTagType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *Dot1QTagType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *Dot1QTagType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *Dot1QTagType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *Dot1QTagType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *Dot1QTagType) UpdateReferences() error {
    return nil
}


