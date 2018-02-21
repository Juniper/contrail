
package models
// ControlTrafficDscpType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propControlTrafficDscpType_control int = iota
    propControlTrafficDscpType_analytics int = iota
    propControlTrafficDscpType_dns int = iota
)

// ControlTrafficDscpType 
type ControlTrafficDscpType struct {

    Control DscpValueType `json:"control,omitempty"`
    Analytics DscpValueType `json:"analytics,omitempty"`
    DNS DscpValueType `json:"dns,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *ControlTrafficDscpType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeControlTrafficDscpType makes ControlTrafficDscpType
func MakeControlTrafficDscpType() *ControlTrafficDscpType{
    return &ControlTrafficDscpType{
    //TODO(nati): Apply default
    Control: MakeDscpValueType(),
        Analytics: MakeDscpValueType(),
        DNS: MakeDscpValueType(),
        
        modified: big.NewInt(0),
    }
}



// MakeControlTrafficDscpTypeSlice makes a slice of ControlTrafficDscpType
func MakeControlTrafficDscpTypeSlice() []*ControlTrafficDscpType {
    return []*ControlTrafficDscpType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *ControlTrafficDscpType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *ControlTrafficDscpType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *ControlTrafficDscpType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *ControlTrafficDscpType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *ControlTrafficDscpType) GetFQName() []string {
    return model.FQName
}

func (model *ControlTrafficDscpType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *ControlTrafficDscpType) GetParentType() string {
    return model.ParentType
}

func (model *ControlTrafficDscpType) GetUuid() string {
    return model.UUID
}

func (model *ControlTrafficDscpType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *ControlTrafficDscpType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *ControlTrafficDscpType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *ControlTrafficDscpType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *ControlTrafficDscpType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propControlTrafficDscpType_dns) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DNS); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DNS as dns")
        }
        msg["dns"] = &val
    }
    
    if model.modified.Bit(propControlTrafficDscpType_control) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Control); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Control as control")
        }
        msg["control"] = &val
    }
    
    if model.modified.Bit(propControlTrafficDscpType_analytics) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Analytics); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Analytics as analytics")
        }
        msg["analytics"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *ControlTrafficDscpType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *ControlTrafficDscpType) UpdateReferences() error {
    return nil
}


