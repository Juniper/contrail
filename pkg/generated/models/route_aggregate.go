
package models
// RouteAggregate



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propRouteAggregate_uuid int = iota
    propRouteAggregate_parent_uuid int = iota
    propRouteAggregate_parent_type int = iota
    propRouteAggregate_fq_name int = iota
    propRouteAggregate_id_perms int = iota
    propRouteAggregate_display_name int = iota
    propRouteAggregate_annotations int = iota
    propRouteAggregate_perms2 int = iota
)

// RouteAggregate 
type RouteAggregate struct {

    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`

    ServiceInstanceRefs []*RouteAggregateServiceInstanceRef `json:"service_instance_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// RouteAggregateServiceInstanceRef references each other
type RouteAggregateServiceInstanceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
    Attr *ServiceInterfaceTag
    
}


// String returns json representation of the object
func (model *RouteAggregate) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeRouteAggregate makes RouteAggregate
func MakeRouteAggregate() *RouteAggregate{
    return &RouteAggregate{
    //TODO(nati): Apply default
    FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        
        modified: big.NewInt(0),
    }
}



// MakeRouteAggregateSlice makes a slice of RouteAggregate
func MakeRouteAggregateSlice() []*RouteAggregate {
    return []*RouteAggregate{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *RouteAggregate) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc4202e4be0])
    fqn := []string{}
    
    fqn = Project{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *RouteAggregate) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-project", "_", "-", -1)
}

func (model *RouteAggregate) GetDefaultName() string {
    return strings.Replace("default-route_aggregate", "_", "-", -1)
}

func (model *RouteAggregate) GetType() string {
    return strings.Replace("route_aggregate", "_", "-", -1)
}

func (model *RouteAggregate) GetFQName() []string {
    return model.FQName
}

func (model *RouteAggregate) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *RouteAggregate) GetParentType() string {
    return model.ParentType
}

func (model *RouteAggregate) GetUuid() string {
    return model.UUID
}

func (model *RouteAggregate) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *RouteAggregate) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *RouteAggregate) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *RouteAggregate) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *RouteAggregate) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propRouteAggregate_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propRouteAggregate_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propRouteAggregate_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propRouteAggregate_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propRouteAggregate_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propRouteAggregate_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propRouteAggregate_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propRouteAggregate_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *RouteAggregate) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *RouteAggregate) UpdateReferences() error {
    return nil
}


