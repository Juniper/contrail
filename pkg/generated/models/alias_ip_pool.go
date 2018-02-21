
package models
// AliasIPPool



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propAliasIPPool_parent_type int = iota
    propAliasIPPool_fq_name int = iota
    propAliasIPPool_id_perms int = iota
    propAliasIPPool_display_name int = iota
    propAliasIPPool_annotations int = iota
    propAliasIPPool_perms2 int = iota
    propAliasIPPool_uuid int = iota
    propAliasIPPool_parent_uuid int = iota
)

// AliasIPPool 
type AliasIPPool struct {

    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`


    AliasIPs []*AliasIP `json:"alias_ips,omitempty"`

    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *AliasIPPool) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeAliasIPPool makes AliasIPPool
func MakeAliasIPPool() *AliasIPPool{
    return &AliasIPPool{
    //TODO(nati): Apply default
    ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        UUID: "",
        
        modified: big.NewInt(0),
    }
}



// MakeAliasIPPoolSlice makes a slice of AliasIPPool
func MakeAliasIPPoolSlice() []*AliasIPPool {
    return []*AliasIPPool{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *AliasIPPool) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[virtual_network:0xc420182b40])
    fqn := []string{}
    
    fqn = VirtualNetwork{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *AliasIPPool) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-virtual_network", "_", "-", -1)
}

func (model *AliasIPPool) GetDefaultName() string {
    return strings.Replace("default-alias_ip_pool", "_", "-", -1)
}

func (model *AliasIPPool) GetType() string {
    return strings.Replace("alias_ip_pool", "_", "-", -1)
}

func (model *AliasIPPool) GetFQName() []string {
    return model.FQName
}

func (model *AliasIPPool) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *AliasIPPool) GetParentType() string {
    return model.ParentType
}

func (model *AliasIPPool) GetUuid() string {
    return model.UUID
}

func (model *AliasIPPool) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *AliasIPPool) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *AliasIPPool) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *AliasIPPool) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *AliasIPPool) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propAliasIPPool_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propAliasIPPool_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propAliasIPPool_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propAliasIPPool_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propAliasIPPool_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propAliasIPPool_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propAliasIPPool_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propAliasIPPool_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *AliasIPPool) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *AliasIPPool) UpdateReferences() error {
    return nil
}


