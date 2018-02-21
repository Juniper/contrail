
package models
// AlarmOrList



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propAlarmOrList_or_list int = iota
)

// AlarmOrList 
type AlarmOrList struct {

    OrList []*AlarmAndList `json:"or_list,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *AlarmOrList) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeAlarmOrList makes AlarmOrList
func MakeAlarmOrList() *AlarmOrList{
    return &AlarmOrList{
    //TODO(nati): Apply default
    
            
                OrList:  MakeAlarmAndListSlice(),
            
        
        modified: big.NewInt(0),
    }
}



// MakeAlarmOrListSlice makes a slice of AlarmOrList
func MakeAlarmOrListSlice() []*AlarmOrList {
    return []*AlarmOrList{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *AlarmOrList) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *AlarmOrList) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *AlarmOrList) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *AlarmOrList) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *AlarmOrList) GetFQName() []string {
    return model.FQName
}

func (model *AlarmOrList) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *AlarmOrList) GetParentType() string {
    return model.ParentType
}

func (model *AlarmOrList) GetUuid() string {
    return model.UUID
}

func (model *AlarmOrList) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *AlarmOrList) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *AlarmOrList) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *AlarmOrList) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *AlarmOrList) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propAlarmOrList_or_list) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.OrList); err != nil {
            return nil, errors.Wrap(err, "Marshal of: OrList as or_list")
        }
        msg["or_list"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *AlarmOrList) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *AlarmOrList) UpdateReferences() error {
    return nil
}


