
package models
// RouteType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propRouteType_prefix int = iota
    propRouteType_next_hop int = iota
    propRouteType_community_attributes int = iota
    propRouteType_next_hop_type int = iota
)

// RouteType 
type RouteType struct {

    Prefix string `json:"prefix,omitempty"`
    NextHop string `json:"next_hop,omitempty"`
    CommunityAttributes *CommunityAttributes `json:"community_attributes,omitempty"`
    NextHopType RouteNextHopType `json:"next_hop_type,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *RouteType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeRouteType makes RouteType
func MakeRouteType() *RouteType{
    return &RouteType{
    //TODO(nati): Apply default
    CommunityAttributes: MakeCommunityAttributes(),
        NextHopType: MakeRouteNextHopType(),
        Prefix: "",
        NextHop: "",
        
        modified: big.NewInt(0),
    }
}



// MakeRouteTypeSlice makes a slice of RouteType
func MakeRouteTypeSlice() []*RouteType {
    return []*RouteType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *RouteType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *RouteType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *RouteType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *RouteType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *RouteType) GetFQName() []string {
    return model.FQName
}

func (model *RouteType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *RouteType) GetParentType() string {
    return model.ParentType
}

func (model *RouteType) GetUuid() string {
    return model.UUID
}

func (model *RouteType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *RouteType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *RouteType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *RouteType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *RouteType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propRouteType_prefix) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Prefix); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Prefix as prefix")
        }
        msg["prefix"] = &val
    }
    
    if model.modified.Bit(propRouteType_next_hop) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.NextHop); err != nil {
            return nil, errors.Wrap(err, "Marshal of: NextHop as next_hop")
        }
        msg["next_hop"] = &val
    }
    
    if model.modified.Bit(propRouteType_community_attributes) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.CommunityAttributes); err != nil {
            return nil, errors.Wrap(err, "Marshal of: CommunityAttributes as community_attributes")
        }
        msg["community_attributes"] = &val
    }
    
    if model.modified.Bit(propRouteType_next_hop_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.NextHopType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: NextHopType as next_hop_type")
        }
        msg["next_hop_type"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *RouteType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *RouteType) UpdateReferences() error {
    return nil
}


