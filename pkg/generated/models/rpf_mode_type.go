
package models
// RpfModeType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type RpfModeType string

// MakeRpfModeType makes RpfModeType
func MakeRpfModeType() RpfModeType {
    var data RpfModeType
    return data
}



// MakeRpfModeTypeSlice makes a slice of RpfModeType
func MakeRpfModeTypeSlice() []RpfModeType {
    return []RpfModeType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *RpfModeType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *RpfModeType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *RpfModeType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *RpfModeType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *RpfModeType) GetFQName() []string {
    return model.FQName
}

func (model *RpfModeType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *RpfModeType) GetParentType() string {
    return model.ParentType
}

func (model *RpfModeType) GetUuid() string {
    return model.UUID
}

func (model *RpfModeType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *RpfModeType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *RpfModeType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *RpfModeType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *RpfModeType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *RpfModeType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *RpfModeType) UpdateReferences() error {
    return nil
}


