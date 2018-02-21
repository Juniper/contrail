
package models
// SessionPersistenceType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type SessionPersistenceType string

// MakeSessionPersistenceType makes SessionPersistenceType
func MakeSessionPersistenceType() SessionPersistenceType {
    var data SessionPersistenceType
    return data
}



// MakeSessionPersistenceTypeSlice makes a slice of SessionPersistenceType
func MakeSessionPersistenceTypeSlice() []SessionPersistenceType {
    return []SessionPersistenceType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *SessionPersistenceType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *SessionPersistenceType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *SessionPersistenceType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *SessionPersistenceType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *SessionPersistenceType) GetFQName() []string {
    return model.FQName
}

func (model *SessionPersistenceType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *SessionPersistenceType) GetParentType() string {
    return model.ParentType
}

func (model *SessionPersistenceType) GetUuid() string {
    return model.UUID
}

func (model *SessionPersistenceType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *SessionPersistenceType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *SessionPersistenceType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *SessionPersistenceType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *SessionPersistenceType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *SessionPersistenceType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *SessionPersistenceType) UpdateReferences() error {
    return nil
}


