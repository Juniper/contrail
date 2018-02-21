
package models
// EtherType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type EtherType string

// MakeEtherType makes EtherType
func MakeEtherType() EtherType {
    var data EtherType
    return data
}



// MakeEtherTypeSlice makes a slice of EtherType
func MakeEtherTypeSlice() []EtherType {
    return []EtherType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *EtherType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *EtherType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *EtherType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *EtherType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *EtherType) GetFQName() []string {
    return model.FQName
}

func (model *EtherType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *EtherType) GetParentType() string {
    return model.ParentType
}

func (model *EtherType) GetUuid() string {
    return model.UUID
}

func (model *EtherType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *EtherType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *EtherType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *EtherType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *EtherType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *EtherType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *EtherType) UpdateReferences() error {
    return nil
}


