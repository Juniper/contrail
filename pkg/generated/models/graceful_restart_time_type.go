
package models
// GracefulRestartTimeType


import (
    "encoding/json"
    "strings"
    //"math/big"
    //"github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)
type GracefulRestartTimeType int

// MakeGracefulRestartTimeType makes GracefulRestartTimeType
func MakeGracefulRestartTimeType() GracefulRestartTimeType {
    var data GracefulRestartTimeType
    return data
}



// MakeGracefulRestartTimeTypeSlice makes a slice of GracefulRestartTimeType
func MakeGracefulRestartTimeTypeSlice() []GracefulRestartTimeType {
    return []GracefulRestartTimeType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *GracefulRestartTimeType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *GracefulRestartTimeType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *GracefulRestartTimeType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *GracefulRestartTimeType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *GracefulRestartTimeType) GetFQName() []string {
    return model.FQName
}

func (model *GracefulRestartTimeType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *GracefulRestartTimeType) GetParentType() string {
    return model.ParentType
}

func (model *GracefulRestartTimeType) GetUuid() string {
    return model.UUID
}

func (model *GracefulRestartTimeType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *GracefulRestartTimeType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *GracefulRestartTimeType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *GracefulRestartTimeType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *GracefulRestartTimeType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    return json.Marshal(msg)
}

func (model *GracefulRestartTimeType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *GracefulRestartTimeType) UpdateReferences() error {
    return nil
}


