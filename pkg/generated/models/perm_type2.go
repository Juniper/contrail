
package models
// PermType2



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propPermType2_owner int = iota
    propPermType2_owner_access int = iota
    propPermType2_global_access int = iota
    propPermType2_share int = iota
)

// PermType2 
type PermType2 struct {

    GlobalAccess AccessType `json:"global_access,omitempty"`
    Share []*ShareType `json:"share,omitempty"`
    Owner string `json:"owner,omitempty"`
    OwnerAccess AccessType `json:"owner_access,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *PermType2) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakePermType2 makes PermType2
func MakePermType2() *PermType2{
    return &PermType2{
    //TODO(nati): Apply default
    Owner: "",
        OwnerAccess: MakeAccessType(),
        GlobalAccess: MakeAccessType(),
        
            
                Share:  MakeShareTypeSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakePermType2Slice makes a slice of PermType2
func MakePermType2Slice() []*PermType2 {
    return []*PermType2{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *PermType2) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *PermType2) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *PermType2) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *PermType2) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *PermType2) GetFQName() []string {
    return model.FQName
}

func (model *PermType2) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *PermType2) GetParentType() string {
    return model.ParentType
}

func (model *PermType2) GetUuid() string {
    return model.UUID
}

func (model *PermType2) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *PermType2) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *PermType2) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *PermType2) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *PermType2) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propPermType2_share) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Share); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Share as share")
        }
        msg["share"] = &val
    }
    
    if model.modified.Bit(propPermType2_owner) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Owner); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Owner as owner")
        }
        msg["owner"] = &val
    }
    
    if model.modified.Bit(propPermType2_owner_access) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.OwnerAccess); err != nil {
            return nil, errors.Wrap(err, "Marshal of: OwnerAccess as owner_access")
        }
        msg["owner_access"] = &val
    }
    
    if model.modified.Bit(propPermType2_global_access) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.GlobalAccess); err != nil {
            return nil, errors.Wrap(err, "Marshal of: GlobalAccess as global_access")
        }
        msg["global_access"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *PermType2) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *PermType2) UpdateReferences() error {
    return nil
}


