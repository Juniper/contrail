
package models
// InterfaceMirrorType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propInterfaceMirrorType_traffic_direction int = iota
    propInterfaceMirrorType_mirror_to int = iota
)

// InterfaceMirrorType 
type InterfaceMirrorType struct {

    TrafficDirection TrafficDirectionType `json:"traffic_direction,omitempty"`
    MirrorTo *MirrorActionType `json:"mirror_to,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *InterfaceMirrorType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeInterfaceMirrorType makes InterfaceMirrorType
func MakeInterfaceMirrorType() *InterfaceMirrorType{
    return &InterfaceMirrorType{
    //TODO(nati): Apply default
    TrafficDirection: MakeTrafficDirectionType(),
        MirrorTo: MakeMirrorActionType(),
        
        modified: big.NewInt(0),
    }
}



// MakeInterfaceMirrorTypeSlice makes a slice of InterfaceMirrorType
func MakeInterfaceMirrorTypeSlice() []*InterfaceMirrorType {
    return []*InterfaceMirrorType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *InterfaceMirrorType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *InterfaceMirrorType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *InterfaceMirrorType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *InterfaceMirrorType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *InterfaceMirrorType) GetFQName() []string {
    return model.FQName
}

func (model *InterfaceMirrorType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *InterfaceMirrorType) GetParentType() string {
    return model.ParentType
}

func (model *InterfaceMirrorType) GetUuid() string {
    return model.UUID
}

func (model *InterfaceMirrorType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *InterfaceMirrorType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *InterfaceMirrorType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *InterfaceMirrorType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *InterfaceMirrorType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propInterfaceMirrorType_traffic_direction) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.TrafficDirection); err != nil {
            return nil, errors.Wrap(err, "Marshal of: TrafficDirection as traffic_direction")
        }
        msg["traffic_direction"] = &val
    }
    
    if model.modified.Bit(propInterfaceMirrorType_mirror_to) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.MirrorTo); err != nil {
            return nil, errors.Wrap(err, "Marshal of: MirrorTo as mirror_to")
        }
        msg["mirror_to"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *InterfaceMirrorType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *InterfaceMirrorType) UpdateReferences() error {
    return nil
}


