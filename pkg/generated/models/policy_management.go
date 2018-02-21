
package models
// PolicyManagement



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propPolicyManagement_id_perms int = iota
    propPolicyManagement_display_name int = iota
    propPolicyManagement_annotations int = iota
    propPolicyManagement_perms2 int = iota
    propPolicyManagement_uuid int = iota
    propPolicyManagement_parent_uuid int = iota
    propPolicyManagement_parent_type int = iota
    propPolicyManagement_fq_name int = iota
)

// PolicyManagement 
type PolicyManagement struct {

    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`


    AddressGroups []*AddressGroup `json:"address_groups,omitempty"`
    ApplicationPolicySets []*ApplicationPolicySet `json:"application_policy_sets,omitempty"`
    FirewallPolicys []*FirewallPolicy `json:"firewall_policys,omitempty"`
    FirewallRules []*FirewallRule `json:"firewall_rules,omitempty"`
    ServiceGroups []*ServiceGroup `json:"service_groups,omitempty"`

    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *PolicyManagement) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakePolicyManagement makes PolicyManagement
func MakePolicyManagement() *PolicyManagement{
    return &PolicyManagement{
    //TODO(nati): Apply default
    IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        Perms2: MakePermType2(),
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        
        modified: big.NewInt(0),
    }
}



// MakePolicyManagementSlice makes a slice of PolicyManagement
func MakePolicyManagementSlice() []*PolicyManagement {
    return []*PolicyManagement{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *PolicyManagement) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[])
    fqn := []string{}
    
    return fqn
}

func (model *PolicyManagement) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *PolicyManagement) GetDefaultName() string {
    return strings.Replace("default-policy_management", "_", "-", -1)
}

func (model *PolicyManagement) GetType() string {
    return strings.Replace("policy_management", "_", "-", -1)
}

func (model *PolicyManagement) GetFQName() []string {
    return model.FQName
}

func (model *PolicyManagement) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *PolicyManagement) GetParentType() string {
    return model.ParentType
}

func (model *PolicyManagement) GetUuid() string {
    return model.UUID
}

func (model *PolicyManagement) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *PolicyManagement) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *PolicyManagement) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *PolicyManagement) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *PolicyManagement) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propPolicyManagement_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propPolicyManagement_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propPolicyManagement_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propPolicyManagement_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propPolicyManagement_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propPolicyManagement_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propPolicyManagement_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propPolicyManagement_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *PolicyManagement) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *PolicyManagement) UpdateReferences() error {
    return nil
}


