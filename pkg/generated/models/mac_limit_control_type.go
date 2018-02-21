
package models
// MACLimitControlType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propMACLimitControlType_mac_limit int = iota
    propMACLimitControlType_mac_limit_action int = iota
)

// MACLimitControlType 
type MACLimitControlType struct {

    MacLimit int `json:"mac_limit,omitempty"`
    MacLimitAction MACLimitExceedActionType `json:"mac_limit_action,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *MACLimitControlType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeMACLimitControlType makes MACLimitControlType
func MakeMACLimitControlType() *MACLimitControlType{
    return &MACLimitControlType{
    //TODO(nati): Apply default
    MacLimit: 0,
        MacLimitAction: MakeMACLimitExceedActionType(),
        
        modified: big.NewInt(0),
    }
}



// MakeMACLimitControlTypeSlice makes a slice of MACLimitControlType
func MakeMACLimitControlTypeSlice() []*MACLimitControlType {
    return []*MACLimitControlType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *MACLimitControlType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *MACLimitControlType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *MACLimitControlType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *MACLimitControlType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *MACLimitControlType) GetFQName() []string {
    return model.FQName
}

func (model *MACLimitControlType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *MACLimitControlType) GetParentType() string {
    return model.ParentType
}

func (model *MACLimitControlType) GetUuid() string {
    return model.UUID
}

func (model *MACLimitControlType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *MACLimitControlType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *MACLimitControlType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *MACLimitControlType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *MACLimitControlType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propMACLimitControlType_mac_limit) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MacLimit); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MacLimit as mac_limit")
        }
        msg["mac_limit"] = &val
    }
    
    if model.modified.Bit(propMACLimitControlType_mac_limit_action) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MacLimitAction); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MacLimitAction as mac_limit_action")
        }
        msg["mac_limit_action"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *MACLimitControlType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *MACLimitControlType) UpdateReferences() error {
    return nil
}


