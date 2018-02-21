
package models
// RbacPermType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propRbacPermType_role_crud int = iota
    propRbacPermType_role_name int = iota
)

// RbacPermType 
type RbacPermType struct {

    RoleCrud string `json:"role_crud,omitempty"`
    RoleName string `json:"role_name,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *RbacPermType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeRbacPermType makes RbacPermType
func MakeRbacPermType() *RbacPermType{
    return &RbacPermType{
    //TODO(nati): Apply default
    RoleName: "",
        RoleCrud: "",
        
        modified: big.NewInt(0),
    }
}



// MakeRbacPermTypeSlice makes a slice of RbacPermType
func MakeRbacPermTypeSlice() []*RbacPermType {
    return []*RbacPermType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *RbacPermType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *RbacPermType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *RbacPermType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *RbacPermType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *RbacPermType) GetFQName() []string {
    return model.FQName
}

func (model *RbacPermType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *RbacPermType) GetParentType() string {
    return model.ParentType
}

func (model *RbacPermType) GetUuid() string {
    return model.UUID
}

func (model *RbacPermType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *RbacPermType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *RbacPermType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *RbacPermType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *RbacPermType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propRbacPermType_role_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RoleName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RoleName as role_name")
        }
        msg["role_name"] = &val
    }
    
    if model.modified.Bit(propRbacPermType_role_crud) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RoleCrud); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RoleCrud as role_crud")
        }
        msg["role_crud"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *RbacPermType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *RbacPermType) UpdateReferences() error {
    return nil
}


