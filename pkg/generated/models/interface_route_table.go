
package models
// InterfaceRouteTable



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propInterfaceRouteTable_parent_type int = iota
    propInterfaceRouteTable_id_perms int = iota
    propInterfaceRouteTable_annotations int = iota
    propInterfaceRouteTable_perms2 int = iota
    propInterfaceRouteTable_interface_route_table_routes int = iota
    propInterfaceRouteTable_fq_name int = iota
    propInterfaceRouteTable_display_name int = iota
    propInterfaceRouteTable_uuid int = iota
    propInterfaceRouteTable_parent_uuid int = iota
)

// InterfaceRouteTable 
type InterfaceRouteTable struct {

    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    InterfaceRouteTableRoutes *RouteTableType `json:"interface_route_table_routes,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    DisplayName string `json:"display_name,omitempty"`

    ServiceInstanceRefs []*InterfaceRouteTableServiceInstanceRef `json:"service_instance_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// InterfaceRouteTableServiceInstanceRef references each other
type InterfaceRouteTableServiceInstanceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
    Attr *ServiceInterfaceTag
    
}


// String returns json representation of the object
func (model *InterfaceRouteTable) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeInterfaceRouteTable makes InterfaceRouteTable
func MakeInterfaceRouteTable() *InterfaceRouteTable{
    return &InterfaceRouteTable{
    //TODO(nati): Apply default
    Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ParentType: "",
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        UUID: "",
        ParentUUID: "",
        InterfaceRouteTableRoutes: MakeRouteTableType(),
        FQName: []string{},
        
        modified: big.NewInt(0),
    }
}



// MakeInterfaceRouteTableSlice makes a slice of InterfaceRouteTable
func MakeInterfaceRouteTableSlice() []*InterfaceRouteTable {
    return []*InterfaceRouteTable{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *InterfaceRouteTable) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc42024ac80])
    fqn := []string{}
    
    fqn = Project{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *InterfaceRouteTable) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-project", "_", "-", -1)
}

func (model *InterfaceRouteTable) GetDefaultName() string {
    return strings.Replace("default-interface_route_table", "_", "-", -1)
}

func (model *InterfaceRouteTable) GetType() string {
    return strings.Replace("interface_route_table", "_", "-", -1)
}

func (model *InterfaceRouteTable) GetFQName() []string {
    return model.FQName
}

func (model *InterfaceRouteTable) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *InterfaceRouteTable) GetParentType() string {
    return model.ParentType
}

func (model *InterfaceRouteTable) GetUuid() string {
    return model.UUID
}

func (model *InterfaceRouteTable) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *InterfaceRouteTable) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *InterfaceRouteTable) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *InterfaceRouteTable) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *InterfaceRouteTable) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propInterfaceRouteTable_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propInterfaceRouteTable_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propInterfaceRouteTable_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propInterfaceRouteTable_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propInterfaceRouteTable_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propInterfaceRouteTable_interface_route_table_routes) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.InterfaceRouteTableRoutes); err != nil {
            return nil, errors.Wrap(err, "Marshal of: InterfaceRouteTableRoutes as interface_route_table_routes")
        }
        msg["interface_route_table_routes"] = &val
    }
    
    if model.modified.Bit(propInterfaceRouteTable_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propInterfaceRouteTable_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propInterfaceRouteTable_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *InterfaceRouteTable) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *InterfaceRouteTable) UpdateReferences() error {
    return nil
}


