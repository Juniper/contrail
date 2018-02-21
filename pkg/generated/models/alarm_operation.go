
package models
// AlarmOperation


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type AlarmOperation string

// MakeAlarmOperation makes AlarmOperation
func MakeAlarmOperation() AlarmOperation {
    var data AlarmOperation
    return data
}



// MakeAlarmOperationSlice makes a slice of AlarmOperation
func MakeAlarmOperationSlice() []AlarmOperation {
    return []AlarmOperation{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *AlarmOperation) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *AlarmOperation) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *AlarmOperation) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *AlarmOperation) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *AlarmOperation) GetFQName() []string {
    return model.FQName
}

func (model *AlarmOperation) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *AlarmOperation) GetParentType() string {
    return model.ParentType
}

func (model *AlarmOperation) GetUuid() string {
    return model.UUID
}

func (model *AlarmOperation) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *AlarmOperation) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *AlarmOperation) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *AlarmOperation) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *AlarmOperation) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *AlarmOperation) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *AlarmOperation) UpdateReferences() error {
    return nil
}


