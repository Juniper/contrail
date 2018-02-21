
package models
// TelemetryResourceInfo



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propTelemetryResourceInfo_path int = iota
    propTelemetryResourceInfo_rate int = iota
    propTelemetryResourceInfo_name int = iota
)

// TelemetryResourceInfo 
type TelemetryResourceInfo struct {

    Path string `json:"path,omitempty"`
    Rate string `json:"rate,omitempty"`
    Name string `json:"name,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *TelemetryResourceInfo) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeTelemetryResourceInfo makes TelemetryResourceInfo
func MakeTelemetryResourceInfo() *TelemetryResourceInfo{
    return &TelemetryResourceInfo{
    //TODO(nati): Apply default
    Rate: "",
        Name: "",
        Path: "",
        
        modified: big.NewInt(0),
    }
}



// MakeTelemetryResourceInfoSlice makes a slice of TelemetryResourceInfo
func MakeTelemetryResourceInfoSlice() []*TelemetryResourceInfo {
    return []*TelemetryResourceInfo{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *TelemetryResourceInfo) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *TelemetryResourceInfo) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *TelemetryResourceInfo) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *TelemetryResourceInfo) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *TelemetryResourceInfo) GetFQName() []string {
    return model.FQName
}

func (model *TelemetryResourceInfo) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *TelemetryResourceInfo) GetParentType() string {
    return model.ParentType
}

func (model *TelemetryResourceInfo) GetUuid() string {
    return model.UUID
}

func (model *TelemetryResourceInfo) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *TelemetryResourceInfo) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *TelemetryResourceInfo) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *TelemetryResourceInfo) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *TelemetryResourceInfo) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propTelemetryResourceInfo_path) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Path); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Path as path")
        }
        msg["path"] = &val
    }
    
    if model.modified.Bit(propTelemetryResourceInfo_rate) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Rate); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Rate as rate")
        }
        msg["rate"] = &val
    }
    
    if model.modified.Bit(propTelemetryResourceInfo_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Name); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Name as name")
        }
        msg["name"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *TelemetryResourceInfo) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *TelemetryResourceInfo) UpdateReferences() error {
    return nil
}


