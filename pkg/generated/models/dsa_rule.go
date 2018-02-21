
package models
// DsaRule



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propDsaRule_uuid int = iota
    propDsaRule_id_perms int = iota
    propDsaRule_display_name int = iota
    propDsaRule_perms2 int = iota
    propDsaRule_annotations int = iota
    propDsaRule_parent_uuid int = iota
    propDsaRule_dsa_rule_entry int = iota
    propDsaRule_parent_type int = iota
    propDsaRule_fq_name int = iota
)

// DsaRule 
type DsaRule struct {

    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    DsaRuleEntry *DiscoveryServiceAssignmentType `json:"dsa_rule_entry,omitempty"`
    ParentType string `json:"parent_type,omitempty"`



    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *DsaRule) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeDsaRule makes DsaRule
func MakeDsaRule() *DsaRule{
    return &DsaRule{
    //TODO(nati): Apply default
    Perms2: MakePermType2(),
        UUID: "",
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        FQName: []string{},
        Annotations: MakeKeyValuePairs(),
        ParentUUID: "",
        DsaRuleEntry: MakeDiscoveryServiceAssignmentType(),
        ParentType: "",
        
        modified: big.NewInt(0),
    }
}



// MakeDsaRuleSlice makes a slice of DsaRule
func MakeDsaRuleSlice() []*DsaRule {
    return []*DsaRule{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *DsaRule) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[discovery_service_assignment:0xc4201839a0])
    fqn := []string{}
    
    fqn = DiscoveryServiceAssignment{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *DsaRule) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-discovery_service_assignment", "_", "-", -1)
}

func (model *DsaRule) GetDefaultName() string {
    return strings.Replace("default-dsa_rule", "_", "-", -1)
}

func (model *DsaRule) GetType() string {
    return strings.Replace("dsa_rule", "_", "-", -1)
}

func (model *DsaRule) GetFQName() []string {
    return model.FQName
}

func (model *DsaRule) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *DsaRule) GetParentType() string {
    return model.ParentType
}

func (model *DsaRule) GetUuid() string {
    return model.UUID
}

func (model *DsaRule) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *DsaRule) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *DsaRule) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *DsaRule) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *DsaRule) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propDsaRule_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propDsaRule_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propDsaRule_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propDsaRule_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propDsaRule_dsa_rule_entry) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DsaRuleEntry); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DsaRuleEntry as dsa_rule_entry")
        }
        msg["dsa_rule_entry"] = &val
    }
    
    if model.modified.Bit(propDsaRule_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propDsaRule_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propDsaRule_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propDsaRule_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *DsaRule) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *DsaRule) UpdateReferences() error {
    return nil
}


