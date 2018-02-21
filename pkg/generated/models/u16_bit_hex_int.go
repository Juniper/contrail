
package models
// U16BitHexInt


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type U16BitHexInt string

// MakeU16BitHexInt makes U16BitHexInt
func MakeU16BitHexInt() U16BitHexInt {
    var data U16BitHexInt
    return data
}



// MakeU16BitHexIntSlice makes a slice of U16BitHexInt
func MakeU16BitHexIntSlice() []U16BitHexInt {
    return []U16BitHexInt{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *U16BitHexInt) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *U16BitHexInt) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *U16BitHexInt) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *U16BitHexInt) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *U16BitHexInt) GetFQName() []string {
    return model.FQName
}

func (model *U16BitHexInt) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *U16BitHexInt) GetParentType() string {
    return model.ParentType
}

func (model *U16BitHexInt) GetUuid() string {
    return model.UUID
}

func (model *U16BitHexInt) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *U16BitHexInt) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *U16BitHexInt) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *U16BitHexInt) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *U16BitHexInt) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *U16BitHexInt) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *U16BitHexInt) UpdateReferences() error {
    return nil
}


