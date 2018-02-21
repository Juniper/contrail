
package models
// EndOfRibTimeType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type EndOfRibTimeType int

// MakeEndOfRibTimeType makes EndOfRibTimeType
func MakeEndOfRibTimeType() EndOfRibTimeType {
    var data EndOfRibTimeType
    return data
}



// MakeEndOfRibTimeTypeSlice makes a slice of EndOfRibTimeType
func MakeEndOfRibTimeTypeSlice() []EndOfRibTimeType {
    return []EndOfRibTimeType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *EndOfRibTimeType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *EndOfRibTimeType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *EndOfRibTimeType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *EndOfRibTimeType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *EndOfRibTimeType) GetFQName() []string {
    return model.FQName
}

func (model *EndOfRibTimeType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *EndOfRibTimeType) GetParentType() string {
    return model.ParentType
}

func (model *EndOfRibTimeType) GetUuid() string {
    return model.UUID
}

func (model *EndOfRibTimeType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *EndOfRibTimeType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *EndOfRibTimeType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *EndOfRibTimeType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *EndOfRibTimeType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *EndOfRibTimeType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *EndOfRibTimeType) UpdateReferences() error {
    return nil
}


