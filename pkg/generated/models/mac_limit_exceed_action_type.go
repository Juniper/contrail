
package models
// MACLimitExceedActionType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type MACLimitExceedActionType string

// MakeMACLimitExceedActionType makes MACLimitExceedActionType
func MakeMACLimitExceedActionType() MACLimitExceedActionType {
    var data MACLimitExceedActionType
    return data
}



// MakeMACLimitExceedActionTypeSlice makes a slice of MACLimitExceedActionType
func MakeMACLimitExceedActionTypeSlice() []MACLimitExceedActionType {
    return []MACLimitExceedActionType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *MACLimitExceedActionType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *MACLimitExceedActionType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *MACLimitExceedActionType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *MACLimitExceedActionType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *MACLimitExceedActionType) GetFQName() []string {
    return model.FQName
}

func (model *MACLimitExceedActionType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *MACLimitExceedActionType) GetParentType() string {
    return model.ParentType
}

func (model *MACLimitExceedActionType) GetUuid() string {
    return model.UUID
}

func (model *MACLimitExceedActionType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *MACLimitExceedActionType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *MACLimitExceedActionType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *MACLimitExceedActionType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *MACLimitExceedActionType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *MACLimitExceedActionType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *MACLimitExceedActionType) UpdateReferences() error {
    return nil
}


