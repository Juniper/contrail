
package models
// CommunityAttribute


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type CommunityAttribute []string

// MakeCommunityAttribute makes CommunityAttribute
func MakeCommunityAttribute() CommunityAttribute {
    var data CommunityAttribute
    return data
}



// MakeCommunityAttributeSlice makes a slice of CommunityAttribute
func MakeCommunityAttributeSlice() []CommunityAttribute {
    return []CommunityAttribute{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *CommunityAttribute) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *CommunityAttribute) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *CommunityAttribute) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *CommunityAttribute) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *CommunityAttribute) GetFQName() []string {
    return model.FQName
}

func (model *CommunityAttribute) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *CommunityAttribute) GetParentType() string {
    return model.ParentType
}

func (model *CommunityAttribute) GetUuid() string {
    return model.UUID
}

func (model *CommunityAttribute) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *CommunityAttribute) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *CommunityAttribute) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *CommunityAttribute) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *CommunityAttribute) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *CommunityAttribute) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *CommunityAttribute) UpdateReferences() error {
    return nil
}


