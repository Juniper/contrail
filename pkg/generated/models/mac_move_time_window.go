
package models
// MACMoveTimeWindow


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type MACMoveTimeWindow int

// MakeMACMoveTimeWindow makes MACMoveTimeWindow
func MakeMACMoveTimeWindow() MACMoveTimeWindow {
    var data MACMoveTimeWindow
    return data
}



// MakeMACMoveTimeWindowSlice makes a slice of MACMoveTimeWindow
func MakeMACMoveTimeWindowSlice() []MACMoveTimeWindow {
    return []MACMoveTimeWindow{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *MACMoveTimeWindow) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *MACMoveTimeWindow) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *MACMoveTimeWindow) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *MACMoveTimeWindow) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *MACMoveTimeWindow) GetFQName() []string {
    return model.FQName
}

func (model *MACMoveTimeWindow) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *MACMoveTimeWindow) GetParentType() string {
    return model.ParentType
}

func (model *MACMoveTimeWindow) GetUuid() string {
    return model.UUID
}

func (model *MACMoveTimeWindow) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *MACMoveTimeWindow) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *MACMoveTimeWindow) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *MACMoveTimeWindow) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *MACMoveTimeWindow) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *MACMoveTimeWindow) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *MACMoveTimeWindow) UpdateReferences() error {
    return nil
}


