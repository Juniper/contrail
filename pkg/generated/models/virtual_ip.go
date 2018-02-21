
package models
// VirtualIP



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propVirtualIP_virtual_ip_properties int = iota
    propVirtualIP_parent_type int = iota
    propVirtualIP_id_perms int = iota
    propVirtualIP_uuid int = iota
    propVirtualIP_perms2 int = iota
    propVirtualIP_parent_uuid int = iota
    propVirtualIP_fq_name int = iota
    propVirtualIP_display_name int = iota
    propVirtualIP_annotations int = iota
)

// VirtualIP 
type VirtualIP struct {

    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    UUID string `json:"uuid,omitempty"`
    VirtualIPProperties *VirtualIpType `json:"virtual_ip_properties,omitempty"`
    ParentType string `json:"parent_type,omitempty"`

    LoadbalancerPoolRefs []*VirtualIPLoadbalancerPoolRef `json:"loadbalancer_pool_refs,omitempty"`
    VirtualMachineInterfaceRefs []*VirtualIPVirtualMachineInterfaceRef `json:"virtual_machine_interface_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// VirtualIPLoadbalancerPoolRef references each other
type VirtualIPLoadbalancerPoolRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}

// VirtualIPVirtualMachineInterfaceRef references each other
type VirtualIPVirtualMachineInterfaceRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *VirtualIP) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeVirtualIP makes VirtualIP
func MakeVirtualIP() *VirtualIP{
    return &VirtualIP{
    //TODO(nati): Apply default
    VirtualIPProperties: MakeVirtualIpType(),
        ParentType: "",
        IDPerms: MakeIdPermsType(),
        UUID: "",
        ParentUUID: "",
        FQName: []string{},
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        
        modified: big.NewInt(0),
    }
}



// MakeVirtualIPSlice makes a slice of VirtualIP
func MakeVirtualIPSlice() []*VirtualIP {
    return []*VirtualIP{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *VirtualIP) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc4200aebe0])
    fqn := []string{}
    
    fqn = Project{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *VirtualIP) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-project", "_", "-", -1)
}

func (model *VirtualIP) GetDefaultName() string {
    return strings.Replace("default-virtual_ip", "_", "-", -1)
}

func (model *VirtualIP) GetType() string {
    return strings.Replace("virtual_ip", "_", "-", -1)
}

func (model *VirtualIP) GetFQName() []string {
    return model.FQName
}

func (model *VirtualIP) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *VirtualIP) GetParentType() string {
    return model.ParentType
}

func (model *VirtualIP) GetUuid() string {
    return model.UUID
}

func (model *VirtualIP) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *VirtualIP) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *VirtualIP) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *VirtualIP) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *VirtualIP) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propVirtualIP_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propVirtualIP_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propVirtualIP_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propVirtualIP_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propVirtualIP_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propVirtualIP_virtual_ip_properties) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.VirtualIPProperties); err != nil {
            return nil, errors.Wrap(err, "Marshal of: VirtualIPProperties as virtual_ip_properties")
        }
        msg["virtual_ip_properties"] = &val
    }
    
    if model.modified.Bit(propVirtualIP_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propVirtualIP_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propVirtualIP_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *VirtualIP) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *VirtualIP) UpdateReferences() error {
    return nil
}


