
package models
// LogicalInterfaceType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type LogicalInterfaceType string

// MakeLogicalInterfaceType makes LogicalInterfaceType
func MakeLogicalInterfaceType() LogicalInterfaceType {
    var data LogicalInterfaceType
    return data
}



// MakeLogicalInterfaceTypeSlice makes a slice of LogicalInterfaceType
func MakeLogicalInterfaceTypeSlice() []LogicalInterfaceType {
    return []LogicalInterfaceType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *LogicalInterfaceType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *LogicalInterfaceType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *LogicalInterfaceType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *LogicalInterfaceType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *LogicalInterfaceType) GetFQName() []string {
    return model.FQName
}

func (model *LogicalInterfaceType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *LogicalInterfaceType) GetParentType() string {
    return model.ParentType
}

func (model *LogicalInterfaceType) GetUuid() string {
    return model.UUID
}

func (model *LogicalInterfaceType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *LogicalInterfaceType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *LogicalInterfaceType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *LogicalInterfaceType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *LogicalInterfaceType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *LogicalInterfaceType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *LogicalInterfaceType) UpdateReferences() error {
    return nil
}


