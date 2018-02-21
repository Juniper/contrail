
package models
// NHModeType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type NHModeType string

// MakeNHModeType makes NHModeType
func MakeNHModeType() NHModeType {
    var data NHModeType
    return data
}



// MakeNHModeTypeSlice makes a slice of NHModeType
func MakeNHModeTypeSlice() []NHModeType {
    return []NHModeType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *NHModeType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *NHModeType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *NHModeType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *NHModeType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *NHModeType) GetFQName() []string {
    return model.FQName
}

func (model *NHModeType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *NHModeType) GetParentType() string {
    return model.ParentType
}

func (model *NHModeType) GetUuid() string {
    return model.UUID
}

func (model *NHModeType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *NHModeType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *NHModeType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *NHModeType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *NHModeType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *NHModeType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *NHModeType) UpdateReferences() error {
    return nil
}


