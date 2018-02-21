
package models
// TelemetryStateInfo



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propTelemetryStateInfo_resource int = iota
    propTelemetryStateInfo_server_port int = iota
    propTelemetryStateInfo_server_ip int = iota
)

// TelemetryStateInfo 
type TelemetryStateInfo struct {

    Resource []*TelemetryResourceInfo `json:"resource,omitempty"`
    ServerPort int `json:"server_port,omitempty"`
    ServerIP string `json:"server_ip,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *TelemetryStateInfo) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeTelemetryStateInfo makes TelemetryStateInfo
func MakeTelemetryStateInfo() *TelemetryStateInfo{
    return &TelemetryStateInfo{
    //TODO(nati): Apply default
    
            
                Resource:  MakeTelemetryResourceInfoSlice(),
            
        ServerPort: 0,
        ServerIP: "",
        
        modified: big.NewInt(0),
    }
}



// MakeTelemetryStateInfoSlice makes a slice of TelemetryStateInfo
func MakeTelemetryStateInfoSlice() []*TelemetryStateInfo {
    return []*TelemetryStateInfo{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *TelemetryStateInfo) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *TelemetryStateInfo) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *TelemetryStateInfo) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *TelemetryStateInfo) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *TelemetryStateInfo) GetFQName() []string {
    return model.FQName
}

func (model *TelemetryStateInfo) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *TelemetryStateInfo) GetParentType() string {
    return model.ParentType
}

func (model *TelemetryStateInfo) GetUuid() string {
    return model.UUID
}

func (model *TelemetryStateInfo) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *TelemetryStateInfo) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *TelemetryStateInfo) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *TelemetryStateInfo) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *TelemetryStateInfo) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propTelemetryStateInfo_server_ip) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServerIP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServerIP as server_ip")
        }
        msg["server_ip"] = &val
    }
    
    if model.modified.Bit(propTelemetryStateInfo_resource) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Resource); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Resource as resource")
        }
        msg["resource"] = &val
    }
    
    if model.modified.Bit(propTelemetryStateInfo_server_port) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ServerPort); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ServerPort as server_port")
        }
        msg["server_port"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *TelemetryStateInfo) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *TelemetryStateInfo) UpdateReferences() error {
    return nil
}


