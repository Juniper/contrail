
package models
// PhysicalRouterRole


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type PhysicalRouterRole string

// MakePhysicalRouterRole makes PhysicalRouterRole
func MakePhysicalRouterRole() PhysicalRouterRole {
    var data PhysicalRouterRole
    return data
}



// MakePhysicalRouterRoleSlice makes a slice of PhysicalRouterRole
func MakePhysicalRouterRoleSlice() []PhysicalRouterRole {
    return []PhysicalRouterRole{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *PhysicalRouterRole) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *PhysicalRouterRole) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *PhysicalRouterRole) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *PhysicalRouterRole) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *PhysicalRouterRole) GetFQName() []string {
    return model.FQName
}

func (model *PhysicalRouterRole) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *PhysicalRouterRole) GetParentType() string {
    return model.ParentType
}

func (model *PhysicalRouterRole) GetUuid() string {
    return model.UUID
}

func (model *PhysicalRouterRole) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *PhysicalRouterRole) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *PhysicalRouterRole) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *PhysicalRouterRole) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *PhysicalRouterRole) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *PhysicalRouterRole) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *PhysicalRouterRole) UpdateReferences() error {
    return nil
}


