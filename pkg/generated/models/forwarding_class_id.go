
package models
// ForwardingClassId


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type ForwardingClassId int

// MakeForwardingClassId makes ForwardingClassId
func MakeForwardingClassId() ForwardingClassId {
    var data ForwardingClassId
    return data
}



// MakeForwardingClassIdSlice makes a slice of ForwardingClassId
func MakeForwardingClassIdSlice() []ForwardingClassId {
    return []ForwardingClassId{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ForwardingClassId) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *ForwardingClassId) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ForwardingClassId) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *ForwardingClassId) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *ForwardingClassId) GetFQName() []string {
    return model.FQName
}

func (model *ForwardingClassId) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ForwardingClassId) GetParentType() string {
    return model.ParentType
}

func (model *ForwardingClassId) GetUuid() string {
    return model.UUID
}

func (model *ForwardingClassId) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ForwardingClassId) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ForwardingClassId) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ForwardingClassId) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ForwardingClassId) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *ForwardingClassId) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ForwardingClassId) UpdateReferences() error {
    return nil
}


