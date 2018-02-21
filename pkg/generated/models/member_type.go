
package models
// MemberType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propMemberType_role int = iota
)

// MemberType 
type MemberType struct {

    Role string `json:"role,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *MemberType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeMemberType makes MemberType
func MakeMemberType() *MemberType{
    return &MemberType{
    //TODO(nati): Apply default
    Role: "",
        
        modified: big.NewInt(0),
    }
}



// MakeMemberTypeSlice makes a slice of MemberType
func MakeMemberTypeSlice() []*MemberType {
    return []*MemberType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *MemberType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *MemberType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *MemberType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *MemberType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *MemberType) GetFQName() []string {
    return model.FQName
}

func (model *MemberType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *MemberType) GetParentType() string {
    return model.ParentType
}

func (model *MemberType) GetUuid() string {
    return model.UUID
}

func (model *MemberType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *MemberType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *MemberType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *MemberType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *MemberType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propMemberType_role) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Role); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Role as role")
        }
        msg["role"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *MemberType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *MemberType) UpdateReferences() error {
    return nil
}


