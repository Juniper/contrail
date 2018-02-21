
package models
// SecurityGroup



import (
    "encoding/json"
    "strings"
    "math/big"
    "github.com/pkg/errors"
    "github.com/Juniper/contrail/pkg/controller"
)

const (
    propSecurityGroup_fq_name int = iota
    propSecurityGroup_display_name int = iota
    propSecurityGroup_security_group_entries int = iota
    propSecurityGroup_security_group_id int = iota
    propSecurityGroup_perms2 int = iota
    propSecurityGroup_uuid int = iota
    propSecurityGroup_parent_uuid int = iota
    propSecurityGroup_parent_type int = iota
    propSecurityGroup_configured_security_group_id int = iota
    propSecurityGroup_id_perms int = iota
    propSecurityGroup_annotations int = iota
)

// SecurityGroup 
type SecurityGroup struct {

    FQName []string `json:"fq_name,omitempty"`
    DisplayName string `json:"display_name,omitempty"`
    SecurityGroupEntries *PolicyEntriesType `json:"security_group_entries,omitempty"`
    SecurityGroupID SecurityGroupIdType `json:"security_group_id,omitempty"`
    Perms2 *PermType2 `json:"perms2,omitempty"`
    UUID string `json:"uuid,omitempty"`
    ParentUUID string `json:"parent_uuid,omitempty"`
    ParentType string `json:"parent_type,omitempty"`
    ConfiguredSecurityGroupID ConfiguredSecurityGroupIdType `json:"configured_security_group_id,omitempty"`
    IDPerms *IdPermsType `json:"id_perms,omitempty"`
    Annotations *KeyValuePairs `json:"annotations,omitempty"`


    AccessControlLists []*AccessControlList `json:"access_control_lists,omitempty"`

    client controller.ObjectInterface
    modified* big.Int
}



// String returns json representation of the object
func (model *SecurityGroup) String() string {
    b, _ := json.Marshal(model)
    return string(b)
}

// MakeSecurityGroup makes SecurityGroup
func MakeSecurityGroup() *SecurityGroup{
    return &SecurityGroup{
    //TODO(nati): Apply default
    SecurityGroupID: MakeSecurityGroupIdType(),
        Perms2: MakePermType2(),
        UUID: "",
        ParentUUID: "",
        ParentType: "",
        FQName: []string{},
        DisplayName: "",
        SecurityGroupEntries: MakePolicyEntriesType(),
        IDPerms: MakeIdPermsType(),
        Annotations: MakeKeyValuePairs(),
        ConfiguredSecurityGroupID: MakeConfiguredSecurityGroupIdType(),
        
        modified: big.NewInt(0),
    }
}



// MakeSecurityGroupSlice makes a slice of SecurityGroup
func MakeSecurityGroupSlice() []*SecurityGroup {
    return []*SecurityGroup{}
}

// Implementation of IObject interface for contrail controller resources management

func (model *SecurityGroup) GetDefaultParent() []string {
    // PArents: +v%!(EXTRA map[string]*common.Reference=map[project:0xc4202e4f00])
    fqn := []string{}
    
    fqn = Project{}.GetDefaultParent()
    fqn = append(fqn, model.GetDefaultParentName())
    
    
    return fqn
}

func (model *SecurityGroup) GetDefaultParentName() string {
    // This might be wrong for some resources
    return strings.Replace("default-project", "_", "-", -1)
}

func (model *SecurityGroup) GetDefaultName() string {
    return strings.Replace("default-security_group", "_", "-", -1)
}

func (model *SecurityGroup) GetType() string {
    return strings.Replace("security_group", "_", "-", -1)
}

func (model *SecurityGroup) GetFQName() []string {
    return model.FQName
}

func (model *SecurityGroup) GetName() string {
    n := len(model.FQName)
    if (n == 0) {
        return ""
    }
    return model.FQName[n-1]
}

func (model *SecurityGroup) GetParentType() string {
    return model.ParentType
}

func (model *SecurityGroup) GetUuid() string {
    return model.UUID
}

func (model *SecurityGroup) GetHref() string {
    return model.client.GetServerUrl() + model.GetType() + "/" + model.UUID
}

func (model *SecurityGroup) SetName(name string) {
    if len(model.FQName) == 0 {
        fqname := model.GetDefaultParent()
        fqname = append(fqname, name)
        model.SetFQName(model.GetParentType(), fqname)
    } else {
        n := len(model.FQName) -1
        model.FQName[n] = name
    }
}

func (model *SecurityGroup) SetFQName(parent string, fqname []string) {
    model.ParentType = parent
    n := len(fqname) 
    model.FQName = make([]string, 0, n)
    model.FQName = append(model.FQName, fqname...)
}

func (model *SecurityGroup) SetClient(cli controller.ObjectInterface) {
    model.client = cli
}

func (model *SecurityGroup) UpdateObject() ([]byte, error) {
    msg := map[string]*json.RawMessage{}

    if model.modified.Bit(propSecurityGroup_configured_security_group_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ConfiguredSecurityGroupID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ConfiguredSecurityGroupID as configured_security_group_id")
        }
        msg["configured_security_group_id"] = &val
    }
    
    if model.modified.Bit(propSecurityGroup_id_perms) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.IDPerms); err != nil {
            return nil, errors.Wrap(err, "Marshal of: IDPerms as id_perms")
        }
        msg["id_perms"] = &val
    }
    
    if model.modified.Bit(propSecurityGroup_annotations) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Annotations); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Annotations as annotations")
        }
        msg["annotations"] = &val
    }
    
    if model.modified.Bit(propSecurityGroup_security_group_entries) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SecurityGroupEntries); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SecurityGroupEntries as security_group_entries")
        }
        msg["security_group_entries"] = &val
    }
    
    if model.modified.Bit(propSecurityGroup_security_group_id) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.SecurityGroupID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: SecurityGroupID as security_group_id")
        }
        msg["security_group_id"] = &val
    }
    
    if model.modified.Bit(propSecurityGroup_perms2) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.Perms2); err != nil {
            return nil, errors.Wrap(err, "Marshal of: Perms2 as perms2")
        }
        msg["perms2"] = &val
    }
    
    if model.modified.Bit(propSecurityGroup_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.UUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: UUID as uuid")
        }
        msg["uuid"] = &val
    }
    
    if model.modified.Bit(propSecurityGroup_parent_uuid) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentUUID); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentUUID as parent_uuid")
        }
        msg["parent_uuid"] = &val
    }
    
    if model.modified.Bit(propSecurityGroup_parent_type) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.ParentType); err != nil {
            return nil, errors.Wrap(err, "Marshal of: ParentType as parent_type")
        }
        msg["parent_type"] = &val
    }
    
    if model.modified.Bit(propSecurityGroup_fq_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.FQName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: FQName as fq_name")
        }
        msg["fq_name"] = &val
    }
    
    if model.modified.Bit(propSecurityGroup_display_name) != 0 {
        var val json.RawMessage
        if val, err := json.Marshal(&model.DisplayName); err != nil {
            return nil, errors.Wrap(err, "Marshal of: DisplayName as display_name")
        }
        msg["display_name"] = &val
    }
    
    return json.Marshal(msg)
}

func (model *SecurityGroup) UpdateDone() {
    model.modified.SetInt64(0)
}

func (model *SecurityGroup) UpdateReferences() error {
    return nil
}


