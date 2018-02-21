
package models
// QosConfigType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type QosConfigType string

// MakeQosConfigType makes QosConfigType
func MakeQosConfigType() QosConfigType {
    var data QosConfigType
    return data
}



// MakeQosConfigTypeSlice makes a slice of QosConfigType
func MakeQosConfigTypeSlice() []QosConfigType {
    return []QosConfigType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *QosConfigType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *QosConfigType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *QosConfigType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *QosConfigType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *QosConfigType) GetFQName() []string {
    return model.FQName
}

func (model *QosConfigType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *QosConfigType) GetParentType() string {
    return model.ParentType
}

func (model *QosConfigType) GetUuid() string {
    return model.UUID
}

func (model *QosConfigType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *QosConfigType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *QosConfigType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *QosConfigType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *QosConfigType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *QosConfigType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *QosConfigType) UpdateReferences() error {
    return nil
}


