
package models
// AliasIP



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propAliasIP_parent_uuid int = iota
    propAliasIP_fq_name int = iota
    propAliasIP_alias_ip_address_family int = iota
    propAliasIP_annotations int = iota
    propAliasIP_perms2 int = iota
    propAliasIP_id_perms int = iota
    propAliasIP_display_name int = iota
    propAliasIP_alias_ip_address int = iota
    propAliasIP_uuid int = iota
    propAliasIP_parent_type int = iota
)

// AliasIP 
type AliasIP struct {

    ParentType string `json:"parent_type,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    AliasIPAddress IpAddressType `json:"alias_ip_address,omitempty"`
    UUID string `json:"uuid,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    AliasIPAddressFamily IpAddressFamilyType `json:"alias_ip_address_family,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`

    ProjectRefs []*AliasIPProjectRef `json:"project_refs,omitempty"`
    VirtualMachineInterfaceRefs []*AliasIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// AliasIPProjectRef references each other
type AliasIPProjectRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// AliasIPVirtualMachineInterfaceRef references each other
type AliasIPVirtualMachineInterfaceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *AliasIP) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeAliasIP makes AliasIP
func MakeAliasIP() *AliasIP{
    return &AliasIP{
    //TODO(nati): Apply default
    AliasIPAddress: MakeIpAddressType(),
        UUID: "",
        ParentType: "",
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        AliasIPAddressFamily: MakeIpAddressFamilyType(),
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        ParentUUID: "",
        FQName: []string{},
        
        modified: big.NewInt(0),
    }
}



// MakeAliasIPSlice makes a slice of AliasIP
func MakeAliasIPSlice() []*AliasIP {
    return []*AliasIP{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *AliasIP) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[alias_ip_pool:0xc420182d20])
    fqn := []string{}
    
    fqn = AliasIPPool{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *AliasIP) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-alias_ip_pool", "_", "-", -1)
}

func (model *AliasIP) GetDefaultName() string {
    return strings.Replace("default-alias_ip", "_", "-", -1)
}

func (model *AliasIP) GetType() string {
    return strings.Replace("alias_ip", "_", "-", -1)
}

func (model *AliasIP) GetFQName() []string {
    return model.FQName
}

func (model *AliasIP) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *AliasIP) GetParentType() string {
    return model.ParentType
}

func (model *AliasIP) GetUuid() string {
    return model.UUID
}

func (model *AliasIP) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *AliasIP) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *AliasIP) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *AliasIP) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *AliasIP) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propAliasIP_alias_ip_address_family) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AliasIPAddressFamily); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AliasIPAddressFamily as alias_ip_address_family")
        }
        msg["alias_ip_address_family"] = &val
    }
    
    if model.modified.Bit(propAliasIP_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propAliasIP_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propAliasIP_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propAliasIP_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propAliasIP_alias_ip_address) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.AliasIPAddress); err != nil {
            return nil, errors.Wrap(err, "Marshal of: AliasIPAddress as alias_ip_address")
        }
        msg["alias_ip_address"] = &val
    }
    
    if model.modified.Bit(propAliasIP_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propAliasIP_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propAliasIP_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propAliasIP_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *AliasIP) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *AliasIP) UpdateReferences() error {
    return nil
}


