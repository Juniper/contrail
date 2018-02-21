
package models
// U32BitHexInt


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type U32BitHexInt string

// MakeU32BitHexInt makes U32BitHexInt
func MakeU32BitHexInt() U32BitHexInt {
    var data U32BitHexInt
    return data
}



// MakeU32BitHexIntSlice makes a slice of U32BitHexInt
func MakeU32BitHexIntSlice() []U32BitHexInt {
    return []U32BitHexInt{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *U32BitHexInt) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *U32BitHexInt) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *U32BitHexInt) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *U32BitHexInt) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *U32BitHexInt) GetFQName() []string {
    return model.FQName
}

func (model *U32BitHexInt) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *U32BitHexInt) GetParentType() string {
    return model.ParentType
}

func (model *U32BitHexInt) GetUuid() string {
    return model.UUID
}

func (model *U32BitHexInt) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *U32BitHexInt) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *U32BitHexInt) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *U32BitHexInt) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *U32BitHexInt) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *U32BitHexInt) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *U32BitHexInt) UpdateReferences() error {
    return nil
}


