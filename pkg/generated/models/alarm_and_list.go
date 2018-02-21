
package models
// AlarmAndList



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propAlarmAndList_and_list int = iota
)

// AlarmAndList 
type AlarmAndList struct {

    AndList []*AlarmExpression `json:"and_list,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *AlarmAndList) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeAlarmAndList makes AlarmAndList
func MakeAlarmAndList() *AlarmAndList{
    return &AlarmAndList{
    //TODO(nati): Apply default
    
            
                AndList:  MakeAlarmExpressionSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeAlarmAndListSlice makes a slice of AlarmAndList
func MakeAlarmAndListSlice() []*AlarmAndList {
    return []*AlarmAndList{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *AlarmAndList) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *AlarmAndList) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *AlarmAndList) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *AlarmAndList) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *AlarmAndList) GetFQName() []string {
    return model.FQName
}

func (model *AlarmAndList) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *AlarmAndList) GetParentType() string {
    return model.ParentType
}

func (model *AlarmAndList) GetUuid() string {
    return model.UUID
}

func (model *AlarmAndList) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *AlarmAndList) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *AlarmAndList) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *AlarmAndList) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *AlarmAndList) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propAlarmAndList_and_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AndList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AndList as and_list")
        }
        msg["and_list"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *AlarmAndList) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *AlarmAndList) UpdateReferences() error {
    return nil
}


