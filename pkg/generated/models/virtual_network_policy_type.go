
package models
// VirtualNetworkPolicyType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propVirtualNetworkPolicyType_timer int = iota
    propVirtualNetworkPolicyType_sequence int = iota
)

// VirtualNetworkPolicyType 
type VirtualNetworkPolicyType struct {

    Timer *TimerType `json:"timer,omitempty"`
    Sequence *SequenceType `json:"sequence,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *VirtualNetworkPolicyType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeVirtualNetworkPolicyType makes VirtualNetworkPolicyType
func MakeVirtualNetworkPolicyType() *VirtualNetworkPolicyType{
    return &VirtualNetworkPolicyType{
    //TODO(nati): Apply default
    Sequence: MakeSequenceType(),
        Timer: MakeTimerType(),
        
        modified: big.NewInt(0),
    }
}



// MakeVirtualNetworkPolicyTypeSlice makes a slice of VirtualNetworkPolicyType
func MakeVirtualNetworkPolicyTypeSlice() []*VirtualNetworkPolicyType {
    return []*VirtualNetworkPolicyType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VirtualNetworkPolicyType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *VirtualNetworkPolicyType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VirtualNetworkPolicyType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *VirtualNetworkPolicyType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *VirtualNetworkPolicyType) GetFQName() []string {
    return model.FQName
}

func (model *VirtualNetworkPolicyType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VirtualNetworkPolicyType) GetParentType() string {
    return model.ParentType
}

func (model *VirtualNetworkPolicyType) GetUuid() string {
    return model.UUID
}

func (model *VirtualNetworkPolicyType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VirtualNetworkPolicyType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VirtualNetworkPolicyType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VirtualNetworkPolicyType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VirtualNetworkPolicyType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propVirtualNetworkPolicyType_timer) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Timer); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Timer as timer")
        }
        msg["timer"] = &val
    }
    
    if model.modified.Bit(propVirtualNetworkPolicyType_sequence) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Sequence); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Sequence as sequence")
        }
        msg["sequence"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *VirtualNetworkPolicyType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VirtualNetworkPolicyType) UpdateReferences() error {
    return nil
}


