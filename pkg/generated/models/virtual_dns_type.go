
package models
// VirtualDnsType



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propVirtualDnsType_reverse_resolution int = iota
    propVirtualDnsType_default_ttl_seconds int = iota
    propVirtualDnsType_record_order int = iota
    propVirtualDnsType_floating_ip_record int = iota
    propVirtualDnsType_domain_name int = iota
    propVirtualDnsType_external_visible int = iota
    propVirtualDnsType_next_virtual_DNS int = iota
    propVirtualDnsType_dynamic_records_from_client int = iota
)

// VirtualDnsType 
type VirtualDnsType struct {

    FloatingIPRecord FloatingIpDnsNotation `json:"floating_ip_record,omitempty"`
    DomainName string `json:"domain_name,omitempty"`
    ExternalVisible bool `json:"external_visible"`
    NextVirtualDNS string `json:"next_virtual_DNS,omitempty"`
    DynamicRecordsFromClient bool `json:"dynamic_records_from_client"`
    ReverseResolution bool `json:"reverse_resolution"`
    DefaultTTLSeconds int `json:"default_ttl_seconds,omitempty"`
    RecordOrder DnsRecordOrderType `json:"record_order,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *VirtualDnsType) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeVirtualDnsType makes VirtualDnsType
func MakeVirtualDnsType() *VirtualDnsType{
    return &VirtualDnsType{
    //TODO(nati): Apply default
    RecordOrder: MakeDnsRecordOrderType(),
        FloatingIPRecord: MakeFloatingIpDnsNotation(),
        DomainName: "",
        ExternalVisible: false,
        NextVirtualDNS: "",
        DynamicRecordsFromClient: false,
        ReverseResolution: false,
        DefaultTTLSeconds: 0,
        
        modified: big.NewInt(0),
    }
}



// MakeVirtualDnsTypeSlice makes a slice of VirtualDnsType
func MakeVirtualDnsTypeSlice() []*VirtualDnsType {
    return []*VirtualDnsType{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VirtualDnsType) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA <nil>)
    fqn := []string{}
    
    return fqn
}

func (model *VirtualDnsType) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *VirtualDnsType) GetDefaultName() string {
    return strings.Replace("default-", "_", "-", -1)
}

func (model *VirtualDnsType) GetType() string {
    return strings.Replace("", "_", "-", -1)
}

func (model *VirtualDnsType) GetFQName() []string {
    return model.FQName
}

func (model *VirtualDnsType) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VirtualDnsType) GetParentType() string {
    return model.ParentType
}

func (model *VirtualDnsType) GetUuid() string {
    return model.UUID
}

func (model *VirtualDnsType) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VirtualDnsType) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VirtualDnsType) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VirtualDnsType) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VirtualDnsType) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propVirtualDnsType_dynamic_records_from_client) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DynamicRecordsFromClient); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DynamicRecordsFromClient as dynamic_records_from_client")
        }
        msg["dynamic_records_from_client"] = &val
    }
    
    if model.modified.Bit(propVirtualDnsType_reverse_resolution) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ReverseResolution); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ReverseResolution as reverse_resolution")
        }
        msg["reverse_resolution"] = &val
    }
    
    if model.modified.Bit(propVirtualDnsType_default_ttl_seconds) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DefaultTTLSeconds); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DefaultTTLSeconds as default_ttl_seconds")
        }
        msg["default_ttl_seconds"] = &val
    }
    
    if model.modified.Bit(propVirtualDnsType_record_order) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.RecordOrder); err != nil {
            return nil, errors.Wrap(err, "Marshal of: RecordOrder as record_order")
        }
        msg["record_order"] = &val
    }
    
    if model.modified.Bit(propVirtualDnsType_floating_ip_record) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FloatingIPRecord); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FloatingIPRecord as floating_ip_record")
        }
        msg["floating_ip_record"] = &val
    }
    
    if model.modified.Bit(propVirtualDnsType_domain_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DomainName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DomainName as domain_name")
        }
        msg["domain_name"] = &val
    }
    
    if model.modified.Bit(propVirtualDnsType_external_visible) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ExternalVisible); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ExternalVisible as external_visible")
        }
        msg["external_visible"] = &val
    }
    
    if model.modified.Bit(propVirtualDnsType_next_virtual_DNS) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.NextVirtualDNS); err != nil {
            return nil, errors.Wrap(err, "Marshal of: NextVirtualDNS as next_virtual_DNS")
        }
        msg["next_virtual_DNS"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *VirtualDnsType) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VirtualDnsType) UpdateReferences() error {
    return nil
}


