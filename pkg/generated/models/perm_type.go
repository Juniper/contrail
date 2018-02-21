
package models
// PermType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propPermType_group_access int = iota
    propPermType_owner int = iota
    propPermType_owner_access int = iota
    propPermType_other_access int = iota
    propPermType_group int = iota
)

// PermType 
type PermType struct {

    Owner string `json:"owner,omitempty"`
    OwnerAccess AccessType `json:"owner_access,omitempty"`
    OtherAccess AccessType `json:"other_access,omitempty"`
    Group string `json:"group,omitempty"`
    GroupAccess AccessType `json:"group_access,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *PermType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakePermType makes PermType
func MakePermType() *PermType{
    return &PermType{
    //TODO(nati): Apply default
    Owner: "",
        OwnerAccess: MakeAccessType(),
        OtherAccess: MakeAccessType(),
        Group: "",
        GroupAccess: MakeAccessType(),
        
        modified: big.NewInt(0),
    }
}



// MakePermTypeSlice makes a slice of PermType
func MakePermTypeSlice() []*PermType {
    return []*PermType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *PermType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *PermType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *PermType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *PermType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *PermType) GetFQName() []string {
    return model.FQName
}

func (model *PermType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *PermType) GetParentType() string {
    return model.ParentType
}

func (model *PermType) GetUuid() string {
    return model.UUID
}

func (model *PermType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *PermType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *PermType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *PermType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *PermType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propPermType_other_access) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.OtherAccess); err != nil {
            return nil, errors.Wrap(err, "Marshal of: OtherAccess as other_access")
        }
        msg["other_access"] = &val
    }
    
    if model.modified.Bit(propPermType_group) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Group); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Group as group")
        }
        msg["group"] = &val
    }
    
    if model.modified.Bit(propPermType_group_access) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.GroupAccess); err != nil {
            return nil, errors.Wrap(err, "Marshal of: GroupAccess as group_access")
        }
        msg["group_access"] = &val
    }
    
    if model.modified.Bit(propPermType_owner) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Owner); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Owner as owner")
        }
        msg["owner"] = &val
    }
    
    if model.modified.Bit(propPermType_owner_access) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.OwnerAccess); err != nil {
            return nil, errors.Wrap(err, "Marshal of: OwnerAccess as owner_access")
        }
        msg["owner_access"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *PermType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *PermType) UpdateReferences() error {
    return nil
}


