
package models
// FirewallPolicy



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propFirewallPolicy_perms2 int = iota
    propFirewallPolicy_uuid int = iota
    propFirewallPolicy_parent_uuid int = iota
    propFirewallPolicy_parent_type int = iota
    propFirewallPolicy_fq_name int = iota
    propFirewallPolicy_id_perms int = iota
    propFirewallPolicy_display_name int = iota
    propFirewallPolicy_annotations int = iota
)

// FirewallPolicy 
type FirewallPolicy struct {

    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`

    FirewallRuleRefs []*FirewallPolicyFirewallRuleRef `json:"firewall_rule_refs,omitempty"`
    SecurityLoggingObjectRefs []*FirewallPolicySecurityLoggingObjectRef `json:"security_logging_object_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// FirewallPolicyFirewallRuleRef references each other
type FirewallPolicyFirewallRuleRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
    Attr *FirewallSequence
    
}

// FirewallPolicySecurityLoggingObjectRef references each other
type FirewallPolicySecurityLoggingObjectRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
}


// String returns json representation of the object
func (model *FirewallPolicy) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeFirewallPolicy makes FirewallPolicy
func MakeFirewallPolicy() *FirewallPolicy{
    return &FirewallPolicy{
    //TODO(nati): Apply default
    FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        
        modified: big.NewInt(0),
    }
}



// MakeFirewallPolicySlice makes a slice of FirewallPolicy
func MakeFirewallPolicySlice() []*FirewallPolicy {
    return []*FirewallPolicy{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *FirewallPolicy) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc420183cc0 policy_management:0xc420183d60])
    fqn := []string{}
    
    fqn = nil
    
    return fqn
}

func (model *FirewallPolicy) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *FirewallPolicy) GetDefaultName() string {
    return strings.Replace("default-firewall_policy", "_", "-", -1)
}

func (model *FirewallPolicy) GetType() string {
    return strings.Replace("firewall_policy", "_", "-", -1)
}

func (model *FirewallPolicy) GetFQName() []string {
    return model.FQName
}

func (model *FirewallPolicy) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *FirewallPolicy) GetParentType() string {
    return model.ParentType
}

func (model *FirewallPolicy) GetUuid() string {
    return model.UUID
}

func (model *FirewallPolicy) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *FirewallPolicy) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *FirewallPolicy) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *FirewallPolicy) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *FirewallPolicy) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propFirewallPolicy_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propFirewallPolicy_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propFirewallPolicy_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propFirewallPolicy_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propFirewallPolicy_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propFirewallPolicy_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propFirewallPolicy_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propFirewallPolicy_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *FirewallPolicy) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *FirewallPolicy) UpdateReferences() error {
    return nil
}


