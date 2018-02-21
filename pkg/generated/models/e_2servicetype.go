
package models
// E2servicetype


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type E2servicetype string

// MakeE2servicetype makes E2servicetype
func MakeE2servicetype() E2servicetype {
    var data E2servicetype
    return data
}



// MakeE2servicetypeSlice makes a slice of E2servicetype
func MakeE2servicetypeSlice() []E2servicetype {
    return []E2servicetype{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *E2servicetype) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *E2servicetype) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *E2servicetype) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *E2servicetype) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *E2servicetype) GetFQName() []string {
    return model.FQName
}

func (model *E2servicetype) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *E2servicetype) GetParentType() string {
    return model.ParentType
}

func (model *E2servicetype) GetUuid() string {
    return model.UUID
}

func (model *E2servicetype) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *E2servicetype) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *E2servicetype) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *E2servicetype) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *E2servicetype) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *E2servicetype) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *E2servicetype) UpdateReferences() error {
    return nil
}


