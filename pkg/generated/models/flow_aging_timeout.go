
package models
// FlowAgingTimeout



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propFlowAgingTimeout_timeout_in_seconds int = iota
    propFlowAgingTimeout_protocol int = iota
    propFlowAgingTimeout_port int = iota
)

// FlowAgingTimeout 
type FlowAgingTimeout struct {

    TimeoutInSeconds int `json:"timeout_in_seconds,omitempty"`
    Protocol string `json:"protocol,omitempty"`
    Port int `json:"port,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *FlowAgingTimeout) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeFlowAgingTimeout makes FlowAgingTimeout
func MakeFlowAgingTimeout() *FlowAgingTimeout{
    return &FlowAgingTimeout{
    //TODO(nati): Apply default
    TimeoutInSeconds: 0,
        Protocol: "",
        Port: 0,
        
        modified: big.NewInt(0),
    }
}



// MakeFlowAgingTimeoutSlice makes a slice of FlowAgingTimeout
func MakeFlowAgingTimeoutSlice() []*FlowAgingTimeout {
    return []*FlowAgingTimeout{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *FlowAgingTimeout) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *FlowAgingTimeout) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *FlowAgingTimeout) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *FlowAgingTimeout) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *FlowAgingTimeout) GetFQName() []string {
    return model.FQName
}

func (model *FlowAgingTimeout) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *FlowAgingTimeout) GetParentType() string {
    return model.ParentType
}

func (model *FlowAgingTimeout) GetUuid() string {
    return model.UUID
}

func (model *FlowAgingTimeout) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *FlowAgingTimeout) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *FlowAgingTimeout) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *FlowAgingTimeout) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *FlowAgingTimeout) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propFlowAgingTimeout_timeout_in_seconds) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.TimeoutInSeconds); err != nil {
            return nil, errors.Wrap(err, "Marshal of: TimeoutInSeconds as timeout_in_seconds")
        }
        msg["timeout_in_seconds"] = &val
    }
    
    if model.modified.Bit(propFlowAgingTimeout_protocol) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Protocol); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Protocol as protocol")
        }
        msg["protocol"] = &val
    }
    
    if model.modified.Bit(propFlowAgingTimeout_port) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Port); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Port as port")
        }
        msg["port"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *FlowAgingTimeout) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *FlowAgingTimeout) UpdateReferences() error {
    return nil
}


