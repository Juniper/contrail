
package models
// MACMoveLimitControlType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propMACMoveLimitControlType_mac_move_limit_action int = iota
    propMACMoveLimitControlType_mac_move_time_window int = iota
    propMACMoveLimitControlType_mac_move_limit int = iota
)

// MACMoveLimitControlType 
type MACMoveLimitControlType struct {

    MacMoveTimeWindow MACMoveTimeWindow `json:"mac_move_time_window,omitempty"`
    MacMoveLimit int `json:"mac_move_limit,omitempty"`
    MacMoveLimitAction MACLimitExceedActionType `json:"mac_move_limit_action,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *MACMoveLimitControlType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeMACMoveLimitControlType makes MACMoveLimitControlType
func MakeMACMoveLimitControlType() *MACMoveLimitControlType{
    return &MACMoveLimitControlType{
    //TODO(nati): Apply default
    MacMoveLimitAction: MakeMACLimitExceedActionType(),
        MacMoveTimeWindow: MakeMACMoveTimeWindow(),
        MacMoveLimit: 0,
        
        modified: big.NewInt(0),
    }
}



// MakeMACMoveLimitControlTypeSlice makes a slice of MACMoveLimitControlType
func MakeMACMoveLimitControlTypeSlice() []*MACMoveLimitControlType {
    return []*MACMoveLimitControlType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *MACMoveLimitControlType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *MACMoveLimitControlType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *MACMoveLimitControlType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *MACMoveLimitControlType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *MACMoveLimitControlType) GetFQName() []string {
    return model.FQName
}

func (model *MACMoveLimitControlType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *MACMoveLimitControlType) GetParentType() string {
    return model.ParentType
}

func (model *MACMoveLimitControlType) GetUuid() string {
    return model.UUID
}

func (model *MACMoveLimitControlType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *MACMoveLimitControlType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *MACMoveLimitControlType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *MACMoveLimitControlType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *MACMoveLimitControlType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propMACMoveLimitControlType_mac_move_limit_action) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MacMoveLimitAction); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MacMoveLimitAction as mac_move_limit_action")
        }
        msg["mac_move_limit_action"] = &val
    }
    
    if model.modified.Bit(propMACMoveLimitControlType_mac_move_time_window) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MacMoveTimeWindow); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MacMoveTimeWindow as mac_move_time_window")
        }
        msg["mac_move_time_window"] = &val
    }
    
    if model.modified.Bit(propMACMoveLimitControlType_mac_move_limit) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MacMoveLimit); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MacMoveLimit as mac_move_limit")
        }
        msg["mac_move_limit"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *MACMoveLimitControlType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *MACMoveLimitControlType) UpdateReferences() error {
    return nil
}


