
package models
// AlarmOperand2



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propAlarmOperand2_uve_attribute int = iota
    propAlarmOperand2_json_value int = iota
)

// AlarmOperand2 
type AlarmOperand2 struct {

    JSONValue string `json:"json_value,omitempty"`
    UveAttribute string `json:"uve_attribute,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *AlarmOperand2) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeAlarmOperand2 makes AlarmOperand2
func MakeAlarmOperand2() *AlarmOperand2{
    return &AlarmOperand2{
    //TODO(nati): Apply default
    UveAttribute: "",
        JSONValue: "",
        
        modified: big.NewInt(0),
    }
}



// MakeAlarmOperand2Slice makes a slice of AlarmOperand2
func MakeAlarmOperand2Slice() []*AlarmOperand2 {
    return []*AlarmOperand2{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *AlarmOperand2) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *AlarmOperand2) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *AlarmOperand2) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *AlarmOperand2) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *AlarmOperand2) GetFQName() []string {
    return model.FQName
}

func (model *AlarmOperand2) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *AlarmOperand2) GetParentType() string {
    return model.ParentType
}

func (model *AlarmOperand2) GetUuid() string {
    return model.UUID
}

func (model *AlarmOperand2) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *AlarmOperand2) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *AlarmOperand2) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *AlarmOperand2) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *AlarmOperand2) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propAlarmOperand2_uve_attribute) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UveAttribute); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UveAttribute as uve_attribute")
        }
        msg["uve_attribute"] = &val
    }
    
    if model.modified.Bit(propAlarmOperand2_json_value) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.JSONValue); err != nil {
            return nil, errors.Wrap(err, "Marshal of: JSONValue as json_value")
        }
        msg["json_value"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *AlarmOperand2) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *AlarmOperand2) UpdateReferences() error {
    return nil
}


