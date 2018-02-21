
package models
// Keypair



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propKeypair_name int = iota
    propKeypair_parent_type int = iota
    propKeypair_perms2 int = iota
    propKeypair_uuid int = iota
    propKeypair_parent_uuid int = iota
    propKeypair_public_key int = iota
    propKeypair_fq_name int = iota
    propKeypair_id_perms int = iota
    propKeypair_display_name int = iota
    propKeypair_annotations int = iota
)

// Keypair 
type Keypair struct {

    Name string `json:"name,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    PublicKey string `json:"public_key,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *Keypair) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeKeypair makes Keypair
func MakeKeypair() *Keypair{
    return &Keypair{
    //TODO(nati): Apply default
    Perms2: MakePermType2(),
        UUID: "",
        Name: "",
        ParentType: "",
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        ParentUUID: "",
        PublicKey: "",
        FQName: []string{},
        
        modified: big.NewInt(0),
    }
}



// MakeKeypairSlice makes a slice of Keypair
func MakeKeypairSlice() []*Keypair {
    return []*Keypair{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *Keypair) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *Keypair) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *Keypair) GetDefaultName() string {
    return strings.Replace("default-keypair", "_", "-", -1)
}

func (model *Keypair) GetType() string {
    return strings.Replace("keypair", "_", "-", -1)
}

func (model *Keypair) GetFQName() []string {
    return model.FQName
}

func (model *Keypair) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *Keypair) GetParentType() string {
    return model.ParentType
}

func (model *Keypair) GetUuid() string {
    return model.UUID
}

func (model *Keypair) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *Keypair) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *Keypair) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *Keypair) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *Keypair) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propKeypair_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propKeypair_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propKeypair_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propKeypair_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Name); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Name as name")
        }
        msg["name"] = &val
    }
    
    if model.modified.Bit(propKeypair_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propKeypair_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propKeypair_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propKeypair_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propKeypair_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propKeypair_public_key) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.PublicKey); err != nil {
            return nil, errors.Wrap(err, "Marshal of: PublicKey as public_key")
        }
        msg["public_key"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *Keypair) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *Keypair) UpdateReferences() error {
    return nil
}


