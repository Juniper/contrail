
package models
// VirtualDNS



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propVirtualDNS_uuid int = iota
    propVirtualDNS_virtual_DNS_data int = iota
    propVirtualDNS_fq_name int = iota
    propVirtualDNS_id_perms int = iota
    propVirtualDNS_perms2 int = iota
    propVirtualDNS_parent_uuid int = iota
    propVirtualDNS_parent_type int = iota
    propVirtualDNS_display_name int = iota
    propVirtualDNS_annotations int = iota
)

// VirtualDNS 
type VirtualDNS struct {

    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    VirtualDNSData *VirtualDnsType `json:"virtual_DNS_data,omitempty"`
    FQName []string `json:"fq_name,omitempty"`


    VirtualDNSRecords []*VirtualDNSRecord `json:"virtual_DNS_records,omitempty"`

    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *VirtualDNS) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeVirtualDNS makes VirtualDNS
func MakeVirtualDNS() *VirtualDNS{
    return &VirtualDNS{
    //TODO(nati): Apply default
    VirtualDNSData: MakeVirtualDnsType(),
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        Perms2: MakePermType2(),
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        
        modified: big.NewInt(0),
    }
}



// MakeVirtualDNSSlice makes a slice of VirtualDNS
func MakeVirtualDNSSlice() []*VirtualDNS {
    return []*VirtualDNS{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VirtualDNS) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[domain:0xc4200ae820])
    fqn := []string{}
    
    fqn = Domain{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *VirtualDNS) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-domain", "_", "-", -1)
}

func (model *VirtualDNS) GetDefaultName() string {
    return strings.Replace("default-virtual_DNS", "_", "-", -1)
}

func (model *VirtualDNS) GetType() string {
    return strings.Replace("virtual_DNS", "_", "-", -1)
}

func (model *VirtualDNS) GetFQName() []string {
    return model.FQName
}

func (model *VirtualDNS) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VirtualDNS) GetParentType() string {
    return model.ParentType
}

func (model *VirtualDNS) GetUuid() string {
    return model.UUID
}

func (model *VirtualDNS) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VirtualDNS) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VirtualDNS) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VirtualDNS) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VirtualDNS) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propVirtualDNS_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propVirtualDNS_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propVirtualDNS_virtual_DNS_data) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualDNSData); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualDNSData as virtual_DNS_data")
        }
        msg["virtual_DNS_data"] = &val
    }
    
    if model.modified.Bit(propVirtualDNS_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propVirtualDNS_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propVirtualDNS_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propVirtualDNS_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propVirtualDNS_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propVirtualDNS_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *VirtualDNS) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VirtualDNS) UpdateReferences() error {
    return nil
}


