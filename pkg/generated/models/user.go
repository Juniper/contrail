
package models
// User



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propUser_perms2 int = iota
    propUser_uuid int = iota
    propUser_fq_name int = iota
    propUser_password int = iota
    propUser_id_perms int = iota
    propUser_display_name int = iota
    propUser_annotations int = iota
    propUser_parent_uuid int = iota
    propUser_parent_type int = iota
)

// User 
type User struct {

    FQName []string `json:"fq_name,omitempty"`
    Password string `json:"password,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *User) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeUser makes User
func MakeUser() *User{
    return &User{
    //TODO(nati): Apply default
    ParentUUID: "",
        ParentType: "",
        Password: "",
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        UUID: "",
        FQName: []string{},
        
        modified: big.NewInt(0),
    }
}



// MakeUserSlice makes a slice of User
func MakeUserSlice() []*User {
    return []*User{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *User) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc4200ae0a0])
    fqn := []string{}
    
    fqn = Project{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *User) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-project", "_", "-", -1)
}

func (model *User) GetDefaultName() string {
    return strings.Replace("default-user", "_", "-", -1)
}

func (model *User) GetType() string {
    return strings.Replace("user", "_", "-", -1)
}

func (model *User) GetFQName() []string {
    return model.FQName
}

func (model *User) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *User) GetParentType() string {
    return model.ParentType
}

func (model *User) GetUuid() string {
    return model.UUID
}

func (model *User) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *User) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *User) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *User) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *User) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propUser_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propUser_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propUser_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propUser_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propUser_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propUser_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propUser_password) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Password); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Password as password")
        }
        msg["password"] = &val
    }
    
    if model.modified.Bit(propUser_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propUser_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *User) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *User) UpdateReferences() error {
    return nil
}


