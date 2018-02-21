
package models
// VrfAssignTableType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propVrfAssignTableType_vrf_assign_rule int = iota
)

// VrfAssignTableType 
type VrfAssignTableType struct {

    VRFAssignRule []*VrfAssignRuleType `json:"vrf_assign_rule,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *VrfAssignTableType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeVrfAssignTableType makes VrfAssignTableType
func MakeVrfAssignTableType() *VrfAssignTableType{
    return &VrfAssignTableType{
    //TODO(nati): Apply default
    
            
                VRFAssignRule:  MakeVrfAssignRuleTypeSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeVrfAssignTableTypeSlice makes a slice of VrfAssignTableType
func MakeVrfAssignTableTypeSlice() []*VrfAssignTableType {
    return []*VrfAssignTableType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VrfAssignTableType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *VrfAssignTableType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VrfAssignTableType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *VrfAssignTableType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *VrfAssignTableType) GetFQName() []string {
    return model.FQName
}

func (model *VrfAssignTableType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VrfAssignTableType) GetParentType() string {
    return model.ParentType
}

func (model *VrfAssignTableType) GetUuid() string {
    return model.UUID
}

func (model *VrfAssignTableType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VrfAssignTableType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VrfAssignTableType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VrfAssignTableType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VrfAssignTableType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propVrfAssignTableType_vrf_assign_rule) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VRFAssignRule); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VRFAssignRule as vrf_assign_rule")
        }
        msg["vrf_assign_rule"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *VrfAssignTableType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VrfAssignTableType) UpdateReferences() error {
    return nil
}


