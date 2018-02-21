
package models
// SecurityLoggingObject



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propSecurityLoggingObject_security_logging_object_rate int = iota
    propSecurityLoggingObject_perms2 int = iota
    propSecurityLoggingObject_uuid int = iota
    propSecurityLoggingObject_parent_uuid int = iota
    propSecurityLoggingObject_display_name int = iota
    propSecurityLoggingObject_annotations int = iota
    propSecurityLoggingObject_security_logging_object_rules int = iota
    propSecurityLoggingObject_parent_type int = iota
    propSecurityLoggingObject_fq_name int = iota
    propSecurityLoggingObject_id_perms int = iota
)

// SecurityLoggingObject 
type SecurityLoggingObject struct {

    SecurityLoggingObjectRate int `json:"security_logging_object_rate,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    SecurityLoggingObjectRules *SecurityLoggingObjectRuleListType `json:"security_logging_object_rules,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    FQName []string `json:"fq_name,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`

    SecurityGroupRefs []*SecurityLoggingObjectSecurityGroupRef `json:"security_group_refs,omitempty"`
    NetworkPolicyRefs []*SecurityLoggingObjectNetworkPolicyRef `json:"network_policy_refs,omitempty"`


    client controller.ObjectInterface
    modified* big.Int
}


// SecurityLoggingObjectSecurityGroupRef references each other
type SecurityLoggingObjectSecurityGroupRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
    Attr *SecurityLoggingObjectRuleListType
    
}

// SecurityLoggingObjectNetworkPolicyRef references each other
type SecurityLoggingObjectNetworkPolicyRef struct {
    UUID string `json:"uuid"`
    To   []string `json:"to"`//FQDN
    
    Attr *SecurityLoggingObjectRuleListType
    
}


// String returns json representation of the object
func (model *SecurityLoggingObject) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeSecurityLoggingObject makes SecurityLoggingObject
func MakeSecurityLoggingObject() *SecurityLoggingObject{
    return &SecurityLoggingObject{
    //TODO(nati): Apply default
    SecurityLoggingObjectRules: MakeSecurityLoggingObjectRuleListType(),
        ParentType: "",
        FQName: []string{},
        IDPerms: MakeIdPermsType(),
        DisplayName: "",
        Annotations: MakeKeyValuePairs(),
        SecurityLoggingObjectRate: 0,
        Perms2: MakePermType2(),
        UUID: "",
        ParentUUID: "",
        
        modified: big.NewInt(0),
    }
}



// MakeSecurityLoggingObjectSlice makes a slice of SecurityLoggingObject
func MakeSecurityLoggingObjectSlice() []*SecurityLoggingObject {
    return []*SecurityLoggingObject{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *SecurityLoggingObject) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[global_vrouter_config:0xc4202e5180 project:0xc4202e50e0])
    fqn := []string{}
    
    fqn = nil
    
    return fqn
}

func (model *SecurityLoggingObject) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("", "_", "-", -1)
}

func (model *SecurityLoggingObject) GetDefaultName() string {
    return strings.Replace("default-security_logging_object", "_", "-", -1)
}

func (model *SecurityLoggingObject) GetType() string {
    return strings.Replace("security_logging_object", "_", "-", -1)
}

func (model *SecurityLoggingObject) GetFQName() []string {
    return model.FQName
}

func (model *SecurityLoggingObject) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *SecurityLoggingObject) GetParentType() string {
    return model.ParentType
}

func (model *SecurityLoggingObject) GetUuid() string {
    return model.UUID
}

func (model *SecurityLoggingObject) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *SecurityLoggingObject) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *SecurityLoggingObject) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *SecurityLoggingObject) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *SecurityLoggingObject) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propSecurityLoggingObject_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propSecurityLoggingObject_security_logging_object_rules) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SecurityLoggingObjectRules); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SecurityLoggingObjectRules as security_logging_object_rules")
        }
        msg["security_logging_object_rules"] = &val
    }
    
    if model.modified.Bit(propSecurityLoggingObject_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propSecurityLoggingObject_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propSecurityLoggingObject_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propSecurityLoggingObject_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    if model.modified.Bit(propSecurityLoggingObject_security_logging_object_rate) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SecurityLoggingObjectRate); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SecurityLoggingObjectRate as security_logging_object_rate")
        }
        msg["security_logging_object_rate"] = &val
    }
    
    if model.modified.Bit(propSecurityLoggingObject_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propSecurityLoggingObject_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propSecurityLoggingObject_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *SecurityLoggingObject) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *SecurityLoggingObject) UpdateReferences() error {
    return nil
}


