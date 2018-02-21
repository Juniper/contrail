
package models
// LoadbalancerMember



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propLoadbalancerMember_perms2 int = iota
    propLoadbalancerMember_parent_uuid int = iota
    propLoadbalancerMember_fq_name int = iota
    propLoadbalancerMember_display_name int = iota
    propLoadbalancerMember_annotations int = iota
    propLoadbalancerMember_loadbalancer_member_properties int = iota
    propLoadbalancerMember_uuid int = iota
    propLoadbalancerMember_parent_type int = iota
    propLoadbalancerMember_id_perms int = iota
)

// LoadbalancerMember 
type LoadbalancerMember struct {

    ParentType string `json:"parent_type,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    LoadbalancerMemberProperties *LoadbalancerMemberType `json:"loadbalancer_member_properties,omitempty"`
    UUID string `json:"uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *LoadbalancerMember) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeLoadbalancerMember makes LoadbalancerMember
func MakeLoadbalancerMember() *LoadbalancerMember{
    return &LoadbalancerMember{
    //TODO(nati): Apply default
    Perms2: MakePermType2(),
        ParentUUID: "",
        FQName: []string{},
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        LoadbalancerMemberProperties: MakeLoadbalancerMemberType(),
        UUID: "",
        ParentType: "",
        IDPerms: MakeIdPermsType(),
        
        modified: big.NewInt(0),
    }
}



// MakeLoadbalancerMemberSlice makes a slice of LoadbalancerMember
func MakeLoadbalancerMemberSlice() []*LoadbalancerMember {
    return []*LoadbalancerMember{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *LoadbalancerMember) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[loadbalancer_pool:0xc42024af00])
    fqn := []string{}
    
    fqn = LoadbalancerPool{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *LoadbalancerMember) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-loadbalancer_pool", "_", "-", -1)
}

func (model *LoadbalancerMember) GetDefaultName() string {
    return strings.Replace("default-loadbalancer_member", "_", "-", -1)
}

func (model *LoadbalancerMember) GetType() string {
    return strings.Replace("loadbalancer_member", "_", "-", -1)
}

func (model *LoadbalancerMember) GetFQName() []string {
    return model.FQName
}

func (model *LoadbalancerMember) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *LoadbalancerMember) GetParentType() string {
    return model.ParentType
}

func (model *LoadbalancerMember) GetUuid() string {
    return model.UUID
}

func (model *LoadbalancerMember) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *LoadbalancerMember) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *LoadbalancerMember) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *LoadbalancerMember) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *LoadbalancerMember) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propLoadbalancerMember_loadbalancer_member_properties) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LoadbalancerMemberProperties); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LoadbalancerMemberProperties as loadbalancer_member_properties")
        }
        msg["loadbalancer_member_properties"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerMember_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerMember_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerMember_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerMember_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerMember_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerMember_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerMember_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerMember_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *LoadbalancerMember) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *LoadbalancerMember) UpdateReferences() error {
    return nil
}


