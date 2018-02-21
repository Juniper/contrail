
package models
// TrafficDirectionType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type TrafficDirectionType string

// MakeTrafficDirectionType makes TrafficDirectionType
func MakeTrafficDirectionType() TrafficDirectionType {
    var data TrafficDirectionType
    return data
}



// MakeTrafficDirectionTypeSlice makes a slice of TrafficDirectionType
func MakeTrafficDirectionTypeSlice() []TrafficDirectionType {
    return []TrafficDirectionType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *TrafficDirectionType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *TrafficDirectionType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *TrafficDirectionType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *TrafficDirectionType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *TrafficDirectionType) GetFQName() []string {
    return model.FQName
}

func (model *TrafficDirectionType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *TrafficDirectionType) GetParentType() string {
    return model.ParentType
}

func (model *TrafficDirectionType) GetUuid() string {
    return model.UUID
}

func (model *TrafficDirectionType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *TrafficDirectionType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *TrafficDirectionType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *TrafficDirectionType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *TrafficDirectionType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *TrafficDirectionType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *TrafficDirectionType) UpdateReferences() error {
    return nil
}


