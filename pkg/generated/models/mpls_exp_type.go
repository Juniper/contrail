
package models
// MplsExpType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type MplsExpType int

// MakeMplsExpType makes MplsExpType
func MakeMplsExpType() MplsExpType {
    var data MplsExpType
    return data
}



// MakeMplsExpTypeSlice makes a slice of MplsExpType
func MakeMplsExpTypeSlice() []MplsExpType {
    return []MplsExpType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *MplsExpType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *MplsExpType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *MplsExpType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *MplsExpType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *MplsExpType) GetFQName() []string {
    return model.FQName
}

func (model *MplsExpType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *MplsExpType) GetParentType() string {
    return model.ParentType
}

func (model *MplsExpType) GetUuid() string {
    return model.UUID
}

func (model *MplsExpType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *MplsExpType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *MplsExpType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *MplsExpType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *MplsExpType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *MplsExpType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *MplsExpType) UpdateReferences() error {
    return nil
}


