
package models
// RoutingPolicyServiceInstanceType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propRoutingPolicyServiceInstanceType_right_sequence int = iota
    propRoutingPolicyServiceInstanceType_left_sequence int = iota
)

// RoutingPolicyServiceInstanceType 
type RoutingPolicyServiceInstanceType struct {

    RightSequence string `json:"right_sequence,omitempty"`
    LeftSequence string `json:"left_sequence,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *RoutingPolicyServiceInstanceType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeRoutingPolicyServiceInstanceType makes RoutingPolicyServiceInstanceType
func MakeRoutingPolicyServiceInstanceType() *RoutingPolicyServiceInstanceType{
    return &RoutingPolicyServiceInstanceType{
    //TODO(nati): Apply default
    RightSequence: "",
        LeftSequence: "",
        
        modified: big.NewInt(0),
    }
}



// MakeRoutingPolicyServiceInstanceTypeSlice makes a slice of RoutingPolicyServiceInstanceType
func MakeRoutingPolicyServiceInstanceTypeSlice() []*RoutingPolicyServiceInstanceType {
    return []*RoutingPolicyServiceInstanceType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *RoutingPolicyServiceInstanceType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *RoutingPolicyServiceInstanceType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *RoutingPolicyServiceInstanceType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *RoutingPolicyServiceInstanceType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *RoutingPolicyServiceInstanceType) GetFQName() []string {
    return model.FQName
}

func (model *RoutingPolicyServiceInstanceType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *RoutingPolicyServiceInstanceType) GetParentType() string {
    return model.ParentType
}

func (model *RoutingPolicyServiceInstanceType) GetUuid() string {
    return model.UUID
}

func (model *RoutingPolicyServiceInstanceType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *RoutingPolicyServiceInstanceType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *RoutingPolicyServiceInstanceType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *RoutingPolicyServiceInstanceType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *RoutingPolicyServiceInstanceType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propRoutingPolicyServiceInstanceType_left_sequence) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.LeftSequence); err != nil {
            return nil, errors.Wrap(err, "Marshal of: LeftSequence as left_sequence")
        }
        msg["left_sequence"] = &val
    }
    
    if model.modified.Bit(propRoutingPolicyServiceInstanceType_right_sequence) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RightSequence); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RightSequence as right_sequence")
        }
        msg["right_sequence"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *RoutingPolicyServiceInstanceType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *RoutingPolicyServiceInstanceType) UpdateReferences() error {
    return nil
}


