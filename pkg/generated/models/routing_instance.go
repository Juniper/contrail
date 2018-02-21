
package models
// RoutingInstance



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propRoutingInstance_parent_type int = iota
    propRoutingInstance_fq_name int = iota
    propRoutingInstance_id_perms int = iota
    propRoutingInstance_display_name int = iota
    propRoutingInstance_annotations int = iota
    propRoutingInstance_perms2 int = iota
    propRoutingInstance_uuid int = iota
    propRoutingInstance_parent_uuid int = iota
)

// RoutingInstance 
type RoutingInstance struct {

    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *RoutingInstance) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeRoutingInstance makes RoutingInstance
func MakeRoutingInstance() *RoutingInstance{
    return &RoutingInstance{
    //TODO(nati): Apply default
    Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        
        modified: big.NewInt(0),
    }
}



// MakeRoutingInstanceSlice makes a slice of RoutingInstance
func MakeRoutingInstanceSlice() []*RoutingInstance {
    return []*RoutingInstance{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *RoutingInstance) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[virtual_network:0xc4202e4d20])
    fqn := []string{}
    
    fqn = VirtualNetwork{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *RoutingInstance) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-virtual_network", "_", "-", -1)
}

func (model *RoutingInstance) GetDefaultName() string {
    return strings.Replace("default-routing_instance", "_", "-", -1)
}

func (model *RoutingInstance) GetType() string {
    return strings.Replace("routing_instance", "_", "-", -1)
}

func (model *RoutingInstance) GetFQName() []string {
    return model.FQName
}

func (model *RoutingInstance) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *RoutingInstance) GetParentType() string {
    return model.ParentType
}

func (model *RoutingInstance) GetUuid() string {
    return model.UUID
}

func (model *RoutingInstance) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *RoutingInstance) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *RoutingInstance) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *RoutingInstance) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *RoutingInstance) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propRoutingInstance_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propRoutingInstance_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propRoutingInstance_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propRoutingInstance_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propRoutingInstance_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propRoutingInstance_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propRoutingInstance_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propRoutingInstance_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *RoutingInstance) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *RoutingInstance) UpdateReferences() error {
    return nil
}


