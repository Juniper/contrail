
package models
// Namespace



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propNamespace_uuid int = iota
    propNamespace_parent_uuid int = iota
    propNamespace_parent_type int = iota
    propNamespace_annotations int = iota
    propNamespace_id_perms int = iota
    propNamespace_display_name int = iota
    propNamespace_perms2 int = iota
    propNamespace_fq_name int = iota
    propNamespace_namespace_cidr int = iota
)

// Namespace 
type Namespace struct {

    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    NamespaceCidr *SubnetType `json:"namespace_cidr,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    FQName []string `json:"fq_name,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *Namespace) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeNamespace makes Namespace
func MakeNamespace() *Namespace{
    return &Namespace{
    //TODO(nati): Apply default
    Perms2: MakePermType2(),
        FQName: []string{},
        NamespaceCidr: MakeSubnetType(),
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        ParentType: "",
        Annotations: MakeKeyValuePairs(),
        UUID: "",
        ParentUUID: "",
        
        modified: big.NewInt(0),
    }
}



// MakeNamespaceSlice makes a slice of Namespace
func MakeNamespaceSlice() []*Namespace {
    return []*Namespace{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *Namespace) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[domain:0xc42024bcc0])
    fqn := []string{}
    
    fqn = Domain{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *Namespace) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-domain", "_", "-", -1)
}

func (model *Namespace) GetDefaultName() string {
    return strings.Replace("default-namespace", "_", "-", -1)
}

func (model *Namespace) GetType() string {
    return strings.Replace("namespace", "_", "-", -1)
}

func (model *Namespace) GetFQName() []string {
    return model.FQName
}

func (model *Namespace) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *Namespace) GetParentType() string {
    return model.ParentType
}

func (model *Namespace) GetUuid() string {
    return model.UUID
}

func (model *Namespace) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *Namespace) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *Namespace) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *Namespace) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *Namespace) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propNamespace_namespace_cidr) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.NamespaceCidr); err != nil {
            return nil, errors.Wrap(err, "Marshal of: NamespaceCidr as namespace_cidr")
        }
        msg["namespace_cidr"] = &val
    }
    
    if model.modified.Bit(propNamespace_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propNamespace_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propNamespace_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propNamespace_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propNamespace_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propNamespace_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propNamespace_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propNamespace_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *Namespace) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *Namespace) UpdateReferences() error {
    return nil
}


