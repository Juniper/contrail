
package models
// Domain



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propDomain_fq_name int = iota
    propDomain_id_perms int = iota
    propDomain_display_name int = iota
    propDomain_annotations int = iota
    propDomain_domain_limits int = iota
    propDomain_uuid int = iota
    propDomain_parent_type int = iota
    propDomain_parent_uuid int = iota
    propDomain_perms2 int = iota
)

// Domain 
type Domain struct {

    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    DomainLimits *DomainLimitsType `json:"domain_limits,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`


    APIAccessLists []*APIAccessList `json:"api_access_lists,omitempty"`
    Namespaces []*Namespace `json:"namespaces,omitempty"`
    Projects []*Project `json:"projects,omitempty"`
    ServiceTemplates []*ServiceTemplate `json:"service_templates,omitempty"`
    VirtualDNSs []*VirtualDNS `json:"virtual_DNSs,omitempty"`

    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *Domain) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeDomain makes Domain
func MakeDomain() *Domain{
    return &Domain{
    //TODO(nati): Apply default
    ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        DomainLimits: MakeDomainLimitsType(),
        UUID: "",
        ParentUUID: "",
        Perms2: MakePermType2(),
        
        modified: big.NewInt(0),
    }
}



// MakeDomainSlice makes a slice of Domain
func MakeDomainSlice() []*Domain {
    return []*Domain{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *Domain) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[config_root:0xc420183900])
    fqn := []string{}
    
    fqn = ConfigRoot{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *Domain) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-config_root", "_", "-", -1)
}

func (model *Domain) GetDefaultName() string {
    return strings.Replace("default-domain", "_", "-", -1)
}

func (model *Domain) GetType() string {
    return strings.Replace("domain", "_", "-", -1)
}

func (model *Domain) GetFQName() []string {
    return model.FQName
}

func (model *Domain) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *Domain) GetParentType() string {
    return model.ParentType
}

func (model *Domain) GetUuid() string {
    return model.UUID
}

func (model *Domain) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *Domain) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *Domain) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *Domain) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *Domain) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propDomain_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propDomain_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propDomain_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propDomain_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propDomain_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propDomain_domain_limits) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DomainLimits); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DomainLimits as domain_limits")
        }
        msg["domain_limits"] = &val
    }
    
    if model.modified.Bit(propDomain_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propDomain_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propDomain_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *Domain) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *Domain) UpdateReferences() error {
    return nil
}


