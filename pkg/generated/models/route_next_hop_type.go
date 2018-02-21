
package models
// RouteNextHopType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type RouteNextHopType string

// MakeRouteNextHopType makes RouteNextHopType
func MakeRouteNextHopType() RouteNextHopType {
    var data RouteNextHopType
    return data
}



// MakeRouteNextHopTypeSlice makes a slice of RouteNextHopType
func MakeRouteNextHopTypeSlice() []RouteNextHopType {
    return []RouteNextHopType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *RouteNextHopType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *RouteNextHopType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *RouteNextHopType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *RouteNextHopType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *RouteNextHopType) GetFQName() []string {
    return model.FQName
}

func (model *RouteNextHopType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *RouteNextHopType) GetParentType() string {
    return model.ParentType
}

func (model *RouteNextHopType) GetUuid() string {
    return model.UUID
}

func (model *RouteNextHopType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *RouteNextHopType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *RouteNextHopType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *RouteNextHopType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *RouteNextHopType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *RouteNextHopType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *RouteNextHopType) UpdateReferences() error {
    return nil
}


