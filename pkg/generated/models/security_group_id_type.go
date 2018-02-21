
package models
// SecurityGroupIdType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type SecurityGroupIdType int

// MakeSecurityGroupIdType makes SecurityGroupIdType
func MakeSecurityGroupIdType() SecurityGroupIdType {
    var data SecurityGroupIdType
    return data
}



// MakeSecurityGroupIdTypeSlice makes a slice of SecurityGroupIdType
func MakeSecurityGroupIdTypeSlice() []SecurityGroupIdType {
    return []SecurityGroupIdType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *SecurityGroupIdType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *SecurityGroupIdType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *SecurityGroupIdType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *SecurityGroupIdType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *SecurityGroupIdType) GetFQName() []string {
    return model.FQName
}

func (model *SecurityGroupIdType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *SecurityGroupIdType) GetParentType() string {
    return model.ParentType
}

func (model *SecurityGroupIdType) GetUuid() string {
    return model.UUID
}

func (model *SecurityGroupIdType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *SecurityGroupIdType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *SecurityGroupIdType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *SecurityGroupIdType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *SecurityGroupIdType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *SecurityGroupIdType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *SecurityGroupIdType) UpdateReferences() error {
    return nil
}


