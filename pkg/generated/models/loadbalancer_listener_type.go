
package models
// LoadbalancerListenerType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propLoadbalancerListenerType_default_tls_container int = iota
    propLoadbalancerListenerType_protocol int = iota
    propLoadbalancerListenerType_connection_limit int = iota
    propLoadbalancerListenerType_admin_state int = iota
    propLoadbalancerListenerType_sni_containers int = iota
    propLoadbalancerListenerType_protocol_port int = iota
)

// LoadbalancerListenerType 
type LoadbalancerListenerType struct {

    ConnectionLimit int `json:"connection_limit,omitempty"`
    AdminState bool `json:"admin_state"`
    SniContainers []string `json:"sni_containers,omitempty"`
    ProtocolPort int `json:"protocol_port,omitempty"`
    DefaultTLSContainer string `json:"default_tls_container,omitempty"`
    Protocol LoadbalancerProtocolType `json:"protocol,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *LoadbalancerListenerType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeLoadbalancerListenerType makes LoadbalancerListenerType
func MakeLoadbalancerListenerType() *LoadbalancerListenerType{
    return &LoadbalancerListenerType{
    //TODO(nati): Apply default
    ConnectionLimit: 0,
        AdminState: false,
        SniContainers: []string{},
        ProtocolPort: 0,
        DefaultTLSContainer: "",
        Protocol: MakeLoadbalancerProtocolType(),
        
        modified: big.NewInt(0),
    }
}



// MakeLoadbalancerListenerTypeSlice makes a slice of LoadbalancerListenerType
func MakeLoadbalancerListenerTypeSlice() []*LoadbalancerListenerType {
    return []*LoadbalancerListenerType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *LoadbalancerListenerType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *LoadbalancerListenerType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *LoadbalancerListenerType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *LoadbalancerListenerType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *LoadbalancerListenerType) GetFQName() []string {
    return model.FQName
}

func (model *LoadbalancerListenerType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *LoadbalancerListenerType) GetParentType() string {
    return model.ParentType
}

func (model *LoadbalancerListenerType) GetUuid() string {
    return model.UUID
}

func (model *LoadbalancerListenerType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *LoadbalancerListenerType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *LoadbalancerListenerType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *LoadbalancerListenerType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *LoadbalancerListenerType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propLoadbalancerListenerType_default_tls_container) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DefaultTLSContainer); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DefaultTLSContainer as default_tls_container")
        }
        msg["default_tls_container"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerListenerType_protocol) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Protocol); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Protocol as protocol")
        }
        msg["protocol"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerListenerType_connection_limit) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ConnectionLimit); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ConnectionLimit as connection_limit")
        }
        msg["connection_limit"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerListenerType_admin_state) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AdminState); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AdminState as admin_state")
        }
        msg["admin_state"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerListenerType_sni_containers) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SniContainers); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SniContainers as sni_containers")
        }
        msg["sni_containers"] = &val
    }
    
    if model.modified.Bit(propLoadbalancerListenerType_protocol_port) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ProtocolPort); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ProtocolPort as protocol_port")
        }
        msg["protocol_port"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *LoadbalancerListenerType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *LoadbalancerListenerType) UpdateReferences() error {
    return nil
}


