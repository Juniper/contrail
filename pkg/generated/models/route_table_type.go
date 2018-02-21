
package models
// RouteTableType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propRouteTableType_route int = iota
)

// RouteTableType 
type RouteTableType struct {

    Route []*RouteType `json:"route,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *RouteTableType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeRouteTableType makes RouteTableType
func MakeRouteTableType() *RouteTableType{
    return &RouteTableType{
    //TODO(nati): Apply default
    
            
                Route:  MakeRouteTypeSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeRouteTableTypeSlice makes a slice of RouteTableType
func MakeRouteTableTypeSlice() []*RouteTableType {
    return []*RouteTableType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *RouteTableType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *RouteTableType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *RouteTableType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *RouteTableType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *RouteTableType) GetFQName() []string {
    return model.FQName
}

func (model *RouteTableType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *RouteTableType) GetParentType() string {
    return model.ParentType
}

func (model *RouteTableType) GetUuid() string {
    return model.UUID
}

func (model *RouteTableType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *RouteTableType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *RouteTableType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *RouteTableType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *RouteTableType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propRouteTableType_route) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Route); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Route as route")
        }
        msg["route"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *RouteTableType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *RouteTableType) UpdateReferences() error {
    return nil
}


