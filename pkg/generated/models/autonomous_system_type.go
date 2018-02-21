
package models
// AutonomousSystemType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type AutonomousSystemType int

// MakeAutonomousSystemType makes AutonomousSystemType
func MakeAutonomousSystemType() AutonomousSystemType {
    var data AutonomousSystemType
    return data
}



// MakeAutonomousSystemTypeSlice makes a slice of AutonomousSystemType
func MakeAutonomousSystemTypeSlice() []AutonomousSystemType {
    return []AutonomousSystemType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *AutonomousSystemType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *AutonomousSystemType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *AutonomousSystemType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *AutonomousSystemType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *AutonomousSystemType) GetFQName() []string {
    return model.FQName
}

func (model *AutonomousSystemType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *AutonomousSystemType) GetParentType() string {
    return model.ParentType
}

func (model *AutonomousSystemType) GetUuid() string {
    return model.UUID
}

func (model *AutonomousSystemType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *AutonomousSystemType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *AutonomousSystemType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *AutonomousSystemType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *AutonomousSystemType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *AutonomousSystemType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *AutonomousSystemType) UpdateReferences() error {
    return nil
}


