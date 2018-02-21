
package models
// RouteTargetList



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propRouteTargetList_route_target int = iota
)

// RouteTargetList 
type RouteTargetList struct {

    RouteTarget []string `json:"route_target,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *RouteTargetList) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeRouteTargetList makes RouteTargetList
func MakeRouteTargetList() *RouteTargetList{
    return &RouteTargetList{
    //TODO(nati): Apply default
    RouteTarget: []string{},
        
        modified: big.NewInt(0),
    }
}



// MakeRouteTargetListSlice makes a slice of RouteTargetList
func MakeRouteTargetListSlice() []*RouteTargetList {
    return []*RouteTargetList{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *RouteTargetList) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *RouteTargetList) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *RouteTargetList) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *RouteTargetList) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *RouteTargetList) GetFQName() []string {
    return model.FQName
}

func (model *RouteTargetList) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *RouteTargetList) GetParentType() string {
    return model.ParentType
}

func (model *RouteTargetList) GetUuid() string {
    return model.UUID
}

func (model *RouteTargetList) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *RouteTargetList) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *RouteTargetList) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *RouteTargetList) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *RouteTargetList) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propRouteTargetList_route_target) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RouteTarget); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RouteTarget as route_target")
        }
        msg["route_target"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *RouteTargetList) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *RouteTargetList) UpdateReferences() error {
    return nil
}


