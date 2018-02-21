
package models
// UserCredentials



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propUserCredentials_username int = iota
    propUserCredentials_password int = iota
)

// UserCredentials 
type UserCredentials struct {

    Username string `json:"username,omitempty"`
    Password string `json:"password,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *UserCredentials) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeUserCredentials makes UserCredentials
func MakeUserCredentials() *UserCredentials{
    return &UserCredentials{
    //TODO(nati): Apply default
    Username: "",
        Password: "",
        
        modified: big.NewInt(0),
    }
}



// MakeUserCredentialsSlice makes a slice of UserCredentials
func MakeUserCredentialsSlice() []*UserCredentials {
    return []*UserCredentials{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *UserCredentials) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *UserCredentials) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *UserCredentials) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *UserCredentials) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *UserCredentials) GetFQName() []string {
    return model.FQName
}

func (model *UserCredentials) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *UserCredentials) GetParentType() string {
    return model.ParentType
}

func (model *UserCredentials) GetUuid() string {
    return model.UUID
}

func (model *UserCredentials) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *UserCredentials) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *UserCredentials) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *UserCredentials) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *UserCredentials) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propUserCredentials_password) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Password); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Password as password")
        }
        msg["password"] = &val
    }
    
    if model.modified.Bit(propUserCredentials_username) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Username); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Username as username")
        }
        msg["username"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *UserCredentials) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *UserCredentials) UpdateReferences() error {
    return nil
}


