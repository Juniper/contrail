
package models
// LoadbalancerMethodType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type LoadbalancerMethodType string

// MakeLoadbalancerMethodType makes LoadbalancerMethodType
func MakeLoadbalancerMethodType() LoadbalancerMethodType {
    var data LoadbalancerMethodType
    return data
}



// MakeLoadbalancerMethodTypeSlice makes a slice of LoadbalancerMethodType
func MakeLoadbalancerMethodTypeSlice() []LoadbalancerMethodType {
    return []LoadbalancerMethodType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *LoadbalancerMethodType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *LoadbalancerMethodType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *LoadbalancerMethodType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *LoadbalancerMethodType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *LoadbalancerMethodType) GetFQName() []string {
    return model.FQName
}

func (model *LoadbalancerMethodType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *LoadbalancerMethodType) GetParentType() string {
    return model.ParentType
}

func (model *LoadbalancerMethodType) GetUuid() string {
    return model.UUID
}

func (model *LoadbalancerMethodType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *LoadbalancerMethodType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *LoadbalancerMethodType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *LoadbalancerMethodType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *LoadbalancerMethodType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *LoadbalancerMethodType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *LoadbalancerMethodType) UpdateReferences() error {
    return nil
}


