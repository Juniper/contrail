
package models
// TimerType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propTimerType_start_time int = iota
    propTimerType_off_interval int = iota
    propTimerType_on_interval int = iota
    propTimerType_end_time int = iota
)

// TimerType 
type TimerType struct {

    OnInterval string `json:"on_interval,omitempty"`
    EndTime string `json:"end_time,omitempty"`
    StartTime string `json:"start_time,omitempty"`
    OffInterval string `json:"off_interval,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *TimerType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeTimerType makes TimerType
func MakeTimerType() *TimerType{
    return &TimerType{
    //TODO(nati): Apply default
    StartTime: "",
        OffInterval: "",
        OnInterval: "",
        EndTime: "",
        
        modified: big.NewInt(0),
    }
}



// MakeTimerTypeSlice makes a slice of TimerType
func MakeTimerTypeSlice() []*TimerType {
    return []*TimerType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *TimerType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *TimerType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *TimerType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *TimerType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *TimerType) GetFQName() []string {
    return model.FQName
}

func (model *TimerType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *TimerType) GetParentType() string {
    return model.ParentType
}

func (model *TimerType) GetUuid() string {
    return model.UUID
}

func (model *TimerType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *TimerType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *TimerType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *TimerType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *TimerType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propTimerType_end_time) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.EndTime); err != nil {
            return nil, errors.Wrap(err, "Marshal of: EndTime as end_time")
        }
        msg["end_time"] = &val
    }
    
    if model.modified.Bit(propTimerType_start_time) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.StartTime); err != nil {
            return nil, errors.Wrap(err, "Marshal of: StartTime as start_time")
        }
        msg["start_time"] = &val
    }
    
    if model.modified.Bit(propTimerType_off_interval) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.OffInterval); err != nil {
            return nil, errors.Wrap(err, "Marshal of: OffInterval as off_interval")
        }
        msg["off_interval"] = &val
    }
    
    if model.modified.Bit(propTimerType_on_interval) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.OnInterval); err != nil {
            return nil, errors.Wrap(err, "Marshal of: OnInterval as on_interval")
        }
        msg["on_interval"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *TimerType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *TimerType) UpdateReferences() error {
    return nil
}


