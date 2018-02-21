
package models
// RoutingPolicy



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propRoutingPolicy_uuid int = iota
    propRoutingPolicy_parent_uuid int = iota
    propRoutingPolicy_parent_type int = iota
    propRoutingPolicy_fq_name int = iota
    propRoutingPolicy_id_perms int = iota
    propRoutingPolicy_display_name int = iota
    propRoutingPolicy_annotations int = iota
    propRoutingPolicy_perms2 int = iota
)

// RoutingPolicy 
type RoutingPolicy struct {

    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`

    ServiceInstanceRefs []*RoutingPolicyServiceInstanceRef `json:"service_instance_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// RoutingPolicyServiceInstanceRef references each other
type RoutingPolicyServiceInstanceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
    Attr *RoutingPolicyServiceInstanceType
    
}


// String returns json representation of the object
func (model *RoutingPolicy) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeRoutingPolicy makes RoutingPolicy
func MakeRoutingPolicy() *RoutingPolicy{
    return &RoutingPolicy{
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



// MakeRoutingPolicySlice makes a slice of RoutingPolicy
func MakeRoutingPolicySlice() []*RoutingPolicy {
    return []*RoutingPolicy{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *RoutingPolicy) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc4202e4e60])
    fqn := []string{}
    
    fqn = Project{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *RoutingPolicy) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-project", "_", "-", -1)
}

func (model *RoutingPolicy) GetDefaultName() string {
    return strings.Replace("default-routing_policy", "_", "-", -1)
}

func (model *RoutingPolicy) GetType() string {
    return strings.Replace("routing_policy", "_", "-", -1)
}

func (model *RoutingPolicy) GetFQName() []string {
    return model.FQName
}

func (model *RoutingPolicy) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *RoutingPolicy) GetParentType() string {
    return model.ParentType
}

func (model *RoutingPolicy) GetUuid() string {
    return model.UUID
}

func (model *RoutingPolicy) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *RoutingPolicy) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *RoutingPolicy) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *RoutingPolicy) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *RoutingPolicy) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propRoutingPolicy_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propRoutingPolicy_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propRoutingPolicy_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propRoutingPolicy_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propRoutingPolicy_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propRoutingPolicy_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propRoutingPolicy_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propRoutingPolicy_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *RoutingPolicy) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *RoutingPolicy) UpdateReferences() error {
    return nil
}


