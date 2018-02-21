
package models
// VirtualDnsRecordType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propVirtualDnsRecordType_record_type int = iota
    propVirtualDnsRecordType_record_ttl_seconds int = iota
    propVirtualDnsRecordType_record_mx_preference int = iota
    propVirtualDnsRecordType_record_name int = iota
    propVirtualDnsRecordType_record_class int = iota
    propVirtualDnsRecordType_record_data int = iota
)

// VirtualDnsRecordType 
type VirtualDnsRecordType struct {

    RecordMXPreference int `json:"record_mx_preference,omitempty"`
    RecordName string `json:"record_name,omitempty"`
    RecordClass DnsRecordClassType `json:"record_class,omitempty"`
    RecordData string `json:"record_data,omitempty"`
    RecordType DnsRecordTypeType `json:"record_type,omitempty"`
    RecordTTLSeconds int `json:"record_ttl_seconds,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *VirtualDnsRecordType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeVirtualDnsRecordType makes VirtualDnsRecordType
func MakeVirtualDnsRecordType() *VirtualDnsRecordType{
    return &VirtualDnsRecordType{
    //TODO(nati): Apply default
    RecordData: "",
        RecordType: MakeDnsRecordTypeType(),
        RecordTTLSeconds: 0,
        RecordMXPreference: 0,
        RecordName: "",
        RecordClass: MakeDnsRecordClassType(),
        
        modified: big.NewInt(0),
    }
}



// MakeVirtualDnsRecordTypeSlice makes a slice of VirtualDnsRecordType
func MakeVirtualDnsRecordTypeSlice() []*VirtualDnsRecordType {
    return []*VirtualDnsRecordType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VirtualDnsRecordType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *VirtualDnsRecordType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VirtualDnsRecordType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *VirtualDnsRecordType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *VirtualDnsRecordType) GetFQName() []string {
    return model.FQName
}

func (model *VirtualDnsRecordType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VirtualDnsRecordType) GetParentType() string {
    return model.ParentType
}

func (model *VirtualDnsRecordType) GetUuid() string {
    return model.UUID
}

func (model *VirtualDnsRecordType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VirtualDnsRecordType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VirtualDnsRecordType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VirtualDnsRecordType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VirtualDnsRecordType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propVirtualDnsRecordType_record_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RecordName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RecordName as record_name")
        }
        msg["record_name"] = &val
    }
    
    if model.modified.Bit(propVirtualDnsRecordType_record_class) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RecordClass); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RecordClass as record_class")
        }
        msg["record_class"] = &val
    }
    
    if model.modified.Bit(propVirtualDnsRecordType_record_data) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RecordData); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RecordData as record_data")
        }
        msg["record_data"] = &val
    }
    
    if model.modified.Bit(propVirtualDnsRecordType_record_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RecordType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RecordType as record_type")
        }
        msg["record_type"] = &val
    }
    
    if model.modified.Bit(propVirtualDnsRecordType_record_ttl_seconds) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RecordTTLSeconds); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RecordTTLSeconds as record_ttl_seconds")
        }
        msg["record_ttl_seconds"] = &val
    }
    
    if model.modified.Bit(propVirtualDnsRecordType_record_mx_preference) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RecordMXPreference); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RecordMXPreference as record_mx_preference")
        }
        msg["record_mx_preference"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *VirtualDnsRecordType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VirtualDnsRecordType) UpdateReferences() error {
    return nil
}


