
package models
// IdPermsType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propIdPermsType_creator int = iota
    propIdPermsType_user_visible int = iota
    propIdPermsType_last_modified int = iota
    propIdPermsType_permissions int = iota
    propIdPermsType_enable int = iota
    propIdPermsType_description int = iota
    propIdPermsType_created int = iota
)

// IdPermsType 
type IdPermsType struct {

    Enable bool `json:"enable"`
    Description string `json:"description,omitempty"`
    Created string `json:"created,omitempty"`
    Creator string `json:"creator,omitempty"`
    UserVisible bool `json:"user_visible"`
    LastModified string `json:"last_modified,omitempty"`
    Permissions *PermType `json:"permissions,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *IdPermsType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeIdPermsType makes IdPermsType
func MakeIdPermsType() *IdPermsType{
    return &IdPermsType{
    //TODO(nati): Apply default
    Enable: false,
        Description: "",
        Created: "",
        Creator: "",
        UserVisible: false,
        LastModified: "",
        Permissions: MakePermType(),
        
        modified: big.NewInt(0),
    }
}



// MakeIdPermsTypeSlice makes a slice of IdPermsType
func MakeIdPermsTypeSlice() []*IdPermsType {
    return []*IdPermsType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *IdPermsType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *IdPermsType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *IdPermsType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *IdPermsType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *IdPermsType) GetFQName() []string {
    return model.FQName
}

func (model *IdPermsType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *IdPermsType) GetParentType() string {
    return model.ParentType
}

func (model *IdPermsType) GetUuid() string {
    return model.UUID
}

func (model *IdPermsType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *IdPermsType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *IdPermsType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *IdPermsType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *IdPermsType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propIdPermsType_enable) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Enable); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Enable as enable")
        }
        msg["enable"] = &val
    }
    
    if model.modified.Bit(propIdPermsType_description) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Description); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Description as description")
        }
        msg["description"] = &val
    }
    
    if model.modified.Bit(propIdPermsType_created) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Created); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Created as created")
        }
        msg["created"] = &val
    }
    
    if model.modified.Bit(propIdPermsType_creator) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Creator); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Creator as creator")
        }
        msg["creator"] = &val
    }
    
    if model.modified.Bit(propIdPermsType_user_visible) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UserVisible); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UserVisible as user_visible")
        }
        msg["user_visible"] = &val
    }
    
    if model.modified.Bit(propIdPermsType_last_modified) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LastModified); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LastModified as last_modified")
        }
        msg["last_modified"] = &val
    }
    
    if model.modified.Bit(propIdPermsType_permissions) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Permissions); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Permissions as permissions")
        }
        msg["permissions"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *IdPermsType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *IdPermsType) UpdateReferences() error {
    return nil
}


