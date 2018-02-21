
package models
// MACAgingTime


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type MACAgingTime int

// MakeMACAgingTime makes MACAgingTime
func MakeMACAgingTime() MACAgingTime {
    var data MACAgingTime
    return data
}



// MakeMACAgingTimeSlice makes a slice of MACAgingTime
func MakeMACAgingTimeSlice() []MACAgingTime {
    return []MACAgingTime{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *MACAgingTime) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *MACAgingTime) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *MACAgingTime) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *MACAgingTime) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *MACAgingTime) GetFQName() []string {
    return model.FQName
}

func (model *MACAgingTime) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *MACAgingTime) GetParentType() string {
    return model.ParentType
}

func (model *MACAgingTime) GetUuid() string {
    return model.UUID
}

func (model *MACAgingTime) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *MACAgingTime) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *MACAgingTime) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *MACAgingTime) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *MACAgingTime) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *MACAgingTime) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *MACAgingTime) UpdateReferences() error {
    return nil
}


