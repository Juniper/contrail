
package models
// PeeringServiceType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type PeeringServiceType string

// MakePeeringServiceType makes PeeringServiceType
func MakePeeringServiceType() PeeringServiceType {
    var data PeeringServiceType
    return data
}



// MakePeeringServiceTypeSlice makes a slice of PeeringServiceType
func MakePeeringServiceTypeSlice() []PeeringServiceType {
    return []PeeringServiceType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *PeeringServiceType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *PeeringServiceType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *PeeringServiceType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *PeeringServiceType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *PeeringServiceType) GetFQName() []string {
    return model.FQName
}

func (model *PeeringServiceType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *PeeringServiceType) GetParentType() string {
    return model.ParentType
}

func (model *PeeringServiceType) GetUuid() string {
    return model.UUID
}

func (model *PeeringServiceType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *PeeringServiceType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *PeeringServiceType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *PeeringServiceType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *PeeringServiceType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *PeeringServiceType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *PeeringServiceType) UpdateReferences() error {
    return nil
}


