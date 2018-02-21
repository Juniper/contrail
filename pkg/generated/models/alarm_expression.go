
package models
// AlarmExpression



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propAlarmExpression_operations int = iota
    propAlarmExpression_operand1 int = iota
    propAlarmExpression_variables int = iota
    propAlarmExpression_operand2 int = iota
)

// AlarmExpression 
type AlarmExpression struct {

    Operand1 string `json:"operand1,omitempty"`
    Variables []string `json:"variables,omitempty"`
    Operand2 *AlarmOperand2 `json:"operand2,omitempty"`
    Operations AlarmOperation `json:"operations,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *AlarmExpression) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeAlarmExpression makes AlarmExpression
func MakeAlarmExpression() *AlarmExpression{
    return &AlarmExpression{
    //TODO(nati): Apply default
    Variables: []string{},
        Operand2: MakeAlarmOperand2(),
        Operations: MakeAlarmOperation(),
        Operand1: "",
        
        modified: big.NewInt(0),
    }
}



// MakeAlarmExpressionSlice makes a slice of AlarmExpression
func MakeAlarmExpressionSlice() []*AlarmExpression {
    return []*AlarmExpression{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *AlarmExpression) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *AlarmExpression) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *AlarmExpression) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *AlarmExpression) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *AlarmExpression) GetFQName() []string {
    return model.FQName
}

func (model *AlarmExpression) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *AlarmExpression) GetParentType() string {
    return model.ParentType
}

func (model *AlarmExpression) GetUuid() string {
    return model.UUID
}

func (model *AlarmExpression) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *AlarmExpression) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *AlarmExpression) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *AlarmExpression) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *AlarmExpression) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propAlarmExpression_operations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Operations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Operations as operations")
        }
        msg["operations"] = &val
    }
    
    if model.modified.Bit(propAlarmExpression_operand1) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Operand1); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Operand1 as operand1")
        }
        msg["operand1"] = &val
    }
    
    if model.modified.Bit(propAlarmExpression_variables) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Variables); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Variables as variables")
        }
        msg["variables"] = &val
    }
    
    if model.modified.Bit(propAlarmExpression_operand2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Operand2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Operand2 as operand2")
        }
        msg["operand2"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *AlarmExpression) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *AlarmExpression) UpdateReferences() error {
    return nil
}


