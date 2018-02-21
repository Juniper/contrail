
package models
// SimpleActionType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type SimpleActionType string

// MakeSimpleActionType makes SimpleActionType
func MakeSimpleActionType() SimpleActionType {
    var data SimpleActionType
    return data
}



// MakeSimpleActionTypeSlice makes a slice of SimpleActionType
func MakeSimpleActionTypeSlice() []SimpleActionType {
    return []SimpleActionType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *SimpleActionType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *SimpleActionType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *SimpleActionType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *SimpleActionType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *SimpleActionType) GetFQName() []string {
    return model.FQName
}

func (model *SimpleActionType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *SimpleActionType) GetParentType() string {
    return model.ParentType
}

func (model *SimpleActionType) GetUuid() string {
    return model.UUID
}

func (model *SimpleActionType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *SimpleActionType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *SimpleActionType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *SimpleActionType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *SimpleActionType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *SimpleActionType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *SimpleActionType) UpdateReferences() error {
    return nil
}


