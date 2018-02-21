
package models
// LoadbalancerProtocolType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type LoadbalancerProtocolType string

// MakeLoadbalancerProtocolType makes LoadbalancerProtocolType
func MakeLoadbalancerProtocolType() LoadbalancerProtocolType {
    var data LoadbalancerProtocolType
    return data
}



// MakeLoadbalancerProtocolTypeSlice makes a slice of LoadbalancerProtocolType
func MakeLoadbalancerProtocolTypeSlice() []LoadbalancerProtocolType {
    return []LoadbalancerProtocolType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *LoadbalancerProtocolType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *LoadbalancerProtocolType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *LoadbalancerProtocolType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *LoadbalancerProtocolType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *LoadbalancerProtocolType) GetFQName() []string {
    return model.FQName
}

func (model *LoadbalancerProtocolType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *LoadbalancerProtocolType) GetParentType() string {
    return model.ParentType
}

func (model *LoadbalancerProtocolType) GetUuid() string {
    return model.UUID
}

func (model *LoadbalancerProtocolType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *LoadbalancerProtocolType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *LoadbalancerProtocolType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *LoadbalancerProtocolType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *LoadbalancerProtocolType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *LoadbalancerProtocolType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *LoadbalancerProtocolType) UpdateReferences() error {
    return nil
}


