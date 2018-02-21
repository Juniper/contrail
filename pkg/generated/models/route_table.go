
package models
// RouteTable



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propRouteTable_parent_type int = iota
    propRouteTable_fq_name int = iota
    propRouteTable_display_name int = iota
    propRouteTable_annotations int = iota
    propRouteTable_uuid int = iota
    propRouteTable_perms2 int = iota
    propRouteTable_parent_uuid int = iota
    propRouteTable_id_perms int = iota
    propRouteTable_routes int = iota
)

// RouteTable 
type RouteTable struct {

    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    Routes *RouteTableType `json:"routes,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *RouteTable) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeRouteTable makes RouteTable
func MakeRouteTable() *RouteTable{
    return &RouteTable{
    //TODO(nati): Apply default
    FQName: []string{},
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        UUID: "",
        ParentType: "",
        ParentUUID: "",
        IDPerms: MakeIdPermsType(),
        Routes: MakeRouteTableType(),
        Perms2: MakePermType2(),
        
        modified: big.NewInt(0),
    }
}



// MakeRouteTableSlice makes a slice of RouteTable
func MakeRouteTableSlice() []*RouteTable {
    return []*RouteTable{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *RouteTable) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc4202e4c80])
    fqn := []string{}
    
    fqn = Project{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *RouteTable) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-project", "_", "-", -1)
}

func (model *RouteTable) GetDefaultName() string {
    return strings.Replace("default-route_table", "_", "-", -1)
}

func (model *RouteTable) GetType() string {
    return strings.Replace("route_table", "_", "-", -1)
}

func (model *RouteTable) GetFQName() []string {
    return model.FQName
}

func (model *RouteTable) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *RouteTable) GetParentType() string {
    return model.ParentType
}

func (model *RouteTable) GetUuid() string {
    return model.UUID
}

func (model *RouteTable) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *RouteTable) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *RouteTable) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *RouteTable) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *RouteTable) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propRouteTable_routes) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Routes); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Routes as routes")
        }
        msg["routes"] = &val
    }
    
    if model.modified.Bit(propRouteTable_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propRouteTable_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propRouteTable_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propRouteTable_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propRouteTable_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propRouteTable_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propRouteTable_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propRouteTable_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *RouteTable) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *RouteTable) UpdateReferences() error {
    return nil
}


