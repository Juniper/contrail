
package models
// EcmpHashingIncludeFields



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propEcmpHashingIncludeFields_destination_ip int = iota
    propEcmpHashingIncludeFields_ip_protocol int = iota
    propEcmpHashingIncludeFields_source_ip int = iota
    propEcmpHashingIncludeFields_hashing_configured int = iota
    propEcmpHashingIncludeFields_source_port int = iota
    propEcmpHashingIncludeFields_destination_port int = iota
)

// EcmpHashingIncludeFields 
type EcmpHashingIncludeFields struct {

    DestinationIP bool `json:"destination_ip"`
    IPProtocol bool `json:"ip_protocol"`
    SourceIP bool `json:"source_ip"`
    HashingConfigured bool `json:"hashing_configured"`
    SourcePort bool `json:"source_port"`
    DestinationPort bool `json:"destination_port"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *EcmpHashingIncludeFields) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeEcmpHashingIncludeFields makes EcmpHashingIncludeFields
func MakeEcmpHashingIncludeFields() *EcmpHashingIncludeFields{
    return &EcmpHashingIncludeFields{
    //TODO(nati): Apply default
    DestinationPort: false,
        DestinationIP: false,
        IPProtocol: false,
        SourceIP: false,
        HashingConfigured: false,
        SourcePort: false,
        
        modified: big.NewInt(0),
    }
}



// MakeEcmpHashingIncludeFieldsSlice makes a slice of EcmpHashingIncludeFields
func MakeEcmpHashingIncludeFieldsSlice() []*EcmpHashingIncludeFields {
    return []*EcmpHashingIncludeFields{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *EcmpHashingIncludeFields) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *EcmpHashingIncludeFields) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *EcmpHashingIncludeFields) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *EcmpHashingIncludeFields) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *EcmpHashingIncludeFields) GetFQName() []string {
    return model.FQName
}

func (model *EcmpHashingIncludeFields) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *EcmpHashingIncludeFields) GetParentType() string {
    return model.ParentType
}

func (model *EcmpHashingIncludeFields) GetUuid() string {
    return model.UUID
}

func (model *EcmpHashingIncludeFields) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *EcmpHashingIncludeFields) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *EcmpHashingIncludeFields) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *EcmpHashingIncludeFields) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *EcmpHashingIncludeFields) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propEcmpHashingIncludeFields_hashing_configured) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.HashingConfigured); err != nil {
            return nil, errors.Wrap(err, "Marshal of: HashingConfigured as hashing_configured")
        }
        msg["hashing_configured"] = &val
    }
    
    if model.modified.Bit(propEcmpHashingIncludeFields_source_port) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SourcePort); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SourcePort as source_port")
        }
        msg["source_port"] = &val
    }
    
    if model.modified.Bit(propEcmpHashingIncludeFields_destination_port) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DestinationPort); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DestinationPort as destination_port")
        }
        msg["destination_port"] = &val
    }
    
    if model.modified.Bit(propEcmpHashingIncludeFields_destination_ip) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DestinationIP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DestinationIP as destination_ip")
        }
        msg["destination_ip"] = &val
    }
    
    if model.modified.Bit(propEcmpHashingIncludeFields_ip_protocol) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IPProtocol); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IPProtocol as ip_protocol")
        }
        msg["ip_protocol"] = &val
    }
    
    if model.modified.Bit(propEcmpHashingIncludeFields_source_ip) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SourceIP); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SourceIP as source_ip")
        }
        msg["source_ip"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *EcmpHashingIncludeFields) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *EcmpHashingIncludeFields) UpdateReferences() error {
    return nil
}


