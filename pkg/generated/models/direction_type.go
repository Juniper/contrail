
package models
// DirectionType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type DirectionType string

// MakeDirectionType makes DirectionType
func MakeDirectionType() DirectionType {
    var data DirectionType
    return data
}



// MakeDirectionTypeSlice makes a slice of DirectionType
func MakeDirectionTypeSlice() []DirectionType {
    return []DirectionType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *DirectionType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *DirectionType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *DirectionType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *DirectionType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *DirectionType) GetFQName() []string {
    return model.FQName
}

func (model *DirectionType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *DirectionType) GetParentType() string {
    return model.ParentType
}

func (model *DirectionType) GetUuid() string {
    return model.UUID
}

func (model *DirectionType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *DirectionType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *DirectionType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *DirectionType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *DirectionType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *DirectionType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *DirectionType) UpdateReferences() error {
    return nil
}


