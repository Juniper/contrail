
package models
// VirtualDNSRecord



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propVirtualDNSRecord_parent_uuid int = iota
    propVirtualDNSRecord_parent_type int = iota
    propVirtualDNSRecord_fq_name int = iota
    propVirtualDNSRecord_id_perms int = iota
    propVirtualDNSRecord_display_name int = iota
    propVirtualDNSRecord_perms2 int = iota
    propVirtualDNSRecord_virtual_DNS_record_data int = iota
    propVirtualDNSRecord_uuid int = iota
    propVirtualDNSRecord_annotations int = iota
)

// VirtualDNSRecord 
type VirtualDNSRecord struct {

    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    UUID string `json:"uuid,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    VirtualDNSRecordData *VirtualDnsRecordType `json:"virtual_DNS_record_data,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *VirtualDNSRecord) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeVirtualDNSRecord makes VirtualDNSRecord
func MakeVirtualDNSRecord() *VirtualDNSRecord{
    return &VirtualDNSRecord{
    //TODO(nati): Apply default
    VirtualDNSRecordData: MakeVirtualDnsRecordType(),
        UUID: "",
        Annotations: MakeKeyValuePairs(),
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Perms2: MakePermType2(),
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        
        modified: big.NewInt(0),
    }
}



// MakeVirtualDNSRecordSlice makes a slice of VirtualDNSRecord
func MakeVirtualDNSRecordSlice() []*VirtualDNSRecord {
    return []*VirtualDNSRecord{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VirtualDNSRecord) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[virtual_DNS:0xc4200ae780])
    fqn := []string{}
    
    fqn = VirtualDNS{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *VirtualDNSRecord) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-virtual_DNS", "_", "-", -1)
}

func (model *VirtualDNSRecord) GetDefaultName() string {
    return strings.Replace("default-virtual_DNS_record", "_", "-", -1)
}

func (model *VirtualDNSRecord) GetType() string {
    return strings.Replace("virtual_DNS_record", "_", "-", -1)
}

func (model *VirtualDNSRecord) GetFQName() []string {
    return model.FQName
}

func (model *VirtualDNSRecord) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VirtualDNSRecord) GetParentType() string {
    return model.ParentType
}

func (model *VirtualDNSRecord) GetUuid() string {
    return model.UUID
}

func (model *VirtualDNSRecord) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VirtualDNSRecord) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VirtualDNSRecord) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VirtualDNSRecord) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VirtualDNSRecord) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propVirtualDNSRecord_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propVirtualDNSRecord_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propVirtualDNSRecord_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propVirtualDNSRecord_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propVirtualDNSRecord_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propVirtualDNSRecord_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propVirtualDNSRecord_virtual_DNS_record_data) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualDNSRecordData); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualDNSRecordData as virtual_DNS_record_data")
        }
        msg["virtual_DNS_record_data"] = &val
    }
    
    if model.modified.Bit(propVirtualDNSRecord_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propVirtualDNSRecord_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *VirtualDNSRecord) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VirtualDNSRecord) UpdateReferences() error {
    return nil
}


